#!/usr/bin/env bash

# Copyright 2021-2026 Zenauth Ltd.
# SPDX-License-Identifier: Apache-2.0

# Memory provisioning sweep.
# Runs at ONE policy count:
#   Step 0  measure the anchor floor R (Sys-HeapReleased) inline, no load
#   Edge 1  vary GOGC, no limit                          → sizing curve
#   Edge 2  GOGC=off, vary GOMEMLIMIT = mult x R         → backstop cost
#   Valid   GOGC=x + generous GOMEMLIMIT (cap loose)     → reproduces Edge-1
# Per arm: restart Cerbos with the knobs, reset VmHWM, run loadtest.sh -e, capture peak
# RSS (VmHWM), per-phase GC counters, and ghz throughput/p99. Emits the tables.
# See reports/loadtest-memory-plan.md and reports/docs/gc-metrics.md.

set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
# shellcheck source=env.sh
source "${SCRIPT_DIR}/env.sh"

require_running_vms "$PDP_VM" "$CLIENT_VM"

# --- PDP internal IP (used by env.sh helpers via the exported PDP_IP) ---
if [[ -n "${TERRAFORM_DIR:-}" ]]; then
  PDP_IP=$(terraform -chdir="$TERRAFORM_DIR" output -raw pdp_internal_ip)
else
  PDP_IP=$(gcloud compute instances describe "$PDP_VM" \
    --zone="$GCP_ZONE" --project="$GCP_PROJECT" \
    --format='get(networkInterfaces[0].networkIP)')
fi
export PDP_IP
log "PDP internal IP: ${PDP_IP}"

# --- Arms / parameters (overridable) ---
NUM_POLICIES=${NUM_POLICIES:-1000}
read -r -a GOGC_ARMS <<< "${GOGC_ARMS:-100 50 20 10}"
# GOGC_ARMS=()
read -r -a MEMLIMIT_MULTS <<< "${MEMLIMIT_MULTS:-2.0 1.5 1.3 1.15}"
VALID_GOGC=${VALID_GOGC:-50}        # cap-loose validation arm GOGC
RPS=${RPS:-auto}   # per-arm: sustained target = RPS_AUTO_PCT% of that arm's measured throughput
DURATION_SECS=${DURATION_SECS:-120}
ITERATIONS=${ITERATIONS:-1000000}

LOCAL_RESULTS="${SCRIPT_DIR}/../results/gcp/sweep-${NUM_POLICIES}"
rm -rf "$LOCAL_RESULTS"
mkdir -p "$LOCAL_RESULTS"

# run_arm LABEL GOGC GOMEMLIMIT_BYTES [CGROUP_BYTES]
#   Empty GOMEMLIMIT/CGROUP => unset. Restarts Cerbos under the knobs (cgroup = hard
#   MemoryMax when set), resets the VmHWM high-water so it measures the load window
#   (build excluded), runs the load, records peak RSS + outcome (ok/oom/start_failed),
#   and downloads the ghz/GC results.
run_arm() {
  local label="$1" gogc="$2" memlimit="$3" cgroup="${4:-}"
  local armdir="${LOCAL_RESULTS}/${label}"
  mkdir -p "$armdir"
  log "=== arm ${label}: GOGC=${gogc:-default} GOMEMLIMIT=${memlimit:-off} cgroup=${cgroup:-none} ==="
  { echo "label=${label}"; echo "gogc=${gogc}"; echo "memlimit=${memlimit}"; echo "cgroup=${cgroup}"; } > "${armdir}/arm.meta"

  if ! GOGC="$gogc" GOMEMLIMIT="$memlimit" CGROUP_LIMIT="$cgroup" restart_cerbos; then
    err "arm ${label} failed to become healthy — cgroup OOM at build or unhealthy (journalctl -u cerbos-loadtest)"
    echo "start_failed" > "${armdir}/status"
    return 0
  fi
  run_load_and_capture "$armdir"   # load + VmHWM/OOM/scrape/download into the arm dir
}

# --- Step 0: floor quantities, inline, no load (GOGC=100, no limit) ---
# Reads (cold, settled): R = Sys-HeapReleased (runtime-managed floor; the unit GOMEMLIMIT
# is accounted in), settled RSS, the off-runtime offset O = RSS-R (binary/stacks/maps the
# cgroup counts but GOMEMLIMIT does NOT), and BUILD_HWM (build peak RSS). Edge 2 sweeps the
# soft GOMEMLIMIT = mult*R and pairs the hard cgroup = GOMEMLIMIT + O, so the cgroup is
# RSS-correct and sits one offset above the soft cap (soft bites first). See plan §4.4.
log "Step 0: measuring floor quantities (GOGC=100, no limit, no load)..."
GOGC=100 GOMEMLIMIT="" CGROUP_LIMIT="" restart_cerbos
BUILD_HWM=$(pdp_vmhwm_bytes 2>/dev/null || echo 0)
pdp_scrape "${PDP_FLOOR_METRICS[@]}" > "${LOCAL_RESULTS}/floor.txt"
_sys=$(awk '/^go_memstats_sys_bytes/{print $2}' "${LOCAL_RESULTS}/floor.txt")
_rel=$(awk '/^go_memstats_heap_released_bytes/{print $2}' "${LOCAL_RESULTS}/floor.txt")
_rss=$(awk '/^process_resident_memory_bytes/{print $2}' "${LOCAL_RESULTS}/floor.txt")
R=$(awk -v s="${_sys:-0}" -v r="${_rel:-0}" 'BEGIN{printf "%d", s-r}')
# O = RSS - R, clamped >=0 (R can exceed RSS when arenas are reserved-but-unfaulted).
OFFSET=$(awk -v rss="${_rss:-0}" -v rr="${R:-0}" 'BEGIN{o=rss-rr; printf "%d", (o>0?o:0)}')
log "floor: R(Sys-HeapReleased)=$(humanise "${R}") RSS=$(humanise "${_rss:-0}") offset O=$(humanise "${OFFSET}") build-HWM=$(humanise "${BUILD_HWM}")"
{ echo "R_bytes=${R}"; echo "rss_bytes=${_rss:-0}"; echo "offset_bytes=${OFFSET}"; echo "build_hwm_bytes=${BUILD_HWM}"; } > "${LOCAL_RESULTS}/floor.meta"

if [[ "${R:-0}" -le 0 ]]; then
  err "could not determine floor R (Sys/HeapReleased missing — is the metrics.go extension deployed?)"
  exit 1
fi

# --- Edge 1: vary GOGC, no limit ---
for g in "${GOGC_ARMS[@]}"; do
  run_arm "edge1_gogc${g}" "$g" ""
done

# --- Edge 2: GOGC=off; soft GOMEMLIMIT = mult x R, hard cgroup = GOMEMLIMIT + O (RSS).
#     The cgroup sits one off-runtime offset above the soft cap, so the runtime GCs
#     (graceful) before the kernel OOM-kills, and both limits are RSS-correct (the binary
#     lives in O). Shrinking mult tightens the soft cap toward the live set -> thrash, then
#     OOM at the bottom. Guard: skip any arm whose cgroup is below BUILD_HWM (can't build).
#     NOTE: MEMLIMIT_MULTS are multiples of R (the runtime floor where the thrash physics
#     lives), NOT of the cgroup; extend them downward (e.g. 1.0, 0.85) if no arm thrashes. ---

for mult in "${MEMLIMIT_MULTS[@]}"; do
  gml=$(awk -v r="$R" -v m="$mult" 'BEGIN{printf "%d", r*m}')
  box=$(awk -v g="$gml" -v o="$OFFSET" 'BEGIN{printf "%d", g+o}')
  if [[ "$box" -le "${BUILD_HWM:-0}" ]]; then
    log "skipping mult=${mult} (cgroup $(humanise "$box") <= build-HWM $(humanise "${BUILD_HWM}")) — would OOM the build"
    continue
  fi
  run_arm "edge2_m${mult}" "off" "$gml" "$box"
done

# --- Validation: cap loose (GOGC=x; GOMEMLIMIT = 2xR generous, cgroup = +O; should not bind) ---

_valid_gml=$(awk -v r="$R" 'BEGIN{printf "%d", r*2.0}')
_valid_box=$(awk -v g="$_valid_gml" -v o="$OFFSET" 'BEGIN{printf "%d", g+o}')
run_arm "valid_caploose_gogc${VALID_GOGC}" "$VALID_GOGC" "$_valid_gml" "$_valid_box"

# --- Hard-OOM demo: cgroup just above the build high-water, NO GOMEMLIMIT, GOGC=100.
#     The runtime doesn't know the box, so under load the sawtooth grows past it -> cgroup
#     OOM. Demonstrates why the soft GOMEMLIMIT backstop is needed (plan §4.3). ---

_oom_box=$(awk -v h="${BUILD_HWM:-0}" 'BEGIN{printf "%d", h*1.05}')
run_arm "oom_demo_nolimit" "100" "" "$_oom_box"

# NOTE: forced-overload of the *shipped* config (GOGC=x + box, then drive concurrency/
# live-set up until the soft cap binds and degrades gracefully) is still manual — vary
# CONCURRENCY/NUM_POLICIES against a deployed validation arm and watch GC CPU vs OOM.

# --- Emit the §4.6 tables from the per-arm JSON (jq on ghz output; robust) ---
emit_tables() {
  local out="${LOCAL_RESULTS}/summary.md"
  {
    printf '# Provisioning sweep — %s policies\n\n' "$NUM_POLICIES"
    printf 'Floor (cold, no load): R (Sys-HeapReleased) = %s; settled RSS = %s; off-runtime offset O = %s; build high-water = %s.\n' \
      "$(humanise "$R")" "$(humanise "${_rss:-0}")" "$(humanise "${OFFSET}")" "$(humanise "${BUILD_HWM}")"
    printf 'Edge 2 / validation: soft GOMEMLIMIT = mult x R, hard cgroup MemoryMax = GOMEMLIMIT + O (RSS-correct; soft bites before the kernel OOM-kills).\n\n'

    printf '## Edge 1 — sizing (no limit)\n\n'
    printf '| Arm | RSS peak | GC CPU%% | Throughput | p99 (ms) | outcome |\n|---|--:|--:|--:|--:|---|\n'
    for g in "${GOGC_ARMS[@]}"; do _row "edge1_gogc${g}" "GOGC=${g}"; done

    printf '\n## Edge 2 — backstop cost (GOGC=off; GOMEMLIMIT = mult x R, cgroup = GOMEMLIMIT + O)\n\n'
    printf '| Arm | cgroup (=mult·R+O) | RSS peak | GC CPU%% | Throughput | p99 (ms) | outcome |\n|---|--:|--:|--:|--:|--:|---|\n'
    for mult in "${MEMLIMIT_MULTS[@]}"; do
      local gml box
      gml=$(awk -v r="$R" -v m="$mult" 'BEGIN{printf "%d", r*m}')
      box=$(awk -v g="$gml" -v o="${OFFSET:-0}" 'BEGIN{printf "%d", g+o}')
      [[ "$box" -le "${BUILD_HWM:-0}" ]] && continue
      _row "edge2_m${mult}" "mult=${mult} (GOMEMLIMIT $(humanise "$gml"))" "$box"
    done

    printf '\n## Validation\n\n'
    printf '| Arm | cgroup | RSS peak | GC CPU%% | Throughput | p99 (ms) | outcome |\n|---|--:|--:|--:|--:|--:|---|\n'
    _row "valid_caploose_gogc${VALID_GOGC}" "GOGC=${VALID_GOGC}, GOMEMLIMIT=2R" "$(awk -v r="$R" -v o="${OFFSET:-0}" 'BEGIN{printf "%d", r*2.0+o}')"
    _row "oom_demo_nolimit" "GOGC=100, no GOMEMLIMIT" "$(awk -v h="${BUILD_HWM:-0}" 'BEGIN{printf "%d", h*1.05}')"
  } > "$out"
  log "Summary table: ${out}"
  cat "$out"
}

# _row ARMLABEL DISPLAY [SETPOINT_BYTES]
_row() {
  local armdir="${LOCAL_RESULTS}/$1" disp="$2" setpoint="${3:-}"
  local vmhwm rps p99 gccpu outcome
  outcome=$(cat "${armdir}/status" 2>/dev/null || echo "n/a")
  # loadtest.sh rejected this arm as degenerate (auto RPS below RPS_MIN) — overrides "ok".
  [[ -f "${armdir}/${STORE}_rejected" ]] && outcome="degenerate"
  vmhwm=$(cat "${armdir}/vmhwm_bytes.txt" 2>/dev/null || echo "")
  vmhwm=$(awk -v b="${vmhwm:-0}" 'BEGIN{ if (b>0) printf "%.2f GiB", b/1073741824; else printf "n/a" }')
  rps=$(jq -r '.rps // empty' "${armdir}/disk_throughput.json" 2>/dev/null | awk '{printf "%.0f", $1}')
  p99=$(jq -r '[.latencyDistribution[]? | select(.percentage==99) | .latency][0] // empty' "${armdir}/disk_rps.json" 2>/dev/null | awk '{ if ($1!="") printf "%.2f", $1/1e6 }')
  gccpu=$(jq -r '.gc_cpu_pct // empty' "${armdir}/disk_rps_gc.json" 2>/dev/null)
  if [[ -n "$setpoint" ]]; then
    local sp; sp=$(awk -v b="$setpoint" 'BEGIN{printf "%.2f GiB", b/1073741824}')
    printf '| %s | %s | %s | %s%% | %s | %s | %s |\n' "$disp" "$sp" "${vmhwm}" "${gccpu:-n/a}" "${rps:-n/a}" "${p99:-n/a}" "${outcome}"
  else
    printf '| %s | %s | %s%% | %s | %s | %s |\n' "$disp" "${vmhwm}" "${gccpu:-n/a}" "${rps:-n/a}" "${p99:-n/a}" "${outcome}"
  fi
}

emit_tables
log "Sweep complete — per-arm results in ${LOCAL_RESULTS}/"

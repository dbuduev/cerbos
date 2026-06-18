#!/usr/bin/env bash

# Copyright 2021-2026 Zenauth Ltd.
# SPDX-License-Identifier: Apache-2.0

# Memory provisioning sweep.
# Runs at ONE policy count:
#   Step 0  measure the anchor floor R (Sys-HeapReleased) inline, no load
#   Edge 1  vary GOGC, no limit                          → sizing curve
#   Edge 2  GOGC=off, vary GOMEMLIMIT = mult x R         → backstop cost
#   Valid   GOGC=x + generous GOMEMLIMIT, in-envelope    → reproduces Edge-1
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
read -r -a MEMLIMIT_MULTS <<< "${MEMLIMIT_MULTS:-2.0 1.5 1.3 1.15}"
VALID_GOGC=${VALID_GOGC:-50}        # in-envelope validation arm GOGC
RPS=${RPS:-5000}
DURATION_SECS=${DURATION_SECS:-120}
ITERATIONS=${ITERATIONS:-1000000}

FLOOR_METRICS=(
  process_resident_memory_bytes
  go_memstats_heap_alloc_bytes
  go_memstats_heap_inuse_bytes
  go_memstats_sys_bytes
  go_memstats_heap_released_bytes
)

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
  GSSH "$CLIENT_VM" "rm -rf ${REMOTE_BASE}/results/* 2>/dev/null || true"
  pdp_reset_vmhwm   # peak now measures the load window, not the build

  # CPU monitor (PDP) for context
  GSSH "$PDP_VM" "pkill -f 'mpstat -P ALL' 2>/dev/null || true; setsid mpstat -P ALL 1 > /opt/cerbos-loadtest/results/cpu_usage.log 2>&1 < /dev/null &" || true

  # A mid-load cgroup OOM makes the load run error out — tolerate it and detect below.
  GSSH "$CLIENT_VM" <<ENDSSH || log "load run returned non-zero (arm ${label}) — possible OOM mid-load"
set -uo pipefail
. /nix/var/nix/profiles/default/etc/profile.d/nix-daemon.sh
cd ${REMOTE_BASE}
mkdir -p ${REMOTE_BASE}/results
SERVER="${PDP_IP}:3593" \
METRICS_URL="http://${PDP_IP}:3592/_cerbos/metrics" \
WORK_DIR="${REMOTE_BASE}" \
STORE="${STORE}" \
RPS="${RPS}" DURATION_SECS="${DURATION_SECS}" ITERATIONS="${ITERATIONS}" \
${CONCURRENCY:+CONCURRENCY="${CONCURRENCY}"} \
${CONNECTIONS:+CONNECTIONS="${CONNECTIONS}"} \
${REQ_KIND:+REQ_KIND="${REQ_KIND}"} \
${PROTOSET:+PROTOSET="${REMOTE_BASE}/cerbos.protoset"} \
  nix develop --command bash loadtest.sh -e
ENDSSH

  GSSH "$PDP_VM" "pkill -f 'mpstat -P ALL' 2>/dev/null || true" || true

  # Outcome: did the cgroup OOM-kill Cerbos during the load?
  if [[ "$(pdp_cerbos_alive)" == dead ]]; then
    echo "oom" > "${armdir}/status"
    log "arm ${label}: Cerbos died during load — cgroup OOM (MemoryMax=${cgroup:-none})"
  else
    echo "ok" > "${armdir}/status"
  fi

  # Peak RSS (VmHWM) over the load window, and post-load accounting.
  pdp_vmhwm_bytes > "${armdir}/vmhwm_bytes.txt" 2>/dev/null || echo "" > "${armdir}/vmhwm_bytes.txt"
  pdp_scrape "${FLOOR_METRICS[@]}" > "${armdir}/post_metrics.txt" 2>/dev/null || true

  # Download client results (ghz JSON + *_gc.json + summaries) for this arm.
  GSSH "$CLIENT_VM" "tar czf /tmp/arm.tar.gz -C ${REMOTE_BASE}/results ." 2>/dev/null || true
  GSCP "${CLIENT_VM}:/tmp/arm.tar.gz" "/tmp/arm.tar.gz" 2>/dev/null || true
  [[ -f /tmp/arm.tar.gz ]] && { tar xzf /tmp/arm.tar.gz -C "$armdir"; rm -f /tmp/arm.tar.gz; }
}

# --- Step 0: anchor floor R, inline, no load (GOGC=100, no limit) ---
log "Step 0: measuring anchor floor R (GOGC=100, no limit, no load)..."
GOGC=100 GOMEMLIMIT="" CGROUP_LIMIT="" restart_cerbos
BUILD_HWM=$(pdp_vmhwm_bytes 2>/dev/null || echo 0)
pdp_scrape "${FLOOR_METRICS[@]}" > "${LOCAL_RESULTS}/floor.txt"
_sys=$(awk '/^go_memstats_sys_bytes/{print $2}' "${LOCAL_RESULTS}/floor.txt")
_rel=$(awk '/^go_memstats_heap_released_bytes/{print $2}' "${LOCAL_RESULTS}/floor.txt")
R=$(awk -v s="${_sys:-0}" -v r="${_rel:-0}" 'BEGIN{printf "%d", s-r}')
log "anchor R (Sys-HeapReleased) = ${R} bytes; build high-water (VmHWM) = ${BUILD_HWM} bytes"
{ echo "R_bytes=${R}"; echo "build_hwm_bytes=${BUILD_HWM}"; } > "${LOCAL_RESULTS}/floor.meta"

if [[ "${R:-0}" -le 0 ]]; then
  err "could not determine floor R (Sys/HeapReleased missing — is the metrics.go extension deployed?)"
  exit 1
fi

# --- Edge 1: vary GOGC, no limit ---
for g in "${GOGC_ARMS[@]}"; do
  run_arm "edge1_gogc${g}" "$g" ""
done

# --- Edge 2: GOGC=off; cgroup box = mult x R (hard), GOMEMLIMIT = 0.9 x box (soft).
#     Production-faithful pairing; shrinking the box finds the floor (OOM at the bottom).
#     Skip arms whose box is below the build high-water (would OOM during the build). ---
for mult in "${MEMLIMIT_MULTS[@]}"; do
  box=$(awk -v r="$R" -v m="$mult" 'BEGIN{printf "%d", r*m}')
  if [[ "$box" -le "${BUILD_HWM:-0}" ]]; then
    log "skipping ${mult}xR box (${box} B) — below build high-water ${BUILD_HWM} B (would OOM the build)"
    continue
  fi
  gml=$(awk -v b="$box" 'BEGIN{printf "%d", b*0.9}')
  run_arm "edge2_m${mult}" "off" "$gml" "$box"
done

# --- Validation: in-envelope (GOGC=x, generous box ~2xR, GOMEMLIMIT 0.9x box; should not bind) ---
_valid_box=$(awk -v r="$R" 'BEGIN{printf "%d", r*2.0}')
_valid_gml=$(awk -v b="$_valid_box" 'BEGIN{printf "%d", b*0.9}')
run_arm "valid_inenvelope_gogc${VALID_GOGC}" "$VALID_GOGC" "$_valid_gml" "$_valid_box"

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
    printf 'Anchor floor R (Sys-HeapReleased) = %s bytes; build high-water = %s bytes.\n' "$R" "$BUILD_HWM"
    printf 'Edge 2 / validation arms run under a cgroup MemoryMax (hard) with GOMEMLIMIT ~0.9x it (soft).\n\n'

    printf '## Edge 1 — sizing (no limit)\n\n'
    printf '| Arm | RSS peak | GC CPU%% | Throughput | p99 (ms) | outcome |\n|---|--:|--:|--:|--:|---|\n'
    for g in "${GOGC_ARMS[@]}"; do _row "edge1_gogc${g}" "GOGC=${g}"; done

    printf '\n## Edge 2 — backstop cost (GOGC=off; cgroup box = mult x R, GOMEMLIMIT ~0.9x box)\n\n'
    printf '| Arm | box (cgroup) | RSS peak | GC CPU%% | Throughput | p99 (ms) | outcome |\n|---|--:|--:|--:|--:|--:|---|\n'
    for mult in "${MEMLIMIT_MULTS[@]}"; do
      local box; box=$(awk -v r="$R" -v m="$mult" 'BEGIN{printf "%d", r*m}')
      [[ "$box" -le "${BUILD_HWM:-0}" ]] && continue
      _row "edge2_m${mult}" "${mult}xR" "$box"
    done

    printf '\n## Validation\n\n'
    printf '| Arm | box (cgroup) | RSS peak | GC CPU%% | Throughput | p99 (ms) | outcome |\n|---|--:|--:|--:|--:|--:|---|\n'
    _row "valid_inenvelope_gogc${VALID_GOGC}" "GOGC=${VALID_GOGC}, box~2R" "$(awk -v r="$R" 'BEGIN{printf "%d", r*2.0}')"
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

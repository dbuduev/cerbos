#!/usr/bin/env bash

# Copyright 2021-2026 Zenauth Ltd.
# SPDX-License-Identifier: Apache-2.0

# Shared configuration for GCP load testing scripts.
# All values can be overridden via environment variables.

# When TERRAFORM_DIR is set, read infrastructure values from Terraform outputs.
# Otherwise, fall back to environment variables / gcloud defaults.
if [[ -n "${TERRAFORM_DIR:-}" ]]; then
  _tf_output() { terraform -chdir="$TERRAFORM_DIR" output -raw "$1"; }
  GCP_PROJECT=$(_tf_output project)
  GCP_ZONE=$(_tf_output zone)
  PDP_VM=$(_tf_output pdp_vm_name)
  CLIENT_VM=$(_tf_output client_vm_name)
  # Tolerant: the staging_bucket output may be absent in older state (pre-apply).
  STAGING_BUCKET=$(_tf_output staging_bucket 2>/dev/null || true)
  unset -f _tf_output
fi

# GCP settings
GCP_PROJECT=${GCP_PROJECT:-$(gcloud config get-value project 2>/dev/null)}
GCP_ZONE=${GCP_ZONE:?"Error: GCP_ZONE is not set"}
GCP_REGION=${GCP_REGION:-"${GCP_ZONE%-*}"}

# Resource naming
NAME_PREFIX=${NAME_PREFIX:-"cerbos-loadtest"}
NETWORK_NAME="${NAME_PREFIX}-net"
SUBNET_NAME="${NAME_PREFIX}-subnet"
PDP_VM=${PDP_VM:-"${NAME_PREFIX}-pdp"}
CLIENT_VM=${CLIENT_VM:-"${NAME_PREFIX}-client"}

# VM configuration
PDP_MACHINE_TYPE=${PDP_MACHINE_TYPE:-"c3-standard-4"}
CLIENT_MACHINE_TYPE=${CLIENT_MACHINE_TYPE:-"e2-standard-4"}
BOOT_DISK_SIZE=${BOOT_DISK_SIZE:-"50GB"}

# Cerbos configuration
CERBOS_VERSION=${CERBOS_VERSION:-"latest"}
STORE=${STORE:-"disk"}
AUDIT_ENABLED=${AUDIT_ENABLED:-"false"}
SCHEMA_ENFORCEMENT=${SCHEMA_ENFORCEMENT:-"none"}

# Optional GCS bucket for staging bulk uploads, e.g. STAGING_BUCKET=gs://my-bucket. When
# set, deploy uploads go local -> GCS -> VM (gcloud storage cp: reliable + parallel,
# bypassing the throughput-limited / stall-prone IAP tunnel) instead of scp over IAP.
# The VMs' service account needs roles/storage.objectViewer on the bucket.
STAGING_BUCKET=${STAGING_BUCKET:-}

# Paths
REMOTE_BASE=${REMOTE_BASE:-"/opt/cerbos-loadtest"}
WORK_DIR=${WORK_DIR:-"$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)/work"}

# Helper functions
GSSH() {
  local vm="$1"
  shift
  gcloud compute ssh "$vm" --zone="$GCP_ZONE" --project="$GCP_PROJECT" --tunnel-through-iap -- "$@"
}

GSCP() {
  gcloud compute scp --zone="$GCP_ZONE" --project="$GCP_PROJECT" --tunnel-through-iap "$@"
}

# Upload a local file to a path on a VM by staging through GCS (gcloud storage cp —
# reliable + parallel; the VM pulls bucket->dest over Google's internal network). This
# avoids the IAP tunnel, which is throughput-limited and stalls on large files, so we
# require STAGING_BUCKET rather than silently falling back to that broken path.
# Args: $1=local_file  $2=vm  $3=remote_dest_path (a file path, not a directory).
upload_to_vm() {
  local src="$1" vm="$2" dest="$3"
  : "${STAGING_BUCKET:?required for uploads — set it, or apply Terraform (staging_bucket output) with TERRAFORM_DIR}"
  local obj="${STAGING_BUCKET%/}/deploy/$(basename "$src")"
  log "Staging $(basename "$src") -> ${obj} -> ${vm}:${dest}"
  gcloud storage cp "$src" "$obj"
  GSSH "$vm" "gcloud storage cp '$obj' '$dest'"
  gcloud storage rm "$obj" 2>/dev/null || true
}

log() {
  printf "[%s] %s\n" "$(date '+%H:%M:%S')" "$*"
}

err() {
  printf "[%s] ERROR: %s\n" "$(date '+%H:%M:%S')" "$*" >&2
}

require_running_vms() {
  local vms=("$@")
  for vm in "${vms[@]}"; do
    local status
    status=$(gcloud compute instances describe "$vm" \
      --zone="$GCP_ZONE" --project="$GCP_PROJECT" \
      --format='get(status)' 2>/dev/null) || { err "VM $vm not found"; exit 1; }
    if [[ "$status" != "RUNNING" ]]; then
      log "VM $vm is $status — starting it..."
      gcloud compute instances start "$vm" --zone="$GCP_ZONE" --project="$GCP_PROJECT"
    fi
  done
}

# Restart Cerbos on the PDP VM. Honours env: GOMAXPROCS, GOGC, GOMEMLIMIT, and
# CGROUP_LIMIT. When CGROUP_LIMIT (bytes) is set, Cerbos runs under a transient systemd
# scope with MemoryMax = CGROUP_LIMIT and swap disabled — a *hard* cap (cgroup OOM-kills
# on breach), the production-faithful pairing for GOMEMLIMIT (soft) a few % below it.
# When unset, the plain nohup path is used.
restart_cerbos() {
  log "Restarting Cerbos on PDP VM (GOGC=${GOGC:-default} GOMEMLIMIT=${GOMEMLIMIT:-off} cgroup=${CGROUP_LIMIT:-none})..."
  GSSH "$PDP_VM" <<ENDSSH
set -euo pipefail
sudo systemctl stop cerbos-loadtest 2>/dev/null || true
sudo systemctl reset-failed cerbos-loadtest 2>/dev/null || true
pkill -f "${REMOTE_BASE}/bin/cerbos" 2>/dev/null || true
sleep 1
echo "Starting Cerbos..."
if [ -n "${CGROUP_LIMIT:-}" ]; then
  sudo systemd-run --collect --unit=cerbos-loadtest \
    -p MemoryMax=${CGROUP_LIMIT:-} -p MemorySwapMax=0 \
    --setenv=STORE=${STORE} --setenv=AUDIT_ENABLED=${AUDIT_ENABLED} --setenv=SCHEMA_ENFORCEMENT=${SCHEMA_ENFORCEMENT} \
    ${GOMAXPROCS:+--setenv=GOMAXPROCS=${GOMAXPROCS}} \
    ${GOGC:+--setenv=GOGC=${GOGC}} \
    ${GOMEMLIMIT:+--setenv=GOMEMLIMIT=${GOMEMLIMIT}} \
    ${REMOTE_BASE}/bin/cerbos server \
      --debug-listen-addr=:6666 --config=${REMOTE_BASE}/conf/cerbos.yaml --log-level=warn
  echo "Started under systemd cgroup (MemoryMax=${CGROUP_LIMIT:-})"
else
  STORE=${STORE} AUDIT_ENABLED=${AUDIT_ENABLED} SCHEMA_ENFORCEMENT=${SCHEMA_ENFORCEMENT} \
    ${GOMAXPROCS:+GOMAXPROCS=${GOMAXPROCS}} \
    ${GOGC:+GOGC=${GOGC}} \
    ${GOMEMLIMIT:+GOMEMLIMIT=${GOMEMLIMIT}} \
    nohup ${REMOTE_BASE}/bin/cerbos server \
   --debug-listen-addr=:6666 \
   --config=${REMOTE_BASE}/conf/cerbos.yaml \
    --log-level=warn \
    > ${REMOTE_BASE}/cerbos.log 2>&1 &
  echo "Cerbos PID: \$!"
fi

echo "Waiting for Cerbos to become healthy..."
healthy=false
for i in \$(seq 1 30); do
  if curl -sf http://localhost:3592/_cerbos/health >/dev/null 2>&1; then
    echo "Cerbos is healthy"
    healthy=true
    break
  fi
  sleep 2
done
if [ "\$healthy" != "true" ]; then
  echo "ERROR: Cerbos health check failed after 30 attempts" >&2
  journalctl -u cerbos-loadtest -n 20 --no-pager 2>/dev/null || tail -20 ${REMOTE_BASE}/cerbos.log 2>/dev/null >&2
  exit 1
fi
ENDSSH
}

check_policies() {
  if [[ ! -d "${WORK_DIR}/policies" ]]; then
    err "Missing ${WORK_DIR}/policies — generate test data first:"
    err "  cd hack/loadtest"
    err "  NUM_POLICIES=1000 ./loadtest.sh -g"
    exit 1
  fi
}

# Peak-RSS helpers
# The metrics endpoint only exposes *current* RSS, so peak RSS is read host-side from
# /proc/<pid>/status on the PDP VM. clear_refs (write "5") resets the high-water so a
# subsequent read measures a fresh window (e.g. the load phase, excluding the build).
_pdp_cerbos_pid_expr="\$(pgrep -f '${REMOTE_BASE}/bin/cerbos server' | head -1)"

# Echo the running Cerbos VmHWM in bytes (peak RSS since last reset / process start).
pdp_vmhwm_bytes() {
  GSSH "$PDP_VM" "awk '/^VmHWM:/{print \$2*1024}' /proc/${_pdp_cerbos_pid_expr}/status"
}

# Reset the VmHWM high-water of the running Cerbos process. clear_refs is owner-writable
# only, and under the cgroup path Cerbos runs as root (sudo systemd-run), so write via
# `sudo tee` — correct whether Cerbos is root- or user-owned. (Reading VmHWM from
# /proc/<pid>/status needs no privilege; status is world-readable.)
pdp_reset_vmhwm() {
  GSSH "$PDP_VM" "echo 5 | sudo tee /proc/${_pdp_cerbos_pid_expr}/clear_refs >/dev/null" 2>/dev/null || \
    err "could not reset VmHWM (clear_refs) — continuing with lifetime peak (RSS peak will include the build)"
}

# Echo "metric value" lines for the named PDP metrics. The endpoint is on the PDP's
# private VPC IP, unreachable from the orchestrator, so the curl runs on the PDP itself
# over SSH (against localhost). Args: metric names. (Counterpart to loadtest.sh's
# client-side scrapeMetrics, which can curl the PDP directly — kept separate by design.)
pdp_scrape() {
  local raw
  raw=$(GSSH "$PDP_VM" "curl -sf http://localhost:3592/_cerbos/metrics") || return 1
  local m val
  for m in "$@"; do
    val=$(echo "$raw" | grep "^${m} " | head -1 | awk '{print $2}')
    [[ -n "$val" ]] && printf '%s %s\n' "$m" "$val"
  done
}

# Echo "running" if the Cerbos process is alive on the PDP, else "dead" (cgroup OOM).
pdp_cerbos_alive() {
  if GSSH "$PDP_VM" "pgrep -f '${REMOTE_BASE}/bin/cerbos server' >/dev/null"; then
    echo running
  else
    echo dead
  fi
}

# PDP footprint gauges scraped post-load / at the settled floor (Sys-HeapReleased
# accounting + resident footprint). Used by run_load_and_capture and the sweep's Step 0.
PDP_FLOOR_METRICS=(
  process_resident_memory_bytes
  go_memstats_heap_alloc_bytes
  go_memstats_heap_inuse_bytes
  go_memstats_sys_bytes
  go_memstats_heap_released_bytes
)

# Print an mpstat CPU-utilization summary (avg/max %used from the "all" rows).
# Args: $1=label  $2=logfile
cpu_summary() {
  local label="$1" logfile="$2"
  if [[ ! -f "$logfile" ]]; then
    printf "  %-10s (no data)\n" "$label"
    return
  fi
  awk '/^ *[0-9].*all/ { idle = $NF; sum += idle; n++; if (n == 1 || idle < min_idle) min_idle = idle }
       END { if (n > 0) printf "  %-10s avg %5.1f%%   max %5.1f%%   (%d samples)\n", label, 100 - sum/n, 100 - min_idle, n }' \
    label="$label" "$logfile"
}

# Run the load (warmup + sustained + throughput via loadtest.sh -e on the client) against
# the already-running PDP, capturing per-run signals into RESULT_DIR: peak RSS (VmHWM,
# build excluded), ghz JSON + GC counters (from loadtest.sh), post-load accounting,
# liveness/OOM status, both VMs' mpstat logs, and a CPU summary. Cerbos must already be
# running (caller restarts it with any knobs first); reads load knobs (RPS/DURATION_SECS/
# ITERATIONS/CONCURRENCY/...) and the global PDP_IP from the environment.
# Args: $1=result_dir (local).
run_load_and_capture() {
  local result_dir="$1"
  mkdir -p "$result_dir"

  # Fresh remote results, and reset the peak so VmHWM measures the load window (build
  # excluded).
  GSSH "$PDP_VM" "rm -rf ${REMOTE_BASE}/results/* 2>/dev/null || true"
  GSSH "$CLIENT_VM" "rm -rf ${REMOTE_BASE}/results/* 2>/dev/null || true"
  pdp_reset_vmhwm

  # CPU monitors on both VMs.
  GSSH "$PDP_VM" "pkill -f 'mpstat -P ALL' 2>/dev/null || true; setsid mpstat -P ALL 1 > ${REMOTE_BASE}/results/cpu_usage.log 2>&1 < /dev/null &" || true
  GSSH "$CLIENT_VM" "pkill -f 'mpstat -P ALL' 2>/dev/null || true; setsid mpstat -P ALL 1 > ${REMOTE_BASE}/results/client_cpu_usage.log 2>&1 < /dev/null &" || true

  log "Running load on Client VM (${CLIENT_VM})..."
  # A mid-load cgroup OOM makes loadtest.sh error out — tolerate it and detect below.
  GSSH "$CLIENT_VM" <<ENDSSH || log "load run returned non-zero — possible OOM mid-load"
set -uo pipefail
. /nix/var/nix/profiles/default/etc/profile.d/nix-daemon.sh
cd ${REMOTE_BASE}
mkdir -p ${REMOTE_BASE}/results
SERVER="${PDP_IP}:3593" \
METRICS_URL="http://${PDP_IP}:3592/_cerbos/metrics" \
WORK_DIR="${REMOTE_BASE}" \
STORE="${STORE}" \
${RPS:+RPS="${RPS}"} \
${DURATION_SECS:+DURATION_SECS="${DURATION_SECS}"} \
${ITERATIONS:+ITERATIONS="${ITERATIONS}"} \
${CONCURRENCY:+CONCURRENCY="${CONCURRENCY}"} \
${CONNECTIONS:+CONNECTIONS="${CONNECTIONS}"} \
${REQ_KIND:+REQ_KIND="${REQ_KIND}"} \
${NUM_POLICIES:+NUM_POLICIES="${NUM_POLICIES}"} \
${PROTOSET:+PROTOSET="${REMOTE_BASE}/cerbos.protoset"} \
  nix develop --command bash loadtest.sh -e
ENDSSH

  GSSH "$PDP_VM" "pkill -f 'mpstat -P ALL' 2>/dev/null || true" || true
  GSSH "$CLIENT_VM" "pkill -f 'mpstat -P ALL' 2>/dev/null || true" || true

  # Per-run captures: OOM status, peak RSS, post-load accounting.
  if [[ "$(pdp_cerbos_alive)" == dead ]]; then
    echo oom > "${result_dir}/status"
    log "Cerbos died during load — likely cgroup OOM"
  else
    echo ok > "${result_dir}/status"
  fi
  pdp_vmhwm_bytes > "${result_dir}/vmhwm_bytes.txt" 2>/dev/null || echo "" > "${result_dir}/vmhwm_bytes.txt"
  pdp_scrape "${PDP_FLOOR_METRICS[@]}" > "${result_dir}/post_metrics.txt" 2>/dev/null || true

  # Download PDP cpu log + client results (ghz JSON, *_gc.json, summaries) into result_dir.
  GSSH "$PDP_VM" "tar czf /tmp/pdp-results.tar.gz -C ${REMOTE_BASE}/results cpu_usage.log" 2>/dev/null || true
  GSCP "${PDP_VM}:/tmp/pdp-results.tar.gz" "/tmp/pdp-results.tar.gz" 2>/dev/null || true
  [[ -f /tmp/pdp-results.tar.gz ]] && { tar xzf /tmp/pdp-results.tar.gz -C "$result_dir"; mv -f "${result_dir}/cpu_usage.log" "${result_dir}/pdp_cpu_usage.log" 2>/dev/null; rm -f /tmp/pdp-results.tar.gz; }

  GSSH "$CLIENT_VM" "tar czf /tmp/client-results.tar.gz -C ${REMOTE_BASE}/results ." 2>/dev/null || true
  GSCP "${CLIENT_VM}:/tmp/client-results.tar.gz" "/tmp/client-results.tar.gz" 2>/dev/null || true
  [[ -f /tmp/client-results.tar.gz ]] && { tar xzf /tmp/client-results.tar.gz -C "$result_dir"; rm -f /tmp/client-results.tar.gz; }

  printf "\nCPU utilization (%% of all cores):\n"
  cpu_summary "PDP" "${result_dir}/pdp_cpu_usage.log"
  cpu_summary "Client" "${result_dir}/client_cpu_usage.log"
}

check_print_summary() {
  if [[ "$POLICIES_ONLY" == false ]] && [[ ! -f "${WORK_DIR}/printsummary" ]]; then
    log "Building printsummary..."
    pushd "${SCRIPT_DIR}/.." > /dev/null
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -tags printsummary -o work/printsummary .
    popd > /dev/null
  fi
}

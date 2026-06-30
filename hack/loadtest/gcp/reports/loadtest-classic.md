# Load Test: Classic Policies Scaling (800 to 40K)

**Date:** 2026-06-30  
**Environment:** GCP c3-standard-4 PDP (4 vCPU, 16 GB RAM), e2-standard-4 client, disk storage  
**Cerbos version:** latest  
**Config:** default runtime (`GOGC=100`, no `GOMEMLIMIT`); concurrency=100, connections=5; mixed requests (cr_req01 + cr_req02); throughput probe (max), then sustained-rate at ~85% of the measured ceiling for 120s.

## Summary

- Throughput degrades gradually from 800 to 40K policies, with no cliff: max RPS falls from 8.6K to 7.7K to 5.7K.
- Under sustained load at ~85% of capacity, p99 grows from 11.6 ms at 800 policies to 40.8 ms at 40K.
- RSS peak is 133 MiB at 800 policies, 224 MiB at 8K, and 748 MiB at 40K, growing sub-linearly with policy count.
- GC CPU stays around 8-9% across the range, with no sustained stalls or throughput gaps.
- To provision for max RPS, run `GOGC=100` and set a cgroup MemoryMax of 256 MiB at 800 policies, 512 MiB at 8K, or 1.5 GiB at 40K, with a `GOMEMLIMIT` backstop below the cap.

## Performance (default config, `GOGC=100`)

| Policies | RSS peak | GC CPU | Max RPS | Sust RPS (~85%) | p99@sust | stalls/gaps |
|---------:|---------:|-------:|--------:|----------------:|---------:|------------:|
| 800 | 133 MiB | 8.5% | 8,638 | 7,300 | 11.6 ms | 1/0 |
| 8K | 224 MiB | 9.0% | 7,700 | 6,500 | 22.1 ms | 1/0 |
| 40K | 748 MiB | 7.7% | 5,736 | 4,900 | 40.8 ms | 0/0 |

RSS peak is `VmHWM` during load. GC CPU is the share of CPU time spent in garbage collection during the sustained run. Max RPS comes from the throughput probe; Sust RPS is the achieved sustained rate, about 85% of Max. Stalls and gaps are 1-second window counts: a stall is a window where more than 10% of requests exceed p95, a gap is a window serving less than 75% of mean throughput. All runs were error-clean apart from a few gRPC `Unavailable` results at connection teardown in the final window, with no application errors.

Throughput declines smoothly with no cliff. The single stall window at 800 and 8K is an isolated warmup blip, and 40K is clean.

## Tail Latencies: Sustained-Rate Test (ms)

| Percentile | 800 (7.3K) | 8K (6.5K) | 40K (4.9K) |
|------------|-----------:|----------:|-----------:|
| p50 | 1.80 | 2.59 | 3.15 |
| p90 | 6.33 | 12.52 | 24.75 |
| p95 | 8.00 | 15.77 | 30.02 |
| p99 | 11.62 | 22.09 | 40.83 |

p50 stays low, between 1.8 and 3.2 ms: with about 15% headroom below capacity, most requests complete quickly. Larger policy sets push the tail up, since each request touches more rules and runs more CEL evaluation, but the rise is gradual and the distribution stays tight with no GC-driven clustering.

## Memory

RSS peak is 133 MiB at 800 policies, 224 MiB at 8K, and 748 MiB at 40K. It grows sub-linearly with policy count: at low counts most of the footprint is fixed runtime and binary overhead, and the per-policy cost of the rule table is small. The index uses lazy sparse bitmaps and a map for fqnBindings, and the build streams one policy at a time, so its peak working set is a single policy.

## Provisioning for Max RPS

For maximum throughput, run the default `GOGC=100` and add a memory backstop: a soft `GOMEMLIMIT` paired with a hard cgroup MemoryMax. The cgroup MemoryMax is the figure to provision.

| Policies | Loaded RSS peak | GOMEMLIMIT | cgroup MemoryMax |
|---------:|----------------:|-----------:|-----------------:|
| 800 | 133 MiB | 160 MiB | 256 MiB |
| 8K | 224 MiB | 320 MiB | 512 MiB |
| 40K | 748 MiB | 1.2 GiB | 1.5 GiB |

`GOMEMLIMIT` is set to about twice the loaded runtime heap (the loaded RSS peak minus ~60 MiB of binary and stack overhead), so it sits well above the working set and does not bind in steady state. Throughput therefore stays at the numbers above; the backstop engages only if the live set later grows past the box, spending some CPU on extra GC to stay under the limit instead of being OOM-killed. The cgroup MemoryMax adds the ~60 MiB back plus a ~20% safety margin, so the soft limit takes effect before the kernel OOM-kills, and it clears the build-time peak as well.

For a policy count between these points, interpolate the loaded RSS peak linearly and apply the same formula.

## Connection Count (HOL Blocking)

Fewer connections is slightly faster: one connection gives the lowest latency and the highest throughput. On a same-zone VPC with near-zero packet loss, HTTP/2 multiplexing works well and head-of-line blocking is not a factor. More connections only add server-side overhead, such as per-connection goroutines, buffers, and flow control, without any benefit. The default of 5 connections is fine, and real deployments do not need connection pooling beyond gRPC defaults.

## Key Takeaways

1. Throughput has no cliff: max RPS falls smoothly from 8.6K at 800 policies to 5.7K at 40K.
2. Under sustained load at about 85% of capacity, p99 rises gradually from 11.6 ms at 800 policies to 40.8 ms at 40K.
3. RSS peak ranges from 133 MiB at 800 policies to 748 MiB at 40K and grows sub-linearly with policy count.
4. GC CPU stays around 8-9% across the range, with no sustained stalls or throughput gaps.
5. For max RPS, run `GOGC=100` and provision a cgroup MemoryMax of 256 MiB at 800 policies, 512 MiB at 8K, or 1.5 GiB at 40K, with a `GOMEMLIMIT` backstop sized just below the cap.

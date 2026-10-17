# Engineering roadmap

Implemented: module-aware build/test commands, broker input regression tests, and a repeatable request-decoding microbenchmark.

Planned:
- Restore a complete local deployment and document service discovery behavior.
- Test unavailable dependencies, messaging retries, shutdown and authentication boundaries.
- Add representative service load tests with throughput, p50/p95/p99 latency and error rates.
- Publish hardware, runtime, concurrency, warm-up, workload and repeated-run variability with results.
- Automate deployment smoke checks and cleanup before benchmarking distributed services.

Microbenchmark results do not establish HFT suitability. Performance regression thresholds require a controlled baseline and noise characterization.

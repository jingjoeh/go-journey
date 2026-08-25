# Lab Notes

## Completion Evidence

- **Status:** Completed
- **Date:** 2026-08-25
- **Implementation:** Ordered deduplication completed by the learner.
- **Verification:** `go test -count=1 ./...` passed.
- **Benchmark evidence:** Before/after benchmark and measured allocation improvement reviewed.
- **Mentor review:** PASS.

## Reviewed Scope

- Functional deduplication behavior while preserving first-seen order.
- Input ownership and non-mutation.
- Benchmarking before and after the optimization.
- Allocation improvement measured rather than assumed.
- Preallocation justified by evidence for this lab workload.

## Benchmark Verification Snapshot

- **Command:** `go test -run '^$' -bench . -benchmem ./tests`
- **Environment:** Darwin/arm64, Apple M4 Pro.
- **Observed 2026-08-25:** `71.71 ns/op`, `96 B/op`, `1 alloc/op`.
- Benchmark values are environment-dependent and may vary between runs.

## Performance Decision and Limitation

- Preallocating the result slice is retained because the lab's before/after measurements showed an allocation improvement.
- The benchmark input is small and represents only one workload shape.
- Do not generalize the observed improvement as the same proportional gain for larger, differently distributed, or production workloads.
- Re-benchmark with representative inputs before applying the same conclusion elsewhere.

## Prediction

## Tests added

## Failed attempts and lessons

## Failure and concurrency behavior

## Production signals

## Trade-offs and deliberate omissions

## After comparing with the reference

Do not complete this section until your own solution passes.

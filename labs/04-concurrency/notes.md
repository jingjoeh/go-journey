# Lab Notes

## Completion Evidence

- **Status:** Completed
- **Date:** 2026-08-22
- **Mentor review:** PASS.
- **Verification reported:** `go test ./...`, `go test -race ./...`, and `go vet ./...` passed.
- **Formatting checks reported:** `gofmt` and `git diff --check` passed.
- **Assessment:** The implementation satisfies the current lab contract. The items below are non-blocking improvements and production considerations.

## Reviewed Scope

- Bounded worker-pool behavior.
- Preserving result order by writing results using the input index.
- Cancellation propagation and goroutine lifecycle.
- Channel ownership and channel closing.
- `WaitGroup` lifecycle.
- Race safety when workers write to distinct result indexes.
- Cancellation semantics when multiple `select` cases are ready.
- Public API surface, including the unnecessary exported `SquareJob` type.
- Resource efficiency when `workers` is much greater than `len(values)`.

## Known Trade-offs and Follow-up Improvements

- `SquareJob` is exported even though the intended public contract centers on `SquareAll`. Reducing its visibility would keep the package API narrower.
- The current tests do not explicitly cover a context canceled before the call, empty input with a valid worker count, `workers == 1`, or more workers than input values.
- The implementation starts the configured number of workers even when the input is empty or much smaller. Extreme worker counts can therefore consume unnecessary resources.
- After cancellation, a `select` may choose another ready case before observing `ctx.Done()`. A small amount of additional work may start, so cancellation means prompt cooperative stopping rather than a strict instantaneous boundary.
- These points are documented as non-blocking production considerations; no implementation change was required for Lab 04 completion.

## Prediction

## Tests added

## Failed attempts and lessons

## Failure and concurrency behavior

## Production signals

## Trade-offs and deliberate omissions

## After comparing with the reference

Do not complete this section until your own solution passes.

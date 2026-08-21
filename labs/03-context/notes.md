# Lab Notes

## Completion Evidence

- **Status:** Completed
- **Date:** 2026-08-21
- **Implementation:** `Wait` was manually attempted and iteratively corrected by the learner.
- **Verification:** `go test ./tests` and `go test ./starter` passed.
- **Mentor review:** PASS.
- **Race/testing review:** Completed as applicable.

## Reviewed Scope

- Correct cancellation propagation through `ctx.Done()`.
- Correctly returns `ctx.Err()` when cancellation or deadline wins.
- Distinguishes `context.Canceled` from `context.DeadlineExceeded`.
- Understands that `ctx.Done()` represents cancellation or deadline, not successful completion.
- Uses `time.NewTimer` instead of `time.Sleep` so context cancellation can interrupt the wait.
- Uses `defer timer.Stop()` to release timer resources when cancellation wins.
- Starts no goroutine.
- Understands that `time.Sleep` blocks the goroutine, not the OS thread, and is not context-cancellable.
- Reviewed `select` behavior between `timer.C` and `ctx.Done()`.

## Prediction

## Tests added

## Failed attempts and lessons

## Failure and concurrency behavior

## Production signals

## Trade-offs and deliberate omissions

## After comparing with the reference

Do not complete this section until your own solution passes.

# Lab Notes

## Completion Evidence

- **Status:** Completed
- **Date:** 2026-08-23
- **Implementation:** Atomic transfer orchestration completed by the learner.
- **Verification:** `go test -count=1 ./...`, `go test -count=1 -race ./...`, and `go vet ./...` passed.
- **Race result:** No data race reported.
- **Mentor review:** PASS.

## Reviewed Scope

- Validates that `amount > 0` before entering the transaction.
- Uses a transaction callback to keep related balance changes in one transaction boundary.
- Performs debit before credit within the same callback.
- Fails fast when either debit or credit returns an error.
- Uses `%w` so callers can inspect the underlying error cause.
- Preserves the primary invariant: a successful debit followed by a failed credit must not be committed.

## Transaction Decision

- The callback returns an error immediately when debit or credit fails.
- That error is the signal the transaction layer needs to roll back instead of committing partial state.
- Error context may be added while `%w` preserves the original cause.

## Prediction

## Tests added

## Failed attempts and lessons

## Failure and concurrency behavior

## Production signals

## Trade-offs and deliberate omissions

## After comparing with the reference

Do not complete this section until your own solution passes.

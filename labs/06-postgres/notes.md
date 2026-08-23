# Lab Notes

## Completion Evidence

- **Status:** Completed
- **Date:** 2026-08-23
- **Implementation:** Safe pagination query construction completed by the learner.
- **Verification:** `go test -count=1 ./...`, `go test -count=1 -race ./...`, and `go vet ./...` passed.
- **Race result:** No data race reported.
- **Mentor review:** PASS.

## Reviewed Scope

- Parameterized SQL using PostgreSQL placeholders `$1` and `$2`.
- Separation of static query structure from runtime values using `[]any` arguments.
- Default limit of `20` when the supplied limit is below the valid range.
- Maximum limit clamped to `100`.
- Negative offsets rejected with `ErrInvalidOffset`.
- Deliberately does not connect to a real database because the lab contract is limited to constructing a query and its arguments.

## SQL Decision

- Runtime values are not interpolated into SQL with `fmt.Sprintf`.
- The function returns SQL placeholders and runtime arguments separately.
- This preserves query structure and supports safe parameter binding at the database boundary.

## Prediction

## Tests added

## Failed attempts and lessons

## Failure and concurrency behavior

## Production signals

## Trade-offs and deliberate omissions

## After comparing with the reference

Do not complete this section until your own solution passes.

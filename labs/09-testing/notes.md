# Lab Notes

## Completion Evidence

- **Status:** Completed
- **Date:** 2026-08-24
- **Implementation and tests:** Completed by the learner.
- **Verification:** `go test -count=1 ./...`, `go test -count=1 -race ./...`, and `go vet ./...` passed.
- **Race result:** No data race reported.
- **Mentor review:** PASS.

## Reviewed Scope

- Table-driven testing.
- Deriving boundary cases from a behavioral contract.
- Trimming surrounding whitespace with `strings.TrimSpace`.
- Preserving the local-part letter case.
- Lowercasing only the domain.
- Validating the number and placement of `@` separators.
- Checking `len(parts)` before indexing to prevent a panic.

## Debugging Evidence

- The first attempt contained a bug.
- The learner identified and corrected it through reasoning and test feedback.
- The reference solution was not opened before the correction.
- This demonstrates a useful live-coding skill: diagnosing and recovering from an incorrect first attempt.

## Prediction

## Tests added

## Failed attempts and lessons

## Failure and concurrency behavior

## Production signals

## Trade-offs and deliberate omissions

## After comparing with the reference

Do not complete this section until your own solution passes.

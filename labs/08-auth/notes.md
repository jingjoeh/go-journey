# Lab Notes

## Completion Evidence

- **Status:** Completed
- **Date:** 2026-08-24
- **Implementation:** `Sign` and `Verify` completed by the learner.
- **Verification:** `go test -count=1 ./...` passed.
- **Mentor review:** PASS.

## Reviewed Scope

- Implements both token signing and verification.
- Rejects an empty subject.
- Reports malformed or modified tokens as `ErrInvalidToken`.
- Uses HMAC-SHA256 for message authentication.
- Uses URL-safe Base64 for the token representation.
- Uses `hmac.Equal` for constant-time signature comparison.

## Security Properties

- HMAC provides integrity and authenticity when the secret remains protected.
- HMAC does not provide encryption or confidentiality.
- Base64 is an encoding, not encryption.
- Token contents encoded with Base64 remain readable and must not be treated as secret.

## Prediction

## Tests added

## Failed attempts and lessons

## Failure and concurrency behavior

## Production signals

## Trade-offs and deliberate omissions

## After comparing with the reference

Do not complete this section until your own solution passes.

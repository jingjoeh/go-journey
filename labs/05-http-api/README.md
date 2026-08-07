# HTTP API: Task Handler

> Work in `starter/`. Do not open `solution/` until you have a tested attempt.

## Objective

Build a small net/http endpoint with a stable JSON contract and method handling.

## Requirements

- GET /health returns 200 and JSON status.
- Other methods return 405 with Allow.
- Unknown paths return 404.

## Constraints

- Use the Go standard library.
- Keep the public contract focused on `func Handler() http.Handler`.
- Make cancellation, errors, and ownership explicit where applicable.
- Do not add infrastructure until a test demonstrates the need.

## Tasks

1. Read the tests and add any missing boundary cases.
2. Implement the contract in `starter/`.
3. Run `go test ./tests`.
4. Run `go test -race ./tests` for concurrent code.
5. Record decisions and failed attempts in `notes.md`.

## Bonus Tasks

- Add a property, fuzz, benchmark, or failure-injection test.
- Explain how the design changes in a long-running service.

## Testing

From this directory:

```sh
go test ./tests
go test ./starter
```

After completing your attempt, compare and verify the reference with `go test ./solution`.

## Senior Review Questions

- Which invariant is most important?
- What happens under concurrent use, cancellation, or repeated execution?
- Which error should cross the boundary and which should be translated?
- How would you observe a failure in production?
- What did you deliberately keep simple?


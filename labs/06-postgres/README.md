# PostgreSQL: Safe Pagination Query

> Work in `starter/`. Do not open `solution/` until you have a tested attempt.

## Objective

Build a parameterized PostgreSQL query and enforce bounded pagination.

## Requirements

- Use placeholders for values.
- Clamp limit to 1-100 with a default of 20.
- Reject negative offsets.

## Constraints

- Use the Go standard library.
- Keep the public contract focused on `func ListTasksQuery(limit, offset int) (string, []any, error)`.
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


# Executable Labs

Each lab is an isolated Go module so its starter can be changed without affecting other work.

## Workflow

1. Read the lab README and predict difficult cases.
2. Run `go test ./tests` from the lab directory to see the initial failure.
3. Implement only inside `starter/`; rerun tests and the race detector where relevant.
4. Answer the review questions in `notes.md`.
5. After finishing, compare against `solution/`. The reference is one design, not the only design.
6. Run `go test ./solution` to verify the reference independently.

The labs intentionally start with a compiling `panic("TODO")`. A red learner test is the starting signal, not a broken repository.


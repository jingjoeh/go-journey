# Learning Progress

Update weekly. Confidence must be backed by a recent implementation, debugging session, review, or explanation.

**Last updated:** 2026-08-23

| Module | Status | Confidence | Exercises | Notes |
|---|---|---:|---:|---|
| Go Fundamentals | Review | 1/5 | 3/10 | Labs 01-03 completed with passing tests and mentor review; broader evidence pending |
| Concurrency | Practicing | 1/5 | 0/10 | Fundamentals covered; production-grade design is current focus |
| HTTP Backend | Learning | 1/5 | 1/10 | Lab 05 completed with passing tests, mentor review, and reference comparison; broader evidence pending |
| PostgreSQL | Learning | 1/5 | 2/10 | Labs 06-07 completed with uncached tests, race tests, vet, and mentor review; broader database evidence pending |
| Architecture | Not Started | 1/5 | 0/10 | |
| Testing | Not Started | 1/5 | 0/10 | |
| Distributed Systems | Not Started | 1/5 | 0/10 | |
| Production Engineering | Not Started | 1/5 | 0/10 | |
| Performance | Not Started | 1/5 | 0/10 | |
| System Design | Not Started | 1/5 | 0/10 | |

Confidence remains unchanged until the evidence demonstrates the capability across the broader module.

## Current Focus

- [ ] Production-grade concurrency design
  - [ ] Worker-pool error and cancellation design
  - [ ] Preventing goroutine leaks and deadlocks
  - [ ] Choosing worker pool vs `errgroup`
  - [ ] Code-review and reasoning exercises

## Needs Reinforcement

- [ ] Method Set vs Method Call distinction
- [ ] Closure state and lifetime
- [ ] Designing concurrent flows without overcomplicating primitives

## Coverage Snapshot — 2026-08-19

The items in this snapshot are learner-reported topic coverage. Checked items mean “studied or practiced,” not “mentor-reviewed mastery.”

### Go Foundations / Type System

- [x] Value semantics / function arguments copy values
- [x] Array vs Slice
- [x] Slice descriptor / backing array / len / cap / reslice / append
- [x] Defined types / aliases / untyped constants
- [x] Struct value semantics
- [x] Pointer basics
- [x] Value receiver vs pointer receiver
- [x] Method Set
- [x] Interface satisfaction / implicit implementation
- [x] Interface static type / concrete type / concrete value
- [x] Type assertion
- [x] Type switch
- [x] nil interface / typed nil

### Error Handling

- [x] `error` as interface
- [x] Sentinel errors
- [x] Custom error types
- [x] `fmt.Errorf`
- [x] `%w` error wrapping
- [x] `errors.Is`
- [x] `errors.As`
- [x] Error translation across repository/service/handler
- [x] Error handling vs logging responsibility

### Functions / Runtime Semantics

- [x] `defer`
- [x] `defer` LIFO
- [x] `defer` argument evaluation
- [x] `defer` + pointer / closure
- [x] named returns + `defer`
- [x] `panic` / `recover`
- [x] error vs panic
- [x] Function values
- [x] Anonymous functions
- [x] Closures / captured variables
- [x] Callback pattern
- [x] Transaction-style callbacks

### Memory

- [x] Stack vs Heap mental model
- [x] Escape Analysis
- [x] Pointer does not imply heap
- [x] Closure escape
- [x] GC reachability
- [x] GC roots mental model
- [x] Reference cycles
- [x] Memory retention via slices/maps

### Concurrency

- [x] Goroutines
- [x] `WaitGroup`
- [x] Unbuffered channels
- [x] Buffered channels
- [x] `close` / comma-ok / `range`
- [x] Channel ownership
- [x] nil channels
- [x] `select` / `default` / timeout
- [x] `context.WithCancel`
- [x] `context.WithTimeout` / deadlines
- [x] Context propagation / parent-child cancellation
- [x] Data races
- [x] `Mutex`
- [x] `RWMutex`
- [x] Atomic basics / CAS
- [x] `sync.Once`
- [x] `sync.Cond` basics
- [x] Race detector
- [x] Happens-before basics
- [x] Goroutine leak / deadlock / data race distinction
- [x] Worker Pool fundamentals
- [x] Bounded concurrency / backpressure
- [x] `errgroup` fundamentals

## Evidence Register

| Area | Evidence | Review status |
|---|---|---|
| Go Foundations / Type System | [Lab 01 — learner implementation](labs/01-go-basics/starter/starter.go); `go test ./tests` passed 2026-08-20; mentor reviewed nil/empty input, non-mutation, and idiomatic `range` | Lab accepted |
| Error Handling | [Lab 02 — learner implementation](labs/02-error-handling/starter/starter.go); `go test ./tests ./starter` and `go vet ./...` passed 2026-08-20; mentor reviewed bounds, sentinel wrapping, and `errors.Is` classification | Lab accepted |
| Context | [Lab 03 — learner implementation](labs/03-context/starter/starter.go); `go test ./tests ./starter` passed 2026-08-21; mentor reviewed cancellation, deadlines, timer cleanup, and `select` behavior | Lab accepted |
| HTTP API | [Lab 05 — learner implementation](labs/05-http-api/starter/starter.go); `go test ./tests ./starter` passed 2026-08-23; mentor reviewed HTTP contracts, JSON responses, routing semantics, and API trade-offs | Lab accepted |
| PostgreSQL Query Construction | [Lab 06 — learner implementation](labs/06-postgres/starter/starter.go); uncached tests, race test, and vet passed 2026-08-23; mentor reviewed parameterization, pagination bounds, and lab scope | Lab accepted |
| Transactions | [Lab 07 — learner implementation](labs/07-transactions/starter/starter.go); uncached tests, race test, and vet passed 2026-08-23; mentor reviewed validation, transaction callback behavior, error propagation, and atomic transfer invariant | Lab accepted |
| Functions / Runtime Semantics | Add links to learner code, tests, or reflection | Pending |
| Memory | Add links to learner experiment, escape-analysis output, or reflection | Pending |
| Concurrency | Add links to learner code, race-detector output, and design review | Pending |

## Completed Lab Evidence

| Lab | Date | Learner evidence | Verification | Mentor review |
|---|---|---|---|---|
| [Lab 01 — Go Basics](labs/01-go-basics/) | 2026-08-20 | Implemented `SumPositive` manually | `go test ./tests` passed | Completed: nil/empty slice behavior, non-mutating slice iteration, idiomatic `range` |
| [Lab 02 — Error Handling](labs/02-error-handling/) | 2026-08-20 | Implemented `ParsePort` manually | `go test ./tests ./starter` and `go vet ./...` passed | Completed: decimal parsing, inclusive bounds, sentinel wrapping, and `errors.Is` classification |
| [Lab 03 — Context: Cancellable Wait](labs/03-context/) | 2026-08-21 | Implemented and iteratively corrected `Wait` manually | `go test ./tests ./starter` passed | PASS: cancellation/deadline semantics, `ctx.Err()`, interruptible timer, cleanup, no goroutine, and `select` behavior |
| [Lab 05 — HTTP API: Task Handler](labs/05-http-api/) | 2026-08-23 | Implemented HTTP routing and JSON health response | `go test ./tests ./starter` passed; reference comparison completed | PASS: handler semantics, response ordering, status codes, `Allow`, `ServeMux`, and method-aware routing |
| [Lab 06 — PostgreSQL: Safe Pagination Query](labs/06-postgres/) | 2026-08-23 | Implemented parameterized query construction and pagination validation | `go test -count=1 ./...`, `go test -count=1 -race ./...`, and `go vet ./...` passed | PASS: placeholders, separate args, default/clamped limit, negative-offset rejection, and deliberate no-database scope |
| [Lab 07 — Transactions: Atomic Transfer](labs/07-transactions/) | 2026-08-23 | Implemented validation and transfer orchestration through a transaction callback | `go test -count=1 ./...`, `go test -count=1 -race ./...`, and `go vet ./...` passed | PASS: amount validation, debit-before-credit ordering, fail-fast errors, wrapped causes, rollback signaling, and atomicity invariant |

## Status values

`Not Started` · `Learning` · `Practicing` · `Review` · `Completed`

## Confidence scale

1. Heard about it
2. Understand the concept
3. Can implement it
4. Can debug it
5. Can teach it or design a production solution

## Weekly review

- What did I build or debug?
- What evidence changed my confidence?
- Which assumption was wrong?
- What failure mode did I test?
- What will I deliberately focus on next?

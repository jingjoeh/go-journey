# Code Review Checklist

## Behavior

- [ ] Requirements, invariants, and edge cases are clear.
- [ ] Errors preserve useful context and are handled at the correct boundary.
- [ ] Cancellation, deadlines, resource cleanup, and concurrency ownership are correct.
- [ ] Repeated or concurrent operations have defined behavior.

## Design

- [ ] Names and package boundaries communicate intent.
- [ ] Dependencies and abstractions are necessary and narrow.
- [ ] Data and transaction boundaries match business consistency needs.
- [ ] The change is small enough to review and reverse safely.

## Evidence

- [ ] Tests cover important behavior and failure paths without copying implementation.
- [ ] Race checks, integration tests, benchmarks, or query plans are included when relevant.
- [ ] Logs and metrics are useful, bounded, and free of sensitive data.
- [ ] Documentation explains non-obvious trade-offs.

## Risk

- [ ] Inputs, authorization, secrets, and data exposure were considered.
- [ ] Deployment, compatibility, migration, rollback, and overload behavior are understood.

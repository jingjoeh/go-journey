# Project 03 — Production Backend

## Purpose

Simulate senior ownership of an e-commerce, booking, or payment-like backend from requirements through operation.

## Required capabilities

- Modular architecture with explicit domain and data ownership
- PostgreSQL, Redis, and a durable queue
- Retry strategy, idempotency, duplicate handling, and distributed failure analysis
- Structured logs, metrics, tracing, health signals, and actionable alerts
- Rate limiting, graceful shutdown, configuration, and secret handling
- Load tests, Go benchmarks, and profiles tied to a performance decision
- Docker Compose and CI checks
- Architecture overview, diagrams where useful, and ADRs
- Completed production-readiness checklist and runbooks
- Design trade-offs and deliberate non-goals

## Delivery stages

1. Frame requirements, capacity estimates, invariants, threats, and non-goals.
2. Build one end-to-end path in a modular monolith.
3. Make data integrity and repeated execution correct.
4. Add asynchronous work and test partial failures.
5. Add telemetry and diagnose an injected incident using it.
6. Load test, profile, improve one measured bottleneck, and preserve evidence.
7. Rehearse deployment, migration, rollback, degradation, and recovery.

## Required review artifacts

- `docs/architecture.md`
- `docs/decisions/` containing ADRs based on `ADR-template.md`
- `docs/runbook.md`
- a filled copy of the repository production-readiness checklist
- load-test method and results
- a final “what I would change at 10× scale” review

## Definition of Done

Another engineer can run, test, operate, and review the service from its documentation. Important guarantees are proven by tests or operational exercises, and every significant complexity has a stated benefit and cost.

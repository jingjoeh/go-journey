# Production Readiness Checklist

## Service and data

- [ ] Owners, dependencies, SLOs, and critical user journeys are documented.
- [ ] Schema changes are backward compatible and have a rollback or roll-forward plan.
- [ ] Transactions, idempotency, retention, backup, and restoration are tested.

## Reliability and capacity

- [ ] Timeouts form a coherent budget; retries are bounded and jittered.
- [ ] Load is bounded with rate limits, queues, concurrency limits, or backpressure.
- [ ] Startup, shutdown, dependency loss, duplicate work, and partial failure are tested.
- [ ] Capacity assumptions are supported by load tests and headroom.

## Operations

- [ ] Logs, metrics, traces, dashboards, and alerts identify user impact and cause.
- [ ] Liveness and readiness checks have distinct, correct semantics.
- [ ] Runbooks cover common alerts, degradation, rollback, and recovery.

## Security and delivery

- [ ] Authentication, authorization, input limits, secrets, and dependency risks are reviewed.
- [ ] Builds are reproducible; CI gates tests and static analysis.
- [ ] Rollout is gradual and observable; rollback has been rehearsed.

## Decision

- [ ] Remaining risks have named owners and explicit acceptance.

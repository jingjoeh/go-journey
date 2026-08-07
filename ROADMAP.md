# Senior Go Backend Roadmap

Treat each phase as a capability gate, not a calendar deadline. Move forward when its goal is demonstrable, while revisiting earlier phases during projects.

## Phase 1 — Strong Go Fundamentals

**Topics:** types, pointers, slices, maps, interfaces, errors, `defer`, panic/recover, context, modules, and dependency management.

**Goal:** Write idiomatic Go without depending on frameworks.

**Evidence:** Build a small package with a narrow API, explicit errors, cancellation, table-driven tests, and documentation explaining ownership and edge cases.

## Phase 2 — Concurrency

**Topics:** goroutines, channels, `select`, mutexes, `RWMutex`, `WaitGroup`, atomics, worker pools, races, deadlocks, and cancellation.

**Goal:** Understand when concurrency helps and when it makes systems worse.

**Evidence:** Implement a bounded worker pool, test cancellation and leaks, run the race detector, and benchmark against a sequential version.

## Phase 3 — HTTP Backend

**Topics:** `net/http`, routing, middleware, request lifecycle, validation, error handling, pagination, and API versioning.

**Goal:** Build production-quality HTTP services.

**Evidence:** Deliver a versioned API with consistent errors, limits, request IDs, timeouts, graceful shutdown, and handler tests.

## Phase 4 — PostgreSQL

**Topics:** schema design, indexes, `EXPLAIN`, `EXPLAIN ANALYZE`, transactions, isolation, locking, deadlocks, connection pooling, and migrations.

**Goal:** Diagnose database problems, not just write SQL.

**Evidence:** Explain a real query plan, improve it measurably, reproduce a lock conflict, and choose transaction boundaries intentionally.

## Phase 5 — Architecture

**Topics:** package design, modular monoliths, clean and hexagonal architecture, dependency inversion, and domain boundaries.

**Goal:** Choose architecture based on trade-offs rather than patterns.

**Evidence:** Produce two viable designs for the same service, implement one, and document why its complexity is justified.

## Phase 6 — Testing

**Topics:** table-driven tests, helpers, mocks, fakes, repository tests, integration tests, API tests, and testcontainers.

**Goal:** Design code that is naturally testable.

**Evidence:** Create a useful test pyramid that catches a seeded regression without over-mocking implementation details.

## Phase 7 — Distributed Systems

**Topics:** retries, idempotency, queues, eventual consistency, distributed locks, duplicate messages, at-least-once delivery, and rate limiting.

**Goal:** Understand failure scenarios.

**Evidence:** Demonstrate a duplicate delivery, partial failure, and retry storm; then implement bounded, observable mitigations.

## Phase 8 — Production Engineering

**Topics:** structured logging, metrics, tracing, health checks, graceful shutdown, configuration, secrets, Docker, and CI/CD.

**Goal:** Operate the service, not only write it.

**Evidence:** Use telemetry to diagnose a deliberately injected failure and complete the production-readiness checklist.

## Phase 9 — Performance

**Topics:** benchmarks, pprof, CPU and heap profiles, allocation analysis, database optimization, and load testing.

**Goal:** Measure before optimizing.

**Evidence:** Preserve a benchmark, profile, hypothesis, change, and before/after result for one genuine bottleneck.

## Phase 10 — System Design

**Practice:** booking, payment, notification, marketplace, and file-processing systems.

**Goal:** Explain requirements, constraints, trade-offs, bottlenecks, failure modes, and scaling strategy.

**Evidence:** Complete a 45-minute design, defend it under changing requirements, and identify what you would validate before building.

## Capstone gate

Complete Project 03 and conduct a mock production review covering correctness, data integrity, security, reliability, observability, performance, delivery, and deliberate omissions.

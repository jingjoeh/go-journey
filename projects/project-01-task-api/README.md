# Project 01 — Task API

## Purpose

Learn clean Go API fundamentals with a deliberately simple architecture.

## Required capabilities

- REST API built with `net/http` or a minimal router
- PostgreSQL schema and reversible migrations
- Create, read, update, delete, list, validation, and bounded pagination
- Stable structured error responses
- Unit, handler, repository, and integration tests
- Docker image and local Compose environment

## Suggested milestones

1. Write the API contract, task invariants, schema, and migration.
2. Implement an in-memory vertical slice and handler tests.
3. Add PostgreSQL through a narrow repository boundary.
4. Add integration tests for constraints, pagination, and concurrent updates.
5. Add configuration, timeouts, logs, health checks, graceful shutdown, and Docker.

## Deliberate limits

Keep one deployable service, synchronous behavior, and simple package boundaries. Do not add Redis, queues, generic repositories, or microservices.

## Definition of Done

The service starts predictably, migrates safely, passes tests from a clean environment, rejects invalid input consistently, shuts down cleanly, and includes a review explaining transaction boundaries and omissions.

# Project 02 — Order Service

## Purpose

Introduce production backend patterns through an order lifecycle where concurrency and data integrity matter.

## Required capabilities

- Users, orders, line items, and inventory
- Authentication and resource-level authorization
- PostgreSQL transactions and explicit order state transitions
- Idempotent order creation
- Bounded pagination
- Optimistic and pessimistic locking experiments with recorded evidence
- Redis caching with an invalidation strategy
- Durable background jobs with bounded retry behavior
- Structured logging, metrics, and integration tests

## Invariants to defend

- Inventory cannot be silently oversold.
- An idempotency key cannot create two logical orders.
- Order totals are reproducible from immutable pricing inputs.
- Unauthorized users cannot observe or mutate an order.
- Retried background work cannot duplicate an external effect.

## Failure drills

Simulate concurrent checkout, a database timeout after commit, Redis loss, duplicate jobs, a poison message, and shutdown during work. Record detection, user impact, and recovery.

## Definition of Done

Demonstrate invariants with tests, compare both locking approaches, expose useful telemetry, and document where exactly-once behavior is neither claimed nor required.

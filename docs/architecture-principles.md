# Architecture Principles

1. Start from business invariants and failure modes.
2. Keep dependencies pointing toward stable policy, not infrastructure details.
3. Prefer explicit data flow and ownership over magical reuse.
4. Put transaction boundaries where consistency requirements demand them.
5. Separate deployability only when operational independence is worth its cost.
6. Design interfaces at the consumer and keep them narrow.
7. Make retries, timeouts, idempotency, and backpressure deliberate.
8. Avoid abstractions until at least two real uses reveal the shared shape.
9. Preserve observability across boundaries.
10. Record consequential decisions and the conditions that would reverse them.

## Decision record questions

- What forces are in tension?
- Which options were seriously considered?
- What evidence supports this choice?
- What new risks and operational costs does it introduce?
- How reversible is it?
- What signal would cause us to revisit it?

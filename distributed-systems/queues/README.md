# Queues

## Why this matters

Senior backend engineers need to reason about delivery semantics, ordering, acknowledgements, backpressure, and poison messages. The skill matters because local correctness is not enough: choices must remain understandable under concurrency, failure, growth, and operational pressure.

## Core Concepts

- Define the vocabulary and the guarantees each mechanism does and does not provide.
- Make ownership, lifecycle, boundaries, and failure behavior explicit.
- Separate correctness requirements from performance preferences.
- Use small experiments and production evidence to validate assumptions.

## Production Use Cases

Apply this topic while designing service boundaries, reviewing changes, diagnosing incidents, and deciding whether added complexity is justified. Connect every use to a user-visible goal, an invariant, or an operational risk.

## Common Mistakes

- Copying a pattern without identifying the problem it solves.
- Designing only for the happy path or assuming dependencies are always available.
- Hiding important behavior behind a broad abstraction.
- Optimizing from intuition without a baseline.
- Adding retries or concurrency without bounds, cancellation, and observability.

## Exercises

1. Explain queues in your own words, including one guarantee and one non-guarantee.
2. Implement the smallest example that demonstrates its main behavior.
3. Add a failure, cancellation, or concurrent case and predict the result before running it.
4. Review a production-style design and identify one simpler alternative.
5. Record what you would measure or alert on in production.

See [exercises.md](exercises.md) for the working checklist.

## Senior Questions

- Why would you choose this design, and what did you reject?
- What happens under high concurrency or if a dependency becomes unavailable?
- Can this operation run twice safely?
- How would you observe and debug it in production?
- What changes at 10x traffic?
- What would you deliberately not abstract?

## Definition of Done

You can explain the mechanism without notes, implement a representative use, test its important failure behavior, diagnose a broken version, and defend a production choice with evidence and trade-offs.


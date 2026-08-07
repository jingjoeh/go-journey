# Senior Go Backend Engineer Bootcamp

This repository is a self-directed curriculum for growing from competent Go developer to senior backend engineer. It combines concept study, executable labs, production-style projects, system-design practice, and explicit reflection on trade-offs and failure modes.

## How to use this repository

1. Read the [Learner Contract](LEARNER_CONTRACT.md) and [Learning Rules](LEARNING_RULES.md), then choose the next phase in [ROADMAP.md](ROADMAP.md).
2. Study the relevant topic README and keep your own understanding in `notes.md`.
3. Complete the exercises before opening a challenge solution or asking an AI for complete code.
4. Work in each lab's `starter/` directory. Run its tests, then compare with `solution/` only after making a serious attempt.
5. Apply several modules together in the projects. Record important decisions and their trade-offs.
6. Update [PROGRESS.md](PROGRESS.md) with evidence, not optimism.

## Learning loop

`Learn -> Explain -> Implement -> Test -> Break -> Observe -> Review -> Teach`

The curriculum is intentionally standard-library-first. Libraries and infrastructure are introduced when their trade-offs become visible.

## Repository map

- `foundations/`: Go language, errors, interfaces, context, memory, and concurrency.
- `backend/`, `database/`, `architecture/`: production service fundamentals.
- `distributed-systems/`, `reliability/`, `performance/`, `security/`: operating under failure and scale.
- `testing/`: testing at unit, integration, repository, API, and load levels.
- `system-design/`: repeatable design exercises with capacity and failure analysis.
- `labs/`: focused, executable practice. Do not inspect `solution/` first.
- `projects/`: three progressively more realistic services.
- `interview/`: question banks and mock-interview practice.
- `templates/`: reusable lesson, lab, review, project, and incident formats.
- `LEARNER_CONTRACT.md`: ownership boundaries for the learner, mentor, and Codex.

## Definition of success

Finishing files is not the goal. You are ready for senior-level work when you can independently choose a design, explain what you rejected, implement and test it, diagnose it under failure, observe it in production, and revise it using evidence.

## Suggested local tools

- A recent stable Go toolchain
- PostgreSQL and Redis
- Docker with Docker Compose
- Git and a GitHub account for CI practice

Start with Phase 1 in [ROADMAP.md](ROADMAP.md).

# AI-Assisted Learning Guide

AI should increase the quality of your reasoning, not remove the productive struggle that creates it.

The role boundaries in [LEARNER_CONTRACT.md](LEARNER_CONTRACT.md) govern all AI-assisted work in this repository.

## ChatGPT role

Use ChatGPT as a teacher, mentor, system-design interviewer, code reviewer, concept explainer, and learning planner. Ask it to challenge assumptions, find knowledge gaps, design exercises, and review architecture decisions.

ChatGPT should not immediately provide a complete solution unless you explicitly request one. A useful prompt is: “Ask one diagnostic question at a time; give hints before code; make me defend trade-offs.”

## Codex role

Use Codex as an implementation assistant, repository maintainer, test generator, and refactoring partner. It may create files, implement starter code, create tests, run tests, fix compilation issues, and refactor code.

Codex should not silently change architecture decisions. Ask it to state the reason, impact, alternatives, and migration path before a structural change.

## Prompt pattern

Include:

1. Your goal and current phase
2. What you already believe
3. Your attempt or design
4. Constraints and deliberate exclusions
5. The review lens you want
6. The maximum level of help allowed

Example: “I am practicing idempotency. Review my design for duplicate concurrent requests. Do not write code. Ask questions about races, storage boundaries, expiry, and observability.”

## Verification rules

- Run generated code and tests yourself.
- Verify API and library claims against primary documentation.
- Treat confident explanations as hypotheses until evidence supports them.
- Never paste secrets, customer data, or proprietary incidents into a model.
- Record what you learned, not merely what the AI produced.

## Suggested AI review questions

- What assumption in my design is most dangerous?
- Which failure is untested?
- Where is ownership unclear?
- What becomes invalid at 10× traffic?
- What is the simplest experiment that would disprove my hypothesis?

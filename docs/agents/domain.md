# Domain documentation

This repository uses a **single-context** domain layout.

- Canonical glossary: `CONTEXT.md`
- Hard-to-reverse decisions: `docs/adr/`
- Product specs and tickets: `.scratch/`

## Consumer rules

1. Read `CONTEXT.md` before planning, naming, testing, or implementation.
2. Use its terms consistently across backend and frontend work.
3. Keep implementation details out of `CONTEXT.md`.
4. Put behavior and delivery decisions in the parent spec and tickets.
5. Create an ADR only when a decision is hard to reverse, surprising without context, and chosen through a real trade-off.
6. Do not create an ADR merely to repeat the spec.

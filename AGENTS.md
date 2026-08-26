# Agent instructions

## Product repositories

This repository is the Donetick backend and the canonical planning repository for the gamified household life system.

The paired frontend is `Jayvbooth/frontend`. Keep the existing two-repository architecture. A vertical feature may require one backend PR and one frontend PR with the same feature slug and reciprocal links.

## Agent skills

### Issue tracker

Specs and implementation tickets live as local Markdown under `.scratch/`. See `docs/agents/issue-tracker.md`.

### Domain docs

This is a single-context product. Read `CONTEXT.md` before planning or implementation. Keep it as a glossary only. See `docs/agents/domain.md`.

### Ponytail

Read `.agents/rules/ponytail.md` before designing, coding, choosing dependencies, or reviewing. Reuse Donetick behavior before adding new systems. Ponytail remains active unless the user explicitly disables it.

## Delivery workflow

1. Synthesize an approved parent spec.
2. Split it into vertical tracer-bullet tickets with explicit blockers.
3. Start each ticket from a fresh context and a branch from `develop`.
4. Read the ticket, parent spec, `CONTEXT.md`, and relevant decisions.
5. Implement one behavior at a time through pre-agreed public test seams.
6. Run focused checks during development and the complete relevant suites at the end.
7. Run Matt Pocock Standards and Spec review against the merge base.
8. Run a separate Ponytail deletion-focused review.
9. Fix findings, commit, push, and open a PR.
10. Use paired backend/frontend PRs only when the vertical slice genuinely crosses repositories.

Do not create implementation tickets from unapproved prototype styling. The actual Donetick component system, navigation shell, typography, spacing, and interaction conventions remain the visual foundation.

## Current parent spec

`.scratch/gamified-life-system/spec.md`

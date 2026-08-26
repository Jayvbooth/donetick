# Issue tracker: Local Markdown

GitHub Issues are disabled on this fork, so specs and tickets for this product live in the repository.

## Conventions

- One feature per directory: `.scratch/<feature-slug>/`
- Parent spec: `.scratch/<feature-slug>/spec.md`
- Implementation tickets: `.scratch/<feature-slug>/issues/<NN>-<slug>.md`
- One ticket per file, numbered in dependency order
- Every ticket declares `Blocked by` and `Status`
- Tickets produced from an approved spec start at `Status: ready-for-agent`
- Discussion or review notes append under `## Comments`

## Vertical delivery

Each implementation ticket must deliver a narrow, complete behavior through every layer it needs. Avoid separate "all database," "all API," and "all UI" tickets.

When a slice touches both repositories:

- Use the same feature slug in `Jayvbooth/donetick` and `Jayvbooth/frontend`.
- Open paired PRs and link them to each other.
- Do not merge either PR until the combined behavior is verifiable.

## Publishing

When a skill says "publish to the issue tracker," create or update the appropriate Markdown file under `.scratch/`.

The current parent spec is `.scratch/gamified-life-system/spec.md`.

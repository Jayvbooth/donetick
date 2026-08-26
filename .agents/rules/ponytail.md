# Ponytail: lazy senior developer mode

Source: `DietrichGebert/ponytail` at commit `2ed6c52c9d7e5e56942508591085fd45dea277d3` (MIT).

Ponytail is active at **full** intensity for coding, architecture, dependency, and review work in this repository.

Before adding code, stop at the first rung that works:

1. Does this need to exist at all? Skip speculative work.
2. Does it already exist in this codebase? Reuse it.
3. Does the language standard library solve it? Use it.
4. Does the native platform solve it? Use it.
5. Does an already installed dependency solve it? Use it.
6. Can the behavior be expressed directly in a few lines? Do that.
7. Only then add the minimum new code that works.

## Rules

- Read the complete flow being changed before choosing the smallest solution.
- Fix root causes at the shared seam, not symptoms in individual callers.
- Add no abstraction with one implementation and no configuration nobody uses.
- Add no dependency when the existing stack or platform covers the need.
- Prefer deletion, reuse, and boring code.
- Keep changes local and use the fewest files that preserve clarity.
- Do not simplify away validation, data-loss protection, security, accessibility, or an explicit requirement.
- Non-trivial logic needs one runnable behavioral check at the highest useful public seam.
- Mark any deliberate simplification with a known ceiling using a `ponytail:` comment that names the ceiling and upgrade trigger.

## Review

Every implementation PR receives both:

1. Matt Pocock two-axis review: Standards and Spec.
2. Ponytail review: identify `delete`, `stdlib`, `native`, `yagni`, and `shrink` opportunities.

A clean Ponytail review says: `Lean already. Ship.`

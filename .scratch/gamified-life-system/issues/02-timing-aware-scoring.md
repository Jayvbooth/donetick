# 02: Add timing-aware scoring to the existing Task result flow

**Parent spec:** `../spec.md`  
**Blocked by:** None (can start immediately)  
**Status:** ready-for-agent  
**Repositories:** paired backend and frontend PRs

## What to build

Extend the existing Chore completion, ChoreHistory, and Points flow so a scored Task or Routine earns a clear signed point result based on its timing mode, due time, completion outcome, and recovery state. Show the available reward when configuring and viewing a Task, then show a concise explanation when the result is recorded.

## Acceptance criteria

- [ ] A Task can use one of the approved timing modes: untimed, today, deadline, or window.
- [ ] Existing Tasks preserve current fixed-point behavior by default after migration.
- [ ] On-time completion, useful early completion, staged lateness, a single bounded miss, and later recovery follow the defaults in the parent spec.
- [ ] Early bonus is opt-in and is never applied to an action whose timing mode makes early completion inappropriate.
- [ ] Reminder delivery, opening, and acknowledgement have no scoring effect.
- [ ] The Member's displayed Circle Points never fall below zero because of a penalty.
- [ ] The completion transaction records the final signed point result exactly once while preserving current recurrence, assignment, approval, privacy, and history behavior.
- [ ] ChoreHistory retains only the compact information required to render a human-readable score breakdown; no occurrence or event ledger is added.
- [ ] The existing Task editor exposes only the minimum timing controls needed for these modes and keeps advanced thresholds centralized.
- [ ] Existing Task cards/details show the available reward without replacing the real Donetick visual language.
- [ ] Completion feedback shows a short breakdown such as base reward, timing adjustment, recovery, and total.
- [ ] Worked behavioral tests cover on-time, useful early, grace-period, late-stage, missed, recovered, non-negative balance, and duplicate-submission cases.
- [ ] Additive migrations are verified for supported SQLite and PostgreSQL paths.

## Constraints carried from the parent spec

- Reuse Chore, ChoreHistory, Circle Points, and the current completion transaction.
- Do not add a task-occurrence subsystem, second currency, gamification package, or notification telemetry.
- Use one pure scoring seam with a small input and result.
- No new frontend dependency.
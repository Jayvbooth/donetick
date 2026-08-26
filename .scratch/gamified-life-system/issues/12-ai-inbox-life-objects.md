# 12: Extend AI Inbox proposals to the household life domains

**Parent spec:** `../spec.md`  
**Blocked by:** 07: Extend Projects into Goals with Milestones and rewards; 08: Add Calendar Items, birthdays, and a unified household calendar; 09: Add standalone Reminders and birthday preparation; 10: Add recurring Bills with scoring and household surfaces; 11: Add an AI Inbox for Task and subtask proposals  
**Status:** ready-for-agent  
**Repositories:** paired backend and frontend PRs

## What to build

Extend the approved AI Inbox flow so one messy typed or spoken capture can propose Goals, Milestones, Calendar Items, birthdays, standalone Reminders, Bills, Tasks, and subtasks. Commit each approved proposal through the authoritative domain behavior rather than through a universal record layer.

## Acceptance criteria

- [ ] Structured proposals support Goal, Milestone, Calendar Item, birthday, standalone Reminder, Bill, Task, and subtask types.
- [ ] The proposal preview uses the existing editor patterns for each type and clearly names what will be created.
- [ ] Money, dates, recurrence, responsible Member, assignment, and private/shared visibility require confirmation when ambiguous or low-confidence.
- [ ] Household context is limited to the minimum visible Circle members and records needed for interpretation.
- [ ] The Member can edit, remove, approve individually, or approve selected proposals together.
- [ ] Each approved type is committed through its existing public create behavior so validation, privacy, recurrence, scoring, and history remain authoritative.
- [ ] A repeated commit request remains idempotent across mixed proposal types.
- [ ] Failure of one proposal is reported without silently committing an invalid version; successfully committed proposals are not duplicated on retry.
- [ ] The model never contacts people, schedules without approval, spends money, changes existing records, or makes a high-stakes decision.
- [ ] Raw private capture is not retained permanently by default.
- [ ] Behavioral tests cover mixed captures, type validation, ambiguous fields, partial failure, idempotency, privacy, and no-AI operation.
- [ ] No generic object store, plugin architecture, workflow engine, or agent framework is introduced.

## Constraints carried from the parent spec

- Reuse ticket 11's Inbox and each domain's existing create seam.
- Keep the proposal schema explicit and finite.
- Ponytail review should prefer a direct type switch over speculative polymorphism unless the codebase already has the needed pattern.
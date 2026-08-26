# 07: Extend Projects into Goals with Milestones and rewards

**Parent spec:** `../spec.md`  
**Blocked by:** 02: Add timing-aware scoring to the existing Task result flow  
**Status:** ready-for-agent  
**Repositories:** paired backend and frontend PRs

## What to build

Extend Donetick's existing Project behavior so a Project can act as a Goal, keep its linked Tasks through the existing Project relationship, and contain ordered Milestones with target dates and rewards. Present Goals through the real Project UI patterns rather than creating a second planning system.

## Acceptance criteria

- [ ] Existing Projects continue to behave as they do today after additive migration.
- [ ] A Project can optionally be a Goal with an owner, active/paused/completed/archived status, optional target date, and progress mode.
- [ ] Existing Tasks remain linked through the current Project relationship; no parallel Task-to-Goal join model is introduced.
- [ ] Goal progress can derive from linked Task completion or Milestone completion without duplicate manual bookkeeping.
- [ ] A Goal can contain ordered Milestones with title, optional notes, optional target date, completion state, and base reward.
- [ ] Completing a Milestone records its reward once through the existing Points system and uses the shared scoring seam when timing affects the result.
- [ ] Goal and Milestone privacy follows Circle, creator, owner, and private visibility behavior consistent with existing Tasks.
- [ ] The frontend evolves the current Projects list/detail/edit experience using existing Donetick components, routes, responsive behavior, and localization.
- [ ] Linked Tasks can be opened and managed through their authoritative Task routes.
- [ ] Goal target dates and Milestones expose the minimum date-range data required by the later calendar ticket without duplicating records.
- [ ] Behavioral tests cover legacy Project compatibility, linking, progress, milestone reward idempotency, status changes, and privacy.
- [ ] Migrations are verified on supported SQLite and PostgreSQL paths.

## Constraints carried from the parent spec

- Reuse Project and Chore.ProjectID.
- Do not add a second project manager, roadmapping framework, or generic hierarchy system.
- Do not redesign the application or add a progress/diagram dependency.
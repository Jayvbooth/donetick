# 06: Add Daily Reset and momentum through ordinary Donetick behavior

**Parent spec:** `../spec.md`  
**Blocked by:** 02: Add timing-aware scoring to the existing Task result flow; 03: Turn the root route into a focused Today experience; 05: Add the I'm Stuck, help-request, and supporter-suggestion flow  
**Status:** ready-for-agent  
**Repositories:** paired backend and frontend PRs

## What to build

Add a Daily Reset that helps a Member deliberately resolve unfinished Tasks, review tomorrow, choose priorities, and close the day. Represent the reset itself as an ordinary recurring Donetick Task and derive momentum from its history rather than creating a workflow or streak platform.

## Acceptance criteria

- [ ] A Member can enable a recurring Daily Reset at a chosen local evening time.
- [ ] Opening the reset shows today's unresolved Tasks and tomorrow's currently scheduled Tasks using existing visibility rules.
- [ ] Every unresolved Task can be completed, moved, broken down, sent for help, skipped, or cancelled through existing actions.
- [ ] The reset cannot close while an unresolved Task remains without a deliberate resolution.
- [ ] The Member can identify tomorrow's top priorities using existing Task priority behavior rather than a separate pinning database.
- [ ] A short household closing checklist is represented by ordinary recurring Tasks or subtasks, not a workflow builder.
- [ ] Completing the Daily Reset records its normal scored completion and advances or protects the visible momentum streak.
- [ ] Momentum is derived from Daily Reset completion history; missing an unrelated Task does not independently destroy it.
- [ ] The Member can close the day even when work remains, provided each remaining item has been intentionally moved, delegated, skipped, or cancelled.
- [ ] Today shows the current momentum cue using restrained Donetick styling and reduced-motion-safe feedback.
- [ ] The aggregate is designed so later Calendar Items, Reminders, Milestones, and Bills can appear without changing the reset's core interaction.
- [ ] Behavioral tests cover incomplete resolution, successful close, priority selection, streak calculation, privacy, and timezone day boundaries.

## Constraints carried from the parent spec

- Do not create a separate workflow engine, streak service, or event ledger.
- Reuse Chore, ChoreHistory, existing task actions, and the score from ticket 02.
- Add no gamification or animation dependency.
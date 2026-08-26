# 05: Add the I'm Stuck, help-request, and supporter-suggestion flow

**Parent spec:** `../spec.md`  
**Blocked by:** 04: Add bounded reminder presets and reusable notification routing  
**Status:** ready-for-agent  
**Repositories:** paired backend and frontend PRs

## What to build

Let a Member respond to an unresolved Task by saying what is blocking it and choosing a useful next action. Support explicit help requests and supporter suggestions without giving Supporters a surveillance view or silently assigning private work.

## Acceptance criteria

- [ ] A visible `I'm stuck` action is available from the normal Task surface.
- [ ] The Member can choose a concise Blocker such as unclear start, too large, missing something, overwhelmed, no time, avoiding, or needing help.
- [ ] The flow can reveal the first existing subtask, add or generate a short checklist, create a prerequisite Task, reschedule, add a note, or request help.
- [ ] Only the current Member-selected Blocker and chosen resolution are retained; notification-by-notification behavior is not stored.
- [ ] A help request goes only to a selected Circle Member and contains only the Task information needed to help plus the authoritative route.
- [ ] Existing Task privacy rules prevent a help request from exposing unrelated or private content.
- [ ] A Supporter can propose a shared Task to another Member, but it becomes assigned work only after the recipient accepts or schedules it.
- [ ] The recipient can accept, schedule, decline, or discuss a suggestion.
- [ ] Existing Circle roles and notification behavior are reused before any new permission framework or messaging system is added.
- [ ] Blockers, suggestions, and help requests do not automatically add or subtract Points.
- [ ] Behavioral tests cover private-item denial, explicit help, suggestion acceptance/decline, and each resolution path at the highest practical seam.
- [ ] The UI uses current Donetick sheets/dialogs, typography, localization, keyboard behavior, and mobile patterns.

## Constraints carried from the parent spec

- No Supporter dashboard of ignored reminders, snoozes, failures, or mistakes.
- Prefer existing notes, subtasks, Task creation, activity, and notifications; add one small record only if the accepted behavior cannot be represented safely without it.
- Do not add chat, a social feed, or a generic approval framework.
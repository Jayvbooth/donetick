# 04: Add bounded reminder presets and reusable notification routing

**Parent spec:** `../spec.md`  
**Blocked by:** None (can start immediately)  
**Status:** ready-for-agent  
**Repositories:** paired backend and frontend PRs

## What to build

Make persistent reminders easy to configure by adding a few bounded presets that populate Donetick's existing notification templates. Generalize the notification target just enough that the same delivery path can later open Tasks, Bills, Reminders, Calendar Items, and Daily Reset.

## Acceptance criteria

- [ ] The Task editor offers clear normal, persistent, and important reminder presets.
- [ ] Each preset maps to no more than the existing maximum number of notification templates and remains editable through current controls.
- [ ] Presets use existing before-due, due, and after-due scheduling rather than a new retry loop or rapid poller.
- [ ] Completing, rescheduling, skipping, cancelling, archiving, or otherwise resolving a Task prevents obsolete future nudges through the current planner lifecycle.
- [ ] Notification payloads support a small authoritative route target while preserving existing Task deep links.
- [ ] Tapping a Task notification opens the correct Task in web/PWA and existing native paths.
- [ ] No acknowledgement, notification-open, or snooze-history ledger is created.
- [ ] Important behavior is visibly described as bounded and is not presented as an Apple Critical Alert.
- [ ] Planner tests verify exact scheduled outputs, cancellation after resolution, maximum-template limits, and route payloads.
- [ ] Frontend checks verify preset selection, custom editing, and current Donetick responsive/accessibility patterns.

## Constraints carried from the parent spec

- Reuse NotificationMetadata, NotificationTemplate, the existing planner, scheduler, and delivery adapters.
- Do not add a generic automation builder or a second reminder scheduler.
- Add no dependency for preset generation or routing.
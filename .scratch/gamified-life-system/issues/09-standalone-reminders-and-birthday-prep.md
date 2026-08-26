# 09: Add standalone Reminders and birthday preparation

**Parent spec:** `../spec.md`  
**Blocked by:** 04: Add bounded reminder presets and reusable notification routing; 08: Add Calendar Items, birthdays, and a unified household calendar  
**Status:** ready-for-agent  
**Repositories:** paired backend and frontend PRs

## What to build

Add the smallest dedicated Reminder record needed for prompts that are not completion-based Tasks. Deliver those prompts through the existing notification architecture, show them in Today, Calendar, and Daily Reset, and use them to support advance birthday preparation without automatically creating obligations.

## Acceptance criteria

- [ ] A standalone Reminder supports title, optional notes, owner, schedule, recurrence required by accepted stories, active state, private/shared visibility, and an optional relation to another record.
- [ ] Members can create, edit, view, deactivate, and remove a Reminder using current Donetick UI patterns.
- [ ] Reminder scheduling reuses the existing offset vocabulary, scheduler, and delivery adapters rather than creating a second scheduler.
- [ ] A Reminder notification routes to the correct authoritative Reminder or related record using ticket 04's target shape.
- [ ] Today, the unified calendar, and Daily Reset show visible Reminders at the appropriate time.
- [ ] Reminder delivery and opening do not change Points and are not recorded as behavioral analytics.
- [ ] A birthday can schedule useful advance prompts such as think about a gift, buy a gift, or call the person.
- [ ] Birthday preparation becomes a real Task only after the Member approves the suggestion.
- [ ] Dismissing, deactivating, or deleting a Reminder cancels obsolete future delivery.
- [ ] Private Reminders and related private records are protected at the server seam.
- [ ] Behavioral tests cover recurrence, cancellation, relation routing, birthday rollover, approval-before-Task creation, and privacy.
- [ ] Additive migrations are verified for supported SQLite and PostgreSQL paths.

## Constraints carried from the parent spec

- No acknowledgement ledger, unlimited retry loop, generic rule builder, People CRM, or automatic gift purchasing.
- Reuse existing notification infrastructure and installed frontend stack.
- Keep important prompts bounded enough that Members will not disable the app.
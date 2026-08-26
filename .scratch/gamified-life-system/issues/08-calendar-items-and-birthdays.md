# 08: Add Calendar Items, birthdays, and a unified household calendar

**Parent spec:** `../spec.md`  
**Blocked by:** 06: Add Daily Reset and momentum through ordinary Donetick behavior; 07: Extend Projects into Goals with Milestones and rewards  
**Status:** ready-for-agent  
**Repositories:** paired backend and frontend PRs

## What to build

Add a minimal Calendar Item domain for non-task facts tied to time, then extend the existing Donetick calendar to display visible Tasks, Routines, Goal Milestones, appointments, events, birthdays, work shifts, and family events through one date-range read.

## Acceptance criteria

- [ ] A Calendar Item can represent an appointment, event, birthday, work shift, or family event.
- [ ] Calendar Items support title, notes, start, optional end, all-day state, the recurrence needed by accepted stories, owner, Circle, active state, and private/shared visibility.
- [ ] Birthdays are yearly all-day Calendar Items; no People CRM is introduced.
- [ ] Members can create, edit, view, and remove Calendar Items through current Donetick form, dialog, and route patterns.
- [ ] A date-range query returns only visible Tasks, Routines, Calendar Items, and Goal Milestones needed for that range.
- [ ] The unified result is a read model; source records are not copied into a universal calendar table.
- [ ] Selecting an entry opens its authoritative Task, Goal/Milestone, or Calendar Item route.
- [ ] The current `react-calendar` and Donetick component system are extended rather than replaced by another calendar package.
- [ ] The calendar offers a useful month plus agenda/list experience on desktop and mobile.
- [ ] Daily Reset shows tomorrow's visible Calendar Items and Milestones through the same read behavior.
- [ ] Private entries are excluded for unauthorized Circle Members at the server seam, not merely hidden in the client.
- [ ] Behavioral tests cover range boundaries, timezone handling, recurrence, birthday rollover, source routing, and privacy.
- [ ] Additive migrations are verified for supported SQLite and PostgreSQL paths.

## Constraints carried from the parent spec

- No external calendar sync in this ticket.
- No universal object framework, copied event store, or new date/calendar dependency.
- Preserve current Donetick navigation, density, themes, accessibility, and localization.
# 03: Turn the root route into a focused Today experience

**Parent spec:** `../spec.md`  
**Blocked by:** 02: Add timing-aware scoring to the existing Task result flow  
**Status:** ready-for-agent  
**Repositories:** primarily `Jayvbooth/frontend`; backend only for a demonstrated missing query

## What to build

Use the current Donetick data and design system to make the root route a focused Today surface. It should reduce choice by separating urgent commitments, a small automatically selected priority set, Routines, and flexible work while keeping the full Task database available at `/chores`.

## Acceptance criteria

- [ ] The root route shows a focused Today experience inside the existing Donetick shell.
- [ ] The complete Task database remains available through the existing Tasks route and behavior.
- [ ] Today separates overdue/urgent items, up to three current priorities, due Routines, and flexible items.
- [ ] The initial top priorities are derived from existing due date and priority data; this ticket does not add a separate pinning database.
- [ ] Task cards retain their real Donetick structure and actions while adding restrained reward and timing information from ticket 02.
- [ ] A Member can complete, start, reschedule, skip, or open a Task through existing behavior.
- [ ] Empty, loading, error, long-list, and mobile states are handled using current Donetick patterns.
- [ ] The existing Points data can show an optional cooperative weekly target as shared progress without ranking Members.
- [ ] No public individual leaderboard or failure dashboard is added.
- [ ] New strings use the current localization system and interactions remain keyboard-accessible, labeled, responsive, and reduced-motion safe.
- [ ] Focused interaction checks verify section assignment, priority selection, normal Task actions, and navigation to the full Task list.

## Constraints carried from the parent spec

- The discarded standalone prototype is not a visual reference.
- Reuse current queries and client-side derivation before adding a Today aggregate endpoint.
- Add no new state, UI, chart, or gamification dependency.
- Do not hide important Tasks merely because they fall outside the top three.
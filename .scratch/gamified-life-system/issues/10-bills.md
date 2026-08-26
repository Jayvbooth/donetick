# 10: Add recurring Bills with scoring and household surfaces

**Parent spec:** `../spec.md`  
**Blocked by:** 02: Add timing-aware scoring to the existing Task result flow; 03: Turn the root route into a focused Today experience; 06: Add Daily Reset and momentum through ordinary Donetick behavior; 08: Add Calendar Items, birthdays, and a unified household calendar  
**Status:** ready-for-agent  
**Repositories:** paired backend and frontend PRs

## What to build

Add a compact Bill domain for one-time and recurring household obligations. Let a responsible Member record, pay, verify autopay, skip, or deactivate a Bill; preserve compact history; apply deadline scoring once; and surface Bills in Today, Calendar, and Daily Reset.

## Acceptance criteria

- [ ] A Bill supports name, optional notes/payee, amount, currency, next due date, recurrence required by accepted stories, responsible Member, autopay state, active state, base reward, and private/shared visibility.
- [ ] Members can create, edit, view, and deactivate Bills using current Donetick UI patterns.
- [ ] A Bill can be upcoming, due, paid, overdue, skipped, or inactive.
- [ ] Paying or deliberately resolving a recurring Bill creates one compact history result and advances the next due date correctly.
- [ ] Repeating the same payment request cannot apply Points, history, or recurrence advancement twice.
- [ ] Bill payment uses ticket 02's deadline scoring seam for on-time, late, missed, and recovered results.
- [ ] Autopay changes the action language to verification rather than pretending the Member must manually pay it.
- [ ] Today shows Bills before and when they become due.
- [ ] The unified calendar displays Bills from their authoritative records and opens the Bill route.
- [ ] Daily Reset includes tomorrow's and unresolved visible Bills.
- [ ] Bill privacy and responsibility are enforced at the server seam.
- [ ] Behavioral tests cover recurrence, month-end dates, amount/history preservation, idempotent payment, scoring, autopay verification, visibility, and timezone handling.
- [ ] Additive migrations are verified for supported SQLite and PostgreSQL paths.

## Constraints carried from the parent spec

- No bank, card, payment-provider, transaction-import, budgeting, or autopay-control integration.
- Do not disguise Bills as Tasks or Calendar Items.
- Reuse the existing scoring, notification routing, Points, and Donetick UI stack.
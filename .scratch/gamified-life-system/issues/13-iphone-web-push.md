# 13: Deliver self-hosted iPhone Web Push and Focus setup

**Parent spec:** `../spec.md`  
**Blocked by:** 04: Add bounded reminder presets and reusable notification routing; 08: Add Calendar Items, birthdays, and a unified household calendar; 09: Add standalone Reminders and birthday preparation; 10: Add recurring Bills with scoring and household surfaces  
**Status:** ready-for-agent  
**Repositories:** paired backend and frontend PRs

## What to build

Let a Member install the self-hosted Donetick PWA on an iPhone, subscribe that Home Screen app to Web Push, receive the existing bounded notifications, and open the correct authoritative record. Include a clear setup path for permission and allowing the app through selected iPhone Focus modes.

## Acceptance criteria

- [ ] The installed Home Screen PWA can request notification permission and create a Web Push subscription on supported iOS versions.
- [ ] The server stores only the subscription data needed to deliver to that Member and supports removing a device.
- [ ] The existing notification scheduler feeds a small Web Push delivery adapter rather than a separate scheduling system.
- [ ] Use one established Web Push library only if protocol encryption/signing cannot be safely provided by the standard library or existing dependencies; do not implement cryptography manually.
- [ ] Invalid or expired subscriptions are cleaned up after definitive delivery failure.
- [ ] Task, Calendar Item, Reminder, Bill, and Daily Reset notifications open the correct route using ticket 04's target shape.
- [ ] Normal, persistent, and important presets deliver as bounded notifications without an unlimited retry loop.
- [ ] The setup screen explains Home Screen installation, permission, device status, test notification, and how to allow the PWA in selected Focus modes.
- [ ] The product clearly states that Web Push is not an Apple Critical Alert and may not override every silent-device condition.
- [ ] No SMS, Pushover, App Store entitlement project, or guaranteed mute-switch bypass is added.
- [ ] Service-worker changes preserve current PWA caching and update behavior.
- [ ] Behavioral checks cover subscription validation, ownership, delivery payloads, route targets, removal, and expired endpoint cleanup.
- [ ] A manual iPhone smoke-test checklist covers install, permission, Focus allowance, foreground/background delivery, tap routing, and known silent-mode limits.
- [ ] Setup and notification actions are localized, keyboard-accessible where applicable, and reduced-motion safe.

## Constraints carried from the parent spec

- Reuse the current PWA, service worker, scheduler, notification records, and route behavior.
- Add no second notification service or mobile rewrite.
- Keep any new dependency narrowly justified in the PR and subject to Ponytail review.
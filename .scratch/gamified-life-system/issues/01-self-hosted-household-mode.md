# 01: Simplify the self-hosted household account experience

**Parent spec:** `../spec.md`  
**Blocked by:** None (can start immediately)  
**Status:** ready-for-agent  
**Repositories:** `Jayvbooth/frontend`; backend only if an existing capability flag cannot express the behavior

## What to build

Make the self-hosted product present the account experience this household actually needs: separate username/password accounts, Circle setup, invitations, joining, and member management. Hide irrelevant hosted-product and advanced account surfaces without rewriting or weakening Donetick authentication.

## Acceptance criteria

- [ ] Existing username/password signup, login, logout, password reset, and session behavior continue to work.
- [ ] Existing Circle creation, invite/join, and member management remain available.
- [ ] In the household self-hosted configuration, unused social-login, payment/subscription, MFA, API-token, analytics, and developer surfaces are hidden from normal navigation and settings.
- [ ] Existing optional capabilities remain available when explicitly enabled through current configuration or capability checks; no feature is deleted from upstream code merely to hide it here.
- [ ] No new authentication model, identity provider, role system, or dependency is introduced.
- [ ] The resulting settings and onboarding screens use the current Donetick shell, MUI components, spacing, themes, mobile behavior, and localization.
- [ ] Focused route/component checks prove the household mode exposes the simple account path and hides disabled surfaces.

## Constraints carried from the parent spec

- Preserve Donetick security, validation, and password handling.
- Reuse existing feature/config checks before adding a new flag.
- This ticket must not become an authentication refactor.
- Ponytail review should find no speculative account abstraction.
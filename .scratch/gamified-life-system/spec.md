# Gamified Household Life System

**Status:** awaiting-user-review  
**Planning repository:** `Jayvbooth/donetick`  
**Paired frontend repository:** `Jayvbooth/frontend`  
**Target branch:** `develop`

## Problem Statement

Donetick already provides a strong shared task and chore foundation, but the intended household needs more than a list of recurring chores. One Member has persistent executive-function weaknesses involving initiation, planning, sequencing, working memory, task completion, error recognition, and judgment under stress. The household needs an external support system that makes everyday responsibilities easier to start and finish without turning family members into permanent managers.

The current product also lacks a unified way to handle goals, appointments, birthdays, bills, standalone reminders, end-of-day preparation, and raw thought capture. These needs should live in one coherent household experience rather than being scattered across texts, calendars, notes, and verbal reminders.

The solution must remain respectful and adult. It must not become surveillance software, expose detailed notification behavior, shame a Member, or visually replace Donetick with a childish game interface. Gamification should make completion, recovery, and preparation feel rewarding while preserving the actual Donetick UI, component system, and interaction language.

The implementation must remain maintainable for a solo developer using agents. It should reuse Donetick's existing Chore, ChoreHistory, Circle, Points, Projects, notification, PWA, voice, and frontend infrastructure before adding any new model or dependency.

## Solution

Extend Donetick into a shared household life system while retaining the existing product shell and visual design.

The product will provide:

- A focused Today experience built from existing Tasks, Routines, Calendar Items, Bills, and Goal Milestones.
- Timing-aware Points and XP that reward completion, sensible early action, recovery, and a consistent Daily Reset.
- Persistent but bounded nudges, plus an "I'm stuck" path that helps a Member identify the next useful action.
- A Daily Reset that resolves unfinished work and prepares tomorrow without requiring a perfect day.
- Goals and Milestones linked to existing Tasks.
- A shared calendar containing Tasks, Routines, appointments, events, birthdays, Bills, Reminders, and Goal Milestones.
- Bill tracking with due, paid, and overdue states.
- An AI Inbox that converts typed or spoken raw captures into editable proposals and never commits them without approval.
- Separate Member accounts using the existing simple username/password flow.
- Existing shared and private visibility behavior, with help requested explicitly rather than inferred from hidden telemetry.
- iPhone PWA notifications using the existing self-hosted product as far as the browser and current notification infrastructure allow.

Gamification will appear as restrained additions to the real Donetick experience: point values, progress, completion feedback, streak or momentum cues, and cooperative household progress. It will not introduce a fantasy dashboard, a separate game navigation shell, or an individual public leaderboard.

## User Stories

### Everyday Tasks and Today

1. As a Member, I want one Today view to show what matters now, so that I do not need to choose from the entire task database.
2. As a Member, I want Today to separate urgent commitments from flexible work, so that deadlines do not get buried.
3. As a Member, I want recurring household responsibilities to appear automatically, so that noticing the need is not required.
4. As a Member, I want a Task to show its next physical action or checklist, so that vague responsibilities become startable.
5. As a Member, I want long checklists to reveal one current step at a time when useful, so that the full task does not feel overwhelming.
6. As a Member, I want to start, complete, reschedule, skip, or mark a Task as stuck from its normal Donetick surface.
7. As a Member, I want overdue Tasks to remain resolvable rather than accumulating forever as red clutter.
8. As a Member, I want a small set of top priorities, so that the app guides attention without hiding the rest of my commitments.
9. As a Member, I want shared household Tasks and private Tasks to coexist, so that one app can support home and personal responsibilities.
10. As a Circle admin, I want to assign or rotate household Routines using Donetick's existing behavior, so that the extension does not replace working chore management.

### Points, XP, and Motivation

11. As a Member, I want each scored Task to show its available reward, so that effort has an immediate visible payoff.
12. As a Member, I want on-time completion to award the full base reward.
13. As a Member, I want a sensible early bonus only on Tasks where early completion is actually useful.
14. As a Member, I do not want an early bonus for actions that should happen inside a specific time window, such as medication or leaving for an appointment.
15. As a Member, I want lateness to reduce the reward in clear stages, so that the consequence is understandable rather than arbitrary.
16. As a Member, I want a missed commitment to cause at most one bounded point loss, so that one Task cannot drain points forever.
17. As a Member, I want completing an overdue Task to earn recovery credit, so that returning to a failure remains worthwhile.
18. As a Member, I want the completion feedback to explain the final score in a few human-readable parts.
19. As a Member, I want my total Points or XP never to be driven below zero by penalties.
20. As a Member, I want reminders themselves to have no point penalty, so that I am scored on the result rather than notification taps.
21. As a Member, I want a momentum or streak cue based primarily on completing the Daily Reset, so that consistency matters more than perfection.
22. As a Member, I want Goal Milestones to award larger rewards than ordinary small Tasks.
23. As a household, we want an optional cooperative weekly target, so that shared progress feels collaborative rather than competitive.
24. As a Member, I do not want a public ranking that labels one person as the worst performer.
25. As a Member, I want the existing Points history to show concise score outcomes without exposing a behavioral surveillance log.

### Reminders, Nudges, and Getting Stuck

26. As a Member, I want reminder presets such as normal, persistent, and important, so that I do not configure every offset manually.
27. As a Member, I want a persistent Task to nudge me at bounded intervals until I complete, reschedule, skip, cancel, or mark it stuck.
28. As a Member, I want a nudge to open the relevant item directly.
29. As a Member, I want to snooze or reschedule without falsely marking an item complete.
30. As a Member, I want the app to ask what is blocking me after a Task remains unresolved, rather than sending the same message forever.
31. As a Member, I want blocker choices such as unclear start, too large, missing something, overwhelmed, no time, avoiding, or needing help.
32. As a Member, I want the app to respond to my blocker by showing the first step, creating a prerequisite, breaking down the Task, rescheduling it, or requesting help.
33. As a Member, I want blocker reporting to be optional and concise.
34. As a Member, I do not want the app to record a timeline of every notification I opened or ignored.
35. As a Supporter, I want to receive a help request only when the Member asks or an agreed important-item rule applies.
36. As a Supporter, I want the help request to explain the Task and requested help without exposing unrelated private items.
37. As a Member, I want standalone Reminders for information or prompts that are not completion-based Tasks.
38. As a Member, I want important reminder behavior to remain rare enough that I do not disable notifications entirely.

### Daily Reset

39. As a Member, I want a Daily Reset prompt at a chosen evening time.
40. As a Member, I want the Daily Reset to surface today's unresolved Tasks.
41. As a Member, I want each unresolved Task to be completed, moved, broken down, sent for help, skipped, or cancelled.
42. As a Member, I want tomorrow's appointments, Bills, Tasks, and preparation needs shown during the Daily Reset.
43. As a Member, I want to choose tomorrow's top priorities during the Daily Reset.
44. As a Member, I want a configurable short household closing checklist without a separate workflow-builder system.
45. As a Member, I want completing the Daily Reset to protect or advance my momentum streak.
46. As a Member, I want to close the day even when some Tasks remain unfinished, provided each one has been resolved deliberately.
47. As a household, we want shared closing responsibilities to use ordinary recurring Donetick Tasks.

### Goals and Milestones

48. As a Member, I want to define a Goal with a purpose, owner, status, and optional target date.
49. As a Member, I want existing Tasks to belong to a Goal without duplicating them.
50. As a Member, I want a Goal to contain Milestones that represent meaningful checkpoints.
51. As a Member, I want Goal progress to derive from linked Tasks or Milestones rather than requiring duplicate manual updates.
52. As a Member, I want a Goal to be active, paused, completed, or archived.
53. As a Member, I want Goal target dates and Milestones to appear in the shared calendar according to visibility.
54. As a Member, I want completing a Milestone to award its configured reward.
55. As a Member, I want Goals to reuse Donetick Projects wherever their behavior overlaps.

### Shared Calendar and Birthdays

56. As a Member, I want one calendar to show Tasks, Routines, Calendar Items, Bills, Reminders, and Goal Milestones.
57. As a Member, I want the current Donetick calendar UI to be extended rather than replaced by a new visual system.
58. As a Member, I want appointments and events to have a start, optional end, all-day state, recurrence, notes, and visibility.
59. As a Member, I want birthdays to recur yearly and appear as all-day Calendar Items.
60. As a Member, I want birthday reminders far enough ahead to think about a gift or plan.
61. As a Member, I want an optional birthday preparation suggestion to become a real Task only after approval.
62. As a Member, I want a work shift or family event to be represented without pretending it is a Task.
63. As a Member, I want calendar entries to open the correct underlying record.
64. As a Member, I want private calendar entries hidden from other Circle members.
65. As a Member, I want a date-range calendar query that remains useful without downloading every historical record.

### Bills

66. As a Member, I want to record a Bill with a name, amount, due date, recurrence, responsible Member, and autopay state.
67. As a Member, I want a Bill to be upcoming, due, paid, overdue, skipped, or inactive.
68. As a Member, I want recurring Bills to advance to the next due date after payment or deliberate resolution.
69. As a Member, I want Bill payment history to preserve the due date, amount, payment time, responsible Member, and final point result.
70. As a Member, I want a Bill to appear in Today and the calendar before it becomes overdue.
71. As a Member, I want paying a Bill on time to earn its reward and paying late to use the same clear scoring principles as a deadline Task.
72. As a Member, I want autopay Bills to be marked as such so the app prompts verification rather than manual payment.
73. As a household, we want Bill visibility and responsibility to be controllable without connecting bank accounts.

### AI Inbox and Capture

74. As a Member, I want to type or speak an unstructured brain dump into one Inbox.
75. As a Member, I want the existing Donetick voice capture and deterministic parsers reused before a model is called.
76. As a Member, I want a messy multi-item capture to be proposed as Tasks, subtasks, Calendar Items, Bills, Reminders, Goals, and Milestones.
77. As a Member, I want every proposal to show its interpreted type, title, date, recurrence, assignee, privacy, and relevant details.
78. As a Member, I want to edit, remove, or approve proposals individually or approve them together.
79. As a Member, I want ambiguous dates, money, recurrence, or assignees flagged for confirmation.
80. As a Member, I do not want AI proposals committed automatically.
81. As a Member, I want one approval action to commit the selected proposals without creating duplicates if it is tapped twice.
82. As a Member, I want raw capture text retained only as needed to review the current proposal, not as a permanent behavioral archive by default.
83. As a self-hosting admin, I want AI to use a configured OpenAI-compatible endpoint such as LiteLLM rather than a hard-coded vendor.
84. As a Member on a supported native device, I want the existing local model capability to remain available as a fallback where practical.
85. As a Member, I want AI failure to leave the raw capture intact and editable rather than losing it.
86. As a Member, I want the app to work normally when AI is not configured.

### Household Accounts, Privacy, and Support

87. As a household, we want each person to have a separate username and password.
88. As a Circle admin, I want to invite or create household Members using Donetick's existing account and Circle behavior.
89. As a self-hosting admin, I want optional OAuth, MFA, developer, and subscription surfaces hidden when they are not used, without rewriting authentication.
90. As a Member, I want shared items visible to the Circle and private items protected by existing Donetick privacy rules.
91. As a Member, I want Supporters to suggest a shared Task without silently adding commands to my private list.
92. As a Member, I want to accept, schedule, decline, or discuss a suggestion.
93. As a Member, I want no supporter dashboard showing every ignored reminder, snooze, or mistake.
94. As a Supporter, I want only the minimum information needed to help with a shared Task or explicit help request.
95. As a Circle admin, I want existing admin and manager roles reused before any new permission system is considered.
96. As a Member, I want sensitive work and personal items private by default when I choose that setting.

### PWA, Notifications, and Quality

97. As a Member, I want to install the existing self-hosted app as a Home Screen PWA on iPhone.
98. As a Member, I want the PWA to request notification permission through a clear setup flow.
99. As a Member, I want instructions for allowing the app through selected iPhone Focus modes.
100. As a Member, I understand that ordinary Web Push is not an Apple Critical Alert and may not override every silent-device condition.
101. As a self-hosting admin, I want the existing notification scheduler and delivery architecture reused before adding a new service.
102. As a Member, I want notification actions and links to route to the correct Task, Bill, Reminder, Calendar Item, or Daily Reset.
103. As a Member, I want motion and reward effects to respect reduced-motion preferences.
104. As a keyboard or assistive-technology user, I want all new interactions labeled and operable without relying on color alone.
105. As a household using multiple languages, we want new user-facing strings to use Donetick's existing localization system.
106. As a self-hosting admin, I want additive migrations to work with both supported SQLite and PostgreSQL deployments.
107. As a maintainer, I want the extension to preserve upstream compatibility and minimize edits to unrelated Donetick code.
108. As a maintainer, I want no new frontend dependency when the current MUI, React, calendar, chart, PWA, date, and Capacitor stack can deliver the behavior.

## Implementation Decisions

### Product and repository shape

1. The actual Donetick frontend is the design system. New surfaces will use its existing shell, navigation drawer, MUI Joy and Material components, typography, spacing, dialogs, mobile behavior, themes, and localization patterns.
2. Gamification is an enhancement to existing task cards, completion feedback, Today, Points, and progress surfaces. It is not a replacement visual identity and will not copy the discarded prototype styling.
3. The backend and frontend remain separate upstream-compatible forks. This spec is canonical in the backend planning repository. Cross-repository tickets will use paired branches and linked PRs.
4. No monorepo migration, shared framework, or new product shell is introduced.
5. The first implementation tickets will make the existing behavior easier to extend only where a demonstrated seam is missing. No broad preparatory refactor is authorized by this spec.

### Authentication, users, and privacy

6. Existing Donetick username/password authentication, Circle membership, invitations, and roles remain authoritative.
7. The self-hosted product may hide optional auth, subscription, and developer surfaces that are irrelevant to this household. Authentication internals are not rewritten.
8. Existing shared/private Chore visibility is reused. New record types must apply equivalent creator, owner, Circle, and private visibility rules.
9. Support begins with existing Circle roles, shared items, suggestions, and explicit help requests. A granular permission framework is added only if a concrete accepted story cannot be satisfied without it.
10. The product will not expose notification acknowledgement history, snooze history, or a detailed behavioral dashboard to Supporters.

### Tasks, Routines, and scoring

11. Existing Chore remains the write model for both Tasks and Routines. Existing recurrence, assignment, rotation, subtasks, priority, approval, privacy, notification templates, Projects, timers, and history are reused.
12. Existing ChoreHistory remains the record of each completed, skipped, missed, rejected, or rescheduled result. A new task-occurrence subsystem is explicitly rejected for this scope.
13. Scoring behavior lives behind one small, pure scoring interface. It receives the existing base Points value, Timing Mode, due time, completion or resolution time, result, and minimal streak context. It returns a signed point change and concise player-facing breakdown.
14. A Chore gains only the minimum scoring fields not derivable from current data: Timing Mode and whether a useful early bonus is enabled. Shared default thresholds remain centralized rather than exposing a large per-Task rules editor.
15. Supported Timing Modes are:
    - **untimed:** fixed base reward; time does not change the score.
    - **today:** full reward through the Member's local day; bounded missed and recovery behavior after day close.
    - **deadline:** full reward through a grace period; reward decreases in clear late stages.
    - **window:** completion is expected within a useful time range; an absurdly early completion receives no early bonus and may be blocked by existing completion-window behavior.
16. Default scoring policy:
    - On-time completion earns the full base reward.
    - A useful early completion may earn up to a 20 percent bonus, rounded to whole Points.
    - Completion inside a 15-minute grace period earns the full reward.
    - Completion after grace and up to 2 hours late earns 80 percent of base.
    - Completion 2 to 8 hours late earns 60 percent of base.
    - Completion more than 8 hours late or after the relevant day closes earns 30 percent of base.
    - A still-unresolved missed item may deduct 30 percent of base once, rounded and capped so one item cannot repeatedly drain Points.
    - Later recovery earns at least the late-completion floor and may restore part of the missed deduction.
    - Total household Points displayed for a Member never fall below zero.
17. Reminder delivery, opening, and acknowledgement do not affect scoring.
18. ChoreHistory stores the final signed point result and a compact scoring breakdown sufficient to explain it. No event warehouse is added.
19. The existing Points balance and history are extended rather than replaced by a second currency. "XP" is the user-facing presentation of the same Points.
20. The main continuity cue is tied to deliberately closing the day through Daily Reset, not to completing every Task perfectly.
21. Any cooperative household target is additive and non-ranking. Individual public leaderboards are excluded.

### Nudges, blockers, and Daily Reset

22. Existing notification templates remain the mechanism for before-due, due, and after-due Task nudges.
23. The UI adds bounded presets that populate existing templates. Presets may be customized using existing controls, but no general automation-builder is introduced.
24. A persistent sequence stops when the Task is completed, rescheduled, skipped, cancelled, archived, or marked stuck.
25. "I'm stuck" captures one optional current Blocker and one resolution action. It does not capture a notification-by-notification history.
26. Resolution actions reuse existing behavior wherever possible: reveal or create subtasks, create a prerequisite Task, reschedule, add a note, or notify a selected Supporter.
27. A help request should be represented with the existing Task, note/activity, and notification capabilities unless implementation proves that one small dedicated record is necessary.
28. Daily Reset is represented by an ordinary recurring Donetick Task plus a dedicated aggregate view. It is not a separate workflow engine.
29. The Daily Reset aggregate reads unresolved Tasks and tomorrow's visible Calendar Items, Bills, Reminders, and Milestones. Each unresolved Task must receive a deliberate resolution before the reset can close.
30. Household closing checklist items remain normal recurring Tasks or subtasks.

### Goals, calendar, reminders, and Bills

31. Existing Project is extended minimally to represent both Projects and Goals. A Goal adds an owner, status, optional target date, and progress mode.
32. Existing Chore.ProjectID continues to link Tasks to Goals. No parallel Task-to-Goal relationship is added.
33. Milestones are a small separate record because they carry ordered Goal progress, an optional target date, completion, and a reward without pretending every checkpoint is a Task.
34. Calendar Items use one model with a kind for appointments, events, birthdays, work shifts, and family events. Birthday is a yearly all-day Calendar Item in version one; a People CRM is not introduced.
35. The existing calendar library and Donetick visual component patterns are retained. A unified date-range read returns visible Tasks, Routines, Calendar Items, Bills, standalone Reminders, and Milestones as display entries.
36. The unified calendar is a read model, not a second copy of every source record. Selecting an entry opens its authoritative source.
37. Standalone Reminders receive one small dedicated record because they are not completion-based Tasks. They support title, owner, schedule, recurrence, visibility, active state, and optional source relation. They do not introduce a generic rules engine or acknowledgement ledger.
38. Task reminder offsets continue to use existing Chore notification metadata. New record types use the same offset vocabulary and existing notification delivery where practical rather than creating parallel schedulers.
39. Bills remain separate from Calendar Items because amount, recurrence, responsible Member, autopay, payment, and overdue state are domain behavior, not display metadata.
40. A recurring Bill stores its current next due date. Each payment or deliberate resolution creates compact Bill history and advances the Bill. Bank or card connections are not required.
41. Bill scoring calls the same scoring interface used by deadline Tasks.

### AI Inbox

42. Existing voice capture, deterministic natural-language parsing, Circle-member parsing, label parsing, recurrence parsing, due-date parsing, and current task creation are reused.
43. Deterministic parsing handles simple captures. A model is called only for messy multi-item classification, decomposition, or ambiguity that current parsing cannot resolve.
44. The self-hosted server uses a configured OpenAI-compatible endpoint, including LiteLLM, through the smallest reliable HTTP integration. A vendor-specific agent framework is not introduced.
45. AI returns structured proposals for supported object types. The frontend presents editable proposal cards using existing Donetick components.
46. Nothing is committed until the Member approves it. Money, dates, recurrence, assignment, and private/shared visibility require confirmation when confidence is insufficient.
47. Committing approved proposals is idempotent so retries do not create duplicates.
48. The uncommitted raw capture remains client-side where practical. The server does not retain a permanent raw-thought archive by default.
49. AI is optional. Missing configuration or a model failure leaves ordinary capture and manual creation fully usable.
50. Existing native local-model support remains available where it works, but version one does not require feature parity between local and server models.

### PWA and notifications

51. The existing PWA, service worker, Capacitor, local notification, push notification, and server scheduler foundations are reused.
52. Self-hosted iPhone Web Push uses the browser Push API and service worker. The implementation may use one small established Web Push library if required by protocol encryption rather than implementing cryptography manually.
53. The app provides an explicit setup path for Home Screen installation, notification permission, and allowing the app in chosen iPhone Focus modes.
54. Version one does not claim Apple Critical Alert behavior and does not add SMS or Pushover.
55. Notification payloads use a small common route target so taps can open Tasks, Bills, standalone Reminders, Calendar Items, or Daily Reset without rebuilding the delivery architecture.
56. Persistent nudges are bounded by preset offsets. The system does not poll rapidly or send an unlimited retry loop.

### Maintainability and delivery

57. No new frontend design system, calendar package, chart package, date package, state library, form framework, or gamification package is added unless a ticket demonstrates that the installed stack cannot satisfy an accepted story.
58. New backend models use additive migrations compatible with SQLite and PostgreSQL.
59. New user-facing strings use the existing localization system.
60. New UI preserves keyboard operation, accessible labels, contrast, reduced motion, and existing responsive conventions.
61. New domains are online-first in version one unless the accepted ticket specifically requires offline mutation. Existing Chore offline behavior must not regress. A generic offline-sync rewrite is excluded.
62. Each implementation ticket is a vertical, demoable slice. Any slice crossing repositories uses paired PRs.
63. Every implementation PR receives:
    - focused behavioral tests at the highest useful public seam,
    - Matt Pocock Standards review,
    - Matt Pocock Spec review,
    - Ponytail over-engineering review.
64. No ticket may silently add functionality listed under Out of Scope.

## Testing Decisions

### Agreed public seams

1. **Scoring seam:** Given a Scoring Policy and a completed, missed, or recovered result, return the signed point change and concise breakdown. Tests use worked literal examples from this spec and do not inspect internal helper functions.
2. **Task completion seam:** Through the existing authenticated completion behavior, a Member completes a scored Task and receives the correct ChoreHistory and Points balance result in one transaction.
3. **Recurring Task seam:** Completing, skipping, missing, or rescheduling a scored Routine preserves current Donetick recurrence and assignment behavior while recording the correct score once.
4. **Nudge seam:** Given an unresolved Task and its existing notification templates, the public scheduling behavior produces only the bounded future nudges and stops after a resolving action.
5. **Daily Reset seam:** Through the Daily Reset view, a Member resolves each unfinished Task, reviews tomorrow, and closes the reset; the resulting ordinary Tasks and Points are observable through normal product interfaces.
6. **Goal seam:** Through Goal behavior, a Member links existing Tasks, completes a Milestone, and sees progress and reward update without duplicating the Task.
7. **Calendar seam:** A date-range query for an authenticated Member returns the correct unified entries and excludes records the Member cannot view.
8. **Bill seam:** Paying, skipping, or resolving a recurring Bill records one history result, applies scoring once, and advances the next due date correctly.
9. **AI proposal seam:** Given raw capture and household context, parsing returns structured proposals; no source record exists until the explicit commit behavior succeeds.
10. **AI commit seam:** Repeating the same approved commit request is idempotent.
11. **Frontend interaction seams:** Tests operate through visible routes, controls, and query/mutation interfaces rather than snapshots of implementation details.
12. **Notification routing seam:** A notification target opens the correct authoritative record type.

### Verification approach

- Work test-first in one red-to-green vertical behavior at a time.
- Prefer existing repository test patterns and highest-level handlers or pure interfaces.
- Do not mock internal collaborators merely to assert call order.
- Run focused backend or frontend checks during each slice.
- Run full relevant suites, typechecking, linting, and builds before review.
- Verify migrations on supported SQLite and PostgreSQL paths before a schema PR is ready.
- Perform a manual iPhone Home Screen PWA smoke test for permission, delivery, Focus configuration, tap routing, and silent-mode limitations.
- Perform manual accessibility smoke checks for keyboard use, labels, focus, contrast, and reduced motion on changed surfaces.
- Review every diff against both the parent spec and repository standards.
- Run Ponytail review separately and remove speculative abstractions, unused configuration, duplicate logic, and unnecessary dependencies before merge.

## Out of Scope

1. Replacing Donetick's current UI with the discarded prototype, a fantasy game shell, or a new design system.
2. Medical diagnosis, treatment, clinical monitoring, competency assessment, or claims that this product cures executive dysfunction.
3. A notification acknowledgement ledger, detailed snooze analytics, behavioral event warehouse, or Supporter surveillance dashboard.
4. Public individual leaderboards, shaming, coercive score visibility, or automatic punishment controlled by another Member.
5. A new authentication system, enterprise identity project, passkeys, mandatory MFA, or an OAuth expansion.
6. A monorepo migration or a rewrite of the backend/frontend release architecture.
7. Bank account, card, payroll, budgeting, transaction-import, or payment-provider integration.
8. SMS, Pushover, Apple Critical Alert entitlement, or guaranteed bypass of every device mute state.
9. A full People or relationship CRM. Version one stores birthdays as Calendar Items.
10. An autonomous AI agent that creates records, changes schedules, contacts people, spends money, or makes high-stakes decisions without confirmation.
11. Permanent storage of raw private brain dumps by default.
12. AI-driven psychological profiling or reminder adaptation based on hidden behavior.
13. A generic automation engine, rule builder, event-sourcing platform, or universal object framework.
14. Full offline creation and synchronization for every new domain.
15. Replacing Donetick Projects, Points, notification templates, Circle roles, Chore recurrence, or ChoreHistory with parallel systems.
16. A reward store, virtual currency economy, avatars, cosmetic marketplace, or complex achievements in version one.
17. Automatic external calendar sync in version one.
18. Automatic gift purchasing or high-stakes financial, legal, medical, or safety decisions.

## Further Notes

- This parent spec intentionally covers the complete product direction. The next phase is to split it into small vertical tickets with explicit blockers, not to implement the entire document in one branch.
- The first implementation slice should prove timing-aware scoring through the current Task completion path and existing Donetick UI before broader life objects are added.
- The frontend must be evaluated inside the real Donetick application with real density and existing components. The earlier standalone prototype is not a visual reference.
- Default scoring values are centralized initial policy, not a promise to expose every threshold as user configuration. Change the constants only when real household use demonstrates a need.
- AGPL obligations from the Donetick forks remain applicable.
- Ponytail is active for every ticket: first reuse existing behavior, then add only the smallest missing capability.

# 11: Add an AI Inbox for Task and subtask proposals

**Parent spec:** `../spec.md`  
**Blocked by:** None (can start immediately)  
**Status:** ready-for-agent  
**Repositories:** paired backend and frontend PRs

## What to build

Add one Inbox surface where a Member can type or speak a raw thought dump and receive editable Task and subtask proposals. Reuse Donetick's existing voice capture and deterministic parsing first, call an optional OpenAI-compatible endpoint only when the input needs multi-item classification or decomposition, and never create records without approval.

## Acceptance criteria

- [ ] A Member can type or use existing voice capture to create a raw Inbox draft.
- [ ] Existing deterministic date, recurrence, assignee, label, priority, and Task parsing handles simple captures without a model call.
- [ ] When configured, an OpenAI-compatible endpoint such as LiteLLM can turn a messy capture into one or more proposed Tasks with optional subtasks.
- [ ] The model integration uses the smallest reliable HTTP path and structured output; no vendor-specific agent framework is introduced.
- [ ] Every proposal is shown in editable Donetick form/card patterns with title, due date, recurrence, assignee, privacy, priority, reward, and subtasks as applicable.
- [ ] Ambiguous dates, recurrence, assignment, or visibility are visibly flagged for confirmation.
- [ ] Nothing is committed until the Member approves the proposal or selected proposals.
- [ ] Committing the same approved proposal request twice cannot create duplicate Tasks.
- [ ] Raw capture remains client-side by default and survives parsing or network failure so it can be edited or retried.
- [ ] The app remains fully usable for ordinary manual capture when AI is not configured or fails.
- [ ] Existing native local-model support may be used as a fallback where practical without requiring parity.
- [ ] Behavioral tests cover deterministic parsing, model-output validation, ambiguous proposals, approval, idempotent commit, unavailable AI, and privacy.
- [ ] The Inbox route uses the current shell, localization, keyboard behavior, responsive layout, and no new UI dependency.

## Constraints carried from the parent spec

- No autonomous creation, permanent raw-thought archive, psychological profiling, or background agent.
- Reuse existing capture and Task creation behavior before adding code.
- Validate all model output at the trust boundary.
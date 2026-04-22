# Deliverable Completion Flow

Completion flow when creating documents, code, or configuration inside a skill. After the initial hearing and direction approval, proceed autonomously to completion without mid-process approval gates.

## Principles

- **Approval happens only once, up front**: After hearing gathers requirements and direction (with an optional single approval of the overall plan), run to completion without further approval requests
- **Hearing is consolidated at the start**: Collect unclear requirements, direction, and selection criteria in the opening hearing ([hearing.md](hearing.md))
- **Plan mode's ExitPlanMode** counts as this single up-front approval. Do not repeat approval requests at every phase transition
- **Exceptions**: Confirm only immediately before destructive or irreversible actions (deleting production data, sending to external services, destructive changes to the user environment)

## Flow

1. **Hearing**: Interactively confirm requirements, direction, and selection criteria at the start ([hearing.md](hearing.md))
2. **Direction approval (only when needed)**: Use Plan mode (or equivalent) to get one-shot approval of the overall scope. No further mid-process approvals after this
3. **Autonomous execution**: Produce the deliverable (documents, code, configuration) through to completion
4. **Self-review**: Review your own output. See [../perspectives/design-and-coding.md](../perspectives/design-and-coding.md) / [../perspectives/testing.md](../perspectives/testing.md) / [../perspectives/documentation.md](../perspectives/documentation.md) for perspectives
5. **Report to user**: Briefly present the deliverable's key points, risks, and remaining tasks (not to request approval, but to communicate what was built and what is needed next)

## ❓-Marked Assumptions

Assumptions needing confirmation should be resolved in the opening hearing by default. If a `[❓ needs confirmation]` item must wait, resolve it in a single batch before starting the related task (do not turn it into an approval gate).

## Applicable Situations

- After creating documents (Design Docs, test cases, OKRs, etc.)
- After changing settings (rules, skills, settings, etc.)
- After completing implementation

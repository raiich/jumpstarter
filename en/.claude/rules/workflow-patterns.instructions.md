# Common Workflow Patterns

## Basic Pattern: Deliverable Creation and Review

1. **Claude: Create deliverable**
2. **Claude: Self-review** - Conduct a critical review including impact analysis (described below), and fix as needed
3. **User: Review** - Always conduct user review after self-review
4. **Claude: Revise** - Revise if there is feedback
5. **Repeat**: Repeat the above until the user approves

Where the flow mentions **basic pattern**, follow this pattern.

## Self-Review Principles

### Critical Review

Even without an explicit user request, conduct a critical review and state candid assessments and concerns before proceeding.

- Point out problems and concerns without hesitation
- Don't settle for "looks good" — actively look for weaknesses and risks
- After stating concerns, also present mitigation options
- **Means-ends alignment**: Verify that the chosen means are appropriate for the user's objective. Do not adopt means that don't fit the objective, even if they address gaps. If there is a mismatch, propose alternative means

### Impact Analysis

After changes, conduct impact analysis including reference pattern consistency checks.

- Use Grep to search for locations that reference changed elements (paths, terms, interfaces, etc.)
- Verify that references are consistent with the changes
- If inconsistencies are found, batch-fix them (following fix-guidelines principles)

## Replanning When Stuck

When an approach isn't working, stop and reassess instead of forcing ahead.

**Replanning triggers:**
- Repeating the same error 2+ times
- Workarounds are becoming complex
- Initial assumptions differ significantly from reality

**Response:**
1. Organize the current problems
2. Consider alternative approaches
3. Consult the user if needed

## Context Management

The context window is the most critical resource. Manage it proactively.

- **When switching tasks**: Reset with `/clear`
- **Long sessions**: Summarize with `/compact` (e.g., `Focus on the API changes`)
- **Research tasks**: Delegate to subagents to protect the main context
- **Fix fails twice**: `/clear` and restart with a new prompt that includes what you learned

## Providing Verification Methods

Deliverable quality depends on having verification methods. Secure verification methods at the start of a task.

- **Code changes**: Test commands, expected output, screenshots
- **Documentation changes**: Self-review checklist, comprehension check by target audience
- **Configuration changes**: Commands to compare behavior before and after the change

If no verification method exists, create one first (write tests, define a checklist, etc.).

## Completion Criteria

Do not mark a deliverable as complete until it meets the following.

- **Code changes**: Build succeeds, tests pass, existing tests are not broken
- **Documentation changes**: Self-review checklist passed
- **Common**: Self-review completed, verified with verification methods

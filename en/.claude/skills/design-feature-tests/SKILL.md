---
name: design-feature-tests
description: Plan implementation and design test cases based on a design doc.
allowed-tools: Read, Grep, Glob, Edit, Write, Task, AskUserQuestion, EnterPlanMode, ExitPlanMode, Skill
---

# Design Feature Tests

A skill for creating an implementation plan and designing test cases based on a Design Doc.

## Prerequisites

- Input is the deliverable of the `design-feature` skill (`design-doc.md`) or direct user input
- The deliverable is saved as `test-cases.md` and serves as input for the `implement-feature` skill
- **Approval gate**: Test cases are not considered complete without explicit user approval

## Output Locations

- **Knowledge base**: Topic-specific files under `.local/docs/`
- **Feature-specific documents**: Under `.local/docs/features/[name]/`
  - `test-cases.md` - Test case design (primary deliverable of this skill)

## Flow

### [Input Verification]

#### 1. Verify Design Doc

If `design-doc.md` exists, read it. Otherwise, understand requirements from user input.
Hear from the user if there are unclear points.

**Tools**: Read, AskUserQuestion

### [Implementation Planning Phase] — Plan mode

#### 2. Enter Plan mode

**Tools**: EnterPlanMode

#### 3. Codebase investigation and implementation planning

Based on the Design Doc, investigate the existing codebase and record the implementation plan in the plan file.
Use Task (Explore agent) when broad exploration is needed.

**Tools**: Read, Glob, Grep, Task

**Content to record in the plan file:**
```markdown
# Implementation Plan

## Design Doc Summary
- Summary from design-doc.md

## Impact Scope
- Target files for changes
- Affected existing features

## Implementation Steps Overview
- Step 1: ...
- Step 2: ...
```

#### 4. Self-review and exit Plan mode

Self-review the plan file and request user approval with ExitPlanMode.

**Tools**: ExitPlanMode

### [Test Case Design Phase] — Normal mode (approval gate)

#### 5. Investigate test infrastructure

Before designing test cases, investigate existing test infrastructure.

- Check for mocks, stubs, and test helpers related to the test target (e.g., under `test/`, `mock/`)
- Prioritize reusing existing infrastructure where possible
- If new infrastructure is needed (e.g., delay control, execution order observation), state the reason explicitly

**Tools**: Glob, Read

#### 6. Test case design

Design test cases. After creation, follow the **basic pattern** (self-review -> user review -> revision) to obtain user approval.

**Location**: `.local/docs/features/[name]/test-cases.md`

**Content:**
```markdown
# Test Case Design

## Test Case 1: [feature name] happy path
- Given: [preconditions]
- When: [action to perform]
- Then: [expected result]

## Test Case 2: [feature name] error path
- Given: [preconditions]
- When: [action to perform]
- Then: [expected error handling]
```

**Tools**: Write, Edit, AskUserQuestion

**⛔ Do not proceed without user approval**

### [Completion Phase]

#### 7. Update knowledge base

Save/update findings from investigation and implementation as files under `.local/docs/`.

**Scope**: Requirements, design decisions, alternatives and their rationale, technical constraints, etc.

**Tools**: Write, Edit

#### 8. Run /kaizen

**Tools**: Skill (kaizen)

## Self-Review Criteria

### Plan (implementation planning phase)
- Is the interpretation of the Design Doc correct?
- Are there overlooked areas in the impact scope?

### Test Cases
- Coverage of the feature
- Coverage of happy paths, error paths, and edge cases
- Is the test granularity appropriate?
- Does it follow the conciseness principles in `.claude/rules/writing-style.instructions.md`?

Adjust criteria based on content. Hear from the user when unclear.

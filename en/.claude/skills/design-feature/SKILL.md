---
name: design-feature
description: Starting from a user prompt, run hearing, design, and test design to create design.md / tests.md. No implementation. Resume from test design when design.md already exists.
allowed-tools: Read, Grep, Glob, Edit, Write, Task, AskUserQuestion
effort: high
---

# Design Feature

Starting from a user prompt, run investigation, hearing, design, and test design end-to-end with no approval gates.
Does not implement (use the `develop-feature` skill when implementation is needed).

## Prerequisites

- No approval gates. Once hearing has gathered the necessary information, proceed autonomously until design is complete
- User confirmation is limited to "missing requirements/direction" (no mid-process approval)
- Mark assumed items with ❓ (verification is deferred to the implementation phase in `develop-feature`)
- Deliverables are only `design.md` and `tests.md`. No implementation code
- For lightweight fixes or small new implementations (where a design doc is unnecessary), use the `fix` skill

## Output Locations

- **Feature-specific documents**: `.local/docs/features/[name]/`
  - `design.md` - Design Doc
  - `tests.md` - Test case design

## Entry Branching

Check `.local/docs/features/[name]/` for the target feature name and decide the starting phase:

| State | Starting phase |
|-------|----------------|
| No `design.md` | Phase [1. Investigation, hearing, and design] |
| `design.md` exists, no `tests.md` | Phase [2. Test design] (respect the existing design) |
| Both exist | Complete. Re-run the relevant phase only if updates are requested |

Ask the user if the feature name is unclear.

## Flow

### 1. Investigation, hearing, and design

#### 1.1. Codebase investigation

Investigate the existing codebase and documentation. Use Task (Explore agent) for broad exploration.

**Tools**: Read, Glob, Grep, Task

#### 1.2. Requirements hearing

Based on investigation results, efficiently hear about missing information. Conduct multiple rounds when needed.

Hearing principles and Good/Bad examples: see [../../guidelines/processes/hearing.md](../../guidelines/processes/hearing.md).

**Tools**: AskUserQuestion

#### 1.3. Create Design Doc

**Location**: `.local/docs/features/[name]/design.md`

**Content:**
```markdown
# Design Doc: [name]

## Background & Purpose
- What: What to build
- Why: Why it's needed

## Requirements
- Functional requirements
- Non-functional requirements and constraints

## Scope

## Technical Approach (rationale, alternatives)

## Design (architecture, processing flow)

## Related Code & References
- Target files and functions for changes
- Existing patterns to reference
- Related documents

## Implementation Details
- Interfaces/signatures only. Do not write method bodies
- Show code examples only for important algorithms and logic branches

## Considerations (security, etc.)
```

**Confidence marks**: see the confidence marks section of [../../guidelines/perspectives/documentation.md](../../guidelines/perspectives/documentation.md). Mark targets are **Requirements**, **Design**, **Implementation Details** sections.

**⛔ A Design Doc is NOT a place to write code**: rules on what is allowed/disallowed and Bad/Good examples are in the "Do Not Write Excessive Implementation Code in Design Docs" section of [../../guidelines/perspectives/documentation.md](../../guidelines/perspectives/documentation.md).

**Tools**: Write, Edit

#### 1.4. Design Doc self-review

Review `design.md` before moving to the next phase. If gaps are found, return to 1.3 and revise.

Criteria: [../../rules/self-review.instructions.md](../../rules/self-review.instructions.md) (targets: documentation, code design)

### 2. Test design

#### 2.1. Investigate test infrastructure

Investigate existing test infrastructure and identify what can be reused. Principles: see "Reuse of Existing Test Infrastructure" in [../../guidelines/perspectives/testing.md](../../guidelines/perspectives/testing.md).

**Tools**: Glob, Read

#### 2.2. Design test cases

Design test cases based on the Design Doc.

**Location**: `.local/docs/features/[name]/tests.md`

**Content:**
```markdown
# Test Case Design

## Test Case 1: [Behavior]
- Given: [preconditions]
- When: [action to perform]
- Then: [expected result]

## Test Case 2: [Behavior] error case
- Given: [preconditions]
- When: [action to perform]
- Then: [expected error handling]
```

**Tools**: Write, Edit

#### 2.3. Test case self-review

Review `tests.md` after creation. If gaps are found, return to 2.2 and revise.

Criteria: [../../rules/self-review.instructions.md](../../rules/self-review.instructions.md) (target: testing)

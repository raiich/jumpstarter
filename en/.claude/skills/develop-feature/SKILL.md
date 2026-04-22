---
name: develop-feature
description: Starting from a user prompt, run hearing, design, test design, implementation, and test execution end-to-end with no approval gates. Skip completed phases when design.md / tests.md already exist.
allowed-tools: Read, Grep, Glob, Edit, Write, Bash, Task, AskUserQuestion, TaskCreate, TaskUpdate, TaskList, TaskGet, Skill
effort: high
---

# Develop Feature

Starting from a user prompt, run investigation, hearing, design, test design, implementation, and test execution end-to-end with no approval gates.
The design and test design phases are delegated to the `design-feature` skill.

## Prerequisites

- No approval gates. Once hearing has gathered the necessary information, proceed autonomously to completion
- User confirmation is limited to "missing requirements/direction" and "verification of ❓-marked assumptions" (no mid-process approval)
- Deliverables are `design.md`, `tests.md`, implementation code, and test code
- If design inconsistencies surface during implementation, update `design.md` / `tests.md` before continuing
- When only design documents are needed, use the `design-feature` skill
- For lightweight fixes or small new implementations (where a design doc is unnecessary), use the `fix` skill

## Output Locations

- **Feature-specific documents**: `.local/docs/features/[name]/`
  - `design.md` - Design Doc
  - `tests.md` - Test case design

## Entry Branching

Check `.local/docs/features/[name]/` for the target feature name and decide the starting phase:

| State | Starting phase |
|-------|----------------|
| Either `design.md` or `tests.md` is missing | Phase [1. Design] |
| Both exist | Phase [2. Implementation and test execution] |

Ask the user if the feature name is unclear.

## Flow

### 1. Design

Call the `design-feature` skill to create `design.md` / `tests.md`.
Inside design-feature, the branching "no `design.md` → start from investigation/design" / "no `tests.md` → start from test design" is handled.

**Tools**: Skill

### 2. Implementation and test execution

#### 2.1. Create implementation task list

Create tasks based on the Design Doc and test cases.
If `design.md` contains ❓ (assumed) items, confirm them with the user before implementing the affected tasks.

**Task granularity**:
- 1 task = approximately 1 function, class, or method
- 1 task = approximately 1–5 test cases

**Tools**: TaskCreate

#### 2.2. TDD cycle

For each task, perform the following:

1. **Write test code** — implement tests from the cases in `tests.md`
2. **Implement the feature** — make tests pass
3. **Run tests** — use the test command per project settings
4. **Root cause analysis on failure** — [../../guidelines/processes/root-cause-analysis.md](../../guidelines/processes/root-cause-analysis.md)
5. **Refactor** — after tests pass, eliminate duplication, improve naming, simplify structure. Re-run tests to confirm no regressions
6. **Self-review** — review implementation and test code per [../../rules/self-review.instructions.md](../../rules/self-review.instructions.md). When revising, re-run tests to confirm no regressions
7. **Complete task** — TaskUpdate to completed, move to next task

**Tools**: TaskUpdate, Bash, Edit, Write

#### 2.3. If design inconsistency is discovered during implementation

- Move the affected task back to pending and update `design.md` / `tests.md`
- Briefly record the reason for the update (add confidence marks as appropriate)
- Return the task to in_progress and continue

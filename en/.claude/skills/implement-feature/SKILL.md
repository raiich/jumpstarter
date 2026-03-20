---
name: implement-feature
description: Implement features with TDD workflow based on a design doc and optional test cases.
allowed-tools: Read, Grep, Glob, Edit, Write, Task, AskUserQuestion, TaskCreate, TaskUpdate, TaskList, TaskGet, Skill
---

# Implement Feature

A skill for TDD implementation based on a Design Doc and test case design.

## Prerequisites

- Input is one of the following:
  - Deliverable of the `design-feature` skill: `design.md`
  - Deliverables of `design-feature` + `design-tests` skills: `design.md` + `tests.md`
  - Direct user input
- Run TDD cycles in small increments

## Output Locations

- **Knowledge base**: Topic-specific files under `.local/docs/`

## Task Granularity

- 1 task = approximately 1 function, class, or method
- 1 task = approximately 1-5 test cases

## Flow

### [Input Verification]

#### 1. Verify input

Read `design.md` if it exists. Also read `tests.md` if available.
If neither exists, understand requirements from user input.
Hear from the user if there are unclear points.

**Tools**: Read, AskUserQuestion

### [Implementation Phase] — Repeat per task

#### 2. Create implementation task list

Create tasks based on the Design Doc and test cases.
If the Design Doc contains ❓ (assumed) items, confirm them with the user before implementing the affected tasks.

**Tools**: TaskCreate

#### 3. Execute TDD cycle

**Tools**: TaskUpdate to set task to in_progress

For each task, perform the following:

**3.1. Write test code (basic pattern)**
- If `tests.md` exists: implement test code based on the test case design
- Otherwise: design and implement test cases based on the Design Doc or user input

**3.2. Implement feature (basic pattern)**
- Implement the feature to make the tests pass

**3.3. Run tests and handle failures**

**Tools**: Bash (use the test execution command according to project settings)

On failure, **analyze the root cause before attempting any fix**:

**Root cause analysis (always before fixing):**
1. Identify the direct cause of the error
2. Trace upstream — why did this happen? Is the symptom location the real cause?
3. Consider multiple fix candidates and choose the best approach:
   - Is this fix addressing the root cause, or just patching the symptom?
   - Could this fix affect other tests or features?
   - Does this require a design-level rethink?

**Act based on the analysis:**
- **Implementation bug** -> Fix at the cause location (not necessarily where the symptom appeared) and re-test
- **Test design mistake** -> Review the test case design and update tests.md
- **Design problem** -> Go back to design.md, rethink the design, then redo
- **Cannot determine** -> Hear from the user

**If a fix fails twice**: Do not repeat the same approach. Reconsider the root cause hypothesis or try a different approach. After 3 failures, consult the user

**3.4. Refactoring**
- After tests pass, clean up the implementation code (eliminate duplication, improve naming, simplify structure, etc.)
- Re-run tests after refactoring to confirm existing tests still pass

**3.5. Complete task**

**Tools**: TaskUpdate to set to completed, TaskList to check next task

**Review granularity**:
- Minor tasks: Batch review is acceptable
- Important tasks: Individual review

### [Completion Phase]

#### 4. Update knowledge base

Save/update findings from investigation and implementation as files under `.local/docs/`.

**Scope**: Requirements, design decisions, alternatives and their rationale, technical constraints, etc.

**Tools**: Write, Edit

#### 5. Run /kaizen

**Tools**: Skill (kaizen)

## Self-Review Criteria

### Test Code
- Consistency with test case design
- Coverage of happy paths, error paths, and edge cases
- Is the test granularity appropriate?
- Arrange-Act-Assert is clearly separated
- DAMP > DRY: Prioritize readability of what each test verifies over hiding details in helpers
- Add comments to supplement test names when the name alone is insufficient

### Implementation Code
- Is there over-engineering?
- Are there security concerns?

Adjust criteria based on content. Hear from the user when unclear.

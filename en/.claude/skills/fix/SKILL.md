---
name: fix
description: Starting from user input, hear then apply a fix or lightweight new implementation without approval gates. Does not produce design docs (design.md / tests.md).
allowed-tools: Read, Grep, Glob, Edit, Write, Bash, AskUserQuestion, TaskCreate, TaskUpdate, TaskList
---

# Fix

Apply lightweight fixes or new implementations after hearing, with no approval gates.
Does not produce design documents (design.md / tests.md).

## When to Use

- **fix**: fixes to existing code, localized additions, lightweight new implementations (where a design doc is unnecessary)
- **design-feature**: when only design documents (`design.md` / `tests.md`) are needed. No implementation
- **develop-feature**: when requirements need to be organized first, or the change spans multiple components and requires deliberate design and implementation end-to-end

## Principles

- Avoid band-aid or incremental fixes; decide the approach first, then change in one pass
- Principles (decide approach first, verify impact scope, pattern search): see [../../guidelines/processes/fix-in-one-pass.md](../../guidelines/processes/fix-in-one-pass.md)

## Flow

### 1. Understand input and hear

- Understand the user's instruction; estimate target and impact scope
- Efficiently hear on missing information ([../../guidelines/processes/hearing.md](../../guidelines/processes/hearing.md))

**Tools**: Read, Glob, Grep, AskUserQuestion

### 2. Comprehensive search (for fix-type tasks)

Grep for the same problem elsewhere (see "Pattern Search" in fix-in-one-pass.md).

**Tools**: Grep, Glob

### 3. Implement / Fix

- Fix-type: fix all found locations in one pass
- New implementation: implement only the minimum necessary code within the impact scope

**Tools**: Edit, Write

### 4. Run tests (if tests exist)

Use the test command per project settings. For failure root cause analysis, see [../../guidelines/processes/root-cause-analysis.md](../../guidelines/processes/root-cause-analysis.md).

**Tools**: Bash

### 5. Self-review

- [ ] Did you fix/implement every relevant location?
- [ ] Are the changes correct?
- [ ] Do they follow the writing rules (`.claude/rules/writing-style.instructions.md`) and the self-review perspectives (`.claude/rules/self-review.instructions.md`)?
- [ ] Do related docs/code also need updates?

### 6. Report

Briefly report: files touched, summary, test results (if executed).

# Task Execution Process Guidelines

## Purpose

Always create a plan before implementation or documentation, and execute step by step. Work with user approval to prevent rework.

## Core Principle

**Planning First**: Never implement directly; always go through a planning phase.

## Task Execution Flow

### Required Phases

1. **Planning Phase**: Create and save a plan to a file
2. **Review Phase**: Wait for user approval
3. **Implementation Phase**: Implement according to approved plan
4. **Verification Phase**: Conduct self-review
5. **Completion Phase**: Report to user

### Phase Transition Rules

**Important**: Do not proceed to the next phase without explicit user approval or instruction.

Always wait in the following cases:
- **Under Review**: "Please review", "Please check"
- **Under Consideration**: "Please consider", "Please propose"
- **Awaiting Fix**: After "Please fix", when there's no execution instruction
- **Confirmation Request**: Conditional instructions like "after ~", "before ~"

## Planning Process

### When Planning is Required

Always create a plan for the following tasks:

- Implementing new features
- Large-scale fixes/refactoring
- Changes spanning multiple files
- Creating new documentation or major revisions
- Architecture changes

### When Planning is Not Required

Plans are not needed for minor tasks:
- Typo fixes
- Single-line fixes
- Clear and simple changes (e.g., changing a variable name in one place)
- When user explicitly says "no planning needed"

### 5 Steps for Plan Creation

1. **Understand the Task**: Confirm what to achieve, constraints, and expected deliverables
2. **Create the Plan**: Document implementation details, files to change, execution order, anticipated issues
3. **Save to File**: Save to `plan.md` or under `.claude/`
4. **User Review**: Ask "May I proceed with this plan?"
5. **Start Implementation After Approval**: Implement only after user approval

## Pre-Execution Checklist

### 1. Confirm User Intent

- Is it under review?
- Have I received permission to execute?
- Did the user instruct "don't execute immediately"?

### 2. Environment Check

- Can it be built? (Check with `go build` etc. before execution)
- Are dependencies satisfied?
- Do necessary files exist?

### 3. Failure Plan

Before test execution:
- Have a strategy for handling failures
- Be prepared to confirm with user if needed

## Execution Timing

### When to Execute

- User explicitly instructed execution
- Review is complete and approved
- Build check succeeded

### When to Wait

- Under review
- User is considering
- Build errors remain
- User instructed "don't execute"

## Practical Examples

```
❌ Bad Example:
User: "Add a new feature"
Claude: [Starts implementation immediately] → [Completion report]
Problem: No plan, no user confirmation, rework occurs

✅ Good Example:
User: "Add a new feature"
Claude:
  1. [Understand task]
  2. [Create and save plan]
  3. "I've created a plan. May I proceed with this?"
User: "Approved"
Claude:
  4. [Implementation] → [Self-review] → [Completion report]
```

## Expected Benefits

- Reduce misalignment with user
- Prevent rework
- Improve work transparency
- Efficient task execution

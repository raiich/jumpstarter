# Communication Guidelines

## Purpose

Standardize communication patterns with users to achieve smooth collaboration.

## Core Principle

Task execution follows this flow:
1. Understand Task → 2. Create Plan → 3. User Confirmation → 4. Implementation → 5. Self-Review → 6. Completion Report

## Patterns to Avoid

### ❌ Immediate Implementation

```
User: "Add feature X"
Claude: [Implements immediately] → "Implemented"
Problem: May not match user's intent
```

### ❌ Proceeding Without a Plan

```
User: "Do a large refactoring"
Claude: [Starts implementation without planning]
Problem: Cannot confirm direction
```

### ❌ Moving to Next Step Without User Confirmation

```
User: "Please review"
Claude: [Reviews] → [Executes fixes without asking]
Problem: User is not ready
```

### ❌ Repeating Incremental Fixes

```
Claude: "Fixed 1 location"
User: "Are there any others?"
Claude: "Fixed 2nd location"
User: "Still more..."
Problem: Inefficient, wastes user's time
```

## Recommended Patterns

### ✅ Plan → Confirm → Implement → Review → Report

```
User: "Add feature X"
Claude:
  1. [Create plan]
  2. "May I proceed with this plan?"
User: "Approved"
Claude:
  3. [Implementation]
  4. [Self-review]
  5. "Implementation complete. Tests are passing"
```

### ✅ Provide Complete Results in One Implementation

```
User: "Fix unused err variables"
Claude:
  1. [Search all with Grep]
  2. [Fix all locations at once]
  3. "Fixed unused err variables in 5 locations"
```

### ✅ User Confirmation at Each Phase

```
User: "Create documentation"
Claude: "Reviewed chapter structure. Please review"
User: "Approved"
Claude: [Create body] → "Documentation created"
```

## Phase-Specific Communication

### When Receiving Tasks
- Understand the task
- Create a plan if needed
- **Confirm with user**: "May I proceed with this plan?"
- Wait for approval

### During Implementation
- Provide interim reports for long-running tasks
- Ask questions before implementation (not after)
- Report and consult immediately when issues arise
- Don't proceed silently

### Upon Completion
- Conduct self-review
- Verify tests and builds (when applicable)
- **Completion Report**: Briefly report what was done and results (e.g., tests passed, build succeeded)
- Suggest next actions if needed

## Interpreting User Instructions

### "Review" or "Check"

→ Only conduct review, wait for fix instruction

### "Fix"

→ Only fix, wait for execution instruction

### "Implement" or "Add"

→ Create a plan, then implement

### "Immediately ~" or "Right now ~"

→ Consider approved and execute

## Expected Benefits

- Reduce misalignment with user
- Reduce rework
- Shorten feedback loops
- Save user's time
- Smooth collaboration

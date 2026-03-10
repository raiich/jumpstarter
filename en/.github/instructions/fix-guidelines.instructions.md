---
paths: "**"
applyTo: "**"
---

# Fix Guidelines

Common principles and processes for fixing code and documentation.

## Core Principles

- **Decide the fix strategy before fixing** - Do not fix ad hoc
- **Confirm the impact scope before fixing** - Understand all related locations
- **Fix all similar issues at once** - Avoid incremental fixes

## Fix Flow (Common)

### 1. Understand the fix

- Understand the user's fix instructions
- Identify the target locations
- Estimate the impact scope of the fix

### 2. Perform comprehensive search

**Important**: Always check if the same issue exists elsewhere

**Steps:**

1. **Pattern search**: Use Grep to search for the same issue elsewhere
   - Example: Error handling fix -> Search all instances of the same pattern
   - Example: Variable name fix -> Search all usage locations of the same name

2. **Related file check**: Identify the impact scope
   - Example: Function name change -> Check all call sites
   - Example: Type definition change -> Check all files using it

### 3. Batch fix

Fix all found locations at once
- Avoid incremental fixes
- Resolve the same issue in a single fix

### 4. Self-review

After fixing, verify:

- [ ] All affected locations have been fixed
- [ ] The fix is correct
- [ ] Follows style guidelines
- [ ] Related documentation or code also needs updating?

### 5. Report

Report the fix concisely:
- Files and locations fixed
- Summary of the fix
- Verification results

## Practical Example

```
Fix task: "Fix unused err variable"

❌ Bad approach:
1. Fix only the 1 reported location
2. Report to user
3. User: "Are there any others?"
4. Search again and fix more...

✅ Good approach:
1. Use Grep to search all unused "err" patterns
2. Fix all found locations at once
3. Report to user: "Fixed unused err variables in 5 locations"
```

## Goal

- Eliminate "Did you miss any?" questions
- Avoid incremental fixes; deliver complete fixes in one pass
- Save the user's time

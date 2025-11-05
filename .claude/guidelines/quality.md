# Quality Assurance Guidelines

## Purpose

After creating or modifying code or documentation, always conduct self-review before presenting to users to ensure quality and completeness.

## Post-Implementation Checklist

### 1. Completeness Check

After implementation, ask yourself:

- [ ] Does it meet all requirements?
- [ ] Are there any omissions from user's instructions?
- [ ] Ask yourself "Are there any other omissions?"
- [ ] Have I confirmed there are no similar patterns elsewhere?

**Important**: Aim to solve all problems in one implementation. Complete with comprehensive review at once, rather than repeating incremental fixes.

#### Conducting Comprehensive Search

Always execute the following after fixes/implementation:

**1. Pattern Search**: Search with Grep for similar issues elsewhere
   - Example: Error handling fix → Search for all similar patterns
   - Example: Variable name fix → Search for all usage locations of same name

**2. Related File Check**: Identify scope of impact
   - Example: Function name change → Check all call sites
   - Example: Type definition change → Check all files using it

**3. Batch Fix**: Fix all found locations at once
   - Avoid incremental fixes
   - Resolve all same issues in one fix

**Goal**: Eliminate "Are there any other omissions?" questions from users

#### Practical Example

```
Fix Task: "Fix unused err variables"

❌ Bad Approach:
1. Fix 1 pointed out location
2. Report to user
3. User: "Are there any others?"
4. Re-search and fix more...

✅ Good Approach:
1. Search all unused "err" patterns with Grep
2. Fix all found locations at once
3. Report to user: "Fixed X locations"
```

### 2. Quality Check

Verify implementation quality:

- [ ] **Build Check**: Can the code compile/build?
- [ ] **Test Check**: Do all tests pass?
- [ ] **Consistency Check**: Do documentation and implementation match?
- [ ] **Error Handling**: Is error handling properly implemented?
- [ ] **Naming Conventions**: Does it follow project naming conventions?

### 3. Consistency Check

Verify consistency with entire codebase:

- [ ] Does it follow existing code style?
- [ ] Does it use same patterns as similar features?
- [ ] Are error handling patterns unified?
- [ ] Do naming conventions match other code?

## Response Process for Test Failures

### 1. Analyze Failure Cause

- Read error messages
- Identify which tests failed
- Understand root cause of failure

### 2. Re-plan Fix Strategy

Don't fix immediately, consider:
- Identify locations needing fixes
- Check if similar issues exist elsewhere
- Evaluate scope of impact from fixes

### 3. Confirm with User

Present fix strategy to user:
- Explain failure cause
- Present fix proposal
- Report scope of impact

### 4. Execute Fix

After user approval:
- Fix according to plan
- Fix all related locations at once
- Conduct self-review
- Re-test

## Important Principles

- Don't fix immediately, create a plan first
- Get confirmation from user
- Fix similar issues all at once

## Self-Review Timing

- When code implementation is complete
- When documentation is created
- When bug fix is performed
- When refactoring is done
- Before reporting results to user

## Practical Examples

### For Code Fixes

```
1. Implement fixes
2. Search for similar patterns elsewhere (use Grep)
3. Fix all applicable locations at once
4. Run build and tests
5. Report to user if no issues
```

### For Documentation Creation

```
1. Create documentation
2. Verify consistency with implementation
3. Verify conciseness (remove verbose parts)
4. Present to user
```

## Expected Benefits

- Reduce "Are there any other omissions?" questions from users
- Shorten feedback loops
- Provide more complete results in one implementation
- Save user's time

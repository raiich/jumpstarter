# Root Cause Analysis Process

A process for addressing problems (test failures, bugs, unexpected behavior) without applying band-aid fixes.

## Process

### 1. Identify the direct cause

From error messages, stack traces, and failing assertions, identify the direct cause.

### 2. Trace upstream

- Why did this happen? Is the symptom location really the cause?
- Search for the same problem elsewhere (see pattern search in [fix-in-one-pass.md](fix-in-one-pass.md))

### 3. Consider multiple fix candidates

- Is this fix addressing the root cause, or patching the symptom?
- Could this fix affect other tests or features?
- Does this require a design-level rethink?

### 4. Classify the cause and act

| Classification | Action |
|---|---|
| Implementation bug | Fix at the actual cause (not necessarily where the symptom appeared) and retest |
| Test design mistake | Review the test case design and update tests.md |
| Design problem | Go back to design.md, rethink the design, then redo |
| Cannot determine | Hear from the user |

## Escalation

- **After 2 failed fixes**: Do not repeat the same approach. Try a different hypothesis or approach
- **After 3 failures**: Consult the user

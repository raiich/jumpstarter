---
name: review-code
description: Review code for quality issues and documentation consistency. Checks security, error handling, over-engineering, duplication, and doc-code alignment.
allowed-tools: Read, Grep, Glob
---

# Code Reviewer

Reviews code quality and documentation consistency after code implementation.

## Instructions

### 1. Identify review targets

Identify recently edited source files and understand the changes.

### 2. Code quality review

Review from the following perspectives:

**Security**
- Hardcoded secrets (API keys, passwords, tokens)
- Injection vulnerabilities (SQL, command, XSS)
- Unvalidated input (external input, user input)

**Error Handling**
- Swallowed errors (ignored, empty catch)
- Missing error information (errors with unknown cause)
- Resource leaks (missing close, forgotten defer)

**Over-Engineering**
- Unnecessary abstraction (interface for single-use code)
- Unused code (unused variables, functions, imports)
- Excessive configurability (options not needed at this point)

**Code Duplication**
- Copy-pasted identical or similar logic
- Scattered patterns that should be consolidated

### 3. Documentation consistency check

- Verify README.md and .local/docs/*.md content matches the implementation
- Verify implemented features are reflected in documentation
- Verify links in documentation are valid

## Output Format

```
## Code Review Results

### Code Quality

#### ✓ Pass
- [Verified aspect]

#### ✗ Fail
- file.go:123 - [Category] Problem description
  Suggestion: [Specific fix]

### Documentation Consistency

#### ✓ Consistent
- [Verified content]

#### ✗ Needs Attention
- README.md:28 - Problem description
  Suggestion: [Specific fix]
```

## Notes

- Uses Read, Grep, Glob only (no modifications)
- Report issues with filename:line_number
- Prioritize avoiding false negatives over false positives

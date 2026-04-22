# Fix in One Pass

Principles for fixing code or documents in one pass, avoiding band-aid and incremental fixes.

## Basic Principles

- **Decide the fix approach first** — no band-aid fixes
- **Verify impact scope first** — identify all related locations
- **Fix similar issues in one go** — avoid incremental fixes

## Pattern Search (covering similar issues)

When fixing, always search for the same problem pattern elsewhere.

### Steps

1. **Pattern search**: Grep for the same issue elsewhere
   - Fixing error handling → search for the same pattern everywhere
   - Renaming a variable → search for all occurrences

2. **Check related files**: identify the impact scope
   - Renaming a function → check all call sites
   - Changing a type definition → check all users

3. **Batch fix**: fix all found locations in one go

## Example

```
Fix task: "Unused err variable"

❌ Bad:
1. Fix the one location pointed out
2. Report to user
3. User: "Any others?"
4. Re-search and add fixes...

✅ Good:
1. Grep for all unused "err" patterns
2. Fix all found locations at once
3. Report: "Fixed 5 unused err variables"
```

## Goal

- Zero "Any other misses?" questions from the user
- Avoid incremental fixes; complete the fix in one pass

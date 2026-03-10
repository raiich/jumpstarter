---
name: fix-doc
description: Properly fix documentation issues. Ensures consistency between documentation and code, and applies quality standards.
allowed-tools: Read, Grep, Glob, Edit, Write
---

# Documentation Fix Helper

A skill for properly fixing documentation issues. Follows the core principles in `.claude/rules/fix-guidelines.instructions.md` to perform complete and accurate fixes.

## Fix Flow

Follow the common flow in `.claude/rules/fix-guidelines.instructions.md`.

## Documentation-Specific Fix Guidelines

### Title & Heading Fixes

1. Check if the target title is used elsewhere
2. If anchors (`#anchor` format) change, update referring locations as well
3. Fix all at once

### Code Example Fixes

1. Verify code examples match the implementation
2. Search for the same code example in other documents
3. Check if reference materials and related documents need updating
4. Fix all at once

### Terminology & Wording Fixes

1. Search all occurrences of the target term
2. Determine if similar expressions should also be fixed for consistency
3. Fix all at once

### Feature Description Fixes

1. Verify the description matches the implementation
2. Check related documents (README, .local/docs/*.md, etc.)
3. Check if related descriptions need updating due to implementation changes
4. Fix all at once

## Self-Review (Documentation-Specific)

After fixing, verify the following:

- [ ] Have all affected locations been fixed?
- [ ] Is the fix correct?
- [ ] Does it follow `.claude/rules/writing-style.instructions.md`?
- [ ] Is there consistency between implementation and documentation?
- [ ] Are links valid (do referenced targets exist)?
- [ ] Are there any redundant parts?

## Conciseness Check

For each fix, verify:

- [ ] Contains only necessary information (no redundant explanations?)
- [ ] Avoids explaining assumed knowledge (readers may already know this)
- [ ] No duplicate information
- [ ] Not using long text to explain what code can show
- [ ] If an explanation feels long, considered whether it can be cut in half

## Documentation Fix Example

```markdown
❌ Bad example: Fix only 1 file
- Fix the term "new feature" in README.md

✅ Good example: Fix all related documents
1. Use Grep to search all occurrences of "new feature"
2. Fix README.md, .local/docs/feature-guide.md, .local/docs/api.md simultaneously
3. Verify code examples reflect the latest version
4. Check for broken links
```

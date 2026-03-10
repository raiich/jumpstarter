---
name: review-doc
description: Review all .md files for naming conventions, conciseness, consistency, and completeness. Use when creating or editing any markdown file.
allowed-tools: Read, Grep, Glob
---

# Documentation Reviewer

Reviews the quality of markdown documents.

## Review Criteria

Review based on `.claude/rules/writing-style.instructions.md` from the following perspectives:

- **Naming conventions**: File names use kebab-case
- **Conciseness**: Redundant explanations, duplicate information, long text replaceable by code
- **Structure**: Appropriateness of sections and hierarchy
- **Style**: Consistent tone, consistent terminology
- **Consistency**: Matches implementation, accurate code examples, valid links

## Output Format

```
## Documentation Review Results

### ✓ Pass
- [Verified aspect]

### ✗ Fail
- file.md:line_number - [Category] Problem description
  Suggestion: [Specific fix]
```

## Notes

- Uses Read, Grep, Glob only (no modifications)
- Report issues with filename:line_number
- Prioritize avoiding false negatives over false positives

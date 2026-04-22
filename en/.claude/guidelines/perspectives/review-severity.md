# Review Severity Levels

Classification of review findings by severity. Shared across code, design, and docs reviews.

## Levels

- **Critical**: must fix immediately
  - Code: runtime errors, data loss, security vulnerabilities
  - Design: fundamental problems (must resolve before implementation)
  - Docs: causes users to use the API incorrectly (wrong code examples, nonexistent APIs, incorrect procedures)
- **Major**: fix soon
  - Code: significant design issues, resource leaks
  - Design: important but has workarounds (recommended to address before implementation)
  - Docs: missing important information, vague explanations
- **Minor**: desirable improvement, no functional impact
  - Redundant explanations, structural improvements, can be addressed after implementation
- **Info**: reference information, suggestions, preferences, code style

## Output Format Example

```markdown
| # | Severity | Perspective | Finding | Location |
|---|----------|-------------|---------|----------|
| 1 | Critical | API design | Return value can be null but undocumented | foo.ts:42 |
```

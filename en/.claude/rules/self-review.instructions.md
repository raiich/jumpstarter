# Self-Review

Perform this review after writing or modifying documentation or code.

## Perspectives

Apply according to the target. Check specific criteria at the linked files.

- **Documentation**: [documentation.md](../guidelines/perspectives/documentation.md) — accuracy, completeness, consistency, readability, confidence marks
- **Code design**: [design-and-coding.md](../guidelines/perspectives/design-and-coding.md) — package structure, file granularity, API design, project-specific patterns
- **Testing**: [testing.md](../guidelines/perspectives/testing.md) — coverage, test quality, false positives
- **Severity**: [review-severity.md](../guidelines/perspectives/review-severity.md) — Critical / Major / Minor / Info

## Common checks

Always verify regardless of target.

- **Deletability**: Can it be removed without losing understanding? Is it explaining known information? Is it duplicated? Can it be halved?
- **Critical lens**: Do not settle for "looks good" — actively look for weaknesses and risks. Verify means-ends alignment
- **Impact analysis**: Grep for references to changed elements (function names, file paths, config keys, etc.) and verify consistency

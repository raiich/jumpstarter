# Evaluation Matrix (means-validity assessment)

Matrix for evaluating external sources (articles, best practices, others' proposals) against the user's purpose.

## Axes

- **Relevance to purpose**: Direct / Related / Weak
- **Consistency with official / existing**: Aligned / Partially / Conflicts / N/A
- **Current coverage**: Covered / Partial / None

## Example Table

| Item | Relevance | Official consistency | Current state | Judgment |
|---|---|---|---|---|
| Item A | Direct | Aligned | Covered in rules/xxx.md | Already covered |
| Item B | Related | Aligned | Partial in some skills | Partial |
| Item C | Weak | N/A | None | Out of scope |
| Item D | Direct | Conflicts | — | Not recommended |

## Judgment Priority

1. **Relevance to purpose** (weak items are "out of scope" even if uncovered)
2. **Consistency with official / existing**
3. **Current coverage**

## Handling Mismatches

When an issue directly relates to the purpose but the candidate does not cover it / is not optimal:

- Propose a better means derived from official docs or the existing setup
- Consider adapting the original item rather than rejecting it wholesale

Proposal format:

```
Purpose: improvement of XX
Original approach: △△ (reason for mismatch)
Proposal: □□ (why this is better)
```

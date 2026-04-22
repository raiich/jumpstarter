---
name: review
description: Review code, design docs, or user-facing documentation from multiple independent perspectives. Enumerate perspectives first, then systematically review each one.
allowed-tools: Read, Grep, Glob, Bash, Task, AskUserQuestion, Agent
effort: high
---

# Review

A skill for reviewing code, design docs, and user-facing documentation from independent perspectives.

## Prerequisites

- **Enumerate perspectives first, then review each one without gaps** (avoid "oh, there was also this perspective" after the fact)
- Surface all findings in the initial review. Producing new findings in additional passes is a failure

## Flow

### 1. Identify the Review Target and Its Type

Identify the review target and its type (`code` / `design` / `docs`).

**Type classification:**

| Type | Examples | Typical paths |
|---|---|---|
| `code` | Implementation code, test code | `src/**`, `**/*.test.*`, etc. |
| `design` | Design Doc, test case design | `.local/docs/features/design.md`, `.local/docs/features/tests.md` |
| `docs` | User-facing documentation | `README.md`, `docs/**/*.md` |

Classification logic:
- If the type is explicitly passed as an argument (`code` / `design` / `docs`), follow it
- If no argument is given but the path uniquely implies a type, follow the inference
- If ambiguous, confirm with the user via `AskUserQuestion`
- If no target is specified, infer from `git diff --name-only` and confirm with the user

Loading related files:
- For `code`, load `.local/docs/features/design.md` / `tests.md` if they exist, to verify alignment with the design
- For `design`, also grasp the related codebase structure
- For `docs`, grasp the code that the document references (public APIs, config files, etc.)

**Tools**: Read, Glob, Grep, Bash

### 2. Enumerate Perspectives

Based on the type, determine which perspectives to reference and autonomously select which to apply (do not ask the user in advance). Report the selection as part of the final result.

| Type | Referenced perspectives |
|---|---|
| `code` | [design-and-coding.md](../../guidelines/perspectives/design-and-coding.md), [testing.md](../../guidelines/perspectives/testing.md) |
| `design` | [design-and-coding.md](../../guidelines/perspectives/design-and-coding.md), [documentation.md](../../guidelines/perspectives/documentation.md), [testing.md](../../guidelines/perspectives/testing.md) |
| `docs` | [documentation.md](../../guidelines/perspectives/documentation.md) |

Respect any scope explicitly given in the arguments or the user's initial instructions. Examples:
- "Don't go deep on algorithms" → exclude algorithm correctness
- "Edge cases are handled in tests" → limit boundary conditions to the test side
- "Sensory checks will come later" → exclude UX/appearance perspectives

### 3. Execute Multi-Perspective Review

Select the subagent composition based on the type and run them in parallel.

| Type | Subagent composition | Perspectives per agent |
|---|---|---|
| `code` | `critical-analyst` + `security-reviewer` + `codebase-explorer` | Structural/design issues / security and robustness / test code quality |
| `design` | `critical-analyst` + `constructive-analyst` + `security-reviewer` | Weaknesses and risks / amplifying strengths and opportunities / security |
| `docs` | `critical-analyst` (explicitly instructed to cross-check with code) | Whether the document's claims align with the code |

Information passed to each subagent:
- The list of review target file paths
- The applicable perspective list (finalized in step 2)
- Related documents (`design.md` / `tests.md`, etc., if they exist)

**Type-specific notes:**

- `code`: `security-reviewer` covers input validation, resource leaks, boundary conditions, and type safety. `codebase-explorer` covers false positives in test code, circular logic, missing arguments, and assertion coverage
- `design`: Perform the review from a perspective independent of the author's self-review
- `docs`: Verify "does it work as written?" by cross-checking with code. Do not judge as "looks correct" by guessing. Actually read the code for code examples, API descriptions, and configuration values

**Tools**: Agent, Read, Grep, Glob, Bash

### 4. Consolidate Review Results

Consolidate subagent results and present in the following format. State the **reviewed scope (target files, applied perspectives, excluded perspectives with rationale)** at the top so the user can see exactly what was and wasn't covered.

```markdown
## Review Results

### Review Target
- Files: [list of reviewed files]
- Type: [code / design / docs]

### Perspectives
- Applied: [list]
- Excluded: [list, with rationale (out of scope, per instruction, etc.)]

### Findings

| # | Severity | Perspective | Finding | Location |
|---|----------|-------------|---------|----------|
| 1 | Critical | ... | ... | file:line |

### Summary
[Overall assessment and recommended actions]
```

**Severity criteria**: [../../guidelines/perspectives/review-severity.md](../../guidelines/perspectives/review-severity.md)

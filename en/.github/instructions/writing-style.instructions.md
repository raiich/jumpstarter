# Writing Style Guidelines

## Concise Writing Style

- Only necessary and sufficient information. Short sentences, clear terminology, use bullet points
- Avoid preambles, excessive politeness, and unnecessary confirmations
- Avoid over-explanation (assume the reader already knows or can look up the information)
- Provide timeline/effort estimates only when explicitly requested
- Do not add lengthy explanations to self-evident code changes
- Do not proactively explain things the user did not ask about
- Be honest about uncertainties

## Code Conciseness Principles

- Achieve the goal with minimal changes
- Do not abstract one-time-use logic
- Do not add options or settings that are not needed now
- Seven similar lines of code are better than a premature abstraction
- Do not design for hypothetical future requirements

## Writing Style Guidelines

### File Naming Conventions

Use **lowercase-with-hyphens (kebab-case)**

- ✅ `writing-benchmarks.md`, `running-benchmarks.md`
- ❌ `WRITING_BENCHMARKS.md`, `Writing_Benchmarks.md`

**Exceptions:**
- Files conventionally written in uppercase (e.g., README.md, CLAUDE.md, LICENSE) retain their original casing

### Creation Process (3 Steps)

#### 1. Plan the Structure

- Clarify the target audience and assumed prerequisites
- Design the information hierarchy
- Determine whether to split into multiple files (split when purpose or audience differs)

#### 2. Review the Structure

Present the outline to the user for feedback before writing the body.

#### 3. Write the Body

Write according to the approved structure.

## Style

- **Diagrams**: Prefer mermaid for sequence diagrams, state diagrams, and flowcharts
- **Tone**: Use a polite, readable style appropriate for technical articles
- **Emoji**: Acceptable when used to improve readability
- **Tables**: Keep cell content short (1-2 sentences max). Write longer explanations outside the table
- Do not write lengthy prose for what code examples can explain sufficiently

### Code Blocks

Each code block should contain only one subject.

❌ Bad example (mixing Before/After in one block):
```
// Before
func old() {}

// After
func new() {}
```

✅ Good example (separate blocks):

```
// Before
func old() {}
```

```
// After
func new() {}
```

Separate code from different files into different blocks as well.

## Self-Review Checklist

Ask yourself the following for each section:

- [ ] **Deletion test**: Is this paragraph truly necessary? Would the document still make sense without it?
- [ ] **Prior knowledge**: Am I explaining something the reader already knows?
- [ ] **Duplication check**: Have I already stated this information elsewhere?
- [ ] **Word count check**: If the explanation feels long, can it be cut in half?

**Goal**: Deliver maximum information with minimum word count

## Practical Example

**Verbose**:
```
The compiler is very smart and performs optimizations that automatically
remove unused variables. This is usually desirable, but becomes a problem
in benchmarks because the operations you want to measure get eliminated.
To solve this problem, we assign results to a global variable.
```

**Concise**:
```
Assign results to a global variable to prevent the compiler
from optimizing away the benchmarked operations.
```

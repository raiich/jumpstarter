# Documentation Creation Guidelines

## Scope

This guideline applies to all files with `.md` extension, including those under `.claude/`.

## Purpose

Create concise and necessary documentation while eliminating verbosity.

## File Naming Convention

Use **lowercase-with-hyphens (kebab-case)**

- ✅ `writing-benchmarks.md`, `running-benchmarks.md`
- ❌ `WRITING_BENCHMARKS.md`, `Writing_Benchmarks.md`

**Exceptions:**
- Important files conventionally written in uppercase (e.g., README.md, CLAUDE.md, LICENSE) remain as-is

## Creation Process (3 Steps)

### 1. Consider Chapter Structure

- Clarify target audience and prerequisite knowledge
- Design information hierarchy
- Decide if file splitting is needed (split when purpose or audience differs)

### 2. Review Chapter Structure

Present structure to user for feedback before writing body.

### 3. Write Body

Write according to approved structure.

## Principle of Conciseness

### Things to Avoid

- Excessive explanation (overly detailed basic concepts, information readers already know)
- Duplicate information

### Things to Aim For

- Only necessary and sufficient information
- Short sentences, clear terminology, use of bullet points
- Balance of code and comments

## Self-Review Checklist After Documentation Creation

Ask yourself for each section:

- [ ] **Deletion Test**: Is this paragraph truly necessary? Can it still be understood if deleted?
- [ ] **Prerequisite Knowledge**: Am I explaining content readers already know?
- [ ] **Duplication Check**: Have I already mentioned the same information elsewhere?
- [ ] **Code First**: Am I writing long text for what can be sufficiently explained with code examples?
- [ ] **Character Count Check**: If explanation feels long, consider if it can be cut in half

**Goal**: Provide maximum information with minimum character count for the content

## Practical Examples

**Verbose Example**:
```
The compiler is very smart and performs optimizations that automatically
remove unused variables. This is usually desirable, but in benchmarks it
becomes a problem because the processing we want to measure gets removed.
To solve this problem, we use assignment to global variables.
```

**Concise Example**:
```
To prevent the benchmark target from being removed by optimization,
assign the result to a global variable.
```

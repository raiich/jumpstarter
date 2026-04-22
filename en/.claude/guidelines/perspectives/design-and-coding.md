# Design Perspectives

Shared quality perspectives for design work (Design Doc authoring, implementation, review). Covers requirements, API design, module structure, and implementation-level correctness.

## Requirements & Scope

- Are What / Why clear?
- Are requirements in a testable form ("when X, then Y")?
- Is the scope boundary clear (included / excluded)?
- Any contradictions between requirements?
- Consideration of non-functional requirements (performance, compatibility, security)?

## Technical Approach

- Have alternatives been sufficiently considered?
- Is the rationale for the choice clear?
- Feasibility and risks?
- **Does it align with existing codebase patterns?** (Before introducing a bespoke solution, check whether an existing factory / helper / middleware can be reused)

## API / Interface Design (consumer's view)

- Is the interface intuitive?
- Are type definitions clear to consumers (required / optional distinction, naming)?
- Is naming consistent (same pattern, same naming rule)?
- Is the public / internal boundary clear (no leakage of internals)?
- Are error messages and error-handling policy useful to consumers?
- Are default values reasonable?
- Is customizability neither excessive nor insufficient?
- Are doc comments attached to public methods and types (at implementation time)?

## Modularity

- Cohesion: Does each module have a single responsibility?
- Coupling: Do functions / methods take only the minimum necessary parameters?
- Information hiding: Are implementation details encapsulated?
- Direction of dependencies: One-directional, no cycles?
- Replaceability: Can components be swapped?
- Change propagation: Is the blast radius of a single change reasonable?

## Package Structure

- Are directories split by technical role (models/, views/, controllers/, machines/, etc.)?
- Are files grouped under reasons like "same state machine" or "same layer"?
- Is each feature directory self-contained (no direct references to other features)?

**How to check**: Run `git diff --name-only` to list changed files and inspect the directory structure. Pay special attention to newly created directories.

## File Granularity

- Excessive splitting into "one class = one file"?
- Are related code (types, events, constants, State, etc.) cohered into a single file?
- Can the code flow be read top-to-bottom?

**How to check**: Review line count and content of newly created files. Consider consolidation when multiple files with 50 lines or fewer are created.

## Resource Management

- Are resources (files, connections, timers, etc.) paired between creation and release?
- Are dispose / cleanup covered across all code paths?
- Are externally injected resources not closed unilaterally?
- Are cache size limits and eviction controlled?

## Correctness

- Are there unintended dependencies on external state (randomness, time, etc.)?
- Are numeric boundary conditions handled (overflow, floating point error, divide-by-zero)?
- Is consistency maintained across all state transition paths?
- Are there magic numbers?

## Type Safety

- Does code bypass type safety with `as` casts or `any`?
- Are union types properly discriminated at the consumer side?
- Possibility of missing arguments or type mismatch at runtime?

## Alignment with Design (when design.md exists)

- Are the design.md requirements implemented?
- Are features not in design.md added?
- Does the interface match the design.md definition?

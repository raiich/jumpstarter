# Testing Perspectives

Shared perspectives for test case design, test code authoring, and test review.

## Coverage

- Any gaps between design.md requirements and test cases?
- Balance across normal, error, and edge cases?
- Are important user scenarios covered?
- Do exported functions / methods have tests?

## Test Quality

- **Does it test behavior?**: "User action → result" is verified. Asserting constants, array sizes, enum counts, or property values is not a test
- **Bug detector or change detector?**: If a correct refactor of production code breaks the test, the test is coupled to the implementation
- **Is it just restating constants?**: If changing a production value only requires mechanically changing the test to the same value, the test has no value
- **Are tests derived from functional requirements?**: Derive from the design.md requirements ("when X, then Y"), not from implementation details (type definitions, module structure)
- **Test through the public API?**: Do not test internal functions or internal state directly. Write from the consumer's perspective
- **Do tests document usage?**: Reading tests should show how to use the module
- **Are mocks kept simple?**: Growing mocks signals the test target has too much responsibility. Revisit the test or production design

## False Positive Prevention

- Is it guaranteed that the target of in-loop assertions is non-empty?
- Are there patterns where assertions are skipped due to conditional branches?
- Is there circular logic deriving expected values from the output under test?

## Assertion Validity

- Are expected values derived from an independent basis?
- Is the tolerance appropriate (not so wide that the test is meaningless)?
- Are there cases that check only the upper or lower bound?

## Argument Consistency

- Do function calls pass all required arguments?
- Even with TypeScript type checking passing, could an argument become `undefined` at runtime?

## Test Structure

- Are tests independent (no dependency on shared state or execution order)?
- Is Arrange-Act-Assert clearly separated?
- Are mocks set up / cleaned up appropriately?
- Do test names describe the behavior?
- DAMP > DRY: Prefer readability within each test over heavy helper abstraction

## Reuse of Existing Test Infrastructure

Before designing or implementing test cases, survey existing test infrastructure for reuse.

**Survey targets**
- Mocks, stubs, and test helpers related to the subject under test (under `test/`, `mock/` etc.)
- setUp / tearDown utilities
- Fixtures and test data

**Principles**
- Reuse what exists first
- Reuse production factories and helpers (do not reimplement them in test-helpers)
- New test helpers should contain only test-specific code (time control, mocks)
- When new infrastructure is genuinely needed (e.g., delay control, observation of execution order), state the reason explicitly

**How to check**: Inspect imports of test files. Grep for production-code-like logic in test-helpers.

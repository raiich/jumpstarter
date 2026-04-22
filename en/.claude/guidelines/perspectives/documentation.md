# Documentation Perspectives

Shared quality perspectives for all documents (README, docs/, Design Doc, plan files, etc.).

## Accuracy (alignment with code)

Do the document's claims match the actual code?

- Do referenced class / function / type names exist and are they exported?
- Do referenced method signatures (argument types, count, order, return type) match the implementation?
- Do referenced default values match the implementation?
- Do behavior descriptions match the implementation?
- Do referenced file paths and directory structures match reality?
- Do version numbers and dependencies match package.json?

## Correctness of Code Examples

- Are import paths correct (paths exposed in the `exports` field)?
- Does the API used actually exist (no references to non-existent methods or properties)?
- Are argument types and counts correct?
- Are code examples internally consistent (not showing the same operation in conflicting ways)?
- Do code examples reflect the latest API (no leftover old APIs after refactoring)?

## Coverage (user-facing documentation)

- **Setup**: installation steps, prerequisites, runtime environment
- **Basic usage**: minimal working example (copy-paste runnable)
- **Public API**: explanation of exported classes, functions, and types
- **Configuration / customization**: options, default values, how to customize
- **Constraints / known issues**: limitations the user should know

## Structure & Navigation

- Is the reading order natural (overview → setup → basic use → detailed → reference)?
- Are heading granularities appropriate (neither too deep nor too shallow)?
- Are cross-document links and references correct (no broken links)?
- If there is a TOC, does it match the content?

## Readability

- Does the explanation level match the target audience (assumed background knowledge is appropriate)?
- Any redundant explanations (repeating the same thing)?
- Are there technical terms that need explanation?
- Are tables, bullet lists, and code examples used appropriately?

## Consistency

- Terminology unified (no multiple names for the same concept)?
- Notation unified (uppercase / lowercase, hyphen / camel, etc.)?
- Code example style unified?
- Writing style unified?

## Diagrams

- Prefer Mermaid for sequence diagrams, state diagrams, and flowcharts
- Not expressing things as prose that would be clearer as a diagram

## Confidence Marks (✅ / ⚠️ / ❓)

Annotate requirement and design items with their source reliability.

| Mark | Meaning | Condition |
|---|---|---|
| ✅ | Confirmed | Directly verified from user statements, code, or documentation |
| ⚠️ | Inferred | Reasonable inference from confirmed information |
| ❓ | Assumed | Assumption without a source. Requires user confirmation before implementation |

**Example**

```markdown
## Requirements
- ✅ Authentication uses the existing JWT middleware (confirmed at auth.go:L42)
- ⚠️ Token expiry is 24 hours (inferred from current config)
- ❓ Whether refresh tokens are needed is unconfirmed
```

**Where to apply**

- Requirements / Design / Implementation Details sections of a Design Doc
- Preconditions / missing-info sections of plan files
- Supporting rationale for review findings

**Not needed for**

- Background & Purpose, Scope, Related Code references

**Note**

- If ❓ (assumed) items are dominant, consider resolving them via hearing

## Do Not Write Excessive Implementation Code in Design Docs

The purpose of a Design Doc is to communicate design decisions — not to write implementation code. Only minimal code fragments that supplement design decisions are permitted.

### Basic Checks

- Do total code block lines exceed 1/3 of the document?
- Can each code block be removed while still conveying the design intent? (If so, remove it)
- Are complete class implementations, function bodies, or initialization code present?
- Is the design being expressed entirely through code where text or diagrams would suffice?
- Does each code example focus on a single design decision (no mixed concerns)?

### Allowed / Disallowed Code

- **Allowed**: Type definitions / interfaces, data structures, non-obvious branching logic (a few lines)
- **Disallowed**: Complete class implementations, function bodies, initialization, boilerplate, utility functions

**Judgment criterion**: If removing a code block still conveys the design intent, that code is unnecessary.

**Quantitative guideline**: If total code block lines exceed 1/3 of the entire document, you've written too much.

### Bad Example (inappropriate for a design doc)

```typescript
// ❌ Writing a complete class implementation
class MovingState implements State<GameData> {
  name() { return "Moving"; }
  entry(machine: EntryMachine<GameData>, event: object): void {
    const data = machine.value();
    if (event instanceof MoveEvent) {
      data.moveDirection = event.direction;
      data.facing = event.direction;
    }
    const tick = (m: AfterFuncMachine<GameData>): void => {
      const d = m.value();
      const dx = d.moveDirection === "right" ? SPEED : -SPEED;
      d.playerX = Math.max(0, Math.min(d.playerX + dx, d.sceneWidthPx - W));
      updateCamera(d);
      checkSceneTransition(m, d);
      m.afterFunc(d.dispatcher, TICK_MS, tick);
    };
    machine.afterFunc(data.dispatcher, TICK_MS, tick);
  }
}
```

### Good Example (communicates only the design decision)

> Moving state achieves pixel-level movement via afterFunc chain (16ms interval).
> Self-transition auto-cancels the previous chain, making direction changes and stops safe.

```typescript
// Detect scene transition zones within the movement tick → StopMoveEvent + EnterTownEvent
machine.afterFunc(dispatcher, TICK_MS, tick);
```

### Other Notes

- One code example = one design decision (do not mix multiple concerns)
- Omit obvious logic with comments like `// ...validation...`
- Use text, tables, or Mermaid diagrams when they work better than code

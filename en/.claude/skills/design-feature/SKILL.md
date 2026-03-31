---
name: design-feature
description: Design features through codebase investigation and user hearing, producing a design doc.
allowed-tools: Read, Grep, Glob, Edit, Write, Task, AskUserQuestion, EnterPlanMode, ExitPlanMode, Skill
effort: high
---

# Design Feature

A skill for creating a Design Doc through codebase investigation and user hearing.

## Prerequisites

- Users may not provide all requirements upfront; proactively conduct hearings
- Investigate the codebase and documentation first, then ask efficient questions based on that understanding
- The deliverable is saved as `design.md` and serves as input for the `implement-feature` skill
- Each requirement and design item is annotated with a **confidence mark** to indicate source reliability

## Output Locations

- **Feature-specific documents**: `.local/docs/features/[name]/`
  - `design.md` - Design Doc (primary deliverable of this skill)

## Flow

### [Investigation & Hearing Phase] — Plan mode

#### 1. User: Provides feature overview and direction

#### 2. Enter Plan mode

**Tools**: EnterPlanMode

#### 3. Codebase investigation

Investigate the existing codebase and documentation and record findings in the plan file.
Use Task (Explore agent) when broad exploration is needed.

**Tools**: Read, Glob, Grep, Task

#### 4. Requirements hearing

Based on investigation results, hear from the user about missing information.
Reflect hearing results in the plan file. Conduct multiple rounds of hearing as needed.

**Principles for efficient hearings:**
- Do not ask about things that can be understood by reading the code
- Present specific options based on constraints and patterns identified during investigation
- Proactively confirm aspects the user has not mentioned (edge cases, consistency with existing features, etc.)

**Example: Efficient hearing**

User input: "I want to add a notification feature"

❌ Bad question: "What kind of notification feature?" (too broad)

✅ Good questions (with options based on investigation results):
- "Do you envision event-driven using the existing EventBus (events/bus.go), or direct invocation?"
- "Are there notification targets beyond Slack webhook (slack_url already exists in config)?"

**Tools**: AskUserQuestion

**Content to record in the plan file:**

⚠️ **Important**: Always include a "Deliverable" section at the beginning of the plan file.
This ensures the deliverable type can be correctly identified even if context is lost after exiting plan mode.

```markdown
# Investigation & Hearing Results

## Deliverable
- Type: Design Doc (markdown document)
- Location: .local/docs/features/[name]/design.md
- ⚠️ No code implementation (that is the responsibility of implement-feature)

## Related Existing Features
- Feature A: path/to/file

## Important Architecture Patterns
- Pattern 1: Description

## Hearing Results
- Confirmed requirements and constraints

## Technical Approach Candidates
- Approach A: Overview with pros and cons
- Approach B: Overview with pros and cons
```

#### 5. Self-review and exit Plan mode

Self-review the plan file and request user approval with ExitPlanMode.

**Tools**: ExitPlanMode

### [Design Doc Creation Phase] — Normal mode (approval gate)

#### 6. Create Design Doc

Create the Design Doc based on investigation and hearing results from the plan file.
After creation, follow the **basic pattern** (self-review -> user review -> revision) to obtain user approval.

**Location**: `.local/docs/features/[name]/design.md`

**Content:**
```markdown
# Design Doc: [name]

## Background & Purpose
- What: What to build
- Why: Why it's needed

## Requirements
- Functional requirements
- Non-functional requirements and constraints

## Scope

## Technical Approach (rationale, alternatives)

## Design (architecture, processing flow)

## Related Code & References
- Target files and functions for changes
- Existing patterns to reference
- Related documents

## Implementation Details
- Interfaces/signatures only. Do not write method bodies
- Show code examples only for important algorithms and logic branches

## Considerations (security, etc.)
```

**Confidence marks:**

Annotate each item in the requirements, design, and implementation details sections with a mark indicating source reliability.

- ✅ **Confirmed** — Directly verified from user statements, code, or documentation
- ⚠️ **Inferred** — Reasonable inference from confirmed information
- ❓ **Assumed** — Assumption without a source. Needs user confirmation before implementation

```markdown
## Requirements (example)
- ✅ Authentication uses the existing JWT middleware (confirmed at auth.go:L42)
- ⚠️ Token expiry is 24 hours (inferred from current config)
- ❓ Whether refresh tokens are needed is unconfirmed
```

Mark target sections: **Requirements**, **Design**, **Implementation Details**.
Not needed for: Background & Purpose, Scope, Related Code & References.

**⛔ A Design Doc is NOT a place to write code:**

The purpose of a Design Doc is to communicate design decisions — not to write implementation code.
Only minimal code fragments that supplement design decisions are permitted.

- **Allowed code**: Type definitions/interfaces, data structures, non-obvious branching logic (a few lines)
- **Disallowed code**: Complete class implementations, function bodies, initialization, boilerplate, utility functions

**Judgment criterion**: If removing a code block still conveys the design intent, that code is unnecessary.

**Quantitative guideline**: If total code block lines exceed 1/3 of the entire document, you've written too much.

**Bad (inappropriate for a design doc)**:

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

**Good (communicates only the design decision)**:

> Moving state achieves pixel-level movement via afterFunc chain (16ms interval).
> Self-transition auto-cancels the previous chain, making direction changes and stops safe.

```typescript
// Detect scene transition zones within the movement tick → StopMoveEvent + EnterTownEvent
machine.afterFunc(dispatcher, TICK_MS, tick);
```

**Other notes:**
- One code example = one design decision (do not mix multiple concerns)
- Omit obvious logic with comments like `// ...validation...`
- Use text, tables, or Mermaid diagrams when they work better than code
- Prefer Mermaid format for diagrams

**Tools**: Write, Edit, AskUserQuestion

**⛔ Do not proceed without user approval**

### [Completion Phase]

#### 7. Run /kaizen

**Tools**: Skill (kaizen)

## Self-Review Criteria

### Plan (investigation phase)
- Are there unidentified questions that should be heard?

### Design Doc
- Are What/Why clear?
- Are there any ambiguities remaining in the requirements?
- Are functional requirements in a testable format (clear operation and result, e.g., "when X, then Y")?
- Is the investigation of related code sufficient?
- Is the scope clear?
- Have alternatives been sufficiently considered?
- **⛔ Code volume check**: Do total code block lines exceed 1/3 of the entire document?
- **⛔ Code necessity check**: Can each code block be removed and the design intent still be understood? If so, remove it
- **⛔ Implementation leak check**: Are there complete class implementations, function bodies, or initialization code?
- Is design being expressed entirely through code when text or diagrams would suffice?
- Does each code example focus on a single design decision (no mixed concerns)?
- Does it follow the conciseness principles in `.claude/rules/writing-style.instructions.md`?
- Confidence marks: Can any ❓ (assumed) items be resolved through additional hearing?

Adjust criteria based on content. Hear from the user when unclear.

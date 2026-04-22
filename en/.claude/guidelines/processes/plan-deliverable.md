# Plan File Header Template: Deliverable Section

Section placed at the top of Plan mode plan files.

## Purpose

Ensures the deliverable type is identifiable even if context is lost after exiting Plan mode.

## Template

```markdown
## Deliverable
- Type: [deliverable type] (e.g., Design Doc, OKR, Test Case Design)
- Location: [absolute or relative path]
- ⚠️ [notes / responsibilities out of scope]
```

## Examples

### Design Doc

```markdown
## Deliverable
- Type: Design Doc (markdown document)
- Location: .local/docs/features/[name]/design.md
- ⚠️ No code implementation at this stage
```

### OKR

```markdown
## Deliverable
- Type: OKR document (markdown)
- Location: .local/docs/okr/{year}.md
```

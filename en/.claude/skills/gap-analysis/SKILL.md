---
name: gap-analysis
description: Analyze gaps between the codebase and .claude/ configuration, and propose missing rules, skills, and agents
allowed-tools: Read, Grep, Glob, Edit, Write, AskUserQuestion
---

# Gap Analysis

Cross-references codebase characteristics with the current `.claude/` configuration to identify uncovered areas and propose improvements.

A reverse-proposal skill where AI proactively finds "what's missing" without relying on user feedback.

## Steps

### 1. Scan Codebase Characteristics

Briefly survey the following (don't go too deep):

- Languages and frameworks
- Directory structure characteristics
- Test setup (test framework, test file placement)
- CI/CD presence and configuration
- Major external dependencies
- Documentation structure (README, docs/, design documents)
- Makefile / task runners / automation scripts

### 2. Review Current .claude/ Configuration

- `rules/` (actual files in `.github/instructions/`) contents
- `skills/` list and each skill's purpose
- `agents/` list and each agent's role
- `settings.local.json` permissions

### 3. Identify Gaps

Look for uncovered areas against codebase characteristics from these angles:

- **Missing rules**: Patterns exist in the codebase but no rule governs them
  - Example: Test code exists but no test conventions defined
  - Example: A naming convention is implicitly followed but not formalized
  - Example: Error handling patterns are inconsistent
- **Missing skills**: Repetitive tasks that haven't been turned into skills
- **Missing agents**: Needed specialized perspectives that are absent
- **Missing procedures/automation**: Missing documentation or scripts for needed processes
  - Example: No setup instructions
  - Example: Build/test execution methods undocumented
  - Example: Deploy/release procedures don't exist
- **Configuration inconsistencies**: Contradictions between existing settings, or divergence from codebase reality

### 4. Create and Execute Proposals

- Organize found gaps by priority
- Create concrete improvement proposals for each gap
- Present to user, implement after approval

## Output Format

```
## Gap Analysis Results

### Codebase Characteristics
- [Summary of languages/frameworks/structure]

### Discovered Gaps

#### Gap 1: [Title]
- **Type**: Missing rule / Missing skill / Missing agent / Missing procedure / Configuration inconsistency
- **Basis**: [Which codebase characteristic reveals this gap]
- **Proposal**: [Concrete improvement]
- **Priority**: High / Medium / Low

#### Gap 2: [Title]
...

### Covered Areas (Reference)
- [Areas adequately covered by current settings]
```

## Notes

- Take care not to break existing settings
- Implement changes after user approval
- Avoid excessive rule additions. Propose only what is truly needed

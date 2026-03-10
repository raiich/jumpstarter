---
name: kaizen
description: Claude Code improvement command
allowed-tools: Read, Grep, Glob, Edit, Write, WebFetch, AskUserQuestion, Skill
---

# Claude Code Improvement Command

Extracts feedback from conversation.log, reflects it in learnings.md, and improves Claude Code settings.

## Procedure

### Phase 1: Feedback Extraction

1. **Load and assess logs**
   - Load `.local/claude/conversation.log`
   - If the log doesn't exist or is empty, skip this phase and proceed to Phase 2

2. **Extract feedback**
   - Extract user corrections, fix requests, policy directives, preferences, and decision criteria
   - Exclude simple task instructions ("do X") — extract only reusable insights
   - Exclude content already recorded in `.local/claude/learnings.md`

3. **Append to learnings.md**
   - If no items were extracted, do not update learnings.md and proceed to Phase 2
   - Append to the appropriate category in `.local/claude/learnings.md`:
     - **Policies**: Development policies, decision-making rationale
     - **Perspectives to consider**: Important viewpoints, decision criteria
     - **Workflow tips**: Effective methods, process improvements
     - **Cautions**: Things to avoid, commonly overlooked points
   - Each item should be concise (one line)
   - Present the additions to the user and update after approval

4. **Clear logs**
   - After approval and update, clear `.local/claude/conversation.log` (make it an empty file)

### Phase 2: Claude Code Settings Improvement

1. **Analyze feedback**
   - Load `.local/claude/learnings.md`
   - Identify improvable patterns from recorded feedback

2. **Check Claude Code features**
   - Use WebFetch to check official Claude Code documentation
   - URL: https://code.claude.com/docs/en/claude_code_docs_map.md
   - Also check the changelog (https://code.claude.com/docs/en/changelog.md) for new features and changes
   - Identify Claude Code features that can address feedback patterns
   - Check for underutilized features in current settings

3. **Create improvement proposals**
   - Combine feedback and Claude Code features into improvement proposals
   - Target the following configuration files:
     - `.claude/settings.local.json` (personal settings; preferred since this is a template repository)
     - `.claude/settings.json` (shared template settings)
     - `.claude/agents/`
     - `.claude/skills/`
     - `.claude/commands/`
     - `.claude/rules/` (actual files in `.github/instructions/`; edit the actual files)

4. **Implement improvements**
   - Present improvement proposals to the user
   - Update configuration files after approval

## Output Format

```
## Feedback Analysis Results

### Key Feedback
- [Feedback 1: frequency X times]
- [Feedback 2: frequency Y times]

### Improvement Proposals

#### Proposal 1: [Title]
- **Target**: [filename]
- **Content**: [change description]
- **Effect**: [expected improvement]

#### Proposal 2: [Title]
...

### Leveraging Latest Claude Code Features
- [Feature name]: [how to use it]
```

## Notes

- Be careful not to break existing settings
- Backup before changes is recommended
- Implement changes only after user approval
- **learnings.md is intermediate data exclusive to kaizen**. Do not propose injecting it into regular sessions (e.g., SessionStart hooks). Reflect improvements in `.claude/` configuration files

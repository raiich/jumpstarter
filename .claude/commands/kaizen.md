# Claude Code Improvement Command

Analyzes user feedback from conversation.log and improves Claude Code settings.

## Execution Steps

1. **Feedback Analysis**
   - Read `.claude/logs/conversation.log`
   - Extract user feedback, requests, and complaints
   - Identify repeatedly pointed out issues

2. **Verify Claude Code Features**
   - Check Claude Code official documentation using WebFetch
   - Check URL: https://docs.claude.com/en/docs/claude-code/claude_code_docs_map.md
   - Identify Claude Code features that can resolve feedback patterns
   - Check if there are any unutilized features in existing settings

3. **Create Improvement Proposals**
   - Create improvement proposals combining feedback and Claude Code features
   - Target the following configuration files:
     - `.claude/settings.json`
     - `.claude/agents/`
     - `.claude/skills/`
     - `.claude/commands/`
     - `.claude/guidelines/`

4. **Implement Improvements**
   - Present improvement proposals to user
   - Update configuration files after approval

## Output Format

```
## Feedback Analysis Results

### Key Feedback
- [Feedback 1: Frequency X times]
- [Feedback 2: Frequency Y times]

### Improvement Proposals

#### Proposal 1: [Title]
- **Target**: [Filename]
- **Content**: [Change details]
- **Effect**: [Expected improvement]

#### Proposal 2: [Title]
...

### Utilizing Latest Claude Code Features
- [Feature name]: [How to utilize]
```

## Improvement Examples

This template itself is continuously improved through `/kaizen`:

### 1st Kaizen
- Guidelines: Generated 4 types
- Added conversation log recording feature

### 2nd Kaizen
- Consolidated and simplified Guidelines
- Eliminated verbosity

### 3rd Kaizen
- Sub-Agents: Added 4 types
- Settings: Enhanced Plan Mode and safety restrictions
- Skills: Simplified doc-reviewer

Each improvement extracts user feedback from conversation.log and reflects it in settings.

## Important Notes

- Be careful not to break existing settings
- Backup before changes is recommended
- Implement changes after user approval

---
title: Kaizen - Project Improvement Proposal
description: Analyze GitHub Copilot conversation history and propose improvements to workspace settings and documentation
---

# Kaizen - Project Improvement Proposal

Analyze user feedback from GitHub Copilot conversation history and improve project settings.

## Execution Steps

1. **Analyze Conversation History**
   - Load conversation history
       - `~/Library/Application Support/Code/User/workspaceStorage/*/chatSessions/*.json`
   - Extract user feedback, requests, and complaints
   - Identify repeatedly mentioned issues or frequently asked questions

2. **Review GitHub Copilot Features**
   - Use fetch_webpage to review official GitHub Copilot documentation
   - Review URLs:
     - https://code.visualstudio.com/docs/copilot/overview
     - https://docs.github.com/en/copilot
   - Identify GitHub Copilot features that could resolve feedback patterns
   - Key feature categories to review:
     - Inline suggestions (code completion)
     - Autonomous coding (autonomous coding by Agents)
     - Natural language chat
     - Smart actions
     - Custom instructions
     - Custom agents
     - MCP servers and tools (external tool integration)
   - Check for underutilized features in existing settings

3. **Identify Improvement Targets**
   - Target the following files:
     - `.github/prompts/` - Prompt files
     - `.github/copilot-instructions.md` - Project-specific instructions
     - `.github/instructions/*.instructions.md` - Project-specific instructions
     - `README.md` - Project documentation
     - `.vscode/settings.json` - Workspace settings
     - Other documentation or configuration files

4. **Create Improvement Proposals**
   - Create improvement proposals combining feedback and GitHub Copilot features
   - Check for underutilized features in existing settings
   - Consider creating new prompt files

5. **Implement Improvements**
   - Present improvement proposals to user
   - Update configuration files after approval

## Output Format

```markdown
## Conversation History Analysis Results

### Main Topics
- [Topic 1: Frequency X times]
- [Topic 2: Frequency Y times]

### Recurring Issues
- [Issue 1: Description]
- [Issue 2: Description]

### Improvement Proposals

#### Proposal 1: [Title]
- **Target**: [File name]
- **Content**: [Changes]
- **Effect**: [Expected improvements]

#### Proposal 2: [Title]
...

### New Prompt Proposals
- [Prompt name]: [Purpose and effect]

### Leveraging Latest GitHub Copilot Features
- [Feature name]: [Usage method]
```

## Improvement Examples

This template itself will be continuously improved through `kaizen`.

## Notes

- Be careful not to break existing settings
- Backup before changes is recommended
- Implement changes only after user approval

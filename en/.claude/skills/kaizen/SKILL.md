---
name: kaizen
description: Claude Code improvement command
allowed-tools: Read, Grep, Glob, Edit, Write, Bash, WebFetch, AskUserQuestion, Skill
---

# Claude Code Improvement Command

Extracts feedback from conversation.log, reflects it in learnings.md, and improves Claude Code settings including codebase gap analysis.

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

   **Example: Feedback extraction judgment**

   Log: "I wanted a sequence diagram in the Design Doc"
   → ✅ Extract (reusable insight: include sequence diagrams when the Design Doc involves processing flows)

   Log: "Fix L42 in auth.go"
   → ❌ Do not extract (one-time task instruction)

4. **Clear logs**
   - After approval and update, clear `.local/claude/conversation.log` (make it an empty file)

### Phase 2: Claude Code Settings Improvement

1. **Analyze feedback**
   - Load `.local/claude/learnings.md`
   - Identify improvable patterns from recorded feedback

2. **Codebase gap analysis**
   - Briefly assess the following (don't over-investigate):
     - Languages, frameworks, directory structure
     - Test setup, CI/CD, major external dependencies
     - Documentation structure, Makefile / task runner availability
   - Compare against current `.claude/` settings to identify uncovered areas:
     - **Missing rules**: Patterns exist in codebase but no corresponding rules
     - **Missing skills/agents**: Repetitive tasks or specialized perspectives not covered
     - **Missing procedures/automation**: Lack of setup, build, test, or deploy procedures
     - **Configuration inconsistencies**: Contradictions between settings or divergence from codebase reality

3. **Check Claude Code features**
   - Check the following documentation (cached):
     - `https://code.claude.com/docs/en/claude_code_docs_map.md`
     - `https://code.claude.com/docs/en/changelog.md`
   - **Cache procedure**:
     1. Read cache files from `.local/claude/cache/` using Read
        - `claude_code_docs_map.md`, `changelog.md`
     2. Compare the frontmatter `fetched_at: YYYY-MM-DD` with the current date; if **within 7 days**, use as-is
     3. Only if the file is missing or older than 7 days, fetch via curl and Write in the following format:
        ```
        ---
        fetched_at: YYYY-MM-DD
        ---
        (fetched content)
        ```
     4. **Important**: Save fetched content as-is — do not summarize or rename files
   - Identify Claude Code features that can address feedback patterns
   - Check for underutilized features in current settings

4. **Create improvement proposals**
   - Combine feedback, gap analysis, and Claude Code features into improvement proposals
   - Target the following configuration files:
     - `.claude/settings.local.json`
     - `.claude/agents/`
     - `.claude/skills/`
     - `.claude/rules/`
     - `.claude/guidelines/`
     - `.claude/hooks/`

5. **Implement improvements**
   - Present improvement proposals to the user
   - Update configuration files after approval

## Output Format

```
## Feedback Analysis Results

### Key Feedback
- [Feedback 1: frequency X times]
- [Feedback 2: frequency Y times]

### Codebase Gaps
- [Gap 1: type - rationale]
- [Gap 2: type - rationale]

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
- **learnings.md and memory are independent**. When extracting feedback, do not check whether it is already saved in memory. The two serve different purposes (learnings.md is input for kaizen improvements, memory is behavioral guidance for regular sessions)

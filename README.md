# Jumpstarter - Claude Code Template

[ **English** | [日本語](README.ja.md) ]

## What is Jumpstarter?

A template repository for Claude Code configuration.

Getting the desired output from Claude Code requires providing appropriate prompts.
Setting up Rules, Skills, and Hooks also takes effort.
Jumpstarter eliminates these hurdles — just clone and start using it, so you can focus on development itself.

The goal is for users to **not have to struggle** with Claude Code:

- **No lengthy prompts needed** — Codebase investigation, requirements hearing, and self-review are handled automatically.
- **No complex setups upfront** — MCP and multi-agent can wait. Focus on Rules and Skills.
- **Continuous improvement made easy** — Settings are auto-updated based on improvement suggestions and best practices.

---

## Usage

Copy the template and start using it right away.

```bash
git clone https://github.com/raiich/jumpstarter.git

# Copy .claude/ and .github/ to your existing project (English version)
cp -R jumpstarter/en/.claude "${YOUR_PROJECT}/"
cp -R jumpstarter/en/.github "${YOUR_PROJECT}/"

## Japanese version
# cp -R jumpstarter/.claude "${YOUR_PROJECT}/"
# cp -R jumpstarter/.github "${YOUR_PROJECT}/"
```

Launch Claude Code and use slash commands to run workflows.

```bash
claude
```

### Design and implement new features

No need to write a perfect prompt when adding features.
We provide workflows for proper design and implementation from minimal instructions.

The `/design-feature` skill investigates the codebase and conducts requirements hearings to create a Design Doc:

```
/design-feature I want to add a verbose option to the CLI
```

If you want to design test cases in advance to reduce implementation drift, the `/design-feature-tests` skill is useful:

```
/design-feature-tests .local/docs/features/verbose-option/design-doc.md
```

The `/implement-feature` skill implements based on the Design Doc and test cases:

```
/implement-feature .local/docs/features/verbose-option/design-doc.md
```

Design documents and test cases are organized under `.local/docs/features/[name]/`.
Each step's deliverable serves as a checkpoint, minimizing rework even if implementation drifts.

### Continuously improve Claude Code settings

Poor development experience is often caused by daily feedback not being reflected in settings.
Users don't need to manually take action to improve settings.
The following mechanisms help with continuous improvement.

The `/kaizen` command extracts improvement suggestions from conversation logs recorded by Hooks:

```
/kaizen   Suggest improvements from conversation logs
```

The `/import-best-practices` command easily imports best practices from online articles:

```
/import-best-practices   https://... Import best practices from this article
```

These refine rules, skills, hooks, and more to suit your project. No need to manually review settings.

### Improve code and documentation quality

Discover issues through review, comprehensively search for similar issues, and batch-fix them.
This reduces missed fixes and the effort of pointing out "this wasn't fixed."

```
Code:           /review-code → /fix-code
Documentation:  /review-doc  → /fix-doc
```

You can also use Claude Code's built-in `/simplify` for code quality improvement.

---

## Other Settings

Each element in the `.claude/` directory works together to achieve the "no struggle" experience.

### Rules for not struggling

Rules that ensure consistent quality without users having to give instructions every time (`.claude/rules/`).

- No need to struggle reading verbose output → **writing-style**
- No need to struggle checking for missed fixes → **fix-guidelines**
- No need to struggle specifying workflows → **workflow-patterns**

### Hooks for not struggling

Conversation logs are automatically recorded to `.local/claude/conversation.log` as input data for `/kaizen` improvement suggestions.

---

## Customization

After copying the template, you can adjust it to suit your project.

- **Add skills and rules**: Add files to `.claude/skills/` or `.claude/rules/`
- **Permission settings**: Adjust allowed/denied commands in `.claude/settings.local.json`
- **Auto-improvement**: Extract improvements from conversation logs with `/kaizen`, import best practices from external articles with `/import-best-practices`

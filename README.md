# Jumpstarter - Claude Code Template

[ **English** | [日本語](README.ja.md) ]

## What is Jumpstarter?

A template repository for Claude Code configuration.

Getting the desired output from Claude Code requires providing appropriate prompts.
Setting up Rules, Skills, and Hooks also takes effort.
Jumpstarter eliminates these hurdles — just clone and start using it, so you can focus on development itself.

The goal is for users to **not have to struggle** with Claude Code:

- **No lengthy prompts needed** — Codebase investigation, requirements hearing, and self-review are handled
  automatically.
- **No complex setups upfront** — MCP and multi-agent can wait. Focus on Rules and Skills.
- **Continuous improvement made easy** — Settings are auto-updated based on improvement suggestions and best practices.

---

## Usage

Copy the template and start using it right away.

```bash
git clone https://github.com/raiich/jumpstarter.git

# Copy .claude/ to your existing project (Japanese version)
cp -R jumpstarter/.claude "${YOUR_PROJECT}/"

# Note: an English version of .claude/ is planned for a future release.
```

Launch Claude Code and use slash commands to run workflows.

```bash
claude
```

### Sketch new features

No need to write a perfect prompt when adding features.
From minimal instructions, we provide a workflow that produces a buildable sketch instead of a standalone design doc:
the "why" goes into a thin design doc, the "what / how" into code stubs that compile, leaving interface consistency to
the type checker.

The `/sketch-feature` skill investigates the codebase while hearing requirements, then writes public types and
signatures with doc comments (build passes) plus test cases that fail explicitly until implemented:

```
/sketch-feature I want to add a verbose option to the CLI
```

The thin `design.md` is saved under `.local/docs/features/[name]/`; the design and test stubs live in the real source
tree. You then fill in the stubs to implement.

### Continuously improve Claude Code settings

Poor development experience is often caused by daily feedback not being reflected in settings.
Users don't need to manually take action to improve settings.
The following mechanisms help with continuous improvement.

The `/kaizen` command observes problems that arose during the session, identifies root causes via 5 Whys analysis, and
prevents recurrence by updating settings:

```
/kaizen                        Observe the whole session
/kaizen <problem description>  Focus on a specific problem
```

It also triggers when the user asks "why did this happen?" — responding with root cause analysis and recurrence
prevention instead of surface-level apologies.

The `/import-best-practices` command easily imports best practices from online articles:

```
/import-best-practices   https://... Import best practices from this article
```

These refine rules, skills, hooks, and more to suit your project. No need to manually review settings.

---

## Other Settings

Each element in the `.claude/` directory works together to achieve the "no struggle" experience.

### Rules for not struggling

Rules that ensure consistent quality without users having to give instructions every time (`.claude/rules/`).

- No need to struggle reading verbose output → **writing-style**
- No need to struggle checking for missed fixes → **review**
- No need to struggle correcting flattery / unverified claims → **no-sycophancy**

---

## Customization

After copying the template, you can adjust it to suit your project.

- **Add skills and rules**: Add files to `.claude/skills/` or `.claude/rules/`
- **Auto-improvement**: Analyze session problems for root causes and standardize fixes with `/kaizen`; import best
  practices from external articles with `/import-best-practices`

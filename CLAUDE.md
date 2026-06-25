# Jumpstarter (Development Instructions for This Template Repository)

A template repository for Claude Code configuration. **This file is the development guide for this repository itself** —
it is not the `CLAUDE.md` that ends up in projects copied from this template. For the user-facing overview,
see [README.md](README.md) / [README.ja.md](README.ja.md).

## Target Environments

The representative coding agent is Claude Code; `.claude/` is the edit target. The same assets are designed to also run
under **GitHub Copilot (CLI / VSCode)** — users who run Copilot copy the contents of `.claude/` into `.github/` after
cloning. For the full scope (which Copilot products are in/out) and the feature mapping,
see [docs/development/claude-vs-copilot.md](docs/development/claude-vs-copilot.md).

- This repository does not contain `.github/`.
- Authoring conventions (cross-environment compatibility, template purity) live
  in [docs/development/authoring-guide.md](docs/development/authoring-guide.md).

## Repository Layout

```
.claude/                  # Claude Code config (Japanese)
  agents/                 # Custom agent definitions (*.agent.md)
  rules/                  # Always-loaded rules (*.instructions.md)
  references/             # On-demand reference material (vocabulary, lookups, examples) linked from rules and skills
  skills/<name>/          # Skill definitions
    SKILL.md              #   skill policy and flow (applies across environments)
    environment.md        #   per-environment concrete means (Claude / Copilot sections)
    references/*.md       #   skill-local reference docs (perspective checklists etc.) consulted at task time (optional)
    README.md             #   user-facing usage examples

.devcontainer/        # Dev container config (docker-compose: workspace + egress gateway)
docs/development/     # Repository-development docs in English (not copied into downstream projects)
```

## Language

`.claude/` is written in Japanese. `CLAUDE.md` and `docs/development/*.md` are written directly in English.
`.devcontainer/` contains only JSON / shell configuration (no natural-language content).

## Rule Application Scope

`.claude/rules/*.instructions.md` apply to **every file in this repository, including `.claude/` itself and the rule
files**. When editing skills, agents, rules, or docs, re-check the relevant rule file before finishing the edit —
especially `writing-style.instructions.md` (terseness, terminology, emphasis discipline, and post-edit
self-review against violation signatures).
"It's a `.claude/` config edit, so style rules don't apply" is
not a valid exception.

## Capturing Repo-Specific Conventions

When a session surfaces a convention meant to apply beyond the current edit — a frontmatter rule, authoring pattern,
or empirically verified gotcha — append it to [docs/development/authoring-guide.md](docs/development/authoring-guide.md)
so future sessions don't re-derive it. One-off edits don't need capturing; policy-shaped instructions ("always", "from
now on") and non-obvious findings (schema constraints, runtime behavior) do.

## Related

- [docs/development/authoring-guide.md](docs/development/authoring-guide.md) — Authoring conventions for skills, agents,
  rules, and instructions
- [docs/development/claude-vs-copilot.md](docs/development/claude-vs-copilot.md) — Feature and config mapping between
  Claude Code and GitHub Copilot (CLI / VSCode)
- [docs/development/sensitive-info-review.md](docs/development/sensitive-info-review.md) — This template's pre-vetted
  allowlist and leakage hotspots; the generic scan lives in the `publish-gate` skill

# Jumpstarter

A template repository for Claude Code configuration. See [README.md](README.md) / [README.ja.md](README.ja.md) for details.

## Repository Structure

Bilingual layout: Japanese (root) and English (`en/`).

```
.claude/              # Claude Code config (Japanese)
  agents/             # Custom agent definitions
  guidelines/         # Shared policies, perspectives, and checklists referenced from skills
    perspectives/     # Review/design perspectives (what to check)
    processes/        # Processes and principles (how to proceed)
  hooks/              # Hooks (conversation logging, etc.)
  rules/              # Rules (*.instructions.md)
  skills/             # Skill definitions (SKILL.md per skill)
  settings.local.json

.devcontainer/    # Dev container config

en/                   # English version (mirrors root structure)
  .claude/
  .devcontainer/
```

## JA/EN Sync Rule

When modifying files under `.claude/` or `.devcontainer/`, update the corresponding file under `en/` as well (and vice versa).

- Edit a root file → update the `en/` counterpart
- Edit an `en/` file → update the root counterpart
- Adding a new file → add to both language versions
- Translate content (Japanese ↔ English). Keep config values (JSON, etc.) identical except for the language key
- File names and directory structure are shared between JA/EN (only the body is translated)

## Guidelines vs Rules

- `.claude/rules/` — **norms** (must follow). Style and bash command rules.
- `.claude/guidelines/` — **policies / perspectives** (reference to improve quality). Referenced from SKILL.md via relative links, loaded on demand rather than always-on.

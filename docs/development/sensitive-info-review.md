# Sensitive Information Review

This repository is a **template**: anything committed propagates into every downstream clone, so the bar for "safe to
ship" is stricter than a typical application repo. Review tracked files (mainly `.claude/`, `.devcontainer/`, and the
root `*.md`) before publishing.

The generic scan — categories, grep patterns, triage, and remediation — is the `publish-gate` skill
(`.claude/skills/publish-gate/`). This document records only what is specific to publishing **this** template: the
excluded paths, the pre-vetted allowlist, and where leakage tends to appear here.

## Excluded by `.gitignore`

No need to inspect: `.env*`, `*.key`, `*.pem`, `*.local.json`, `.DS_Store`, `.idea/`, `.vscode/`, `.local/`,
`.claude/agent-memory-local/`. Confirm coverage with `git check-ignore <path>`.

## Pre-Vetted Public References

The repository's known-safe allowlist (the project-specific allowlist the skill defers to). The references below appear
in tracked files and are already judged safe; flag only new occurrences outside this set.

- `api.anthropic.com`, `code.claude.com`, `claude.ai/install.sh` — Claude Code installation / endpoint
- `mcr.microsoft.com/devcontainers/base:debian` — public base image
- `docs.github.com`, `raw.githubusercontent.com`, `github.blog`, `agents.md`, `code.visualstudio.com` — public docs
- `deb.debian.org`, `www.debian.org`, `bugs.debian.org` — public Debian mirrors
- `just.systems`, `playwright.dev` — public tool documentation (devcontainer features)
- `github.com/raiich/jumpstarter`, handle `raiich` — this repository's own clone URL and handle
- `/home/vscode/.claude`, user `vscode` — devcontainer convention, not a real account

Repo-specific allowed exceptions beyond the skill's generic allowlist:

- The placeholder set in `.claude/rules/writing-style.instructions.md`
- Generic business terms ("customer base", "stakeholders") in analysis-agent prompts

## Where Leakage Typically Appears

- `*.agent.md` / `SKILL.md` examples — security-concept examples may drift into using real values
- `.devcontainer/devcontainer.json` — `name`, IDE backend, mount paths often carry maintainer defaults
- `environment.md` — per-environment paths can capture the author's specific machine layout
- Any new `docs/development/*.md` that quotes log output or transcripts — likely to embed local paths

When a scan surfaces a new public reference (e.g., a newly added official docs domain), add it to the allowlist above
so future reviews skip it.

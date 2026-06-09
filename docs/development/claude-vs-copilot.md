# Claude Code vs GitHub Copilot (CLI / VSCode) — Feature & Configuration Mapping

The two tools are conceptually similar but differ in config-file locations, frontmatter, and tool names. This document
captures the mapping alongside the conventions used in this repository.

**Scope note**: the Copilot side refers to **GitHub Copilot CLI and the VS Code IDE**. To keep the surface uniform,
features used in this repo are limited to the **GitHub Copilot CLI**; the VS Code IDE supports the same scope. Other
GitHub Copilot products (PR review, Copilot in GitHub Actions, etc.) and VS Code-only capabilities (IDE-specific UI
affordances) are intentionally out of scope.

Confidence marks:

- ✅ Confirmed in official docs or implementation
- ⚠️ Reasonable inference from official docs
- ❓ Unverified

## How This Repository Maps to Each Tool

This repository (jumpstarter) targets **Claude Code as the primary environment**; `.claude/` is the edit target. Users
who run GitHub Copilot (CLI / VSCode) copy the contents of `.claude/` into `.github/` after cloning. This repository
does not contain `.github/`.

| Concept                | This repo (`.claude/`)                          | User-provided on copy (`.github/`)                                                                |
|------------------------|-------------------------------------------------|---------------------------------------------------------------------------------------------------|
| Repo-wide instructions | `CLAUDE.md` (this template's development guide) | `AGENTS.md` or `.github/copilot-instructions.md`                                                  |
| Path-specific rules    | `.claude/rules/*.instructions.md`               | `.github/instructions/*.instructions.md`                                                          |
| Custom agents          | `.claude/agents/*.agent.md`                     | `.github/agents/*.agent.md`                                                                       |
| Skills                 | `.claude/skills/<name>/SKILL.md`                | `.github/skills/<name>/SKILL.md`                                                                  |
| Hooks                  | `.claude/hooks/` (unused in this repo)          | `.github/hooks/<name>.json` (repo) / `~/.copilot/hooks/` (user) ✅                                 |
| Memory                 | `~/.claude/projects/<id>/memory/` (user home)   | "Copilot Memory" — repo-level managed memory ✅ (semantics differ from Claude's file-based memory) |

Copilot also reads `.claude/skills/` for backward compatibility ✅. When `.github/skills/` and `.claude/skills/` both
contain the same skill, precedence is undocumented ❓. This repo ships only `.claude/skills/`, leaving the placement
decision to the user.

## Skill / Agent Auto-Invocation

Both tools auto-invoke skills/agents by matching the user prompt against the `description`. The `/<name>` syntax is for
explicit invocation; it is not a prerequisite.

- Claude Code: the model selects via internal tools ✅
- GitHub Copilot (CLI / VSCode): auto-inference loads relevant skills/agents ✅. Skills added mid-session can be
  reloaded (on the CLI, via `/skills reload`) ✅

## Config File Locations

### Repository-wide instructions

| Role                                 | Claude Code             | GitHub Copilot (CLI / VSCode)              |
|--------------------------------------|-------------------------|--------------------------------------------|
| Primary instructions (always loaded) | `CLAUDE.md` (repo root) | `AGENTS.md` (repo root, open standard) ✅   |
| Compatible names                     | —                       | `CLAUDE.md` is also picked up if at root ✅ |
| Additional repo instructions         | —                       | `.github/copilot-instructions.md` ✅        |

`AGENTS.md` is a community standard that Copilot also reads ✅.

### Path- or topic-specific instructions

| Role                 | Claude Code                                                | GitHub Copilot (CLI / VSCode)                   |
|----------------------|------------------------------------------------------------|-------------------------------------------------|
| Path/topic-specific  | `.claude/rules/*.instructions.md` (this repo's convention) | `.github/instructions/**/*.instructions.md` ✅   |
| Required frontmatter | `applyTo: "**"` (carried for Copilot; Claude ignores it)   | `applyTo` (glob, e.g. `"**/*.ts,**/*.tsx"`) ✅   |
| Optional frontmatter | `description` (rule summary; ignored by both as an instructions field) | `excludeAgent` (exclude from specific agents) ✅ |

Claude Code has no built-in mechanism for `.instructions.md`; it loads these rules globally and ignores `applyTo`. The
files nonetheless carry `applyTo: "**"` (plus a short `description`) so they drop into Copilot's `.github/instructions/`
unchanged. This repo references them from `CLAUDE.md` or from skills.

### Custom agents

| Item                   | Claude Code                                                            | GitHub Copilot (CLI / VSCode)       |
|------------------------|------------------------------------------------------------------------|-------------------------------------|
| Project location       | `.claude/agents/*.md` ✅ (this repo uses `.agent.md`, accepted in both) | `.github/agents/*.agent.md` ✅       |
| User location          | `~/.claude/agents/` ✅                                                  | `~/.copilot/agents/*.agent.md` ✅    |
| Recommended extension  | `.md`                                                                  | `.agent.md` (`.md` also accepted) ✅ |
| Precedence on conflict | User overrides repo ⚠️                                                 | User-home > Repo ✅ (CLI agents)     |

### Built-in agents

Both environments ship a default agent set. Useful overlap below; reach for these before defining custom equivalents.

| Role                                | Claude Code                          | GitHub Copilot (CLI / VSCode)                                                    |
|-------------------------------------|--------------------------------------|----------------------------------------------------------------------------------|
| Code exploration / read-only search | `Explore` ✅ (built-in subagent)      | `Explore` ✅ (built-in, "quick codebase analysis without adding to main context") |
| General multi-step task             | `general-purpose` ✅                  | `General purpose` ✅                                                              |
| Plan / design                       | `Plan` ✅                             | —                                                                                |
| Command / test / build execution    | — (main agent uses `Bash`)           | `Task` ✅                                                                         |
| Code review                         | — (use custom or `/security-review`) | `Code review` ✅                                                                  |
| Deep web + code research            | —                                    | `Research` ✅                                                                     |
| Constructive feedback               | —                                    | `Rubber duck` ✅ (auto-invoked)                                                   |

Invocation: Claude uses `subagent_type=<name>` via the `Agent` tool. Copilot uses the `agent` tool, the `/agent` slash
command, direct mention in the prompt, or the `--agent=<name>` CLI flag ✅.

The table above lists the agents likely to be reached for in skill flows. Other Claude Code built-ins (setup utilities
like `statusline-setup`, help bots like `claude-code-guide`) are omitted — they're scoped to specific UI / support flows
rather than general work.

### Skills

| Item             | Claude Code                        | GitHub Copilot (CLI / VSCode)                                                                  |
|------------------|------------------------------------|------------------------------------------------------------------------------------------------|
| Project location | `.claude/skills/<name>/SKILL.md` ✅ | `.github/skills/<name>/SKILL.md` (recommended) or `.claude/skills/<name>/` (backward-compat) ✅ |
| User location    | `~/.claude/skills/<name>/` ✅       | `~/.copilot/skills/<name>/` or `~/.agents/skills/<name>/` ✅                                    |

### permissions, MCP, hooks

| Feature               | Claude Code                                             | GitHub Copilot (CLI / VSCode)                                                                                                                                                                                                                                       |
|-----------------------|---------------------------------------------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| Permission allow/deny | `permissions.allow/deny/ask` in settings ✅       | Interactive approval per call ⚠️ (CLI also accepts `--allow-all` / `--yolo` flags)                                                                                                                                                                                  |
| MCP                   | `.mcp.json` / settings ✅                                | `~/.copilot/mcp-config.json`, or agent frontmatter `mcp-servers` ✅                                                                                                                                                                                                  |
| Hooks                 | `.claude/hooks/` (UserPromptSubmit, PreToolUse, etc.) ✅ | `.github/hooks/<name>.json` (repo) / `~/.copilot/hooks/` (user). Triggers: `sessionStart` / `sessionEnd` / `userPromptSubmitted` / `preToolUse` / `postToolUse` / `errorOccurred` / `agentStop`. JSON schema with `version: 1` and `bash` / `powershell` commands ✅ |
| Environment variables | —                                                       | CLI: `COPILOT_HOME` (config dir), `COPILOT_CUSTOM_INSTRUCTIONS_DIRS` (extra instruction dirs) ✅                                                                                                                                                                     |

## Tool Names

Copilot uses **two different naming schemes** for agents and skills.

- **Agent `tools`**: aliases (`execute` / `read` / `edit` / `search` / `agent` / `web` / `todo`), case-insensitive ✅
- **Skill `allowed-tools`**: MCP-style canonical names (`shell`, `bash`, `read_file`, `list_directory`, `search_files`,
  `server-name/tool-name`) ✅

The table below shows the **agent `tools` alias mapping**.

| Claude Code                                       | Copilot (agent alias)                                       |
|---------------------------------------------------|-------------------------------------------------------------|
| `Read`                                            | `read`                                                      |
| `Grep`, `Glob`                                    | `search`                                                    |
| `Edit`, `Write`                                   | `edit`                                                      |
| `Bash`                                            | `execute` (`shell` / `bash` / `powershell` also accepted) ✅ |
| `WebFetch`, `WebSearch`                           | `web`                                                       |
| `Agent`                                           | `agent`                                                     |
| `TaskCreate`, `TaskUpdate`, `TaskList`, `TaskGet` | `todo`                                                      |
| `AskUserQuestion`                                 | No equivalent                                               |
| `Skill`                                           | No equivalent (skills auto-load from search paths)          |
| `EnterPlanMode`, `ExitPlanMode`                   | No equivalent (Shift+Tab toggles mode) ✅                    |
| `NotebookEdit`                                    | No equivalent ⚠️                                            |

To reference an MCP tool in Copilot, use `server-name/tool-name` or `server-name/*` for all.

## Frontmatter

### Agents

| Field                      | Claude Code                         | GitHub Copilot (CLI / VSCode)                           | Notes                                                                                     |
|----------------------------|-------------------------------------|---------------------------------------------------------|-------------------------------------------------------------------------------------------|
| `name`                     | yes                                 | yes                                                     | Copilot can fall back to file name                                                        |
| `description`              | yes                                 | required                                                | Auto-invocation trigger                                                                   |
| `tools`                    | comma-separated string (optional) ✅ | comma-separated string or array of aliases (optional) ✅ | **Value vocabularies differ (Claude tool names vs Copilot aliases) → this repo omits it** |
| `model`                    | yes (`haiku`/`sonnet`/`opus`, etc.) | yes (default if omitted) ✅                              | Copilot does not publish the value list ❓                                                 |
| `memory`                   | yes                                 | absent (ignored) ✅                                      |                                                                                           |
| `maxTurns`                 | yes                                 | absent (ignored) ✅                                      |                                                                                           |
| `disable-model-invocation` | skill only ✅                        | supported on agents (boolean, default `false`) ✅        | Prevents auto-selection. Former `infer` is retired ✅                                      |
| `user-invocable`           | skill only ✅                        | supported on agents (boolean, default `true`) ✅         |                                                                                           |
| `target`                   | —                                   | `vscode` / `github-copilot` ✅ (default: both)           |                                                                                           |
| `mcp-servers`              | —                                   | object ✅ (CLI only; ignored in VS Code)                 |                                                                                           |

### Skills

| Field                      | Claude Code                                | GitHub Copilot (CLI / VSCode)                                    |
|----------------------------|--------------------------------------------|------------------------------------------------------------------|
| `name`                     | yes ✅ (required, matches directory name)   | yes ✅ (required, matches directory name)                         |
| `description`              | yes ✅                                      | yes (required) ✅                                                 |
| `allowed-tools`            | yes ✅ (optional, comma-separated or array) | yes ✅ (optional, array of MCP names)                             |
| `disable-model-invocation` | yes ✅                                      | yes (boolean; `true` makes it explicit-only via `/skill-name`) ✅ |
| `version`                  | yes ✅                                      | undocumented                                                     |
| `user-invocable`           | yes ✅                                      | yes ✅                                                            |
| `mode`                     | yes ✅                                      | undocumented                                                     |
| `license`                  | —                                          | yes ✅                                                            |

### Instructions (path-specific)

| Field          | Claude Code                 | GitHub Copilot (CLI / VSCode) |
|----------------|-----------------------------|-------------------------------|
| `applyTo`      | `"**"` (carried, ignored)   | required (glob)               |
| `excludeAgent` | —                           | optional                      |

## This Repository's Frontmatter Policy

To share files across both environments, fields fall into three groups. This repo runs on devcontainer + auto/edit mode,
so narrowing tool permissions is unnecessary.

### ❌ Do not write (schema incompatible)

| Field                   | Reason                                                                                                                                                                    |
|-------------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| `tools` (agent)         | Both accept comma-separated string and array, but Claude uses tool names (`Read`, `Bash`) while Copilot uses aliases (`read`, `execute`) — values are not interchangeable |
| `allowed-tools` (skill) | Claude uses tool names; Copilot uses MCP canonical names — schemas don't match                                                                                            |

When omitted, the parent agent / default tool set is inherited. Copilot falls back to interactive per-call approval.

### ⭕ Ignored by one side but safe to write

The side that understands it interprets; the other side ignores. Worth keeping for Claude-side optimization.

| Field             | Claude          | Copilot                             |
|-------------------|-----------------|-------------------------------------|
| `model`           | interprets      | interprets (value list unpublished) |
| `memory`          | interprets      | ignored                             |
| `maxTurns`        | interprets      | ignored                             |
| `mode`, `version` | interprets      | ignored                             |
| `license`         | ignored         | interprets                          |

### ✅ Works in both

| Field                      | Notes                                                                         |
|----------------------------|-------------------------------------------------------------------------------|
| `name`                     | Skill: required, matches directory. Agent: Copilot may fall back to file name |
| `description`              | **Auto-invocation trigger.** Include both trigger and non-trigger conditions  |
| `disable-model-invocation` | Claude: skills / Copilot: both agents and skills                              |
| `user-invocable`           | Same as above                                                                 |

## Other Feature Mapping

| Feature            | Claude Code                                                                                       | GitHub Copilot (CLI / VSCode)                                                            |
|--------------------|---------------------------------------------------------------------------------------------------|------------------------------------------------------------------------------------------|
| Plan mode          | `EnterPlanMode` / `ExitPlanMode` tools ✅                                                          | Shift+Tab toggles mode ✅                                                                 |
| Inline shell       | `!<command>` in prompt ✅                                                                          | `!<command>` in prompt ✅                                                                 |
| File reference     | `@path/to/file` ✅                                                                                 | `@path/to/file` ✅                                                                        |
| Session resume     | `--resume` / `--continue` ✅                                                                       | `--resume` / `--continue` (CLI) ✅                                                        |
| Reasoning display  | Auto-shown                                                                                        | Ctrl+T to toggle (CLI) ✅                                                                 |
| Slash commands     | `/login`, `/agent`, `/clear`, `/compact`, `/context`, `/init`, `/model`, `/usage`, `/mcp`, etc. ✅ | `/login`, `/agent`, `/cwd`, `/add-dir`, `/usage`, `/context`, `/compact`, `/mcp`, etc. ✅ |
| Memory             | `~/.claude/projects/<id>/memory/` (file-based, per-project) ✅                                     | "Copilot Memory" — repo-level managed memory ✅ (not file-based; semantics differ)        |
| Background tasks   | `run_in_background` (per tool call) ✅                                                             | No per-tool-call equivalent ❓                                                            |
| Worktree isolation | `isolation: "worktree"` (per agent call) ✅                                                        | Session-level Workspace / Worktree isolation modes (choose at session start) ✅           |

## Constraints & Caveats

- `applyTo` is required by Copilot. This repo's `.claude/rules/*.instructions.md` already carry `applyTo: "**"`, so they
  drop into `.github/instructions/` unchanged — no per-file transform on port
- Claude-specific tools (`AskUserQuestion`, `Skill`, `EnterPlanMode`, etc.) have no Copilot equivalent. Skills that
  depend on them will degrade on Copilot
- Hook scripts don't transfer 1:1: Copilot CLI hooks use a `version: 1` JSON config with separate `bash` / `powershell`
  commands, whereas Claude's `.claude/hooks/` holds standalone scripts. Trigger names also differ (e.g.
  `userPromptSubmitted` vs `UserPromptSubmit`)

## Verification

Official specs change. When adopting new fields or tools:

- Verify against official docs via web fetch (the `import-best-practices` skill helps)
- Mark unverifiable items with `❓`
- Update the confidence marks (✅ / ⚠️ / ❓) in this file

## References

- [AGENTS.md (open standard)](https://agents.md/)
- [Using GitHub Copilot CLI — overview](https://docs.github.com/en/copilot/how-tos/copilot-cli/use-copilot-cli/overview) —
  built-in agents (Explore / Task / General purpose / Code review / Research / Rubber duck)
- [Adding custom instructions for GitHub Copilot CLI](https://docs.github.com/en/copilot/how-tos/copilot-cli/customize-copilot/add-custom-instructions)
- [Creating and using custom agents for GitHub Copilot CLI](https://docs.github.com/en/copilot/how-tos/copilot-cli/customize-copilot/create-custom-agents-for-cli)
- [Adding agent skills for GitHub Copilot CLI](https://docs.github.com/en/copilot/how-tos/copilot-cli/customize-copilot/add-skills)
- [Custom agents configuration reference](https://docs.github.com/en/copilot/reference/custom-agents-configuration)
- [Using hooks with GitHub Copilot CLI](https://docs.github.com/en/copilot/how-tos/copilot-cli/customize-copilot/use-hooks)
- [Copilot CLI sessions in Visual Studio Code](https://code.visualstudio.com/docs/copilot/agents/copilot-cli) —
  Workspace / Worktree isolation
- [Copilot Memory now on by default — GitHub Changelog (2026-03-04)](https://github.blog/changelog/2026-03-04-copilot-memory-now-on-by-default-for-pro-and-pro-users-in-public-preview/)

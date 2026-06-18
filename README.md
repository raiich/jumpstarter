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

It also ships a **dev container** for running the agent autonomously and safely: outbound traffic is
restricted to domains you allow, and runtimes like Node, Python, Go, and .NET come preinstalled, so it
runs in isolation out of the box.

---

## Usage

The main way to use Jumpstarter is to run the `jumpstart` command in your own repo, reusing the
bundled dev container. You can also open jumpstarter itself directly, or just copy `.claude/` into
your project.

### Use the dev container

Clone jumpstarter once, then run `jumpstart` from your target repo to reuse the dev container.
`jumpstart` symlinks jumpstarter's `.devcontainer` into the target repo — a symlink rather than a
copy, so the template's updates apply automatically — and its bind mount exposes jumpstarter's
`.claude/` at that repo's root.

```bash
git clone https://github.com/raiich/jumpstarter.git

# Run from the target repo (jumpstarter/bin on PATH, or call by full path)
cd /path/to/other-repo
jumpstart
```

Reopen the target repo in the dev container afterward (VS Code: **Reopen in Container**; requires
Docker and the Dev Containers extension). The symlink points at a machine-local path, so don't commit
it. The allowed domains live in `.devcontainer/gateway/allowed-domains.acl`, and forwarded ports (5173
for dev servers by default) in `devcontainer.json`.

> **Note.** Claude's auth and settings live in a shared `claude-code-config` volume, so login state
> and config changes (e.g. from `/kaizen`) apply across every project that uses this container.

#### Sibling directories are visible

The container mounts the target repo's *parent* directory, so beyond the workspace you can also reach
sibling directories — handy for using your own libraries kept alongside the repo. They're mounted
read/write, though, so keep repos holding secrets out of the same parent.

#### Open jumpstarter itself

Clone jumpstarter and open it directly in the dev container to use it without `jumpstart`. Suited to
trying out or developing the template itself.

### Copy `.claude/` only

Without the dev container, you can simply copy `.claude/` into an existing project. It's lighter, but
you don't get the network isolation or bundled runtimes.

```bash
git clone https://github.com/raiich/jumpstarter.git

# Copy .claude/ into your project (Japanese; English version in progress)
cp -R jumpstarter/.claude "${YOUR_PROJECT}/"

claude
```

### Develop with Claude Code

Once set up, just launch `claude` and develop as usual — the bundled Rules and Skills kick in
automatically, no special prompts required. For specific situations, the workflows below are also
available.

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
- No need to struggle correcting flattery / unverified claims → **no-sycophancy**

### Skills for not struggling

Beyond the workflows above, these skills (`.claude/skills/`) cut everyday friction.

- No need to struggle with half-finished fixes or ambiguous change requests → **fix-well** (pins the interpretation down first when ambiguous, then fixes the related spots too, consistently)
- No need to struggle thinking from a single angle → **hyper-think** (parallel multi-agent review / ideation)

---

## Customization

After copying the template, you can adjust it to suit your project.

- **Add skills and rules**: Add files to `.claude/skills/` or `.claude/rules/`
- **Auto-improvement**: Analyze session problems for root causes and standardize fixes with `/kaizen`; import best
  practices from external articles with `/import-best-practices`

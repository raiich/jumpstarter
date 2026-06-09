# Authoring Guide for Skills / Agents / Instructions

Conventions for editing skills, agents, and rule files under `.claude/`. This repository targets Claude Code as the
primary environment; the same assets are designed to also run under **GitHub Copilot (CLI / VSCode)** once users copy
`.claude/` into `.github/`.

For the supported-product scope and the full feature / config mapping (paths, tool names, frontmatter details),
see [claude-vs-copilot.md](claude-vs-copilot.md). This file focuses on "how to write." It lives in `docs/development/`
and is not copied into downstream projects.

## Premise

A given user runs either Claude Code or GitHub Copilot (CLI / VSCode) — not both. Author assets so they work in either
environment, and at deploy time **keep only the section for the chosen environment and delete the other**. Duplicate
content across both sections of each `environment.md` is intentional under this premise.

## Template Purity

`.claude/` is the body that downstream projects copy. Keep jumpstarter-specific concerns out of `.claude/` so the copied
template is free of the maintainer's situation.

- **In `.claude/` (`hooks/`, `rules/`, `skills/`, `agents/`, `settings.json`)**: only content that makes sense in any
  project that adopts this template. No references to jumpstarter-specific facts (maintainer's directory layout,
  internal mirrors, etc.).
- **Maintainer-specific guidance**: lives in the root `CLAUDE.md`, `docs/development/`, user-home memory — not in
  tracked `.claude/` files.
- **Authoring test**: "Would this make sense in a project that copied `.claude/` from this repo?" If no, it doesn't
  belong in `.claude/`.

The rule applies inside skills as well: `SKILL.md` bodies, `environment.md`, and `references/` must not assume any
jumpstarter-repo-specific structure.

## File Responsibilities

Each skill lives at `.claude/skills/<name>/` and contains:

| File (path under `.claude/skills/<name>/`) | Write here                                                         | Don't write here                                        |
|--------------------------------------------|--------------------------------------------------------------------|---------------------------------------------------------|
| `SKILL.md`                                 | Skill policy, flow, principles — cross-environment, always kept    | Per-environment tool names / paths, env-specific config |
| `environment.md`                           | Per-environment concrete means, user-specific local paths          | Environment-independent policy (→ `SKILL.md`)           |
| `references/*.md`, `*-mode.md`, etc.       | Supplementary docs split by concern (templates, decision criteria) | The skill's main flow                                   |
| `README.md`                                | User-facing usage examples (concrete invocation samples OK)        | Internal specifications                                 |

## `references/*.md`

`references/` holds **on-demand support material** that a skill loads when a phase needs it. The aim is to stabilize
agent behavior against missed perspectives (e.g., testing depth, error handling, public-API stability, observability)
and against training-data tendencies (e.g., over-confident assertions, scope creep, self-checking tests with the
implementation).

### File-type taxonomy

A `references/` file falls into one of three types. Mixing types in one file makes it hard for an agent to know
*how* to apply the content. Keep the type uniform within a file; if a file drifts to two types, split it.

| Type                | Body shape                                                       | Example                                                                                                                  |
|---------------------|------------------------------------------------------------------|--------------------------------------------------------------------------------------------------------------------------|
| Perspective         | List of "what to consider" / "what to check" — no ordering       | `sketch-feature/references/missable-checklist.md`                                                                       |
| Decision criteria   | Trade-offs and selection guidance for choosing among alternatives | `sketch-feature/references/verification-strategy.md`                                                                    |
| Procedure / process | Ordered steps to follow when a specific condition triggers       | `sketch-feature/references/derive-perspectives.md`, `root-cause-analysis.md`                                            |

Perspective and procedure are both common. A large perspective *catalog* — mostly items the agent already produces
unprompted — is better expressed as a derivation procedure (how to enumerate what matters for the code at hand) plus a
compact perspective file holding only the easily-missed long tail; `sketch-feature/references/derive-perspectives.md` +
`missable-checklist.md` is the canonical split. Decision criteria and procedures otherwise appear when an entry-point in
`SKILL.md` needs a self-contained block that would bloat the body. If you find yourself writing perspective +
procedure + templates in one file, split by type.

### File layout

Every file begins with a single header block so an agent arriving via concept search can triage without reading the
body:

1. **Opening line — purpose and trigger.** One sentence stating what the file covers and the phase / task that loads
   it. Combining the two ("テストケース設計・テストコード作成・テストレビュー共通の観点集") is the common shape;
   purpose alone is fine when the trigger is implied by the topic
2. **Related-files block (when part of a cluster).** Files sharing a skill's `references/` directory usually overlap
   in subject. Right after the opening line, add 1–2 lines naming sibling files and the explicit boundary so an
   agent landing here doesn't re-derive what neighbors already cover. Two pieces compose this block:
   - Redirect to siblings: `観点の導出は [derive-perspectives.md](derive-perspectives.md)、エージェント特有の落とし穴は [agent-pitfalls.md](agent-pitfalls.md) を併用します。`
   - Explicit scope exclusion: `本ファイルはテスト設計・実装レビュー時の具体的観点に絞ります。`

   Single-file references (no siblings) skip this block.
3. **Minimum set (when the body exceeds ~80 lines).** Open with 1 sentence framing what's mandatory regardless of
   task (e.g., `対象タスクに関わらず、まず以下を確認します。網羅は <条件> に該当する場合のみ行います。`). Follow with
   3–5 bullets in the form `**ラベル**: 一行説明（→ 「節名」節）`. Sections in the body that apply only on specific
   triggers (e.g., 非機能テスト, 言語別チェック) are marked *conditional* so an agent skips them by default. See
   `sketch-feature/references/testing.md` as a canonical example
4. **Body.** Sections grouped by sub-topic, headed by search-friendly vocabulary

For files at or under ~80 lines the minimum-set block is overhead — just lead with the opening line, optional
related-files block, and the body.

### Maintenance policy

- **No frontmatter.** Loaded on demand via file-path link; no `applyTo` / `description` is needed
- **No section numbers.** Numbering breaks every time a section is added, removed, or reordered. Use heading text and
  in-file anchors for cross-references. Exception: procedure-type files may number the ordered steps themselves
  (e.g., `## 1. 直接原因の特定`) since the step order is intrinsic to the content
- **Granularity.** Keep each entry to a single perspective. Perspective files default to prose, bullets, or lookup
  tables. Inline code snippets stay rare and minimal; long procedures or worked examples belong in their own file
- **Size guidance.** Aim for ≤ 150 lines per file. Beyond that, the agent either full-reads (cost) or scans the
  minimum-set block (which is why ~80 lines is the trigger for adding one). Files over 200 lines are split candidates;
  a large perspective catalog is often better re-expressed as a derivation procedure plus a long-tail safety net (see
  the taxonomy note above)
- **Handling duplication.** When similar perspectives appear across files, keep the shared method in the derivation
  procedure (e.g., `derive-perspectives.md`) and restrict perspective files to the easily-missed long tail. Pair this with the
  related-files block in the file layout above so the boundary is visible at the top of each cluster member
- **Naming.** kebab-case. Within a skill's `references/`, bare topic words (`security.md`, `testing.md`) suffice —
  the skill name already groups them, so the topic-prefix convention used under `.claude/references/` is unnecessary
- **Adding.** Check that the new heading is not semantically duplicating an existing one and that several proposed items
  cannot be consolidated into one perspective
- **Removing.** Search for cross-references from other files before deleting
- **Searchability.** File names and headings use search-friendly vocabulary so an agent can jump to the relevant chapter
  directly

Place files under the skill that most naturally consults them. Cross-skill reference is allowed (relative path); when
the same perspective is needed from multiple skills, keep it in one location and link.

`references/*.md` differs from `.claude/rules/*.instructions.md`: rules are always loaded and govern surface form
(style, review), while `references/*.md` are loaded on demand and cover design considerations.

## Rule Files (`.claude/rules/`)

Files under `.claude/rules/*.instructions.md` are always loaded into every conversation via the global instruction
mechanism. Do not link to them by file path from `SKILL.md`, `environment.md`, `*-mode.md`, or skill-local
`references/*.md` — the rule is already in context, so the link adds noise without adding signal, and the path also
tends to rot when files are renamed or moved.

Unlike `references/*.md`, each rule carries `applyTo: "**"` (plus a short `description`) frontmatter so it ports into
Copilot's `.github/instructions/` unchanged — Copilot requires `applyTo`; Claude ignores it and loads the rule globally
regardless.

When skill content needs to invoke a rule, reference it by concept (e.g., "セルフレビュー", "記法・文体") rather
than file path. The same skill-isolation rule applies to `.claude/references/*.md` (see next section).

## References under `.claude/references/`

On-demand reference material shared across rules and skills (vocabulary tables, severity classifications, writing
examples, terminology guides). Not loaded into context unconditionally — read when a rule or another root reference
links to it, or when an agent concept-matches the topic from inside a skill.

Use this directory when the content is invoked from multiple rules/skills, or is reference material (vocabulary,
lookup tables, examples) rather than a standing directive that must influence every response.

Three placement locations:

- `.claude/rules/*.instructions.md`
  Must influence every response (style, review, no-sycophancy). Always loaded into context.
- `.claude/skills/<name>/references/*.md`
  Only meaningful inside the owning skill (perspective checklists, mode-specific templates).
- `.claude/references/*.md`
  Shared across rules/skills, or vocabulary / lookups needed only in specific authoring spots.

### Link rules (skill isolation)

Skill assets stay decoupled from rules and root references so a skill can be copied or extracted from the template
without dragging in cross-cutting paths. Rules and root references form their own link graph.

#### ❌ Forbidden — skill isolation

- `.claude/skills/<name>/**/*` → `.claude/rules/*`
- `.claude/skills/<name>/**/*` → `.claude/references/*`

Use concept-name references instead (see below).

#### ❌ Avoid — target already in context

The target is already loaded into context, so a file-path link adds noise without adding signal.

- `.claude/references/*` → `.claude/rules/*`

#### ✅ OK

- `.claude/rules/*` → `.claude/rules/*`
  Both targets are always-loaded, but the link helps maintainers see the relationship and keeps the resolution explicit for the agent. The token cost is small.
- `.claude/rules/*` → `.claude/references/*`
  Rule pulls in shared vocabulary or lookup.
- `.claude/references/*` → `.claude/references/*`
  Root references can cross-link freely.
- `.claude/skills/<a>/**/*` → `.claude/skills/<b>/references/*`
  Perspective sharing across skills (see cross-references section below).

When a skill needs the content of a rule or root reference, reference it by **concept name** (the heading or topic
phrase, e.g., "レビュー深刻度基準", "語彙のトーン", "「技術」とツール/サービスの区別")
and leave path resolution to the agent. The vocabulary itself (e.g., `Critical / Major / Minor / Info`) may be
inlined at the point of use when the reader needs only the labels rather than the full reference.

Do **not** introduce repo-external lookup mechanisms (the agent's `memory/` directory, `agent-memory-local/`, etc.)
to compensate for the missing path. This repository is a template that downstream projects copy as-is; anything
outside the tracked tree is unavailable to those copies.

### Authoring conventions

- **No index file.** Do not create a table-of-contents file under `.claude/references/`. The link graph among rules
  and root references is the index
- **One topic per file.** Split rather than combine when a file starts to mix purposes (e.g., process vs. examples)
- **No frontmatter.** Loaded on demand; no `applyTo` / `description` is needed
- **Opening sentence states when to open it.** Each file begins with a one-line "when this is consulted" so an agent
  arriving via concept search can confirm relevance before reading the body
- **Naming.** kebab-case; group with a topic prefix (e.g., `writing-*`) when related files are likely to grow alongside
  it. Files without natural siblings stay bare (e.g., `wording.md`)

## `environment.md` Structure

Use exactly these two sections. **Do not add a `## Common` section** (the deploy-time removal of one side would leave
that content orphaned).

- `## Claude Code`
- `## GitHub Copilot (CLI / VSCode)`

Write here:

- Means that differ per environment (e.g., structured questioning, parallel exploration)
- User-specific local paths / cache locations (even if identical in both environments, write them in both sections)

Don't write here:

- Environment-independent policy (→ `SKILL.md`, which is always kept regardless of deploy)
- The skill's main flow (→ `SKILL.md`)

"Both environments use the same value, so let's centralize it" is not the right reasoning. Duplication across the two
sections is intentional (one section is dropped on deploy).

## `SKILL.md` Body

- Do not write agent-specific tool names (`Agent`, `AskUserQuestion`, `TaskCreate`, `WebFetch`, `Skill`, etc.) in the
  body. Replace with neutral wording and concentrate concrete means in `environment.md`
- For the neutral-wording / environment-tool-name correspondence,
  see [the tool-name mapping](claude-vs-copilot.md#tool-names) in claude-vs-copilot.md
- A single reference to `environment.md` at the top is easy for the LLM to skip past. Add
  `[environment.md](environment.md)` pinpoint links at every environment-dependent passage
- **Within the same file**, link the same concept only once (at first occurrence or in an upfront definition). Repeating
  the link in every cell of a table or in successive bullets is noise
- **Supplementary files** in the same skill (`*-mode.md`, `references/*.md`, etc.) inherit terminology from `SKILL.md`.
  Don't re-link to `environment.md` for concepts already defined there unless the supplementary file introduces a new
  env-dependent term
- **Overview vs procedure**: the top paragraph (before any `##` heading) states *what the skill is* — purpose, scope,
  output artifacts. Do not embed procedural detail (the order of steps, where user confirmation occurs, how phases
  interact) in the overview — that belongs in the 手順 sections. Procedure stated in both places duplicates content
  and makes the overview hard to skim. The overview should answer "what is this skill," not "how does it run"

## Skill `README.md` Body

`README.md` ships with `.claude/` and is user-facing — it shows what the skill is and a concrete invocation sample.
Follow this shape:

- `# <Display Name>` (one-line title)
- An overview paragraph immediately after the title — one or two short sentences describing what the skill is, ending
  in a noun that names the skill (in JA: `...スキル`). The overview lives here and only here
- `## 実行例` containing a fenced block with `❯ /<name> <example argument>`, an `● <agent intro line>`, and an indented
  output sample. Put no prose between the heading and the fenced block — the example itself communicates the case, and
  a case-description preamble duplicates the overview above
- For multi-mode skills, split the section into `### <mode name>` subsections, each holding its own fenced block

Keep examples concrete but generic — template purity applies, so don't encode jumpstarter-specific scenarios.

## YAML Frontmatter

To share files across environments, fields fall into three groups.
See [This Repository's Frontmatter Policy](claude-vs-copilot.md#this-repositorys-frontmatter-policy) in
claude-vs-copilot.md for the full list and reasoning. Key points:

- Do not write `tools` / `allowed-tools` (schemas are incompatible). Narrowing is unneeded under devcontainer +
  auto/edit mode
- `description` is the **auto-invocation trigger**. Include trigger and non-trigger conditions
- `name` is required on skills and must match the directory name

## Where to Put What

Decide based on which phase you want to affect.

- **Invocation decision** (before the skill starts): only the frontmatter `description` is read. Trigger conditions,
  examples, and exclusions must go in `description`
- **Post-invocation behavior**: write directly in the `SKILL.md` body
- **`references/`**: only read when the body explicitly references them; has no effect on invocation. Split off only
  when multiple skills share the reference or the body has grown too large

## Writing `description`

`description` is the **auto-invocation matcher** — it sits in the system prompt and is read on every user turn. Every
line is permanent context overhead, so treat it like a description gate: each line must help decide whether the skill
fires.

### Step 1: decide auto-invocation vs explicit-only

Before writing the description, decide whether auto-invocation is needed at all:

- **Auto-invocation**: the user may invoke without typing `/<name>`. The skill must match natural-language prompts (
  e.g., "rename X" → `fix-well`).
- **Explicit-only**: the skill is heavyweight, costly, or only useful with specific arguments / commands (e.g., a
  multi-agent thinker, a URL-driven importer). Set `disable-model-invocation: true` and write a single short sentence —
  the user types `/<name>` to invoke.

Auto-invocation is opt-in. If natural-language matching brings little value, prefer explicit-only to avoid false
positives and description bloat.

**Caveat — inter-skill calls**: `disable-model-invocation: true` removes the skill from the available-skills list
entirely, so other skills can't call it via the Skill tool either. If a skill is meant to be invoked from another
skill's body via the Skill tool, do **not** set `disable-model-invocation: true`. Workarounds when you still want to
suppress auto-fire:

1. Leave `disable-model-invocation` unset and keep the description tight (one short sentence) — the matcher cost is
   modest and the skill stays callable.
2. Reference the callee via a Markdown link to its `SKILL.md` (e.g., `[callee-skill](../callee-skill/SKILL.md)
   を読んで...`) instead of "execute the X skill". The caller's Claude reads the file and follows it; this works
   regardless of `disable-model-invocation`.

### Step 2: write the description

For an auto-invocable skill, include:

- **Purpose** (required): what the skill achieves, in one short clause
- **Trigger conditions** (required): which prompts / situations should activate it, with concrete keyword examples
- **Non-trigger conditions** (recommended when ambiguity exists): tempting but should not activate. Add this when
  trigger keywords overlap with another skill or with a "default" workflow; omit when the trigger is self-evident

For an explicit-only skill, one short sentence describing what it does for the user typing `/<name>` is enough. Do
**not** restate "explicit-only" / "no auto-fire" / 「明示呼び出し専用」 in the description text —
`disable-model-invocation: true` already conveys that to the matcher, and the extra prose only adds permanent context
overhead. Also drop trigger / non-trigger examples (e.g., 「発動: ...」 / 「非発動: ...」): they exist to steer
auto-invocation, and the matcher is disabled.

### Step 3: remove execution detail

Every remaining line must help auto-invocation matching. After drafting, delete:

- Implementation steps (verbs describing what the skill does: "creates X", "outputs Y", "performs Z")
- Output file names, internal phase names, or other post-invocation behavior
- Effect / outcome statements ("prevents the user from coming back with X", "saves Y rounds") — they describe results,
  not triggers
- Restating the purpose in multiple ways

Test for each line: "If I deleted this line, would the matcher miss prompts it should catch, or catch prompts it
shouldn't?" If neither, delete the line — its content belongs in `SKILL.md` body.

Example:

> For change requests ("fix it", "rename", "unify"), consider impact scope, similar issues, and consistency with
> callers and docs, pinning the interpretation down first when the request is ambiguous. Also covers fix-style
> instructions issued during feature work or review. Does not trigger for knowledge questions or new creation from
> scratch.

## Auto-Invocation and Explicit Invocation

Both environments auto-invoke on a `description` match. Explicit `/agent` or `/fleet` in Copilot is supplementary.

Explicit invocation uses `/<name>`. Notes when writing:

- **Don't call it "slash command"** — that term refers to built-ins (`/login`, `/clear`, `/help`, etc.) and conflicts
  with skill `/<name>` invocation
- **Avoid hardcoding the skill's own name** — renames cause drift. In internal docs like `environment.md`, write
  `/<name>` as a placeholder
- README usage examples may use the concrete name (for copy-pasteable demos)

## Cross-Skill References

Use these mechanisms to reference other skills:

- **File reference**: relative-path link like `[asking principles](../fix-well/references/ask-the-user.md)`.
  Renames are followed by updating the path
- **`/<name>` invocation**: when describing "call another skill" in `environment.md`, write the `/<name>` form or pair
  it with a file-path link rather than hardcoding a concrete name

Additional constraints:

- Don't introduce concepts foreign to the skill in the `SKILL.md` summary or body without context
- In any file under `.claude/skills/` (`SKILL.md` or supporting `.md`, **including YAML frontmatter such as
  `description`**), don't reference another skill by name unless there's a call relation or a clear file dependency.
  When you want a `description` to express a trigger boundary against another skill, describe what *this* skill does
  and doesn't fire on — don't name the other skill (e.g., 「個別箇所の修正指示は発動しない」, not 「(fix-well
  が優先)」)

### Reference direction: depended-on skills stay caller-agnostic

A skill that is depended on by other skills must **not** name those callers in its own `SKILL.md` / `environment.md` /
supplementary `*-mode.md`. Caller lists rot every time a new caller appears or an existing one is renamed; the caller
side already links forward, so the depended-on side does not need a mirror reference.

- ❌ Bad — `sketch-feature/SKILL.md` lists `kaizen` / `hyper-think` / `import-best-practices` as "skills that consult my references"
- ✅ Good — each caller links to `sketch-feature/references/*.md` from its own files; `sketch-feature` stays caller-agnostic

Apply this even when the caller list looks stable. The goal is that a depended-on skill can be invoked, replaced, or
copied into another project without dragging in caller assumptions.

### `references/*.md` cross-references are explicitly allowed

`references/*.md` files manage **one concern per file** and are designed to be shared. Cross-skill references that
target `references/*.md` are allowed regardless of the runtime call direction — this is the intended sharing
mechanism for shared references.

- ✅ `hyper-think/review-mode.md` → `sketch-feature/references/derive-perspectives.md` (reference reuse, not entry-point call)
- ✅ `kaizen/SKILL.md` body → `sketch-feature/references/verification-strategy.md` (reference reuse, not entry-point call)
- ❌ `kaizen/SKILL.md` body → `sketch-feature/SKILL.md` body (entry-point coupling, not a reference share)

The single-concern-per-file structure makes `references/*.md` stable link targets. The "stay caller-agnostic at
the entry point" rule above still applies to `SKILL.md` / `environment.md` / `*-mode.md` bodies.

## Review Checklist

When editing or reviewing PRs:

| Aspect                     | Check                                                                                                                           |
|----------------------------|---------------------------------------------------------------------------------------------------------------------------------|
| `environment.md` structure | Only `## Claude Code` and `## GitHub Copilot (CLI / VSCode)`; no `## Common`                                                    |
| `environment.md` purity    | Items that have the same means in both environments are moved to `SKILL.md`; only env-specific items remain                     |
| Frontmatter                | No `tools` / `allowed-tools` slipped in                                                                                         |
| Self-references            | The skill's own name is not hardcoded in `environment.md` etc.                                                                  |
| Reference direction        | Depended-on skills don't name their callers in `SKILL.md` / `environment.md`; `references/*.md` cross-refs are OK in both ways  |
| `description`              | (1) auto vs explicit-only decided; (2) purpose + trigger only, no execution detail; (3) non-trigger added when ambiguity exists |
| Terminology                | "Slash command" used only for built-ins, not for skill `/<name>`                                                                |
| Impact scope               | When renaming identifiers / paths, references (links, examples, other skills) follow                                            |
| Rule application           | Relevant `.claude/rules/*.instructions.md` and `.claude/references/*.md` are applied to the edit itself                         |
| Skill isolation            | No file-path link from skill assets to `.claude/rules/*` or `.claude/references/*`; concept-name references used instead        |

## Verification

Official specs change. When adopting new fields or tools:

- Verify against official docs via web fetch
- Mark unverifiable items with `❓`
- Update the confidence marks (✅ / ⚠️ / ❓) in [claude-vs-copilot.md](claude-vs-copilot.md)

## Related Files

- [claude-vs-copilot.md](claude-vs-copilot.md): feature / config / placement / tool-name / frontmatter mapping
- `.claude/skills/*/environment.md`: examples of environment-specific content
- `.claude/skills/*/SKILL.md`: examples of skill bodies

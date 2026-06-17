# Sensitive Information Review

Checklist for reviewing tracked files (mainly `.claude/`, `.devcontainer/`, and the root `*.md`) before publishing
changes. This repository is a **template**: anything committed propagates into every downstream clone, so the bar for
"safe to ship" is stricter than a typical application repo.

Excluded by `.gitignore` (no need to inspect): `.env*`, `*.key`, `*.pem`, `*.local.json`, `.DS_Store`, `.idea/`,
`.vscode/`, `.local/`, `.claude/agent-memory-local/`. Confirm coverage with `git check-ignore <path>`.

## What Counts as Sensitive Here

| Category                | Examples to flag                                                                         | Allowed exception                                                                  |
|-------------------------|------------------------------------------------------------------------------------------|------------------------------------------------------------------------------------|
| Credentials             | API keys, tokens, passwords, private keys, OAuth client secrets                          | The words alone, used as security-concept terminology in skill/agent docs          |
| Personal identifiers    | Real names, real email addresses, phone numbers, account handles other than `raiich`     | `raiich/jumpstarter` (this repo's own public URL)                                  |
| Maintainer-local state  | Absolute paths under `/Users/<name>/` or `/home/<name>/`, local IDE configs, host names  | `/Users/username/...` and similar explicit placeholders                            |
| Internal references     | Private hostnames, internal project codenames, employer/customer mentions                | Generic business terms ("customer base", "stakeholders") in analysis-agent prompts |
| Real IPs                | Public IPs, internal RFC 1918 ranges (10.x, 172.16-31.x, 192.168.x)                      | RFC 5737 / RFC 3849 documentation ranges                                           |
| Real URLs               | Internal wikis, private repos, unpublished service endpoints                             | Public docs (Anthropic, GitHub, Microsoft, Debian, RFC-allocated example domains)  |
| Placeholder hygiene     | Slang, real-resource look-alikes, copy-paste leftovers from other projects               | The placeholder set defined in `.claude/rules/writing-style.instructions.md`       |
| Dev markers             | `TODO` / `FIXME` / `XXX` that leak unfinished context or names                           | Documentation that explains the marker itself                                      |

Placeholder convention (authoritative source: `.claude/rules/writing-style.instructions.md`): name = `foo`, IPv4 =
`192.0.2.x`, IPv6 = `2001:db8::/32`, domain = `example.com`, email = `user@example.com`.

## Pre-Vetted Public References

The following appear in tracked files and have already been judged safe. Re-flagging them on each review wastes time;
flag only new occurrences outside this set.

- `api.anthropic.com`, `code.claude.com`, `claude.ai/install.sh` — Claude Code installation / endpoint
- `mcr.microsoft.com/devcontainers/base:debian` — public base image
- `docs.github.com`, `raw.githubusercontent.com`, `github.blog`, `agents.md`, `code.visualstudio.com` — public docs
- `deb.debian.org`, `www.debian.org`, `bugs.debian.org` — public Debian mirrors
- `just.systems`, `playwright.dev` — public tool documentation (devcontainer features)
- `example.com` — RFC 2606 documentation domain (sample value in docs/rules)
- `github.com/raiich/jumpstarter` — this repository's own clone URL
- `/home/vscode/.claude`, user `vscode` — devcontainer convention, not a real account

## Review Procedure

Run these against tracked files only. Each pattern is intentionally noisy — expect to triage hits against the
allowlist above.

```bash
# 1. Credentials and key material
git ls-files | xargs grep -nEi \
  '(secret|token|api[_-]?key|password|passwd|credential|bearer|BEGIN (RSA|OPENSSH|EC|PRIVATE) (PRIVATE )?KEY)' 2>/dev/null

# 2. Personal identifiers and real emails
git ls-files | xargs grep -nE \
  '@[A-Za-z0-9.-]+\.(com|jp|net|org|io|dev|ai)\b' 2>/dev/null \
  | grep -vE '(example\.com|user@example|noreply@)'

# 3. Local absolute paths
git ls-files | xargs grep -nE '(/Users/|/home/[a-zA-Z]+)' 2>/dev/null \
  | grep -vE '(/Users/username|/home/vscode)'

# 4. Private / internal IPs
git ls-files | xargs grep -nE \
  '\b(10\.|172\.(1[6-9]|2[0-9]|3[01])\.|192\.168\.|127\.0\.0\.|0\.0\.0\.0)' 2>/dev/null

# 5. All outbound URLs (review the unique set against the allowlist)
git ls-files | xargs grep -nE 'https?://[a-zA-Z0-9./_-]+' 2>/dev/null | sort -u

# 6. Forgotten dev markers
git ls-files | xargs grep -nE '\b(TODO|FIXME|XXX|HACK|TEMP)\b' 2>/dev/null

# 7. Verify ignored paths really are ignored
git check-ignore .claude/agent-memory-local .local .idea
```

## Where Leakage Typically Appears

- `*.agent.md` / `SKILL.md` examples — security-concept examples may drift into using real values
- `.devcontainer/devcontainer.json` — `name`, IDE backend, mount paths often carry maintainer defaults
- `environment.md` — per-environment paths can capture the author's specific machine layout
- Any new `docs/development/*.md` that quotes log output or transcripts — likely to embed local paths

## When You Find Something

1. Replace with the placeholder set from `writing-style.instructions.md`.
2. If the value was previously committed, treat `git rm` + new commit as **not sufficient** for true secrets — rotate
   the credential at its source.
3. If the finding is a new public reference (e.g., a newly added official docs domain), add it to the **Pre-Vetted
   Public References** section above so future reviews skip it.

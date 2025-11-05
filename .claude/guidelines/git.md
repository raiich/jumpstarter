# Git Workflow

## Commit Message Format

```bash
git commit -m "$(cat <<'EOF'
[Verb] [Concise description]

[Optional: Detailed explanation in English]

Changes:
- [Change 1]
- [Change 2]

🤖 Generated with [Claude Code](https://claude.com/claude-code)

Co-Authored-By: Claude <noreply@anthropic.com>
EOF
)"
```

## Basic Rules

**Subject Line:**
- Start with a verb: `Fix`, `Add`, `Refactor`, `Update`, `Replace`, etc.
- Keep concise (40-80 characters)
- No period at the end

**Body:**
- Use English primarily
- Use bullet points in `Changes:` / `Fixes:` / `Benefits:` sections
- File-level details (when multiple files are changed)
- Metrics and verification results (when applicable)

**Footer (Required):**
- Claude Code attribution
- Co-Authored-By

## Example

```
Replace all _ = patterns with global variable assignments

Replaced all `_ = result` patterns with global variable assignments in all benchmarks.

Changes:
- time_bench_test.go: Fixed 8 locations
- context_bench_test.go: Fixed 10 locations
- allocation_bench_test.go: Fixed all locations

🤖 Generated with [Claude Code](https://claude.com/claude-code)

Co-Authored-By: Claude <noreply@anthropic.com>
```

## Pre-Commit Checks

```bash
git status && git diff --stat
go build ./... && go test -c ./...
git add <files> && git status
```

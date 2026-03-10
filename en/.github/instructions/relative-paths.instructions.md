# Path Specification Rules

## Principle: Use Relative Paths in Bash Commands

When executing commands with the Bash tool, specify file and directory paths relative to the working directory.

```bash
# ✅ Good
go test ./pkg/...
cat README.md
git diff src/main.go

# ❌ Bad
go test /Users/username/project/pkg/...
cat /Users/username/project/README.md
git diff /Users/username/project/src/main.go
```

## Exceptions

- Read / Edit / Write tools: Absolute paths are required by specification, so use them as-is

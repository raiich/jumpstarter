# Bash Command Rules

## Use Relative Paths

Specify file and directory paths relative to the working directory.

```bash
# ✅ Good
go test ./pkg/...
cat README.md

# ❌ Bad
go test /Users/username/project/pkg/...
cat /Users/username/project/README.md
```

**Exception:** Read / Edit / Write tools require absolute paths by specification, so use them as-is.

## Use git rm for File Deletion

Delete tracked files with `git rm` to stage the deletion automatically.

```bash
# ✅ Good
git rm path/to/file.txt
git rm -r path/to/directory/

# ❌ Bad
rm path/to/file.txt
rm -r path/to/directory/
```

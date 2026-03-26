# Bash コマンドルール

## 相対パスを使用する

ファイルやディレクトリのパスは作業ディレクトリからの相対パスで指定する。

```bash
# ✅ 良い例
go test ./pkg/...
cat README.md

# ❌ 悪い例
go test /Users/username/project/pkg/...
cat /Users/username/project/README.md
```

**対象外:** Read / Edit / Write ツールは仕様上、絶対パスが必須のためそのまま使用。

## ファイル削除は git rm を使用する

Git 追跡下のファイルは `git rm` で削除し、削除を直接ステージングする。

```bash
# ✅ 良い例
git rm path/to/file.txt
git rm -r path/to/directory/

# ❌ 悪い例
rm path/to/file.txt
rm -r path/to/directory/
```

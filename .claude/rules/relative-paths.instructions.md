# パス指定ルール

## 原則: Bash コマンドでは相対パスを使用する

Bash ツールでコマンドを実行する際、ファイルやディレクトリのパスは作業ディレクトリからの相対パスで指定する。

```bash
# ✅ 良い例
go test ./pkg/...
cat README.md
git diff src/main.go

# ❌ 悪い例
go test /Users/username/project/pkg/...
cat /Users/username/project/README.md
git diff /Users/username/project/src/main.go
```

## 対象外

- Read / Edit / Write ツール: 仕様上、絶対パスが必須のためそのまま使用

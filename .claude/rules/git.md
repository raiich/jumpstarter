# Gitワークフロー

## コミットメッセージフォーマット

```bash
git commit -m "$(cat <<'EOF'
[動詞] [簡潔な説明]

[オプション: 日本語での詳細説明]

Changes:
- [変更1]
- [変更2]

🤖 Generated with [Claude Code](https://claude.com/claude-code)

Co-Authored-By: Claude <noreply@anthropic.com>
EOF
)"
```

## 基本ルール

**サブジェクトライン:**
- 動詞で始める: `Fix`, `Add`, `Refactor`, `Update`, `Replace`, etc.
- 簡潔に (40-80文字)
- 末尾にピリオドなし

**ボディ:**
- 日本語を優先的に使用
- `Changes:` / `Fixes:` / `Benefits:` セクションで箇条書き
- ファイル単位の詳細（複数ファイル変更時）
- メトリクスや検証結果（該当時）

**フッター（必須）:**
- Claude Code attribution
- Co-Authored-By

## 例

```
Replace all _ = patterns with global variable assignments

全てのベンチマークで `_ = result` パターンをグローバル変数への代入に置き換えました。

Changes:
- time_bench_test.go: 8箇所修正
- context_bench_test.go: 10箇所修正
- allocation_bench_test.go: 全箇所修正

🤖 Generated with [Claude Code](https://claude.com/claude-code)

Co-Authored-By: Claude <noreply@anthropic.com>
```

## コミット前の確認

```bash
git status && git diff --stat
go build ./... && go test -c ./...
git add <files> && git status
```

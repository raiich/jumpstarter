---
name: auto-reviewer
description: Automatically review Go benchmark code after implementation. Checks compliance with docs/writing-benchmarks.md standards.
allowed-tools: Read, Grep, Glob
---

# Auto Code Reviewer

ベンチマーク実装後、`docs/writing-benchmarks.md` の基準に基づいて自動レビューを実行します。

## Instructions

1. `docs/writing-benchmarks.md` を参照してレビュー基準を確認
2. 最近編集された `*_test.go` ファイルを特定
3. ドキュメントのチェックリストに基づいてレビュー実行

## Output Format

```
## ベンチマークレビュー結果

### ✓ 適合
- グローバル変数への代入
- goroutine クリーンアップ

### ✗ 要修正
- file.go:45 - グローバル変数への代入なし
- file.go:89 - globalResult → GlobalResult（エクスポート必須）

📚 詳細: docs/writing-benchmarks.md
```

## Notes

- Read, Grep, Glob のみ使用（変更不可）
- 問題発見時はファイル名:行番号で報告
- ドキュメント参照を促す

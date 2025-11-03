---
name: doc-checker
description: Automatically check documentation-code consistency after implementation. Verifies README.md and docs/*.md match actual code.
allowed-tools: Read, Grep, Glob
---

# Documentation Consistency Checker

コード実装後、ドキュメントと実装の整合性を自動確認します。

## Instructions

1. README.md と docs/*.md を参照して記載内容を確認
2. 最近編集された Go ファイルを特定
3. 以下を検証：
   - README.md に記載されたベンチマークカテゴリが実装されているか
   - 実装された機能がドキュメントに記載されているか
   - docs/*.md のリンクが有効か

## Output Format

```
## ドキュメント整合性チェック結果

### ✓ 整合性あり
- README.md のカテゴリ記載と実装が一致
- 新機能がドキュメント化済み

### ✗ 要確認
- README.md:28 - 記載されているが未実装: "XXX機能"
- file.go:123 - 実装されているが未ドキュメント: BenchmarkNewFeature

📚 参照: README.md, docs/*.md
```

## Notes

- Read, Grep, Glob のみ使用（変更不可）
- 問題発見時はファイル名:行番号で報告
- ドキュメント参照を促す

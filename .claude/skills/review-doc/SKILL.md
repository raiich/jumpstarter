---
name: review-doc
description: Review all .md files for naming conventions, conciseness, consistency, and completeness. Use when creating or editing any markdown file.
allowed-tools: Read, Grep, Glob
---

# Documentation Reviewer

マークダウンドキュメントの品質をレビューします。

## レビュー観点

`.claude/rules/documentation.instructions.md` に基づき、以下の観点でレビューを実施：

- **命名規則**: ファイル名が kebab-case か
- **簡潔さ**: 冗長な説明、重複情報、コードで代替可能な長文
- **構成**: 章立て・階層の妥当性
- **文体**: ですます調の統一、用語の一貫性
- **整合性**: 実装との一致、コード例の正確さ、リンクの有効性

## Output Format

```
## ドキュメントレビュー結果

### ✓ 適合
- [確認した観点]

### ✗ 要修正
- file.md:行番号 - [カテゴリ] 問題の説明
  修正提案: [具体的な修正内容]
```

## Notes

- Read, Grep, Glob のみ使用（変更不可）
- 問題発見時はファイル名:行番号で報告
- 過検出より見逃し防止を優先

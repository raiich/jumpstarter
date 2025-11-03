---
name: doc-reviewer
description: Review markdown documentation for naming conventions, conciseness, consistency, and completeness. Use when creating or editing documentation files.
allowed-tools: Read, Grep, Glob
---

# Documentation Reviewer

マークダウンドキュメント作成・編集後に品質をレビューします。

## Instructions

1. `.claude/guidelines/documentation.md` を参照してレビュー基準を確認
2. 対象となるマークダウンファイルを特定
3. 以下の順序でレビュー実行：
   - 命名規則チェック
   - 簡潔さチェック
   - 一貫性チェック
   - 完全性チェック

## Review Checklist

### 1. 命名規則チェック
- kebab-case準拠（例外: README.md）
- ファイル名が内容を適切に表現

### 2. 簡潔さチェック
- 過度な説明がないか
- 不要な装飾（マーカー、Phase/Stepなど）がないか
- 重複情報がないか

### 3. 一貫性チェック
- 自己矛盾（ドキュメント自身がルールに違反）
- 他のガイドラインとの矛盾
- 用語の統一

### 4. 完全性チェック
- 必要なセクションの存在
- リンク切れ
- 実装例の有無（必要な場合）

## Output Format

```
## ドキュメントレビュー結果

### ✓ 適合
- kebab-case命名規則
- 簡潔な記述

### ✗ 要修正
- file.md - UPPERCASE使用（kebab-case必須）
- file.md:15 - 冗長な説明
- file.md:28 - 自己矛盾: 装飾禁止を説くが自身で使用

📚 参照: .claude/guidelines/documentation.md
```

## Notes

- Read, Grep, Glob のみ使用（変更不可）
- 問題発見時はファイル名:行番号で報告
- ガイドライン参照を促す

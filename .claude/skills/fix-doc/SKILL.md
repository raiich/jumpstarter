---
name: fix-doc
description: Properly fix documentation issues. Ensures consistency between documentation and code, and applies quality standards.
allowed-tools: Read, Grep, Glob, Edit, Write
---

# Documentation Fix Helper

ドキュメント修正を適切に実施するスキルです。`.claude/rules/fix-guidelines.instructions.md` の基本原則に従い、完全で正確な修正を行います。

## 修正フロー

`.claude/rules/fix-guidelines.instructions.md` の共通フローに従ってください。

## ドキュメント修正固有のガイドライン

### タイトル・見出しの修正

1. 修正対象のタイトルが他で使われていないか確認
2. リンク（`#anchor` 形式）が変わる場合、参照元も更新
3. 全て一度に修正

### コード例の修正

1. コード例が実装と一致しているか確認
2. 同じコード例が他のドキュメントで使われていないか検索
3. 参考資料・関連ドキュメントも更新が必要か確認
4. 全て一度に修正

### 用語・表現の修正

1. 修正対象の用語が他で使われていないか全検索
2. 一貫性のために似た表現も修正対象か判断
3. 全て一度に修正

### 機能説明の修正

1. 説明内容が実装と一致しているか確認
2. 関連する複数のドキュメント（README、docs/*.md など）を確認
3. 実装の変更に伴う関連説明も更新が必要か確認
4. 全て一度に修正

## セルフレビュー（ドキュメント固有）

修正後、以下を確認：

- [ ] 全ての該当箇所を修正したか
- [ ] 修正内容は正確か
- [ ] `.claude/rules/writing-style.instructions.md` に従っているか
- [ ] 実装とドキュメントの整合性はあるか
- [ ] リンクが有効か（参照箇所が存在するか）
- [ ] 冗長な部分はないか

## 簡潔さのチェック

各修正について以下を確認：

- [ ] 必要な情報のみか（冗長な説明はないか）
- [ ] 前提知識の説明は避けているか（読者が既に知っている可能性）
- [ ] 重複情報がないか
- [ ] コードで説明できることを長文で書いていないか
- [ ] 説明が長い場合、半分に削れないか検討した

## ドキュメント修正の参考例

```markdown
❌ 悪い例：1ファイルだけ修正
- README.md で "新機能" という用語を修正

✅ 良い例：関連ドキュメント全て修正
1. Grep ツールで "新機能" を全検索して箇所を特定
2. README.md, docs/feature-guide.md, docs/api.md を同時修正
3. コード例も確認して最新版を反映
4. リンク切れがないか確認
```

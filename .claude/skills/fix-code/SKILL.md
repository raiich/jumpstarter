---
name: fix-code
description: Properly fix code issues by following the quality assurance guidelines. Ensures comprehensive fixes across all related locations.
allowed-tools: Read, Grep, Glob, Edit, Write
---

# Code Fix Helper

コード修正を適切に実施するスキルです。`.claude/rules/fix-guidelines.instructions.md` の基本原則に従い、包括的かつ効率的に修正を行います。

## 修正フロー

`.claude/rules/fix-guidelines.instructions.md` の共通フローに従ってください。

## コード修正固有のガイドライン

### コードの削除・リファクタリング

1. 削除対象のパターンを全検索（Grepで確認）
2. 影響するテスト・ドキュメントを特定
3. 全ての関連箇所を一度に修正
4. ビルド・テスト実行

### 変数名・関数名の変更

1. 変更対象の全使用箇所を検索（Grep + file type filter）
2. 参照ドキュメントも確認
3. 全て一度に変更
4. ビルド確認

### エラーハンドリング・ロジック修正

1. 同じパターンが他にもないか検索
2. 関連する呼び出し箇所を確認
3. 全て一度に修正
4. テスト実行

## セルフレビュー（コード固有）

修正後、以下を確認：

- [ ] 全ての該当箇所を修正したか
- [ ] 修正内容は正しいか
- [ ] ビルドエラーはないか
- [ ] 全テストが通るか
- [ ] 既存のコードスタイルに従っているか
- [ ] ドキュメントも更新が必要か

## 実装の参考例

```
❌ 悪い例：1箇所だけ修正
- 指摘された1箇所のみ修正

✅ 良い例：全箇所を検索して修正
1. Grep ツールで同じパターンを全箇所検索
2. 見つかった全ファイルの同じパターンを一度に修正
3. 関連テストも修正
4. ビルド・テスト実行
```

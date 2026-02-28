---
name: review-code
description: Review code for quality issues and documentation consistency. Checks security, error handling, over-engineering, duplication, and doc-code alignment.
allowed-tools: Read, Grep, Glob
---

# Code Reviewer

コード実装後、コード品質とドキュメント整合性をレビューします。

## Instructions

### 1. レビュー対象の特定

最近編集されたソースファイルを特定し、変更内容を把握する。

### 2. コード品質レビュー

以下の観点でレビューを実施：

**セキュリティ**
- ハードコードされた秘密情報（APIキー、パスワード、トークン）
- インジェクション脆弱性（SQL、コマンド、XSS）
- 入力値の未検証（外部入力、ユーザー入力）

**エラーハンドリング**
- エラーの握りつぶし（無視、空catch）
- エラー情報の欠落（原因不明のエラーになる）
- リソースリーク（close漏れ、defer忘れ）

**過剰な実装**
- 不要な抽象化（1回しか使わないのにインターフェース化）
- 使われていないコード（未使用の変数、関数、import）
- 過度な設定可能性（現時点で不要なオプション）

**コードの重複**
- 同一・類似ロジックのコピペ
- 共通化すべきパターンの散在

### 3. ドキュメント整合性チェック

- README.md や .local/docs/*.md の記載内容と実装の一致
- 実装された機能がドキュメントに反映されているか
- ドキュメント内のリンクが有効か

## Output Format

```
## コードレビュー結果

### コード品質

#### ✓ 問題なし
- [確認した観点]

#### ✗ 要修正
- file.go:123 - [カテゴリ] 問題の説明
  提案: [具体的な修正内容]

### ドキュメント整合性

#### ✓ 整合性あり
- [確認した内容]

#### ✗ 要確認
- README.md:28 - 問題の説明
  提案: [具体的な修正内容]
```

## Notes

- Read, Grep, Glob のみ使用（変更不可）
- 問題発見時はファイル名:行番号で報告
- 過検出より見逃し防止を優先

# Claude Code 自動化実装計画

## 目的

Claude Codeが変更を行う際に、自動的にプラン作成・レビュー・ドキュメント整合性チェックを実行する仕組みを構築する。

## 実装方針

### 自動化の範囲

**Claude Codeが変更を行う場合**（完全自動）:
- 変更前: 自動でプラン作成（TodoWrite + ExitPlanMode）
- 変更後: 自動でレビュー + ドキュメント整合性チェック

**ユーザーが変更を行う場合**:
- Skillsによる自動チェックに依存

## 実装するコンポーネント

### 1. Skills（自動起動）

#### 1-1. code-reviewer Skill
- **ファイル**: `.claude/skills/code-reviewer/SKILL.md`
- **目的**: ドキュメントと実装の整合性を自動確認
- **機能**:
  - README.md と実装の一致確認
  - docs/*.md の記述と実装の照合
  - 未ドキュメント機能の検出
  - 実装されていない記載の検出
- **ツール制限**: Read, Grep, Glob のみ
- **起動タイミング**: コード実装完了後、自動起動

#### 1-2. doc-reviewer Skill
- **ファイル**: `.claude/skills/doc-reviewer/SKILL.md`
- **目的**: マークダウンドキュメントの品質チェック
- **機能**:
  - 命名規則チェック（kebab-case）
  - 簡潔さチェック（冗長性、不要な装飾）
  - 一貫性チェック（自己矛盾、ガイドライン間の矛盾）
  - 完全性チェック（必要セクション、リンク切れ）
- **ツール制限**: Read, Grep, Glob のみ
- **起動タイミング**: ドキュメント作成・編集後、自動起動

### 2. Claude の振る舞い設定

#### 2-1. settings.json の拡張
- **ファイル**: `.claude/settings.json`（既存ファイルに追記）
- **目的**: 私（Claude）の自動プラン作成を有効化
- **内容**:
  - 変更前の自動プラン作成設定
  - TodoWrite の自動使用
  - ExitPlanMode の自動使用
  - Skillsとの連携方針

## 実装順序

1. **Skills の作成**（最優先）
   - code-reviewer Skill
   - doc-reviewer Skill

2. **settings.json の更新**
   - Claude の振る舞い設定を追記

3. **動作確認**
   - テスト実装を実行して自動起動を確認
   - Skills が適切に動作するか検証

## 期待される動作フロー

### シナリオ1: 新機能実装

```
[ユーザー] 新しい機能Xを実装してください

↓ 自動実行

[Claude] プランを作成（TodoWrite）
[Claude] 実装計画を提示（ExitPlanMode）

↓ ユーザー承認

[Claude] 実装実行（Edit/Write）

↓ 実装完了後、自動実行（並列）

[code-reviewer Skill] ドキュメント整合性チェック
[doc-reviewer Skill] ドキュメント品質チェック

↓ 自動レポート

[Claude] チェック結果を報告
  ✓ ドキュメント品質: 問題なし
  ✗ 整合性: README.mdに機能Xの説明が不足
```

## ファイル構成

```
.claude/
├── automation-plan.md           # この計画書
├── settings.json                # 既存ファイル（拡張）
└── skills/
    ├── code-reviewer/
    │   └── SKILL.md            # コードとドキュメントの整合性チェック
    └── doc-reviewer/
        └── SKILL.md            # ドキュメント品質チェック
```

## 技術仕様

### Skills の仕様

- **フォーマット**: YAML frontmatter + Markdown
- **必須フィールド**: name, description
- **ツール制限**: allowed-tools で安全性確保
- **起動方法**: 文脈に基づいて自動判断

### 自動起動の条件

Skills は以下の条件で自動起動：
- **code-reviewer**: "implementation", "documentation-code consistency" などのキーワード
- **doc-reviewer**: "documentation", "markdown", "creating files" などのキーワード

description フィールドにこれらのトリガーワードを含めることで、適切なタイミングで自動起動される。

## 注意事項

- Skills は文脈で自動起動されるため、description の記述が重要
- allowed-tools でツールを制限することでセキュリティを確保
- settings.json の変更は既存の hooks 設定を保持しながら追記

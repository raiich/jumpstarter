# Claude Code 自動化実装計画

## 目的

Claude Codeが変更を行う際に、自動的にプラン作成・レビュー・ドキュメント整合性チェックを実行する仕組みを構築する。

## 実装方針

### 自動化の範囲

**Claude Codeが変更を行う場合**（完全自動）:
- 変更前: 自動でプラン作成（TodoWrite + ExitPlanMode）
- 変更後: 自動でレビュー + ドキュメント整合性チェック

**ユーザーが変更を行う場合**（手動実行）:
- スラッシュコマンド（`/review`、`/check-docs`）で明示的に実行

## 実装するコンポーネント

### 1. Skills（自動起動）

#### 1-1. auto-reviewer Skill
- **ファイル**: `.claude/skills/auto-reviewer/SKILL.md`
- **目的**: コード変更後に自動的に品質チェック
- **機能**:
  - エラーハンドリングの確認
  - 戻り値の無視チェック
  - コーディング規約の確認
  - セキュリティ問題の検出
- **ツール制限**: Read, Grep, Glob のみ
- **起動タイミング**: 私（Claude）がコード実装完了後、自動起動

#### 1-2. doc-checker Skill
- **ファイル**: `.claude/skills/doc-checker/SKILL.md`
- **目的**: ドキュメントと実装の整合性を自動確認
- **機能**:
  - README.md と実装の一致確認
  - docs/*.md の記述と実装の照合
  - 未ドキュメント機能の検出
  - 実装されていない記載の検出
- **ツール制限**: Read, Grep, Glob のみ
- **起動タイミング**: 私（Claude）がコード実装完了後、auto-reviewer と並列で自動起動

### 2. スラッシュコマンド（手動実行）

#### 2-1. /review コマンド
- **ファイル**: `.claude/commands/review.md`
- **目的**: ユーザーが明示的にレビューを要求
- **機能**: 最近の変更をレビュー
- **実行**: ユーザーが `/review` と入力

#### 2-2. /check-docs コマンド
- **ファイル**: `.claude/commands/check-docs.md`
- **目的**: ユーザーが明示的にドキュメントチェックを要求
- **機能**: ドキュメントと実装の整合性確認
- **実行**: ユーザーが `/check-docs` と入力

### 3. Claude の振る舞い設定

#### 3-1. settings.json の拡張
- **ファイル**: `.claude/settings.json`（既存ファイルに追記）
- **目的**: 私（Claude）の自動プラン作成を有効化
- **内容**:
  - 変更前の自動プラン作成設定
  - TodoWrite の自動使用
  - ExitPlanMode の自動使用
  - Skillsとの連携方針

## 実装順序

1. **Skills の作成**（最優先）
   - auto-reviewer Skill
   - doc-checker Skill

2. **スラッシュコマンドの作成**
   - /review コマンド
   - /check-docs コマンド

3. **settings.json の更新**
   - Claude の振る舞い設定を追記

4. **動作確認**
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

[auto-reviewer Skill] コード品質チェック
[doc-checker Skill] ドキュメント整合性チェック

↓ 自動レポート

[Claude] チェック結果を報告
  ✓ コード品質: 問題なし
  ✗ ドキュメント: README.mdに機能Xの説明が不足
```

### シナリオ2: ユーザーが変更した場合

```
[ユーザーがファイルを手動編集]

[ユーザー] /review

↓ 手動実行

[Claude] レビューを実行して結果を報告
```

## ファイル構成

```
.claude/
├── AUTOMATION_PLAN.md           # この計画書
├── settings.json                # 既存ファイル（拡張）
├── skills/
│   ├── auto-reviewer/
│   │   └── SKILL.md            # 自動コードレビュー
│   └── doc-checker/
│       └── SKILL.md            # ドキュメント整合性チェック
└── commands/
    ├── review.md               # 手動レビューコマンド
    └── check-docs.md           # 手動チェックコマンド
```

## 技術仕様

### Skills の仕様

- **フォーマット**: YAML frontmatter + Markdown
- **必須フィールド**: name, description
- **ツール制限**: allowed-tools で安全性確保
- **起動方法**: 文脈に基づいて自動判断

### 自動起動の条件

Skills は以下の条件で自動起動：
- **auto-reviewer**: "code changes", "implementation", "review" などのキーワード
- **doc-checker**: "documentation", "implementation complete", "consistency" などのキーワード

description フィールドにこれらのトリガーワードを含めることで、適切なタイミングで自動起動される。

## 注意事項

- Skills は文脈で自動起動されるため、description の記述が重要
- allowed-tools でツールを制限することでセキュリティを確保
- settings.json の変更は既存の hooks 設定を保持しながら追記

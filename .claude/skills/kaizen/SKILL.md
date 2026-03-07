---
name: kaizen
description: Claude Code 改善コマンド
allowed-tools: Read, Grep, Glob, Edit, Write, WebFetch, AskUserQuestion, Skill
---

# Claude Code 改善コマンド

conversation.log からフィードバックを抽出し、learnings.md に反映した上で、Claude Code 設定を改善します。

## 実行手順

### フェーズ1: フィードバック抽出

1. **ログ読み込みと判定**
   - `.local/claude/conversation.log` を読み込む
   - ログが存在しない、または空の場合はこのフェーズをスキップし、フェーズ2へ進む

2. **フィードバック抽出**
   - ユーザーからの指摘、修正依頼、方針指示、好み・判断基準を抽出
   - 単なる作業指示（「○○して」）は除外し、再利用可能な知見のみ抽出
   - 既に `.local/claude/learnings.md` に記載済みの内容は除外

3. **learnings.md への追記**
   - 抽出結果が0件の場合は learnings.md を更新せず、そのままフェーズ2へ進む
   - `.local/claude/learnings.md` の該当カテゴリに追記:
     - **方針**: 開発方針、意思決定の背景
     - **考慮すべき観点**: 重視する観点、判断基準
     - **進め方のコツ**: 効果的なやり方、段取り
     - **注意点**: 避けるべきこと、見落としやすいポイント
   - 各項目は1行で簡潔に記述
   - 追記内容をユーザーに提示し、承認後に更新

4. **ログクリア**
   - 承認・更新完了後、`.local/claude/conversation.log` をクリア（空ファイルにする）

### フェーズ2: Claude Code 設定の改善

1. **フィードバック分析**
   - `.local/claude/learnings.md` を読み込む
   - 記録されたフィードバックから改善可能なパターンを特定

2. **Claude Code 機能の確認**
   - WebFetchを使用してClaude Code公式ドキュメントを確認
   - 確認URL: https://code.claude.com/docs/en/claude_code_docs_map.md
   - changelog ( https://code.claude.com/docs/en/changelog.md ) で新機能・変更点も確認
   - フィードバックパターンを解決できるClaude Code機能を特定
   - 既存の設定で活用できていない機能がないかチェック

3. **改善案の作成**
   - フィードバックとClaude Code機能を組み合わせた改善案を作成
   - 以下の設定ファイルを対象：
     - `.claude/settings.local.json`（個人設定。テンプレートリポジトリのため優先）
     - `.claude/settings.json`（テンプレート共有設定）
     - `.claude/agents/`
     - `.claude/skills/`
     - `.claude/commands/`
     - `.claude/rules/`（実体は `.github/instructions/`。編集は実体側で行う）

4. **改善の実施**
   - 改善案をユーザーに提示
   - 承認後、設定ファイルを更新

## 出力フォーマット

```
## フィードバック分析結果

### 主要なフィードバック
- [フィードバック1: 頻度X回]
- [フィードバック2: 頻度Y回]

### 改善提案

#### 提案1: [タイトル]
- **対象**: [ファイル名]
- **内容**: [変更内容]
- **効果**: [期待される改善]

#### 提案2: [タイトル]
...

### Claude Code 最新機能の活用
- [機能名]: [活用方法]
```

## 注意事項

- 既存設定を破壊しないよう注意
- 変更前にバックアップを推奨
- ユーザー承認後に変更を実施
- **learnings.md は kaizen 専用の中間データ**。通常セッションへのコンテキスト注入（SessionStart フック等）は提案しない。改善は `.claude/` 設定ファイルに反映する

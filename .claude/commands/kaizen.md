# Claude Code 改善コマンド

conversation.logからユーザーフィードバックを分析し、Claude Code設定を改善します。

## 実行手順

1. **フィードバック分析**
   - `.claude/logs/conversation.log` を読み込む
   - ユーザーからのフィードバック、要望、不満を抽出
   - 繰り返し指摘される問題を特定

2. **Claude Code機能の確認**
   - WebFetchを使用してClaude Code公式ドキュメントを確認
   - 確認URL: https://docs.claude.com/en/docs/claude-code/claude_code_docs_map.md
   - フィードバックパターンを解決できるClaude Code機能を特定
   - 既存の設定で活用できていない機能がないかチェック

3. **改善案の作成**
   - フィードバックとClaude Code機能を組み合わせた改善案を作成
   - 以下の設定ファイルを対象：
     - `.claude/settings.json`
     - `.claude/agents/`
     - `.claude/skills/`
     - `.claude/commands/`
     - `.claude/guidelines/`

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

## 改善の実例

このテンプレート自体も `/kaizen` により継続的に改善していきます:

### 第1回 kaizen
- Guidelines: 4種類を生成
- 対話ログ記録機能を追加

### 第2回 kaizen
- Guidelines を統合・簡素化
- 冗長性を排除

### 第3回 kaizen
- Sub-Agents: 4種類を追加
- Settings: Plan Mode、安全制限を強化
- Skills: doc-reviewer を簡素化

各改善で conversation.log からユーザーフィードバックを抽出し、設定に反映しています。

## 注意事項

- 既存設定を破壊しないよう注意
- 変更前にバックアップを推奨
- ユーザー承認後に変更を実施

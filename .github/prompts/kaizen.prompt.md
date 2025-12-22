---
title: Kaizen - プロジェクト改善提案
description: GitHub Copilot の会話履歴を分析し、ワークスペース設定やドキュメントの改善を提案
---

# Kaizen - プロジェクト改善提案

GitHub Copilot の会話履歴からユーザーフィードバックを分析し、プロジェクト設定を改善します。

## 実行手順

1. **会話履歴の分析**
   - 会話履歴を読み込む
       - `~/Library/Application Support/Code/User/workspaceStorage/*/chatSessions/*.json`
   - ユーザーからのフィードバック、要望、不満を抽出
   - 繰り返し指摘される問題や頻繁に聞かれる質問を特定

2. **GitHub Copilot機能の確認**
   - fetch_webpageを使用してGitHub Copilot公式ドキュメントを確認
   - 確認URL:
     - https://code.visualstudio.com/docs/copilot/overview
     - https://docs.github.com/en/copilot
   - フィードバックパターンを解決できるGitHub Copilot機能を特定
   - 確認すべき主な機能カテゴリ:
     - Inline suggestions（コード補完）
     - Autonomous coding（Agents による自律的なコーディング）
     - Natural language chat（自然言語チャット）
     - Smart actions（スマートアクション）
     - Custom instructions（カスタム指示）
     - Custom agents（カスタムエージェント）
     - MCP servers and tools（外部ツール連携）
   - 既存の設定で活用できていない機能がないかチェック

3. **改善対象の特定**
   - 以下のファイルを対象：
     - `.github/prompts/` - プロンプトファイル
     - `.github/copilot-instructions.md` - プロジェクト固有の指示
     - `.github/instructions/*.instructions.md` - プロジェクト固有の指示
     - `README.md` - プロジェクトドキュメント
     - `.vscode/settings.json` - ワークスペース設定
     - その他のドキュメントや設定ファイル

4. **改善案の作成**
   - フィードバックとGitHub Copilot機能を組み合わせた改善案を作成
   - 既存の設定で活用できていない機能がないかチェック
   - 新しいプロンプトファイルの作成を検討

5. **改善の実施**
   - 改善案をユーザーに提示
   - 承認後、設定ファイルを更新

## 出力フォーマット

```markdown
## 会話履歴分析結果

### 主要なトピック
- [トピック1: 頻度X回]
- [トピック2: 頻度Y回]

### 繰り返し発生する問題
- [問題1: 説明]
- [問題2: 説明]

### 改善提案

#### 提案1: [タイトル]
- **対象**: [ファイル名]
- **内容**: [変更内容]
- **効果**: [期待される改善]

#### 提案2: [タイトル]
...

### 新規プロンプトの提案
- [プロンプト名]: [用途と効果]

### GitHub Copilot 最新機能の活用
- [機能名]: [活用方法]
```

## 改善の実例

このテンプレート自体も `kaizen` により継続的に改善していきます。

## 注意事項

- 既存設定を破壊しないよう注意
- 変更前にバックアップを推奨
- ユーザー承認後に変更を実施

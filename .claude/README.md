# Claude Code 会話ログ

このディレクトリには、Claude Code との会話を記録するための設定が含まれています。

## 設定内容

`config.yaml` に定義されたフックにより、以下が自動的にログに記録されます：

- **ユーザープロンプト**: あなたが Claude Code に送信した質問や指示
- **ツール実行**: Claude Code が使用したツール（ファイル読み込み、編集、コマンド実行など）

## ログファイル

会話ログは `.claude/logs/conversation.log` に保存されます。

```bash
# ログを確認
cat .claude/logs/conversation.log

# リアルタイムでログを監視
tail -f .claude/logs/conversation.log

# 特定の日付のログを検索
grep "2025-11-01" .claude/logs/conversation.log
```

## ログの活用方法

### 1. 過去の会話を検索
```bash
# 特定のキーワードを含む会話を検索
grep -A 5 "キーワード" .claude/logs/conversation.log
```

### 2. 日次レポートの作成
```bash
# 今日の会話をファイルに出力
grep "$(date '+%Y-%m-%d')" .claude/logs/conversation.log > today.log
```

### 3. 会話の統計
```bash
# ユーザープロンプトの数をカウント
grep "USER PROMPT" .claude/logs/conversation.log | wc -l

# 使用されたツールの種類を確認
grep "TOOL:" .claude/logs/conversation.log | sort | uniq -c
```

## 注意事項

- ログファイル (`.claude/logs/`) は `.gitignore` に追加されており、Git にコミットされません
- 個人的な会話履歴が含まれる可能性があるため、共有に注意してください
- ログファイルは自動的にローテーションされないため、定期的に整理することを推奨します

## カスタマイズ

`config.yaml` を編集することで、ログの形式や記録する情報を変更できます。
詳細は [Claude Code ドキュメント](https://docs.claude.com/claude-code) を参照してください。

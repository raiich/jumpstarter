# Jumpstarter - Claude Code 継続的改善テンプレート

**Claude Codeとの対話から学び、開発効率を継続的に改善するテンプレートリポジトリ**

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

---

## 概要

Jumpstarterは、Claude Codeを使った開発で「同じフィードバックを何度も繰り返す」問題を解決するテンプレートリポジトリです。

**コアコンセプト**: 対話ログ（conversation.log）を分析し、`/kaizen`コマンドでClaude Codeの設定を自動最適化

### こんな経験ありませんか？

```
あなた: 「新機能を実装して」
Claude: [実装]
あなた: 「冗長すぎます」
Claude: [修正]
あなた: 「他に漏れはないですか？」
Claude: [追加修正]
あなた: 「まだあります...」 ← 何回目？
```

**Jumpstarterは、この繰り返しを自動で解決します。**

---

## クイックスタート

### 1. このテンプレートを使う

```bash
# GitHub上で "Use this template" をクリック、または
git clone https://github.com/your-org/jumpstarter.git my-project
cd my-project
```

### 2. Claude Codeで開く

```bash
# Claude Code CLIで開く
claude .
```

### 3. 開発を始める

普段通り開発します。conversation.logに対話履歴が自動記録されます。

### 4. 定期的に改善

```bash
# Claude Codeセッション内で
/kaizen
```

conversation.logを分析し、改善提案を自動生成します。

---

## 主な機能

### 📊 対話ログ自動記録

`.claude/settings.json`に設定されたhooksにより、全ての対話を`.claude/logs/conversation.log`に記録。

### 🔍 `/kaizen` コマンド

conversation.logを分析し、繰り返されるフィードバックパターンから改善策を自動生成：

- **ガイドライン**: Claude Codeの振る舞いルール
- **Skills**: 自動実行される品質チェック機能
- **Hooks**: イベント駆動の自動処理

### 📝 充実したテンプレート

すぐに使える設定ファイル：

```
.claude/
├── commands/
│   └── kaizen.md              # 継続的改善コマンド
├── guidelines/
│   ├── self-review.md         # セルフレビュー手順
│   ├── documentation.md       # ドキュメント作成ルール
│   ├── execution.md           # 実行前確認
│   └── test-failure-handling.md # テスト失敗時対応
├── skills/
│   ├── code-reviewer/         # コード整合性チェック
│   └── doc-reviewer/          # ドキュメント品質チェック
└── settings.json              # 対話ログ記録設定
```

---

## 使い方

### 基本ワークフロー

```
┌─────────────┐
│ 開発作業     │
└──────┬──────┘
       ↓
┌─────────────┐
│conversation.log│ ← 自動記録
│に履歴蓄積    │
└──────┬──────┘
       ↓
┌─────────────┐
│ /kaizen実行  │ ← 週次・隔週
└──────┬──────┘
       ↓
┌─────────────┐
│改善策の提案  │
│と自動生成    │
└──────┬──────┘
       ↓
┌─────────────┐
│設定の最適化  │
└─────────────┘
```

### 推奨実行頻度

- **週次**: プロジェクト初期〜中期
- **隔週**: プロジェクト安定期
- **月次**: 設定が成熟した段階

---

## 実証結果

このテンプレートは、**4日間のGoベンチマーク実装プロジェクトで検証されました。**

### Before（改善前）

```
📊 conversation.log分析結果
- "冗長すぎます"      → 15回
- "他に漏れは..."     → 10回
- "実行しないで"      → 3回

合計 217箇所で同じフィードバックを繰り返していた
```

### `/kaizen`実行

```
🤖 自動生成された改善策
- ガイドライン: 4種類
- Skills: 2種類
- 対象課題: 冗長性、完全性、実行タイミング
```

### After（改善後）

```
✅ 効果
- フィードバックループが激減（2-3往復 → 1往復）
- 冗長なドキュメントの自動抑制
- 事前セルフレビュー機能の追加
```

詳細は [プロジェクト履歴](docs/project-history.md) を参照。

---

## ディレクトリ構成

```
jumpstarter/
├── .claude/                    # Claude Code設定
│   ├── commands/              # カスタムコマンド
│   │   └── kaizen.md          # 継続的改善コマンド
│   ├── guidelines/            # 振る舞いルール
│   │   ├── self-review.md
│   │   ├── documentation.md
│   │   ├── execution.md
│   │   └── test-failure-handling.md
│   ├── skills/                # 自動実行機能
│   │   ├── code-reviewer/
│   │   └── doc-reviewer/
│   ├── logs/                  # 対話ログ（自動生成）
│   │   └── conversation.log
│   ├── settings.json          # メイン設定
│   ├── kaizen-introduction.md # /kaizen使い方ガイド
│   └── automation-plan.md     # 自動化の仕様
├── docs/                      # ドキュメント
│   ├── project-history.md     # プロジェクト履歴
│   ├── blog-story.md          # ブログ記事用
│   └── lt-slides-summary.md   # LT発表用
├── benchmarks/                # 検証用ベンチマーク
│   └── *_bench_test.go
└── README.md                  # このファイル
```

---

## カスタマイズ

### 1. プロジェクトに合わせた調整

`.claude/guidelines/`や`.claude/skills/`を、あなたのプロジェクトに合わせて編集：

```bash
# 例: Python向けのガイドライン追加
echo "# Python Coding Standards" > .claude/guidelines/python-style.md
```

### 2. チーム共有

`.claude/`ディレクトリをgitで共有すれば、チーム全体で一貫した設定を使用：

```bash
git add .claude/
git commit -m "Add team guidelines"
git push
```

### 3. 他のAIツールへの応用

このコンセプトは、他のAIツールにも応用可能：

- **Cursor**: `.cursorrules`の自動生成
- **GitHub Copilot**: コメント規約の最適化
- **カスタムLLM**: システムプロンプトの継続的改善

---

## ドキュメント

### 入門ガイド

- **[/kaizenコマンド紹介](.claude/kaizen-introduction.md)** - 使い方と効果
- **[自動化計画](.claude/automation-plan.md)** - 技術仕様

### 背景資料

- **[プロジェクト履歴](docs/project-history.md)** - 開発の経緯と成果
- **[ブログ記事](docs/blog-story.md)** - ストーリー形式の解説
- **[LT資料](docs/lt-slides-summary.md)** - 発表用サマリー

### ベンチマーク（検証用）

- **[ベンチマークの書き方](docs/writing-benchmarks.md)** - 検証方法
- **[実行ガイド](docs/running-benchmarks.md)** - 実行方法

---

## FAQ

### Q1: なぜconversation.logが必要なのですか？

A: AIアシスタントは毎回まっさらな状態で会話を始めます。過去の失敗を覚えていません。conversation.logを分析することで、繰り返される問題を特定し、設定で修正できます。

### Q2: `/kaizen`はどのように動作しますか？

A: 以下のステップで動作します：
1. conversation.logを読み込む
2. フィードバックパターンを抽出（AIで意味を理解）
3. 問題の根本原因を推定
4. Claude Code公式機能を活用した改善策を生成
5. ユーザー承認後、設定ファイルを自動作成

### Q3: 他のプロジェクトでも使えますか？

A: はい。Jumpstarterはテンプレートリポジトリなので、どんなプロジェクトにも適用できます。言語やフレームワークは問いません。

### Q4: conversation.logが巨大になりませんか？

A: 定期的にローテーションするか、直近N日分のみを分析対象にする方法があります。また、`.gitignore`に追加すればリポジトリには含まれません。

### Q5: チーム開発でどう使いますか？

A: `.claude/`ディレクトリをgitで共有し、各メンバーが同じ設定を使用します。チーム全体で一貫した開発体験を実現できます。

---

## 貢献

Issues、Pull Requestsを歓迎します。

### 改善アイデア

- 他の言語・フレームワーク向けのガイドライン
- より高度なSkills
- `/kaizen`の機能拡張

---

## ライセンス

MIT License

---

## 関連リンク

- [Claude Code公式ドキュメント](https://docs.claude.com/en/docs/claude-code)
- [Claude Code GitHub](https://github.com/anthropics/claude-code)

---

## クレジット

このプロジェクトは、Claude Codeとの対話から生まれました。

**コアアイデア**: 「AIに教育される」から「AIを教育する」へのパラダイムシフト

---

**始める準備はできましたか？**

```bash
/kaizen
```

あなたのプロジェクトも、今日から継続的に改善できます。

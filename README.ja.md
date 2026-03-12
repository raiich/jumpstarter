# Jumpstarter - Claude Code テンプレート

[ [English](README.md) | **日本語** ]

## Jumpstarter とは

Claude Code の設定のテンプレートリポジトリです。

Claude Code から期待する出力を得るには、適切なプロンプトを与える必要があります。
また、Rules / Skills / Hooks といった設定の整備にもコストがかかります。
Jumpstarter は、clone して使い始めるだけでこれらの課題を解消し、開発そのものに集中できるようにします。

Claude Code の利用で**ユーザーが頑張らない**ことを目指します:

- **長いプロンプト不要** — コードベース調査・要件のヒアリング・自己レビューを適切に実行します。
- **凝った仕組みは後回し** — MCP やマルチエージェントは後回し。Rules / Skillsをメインに使います。
- **継続的改善も手軽** — 改善点やベストプラクティスにもとづき、設定を自動更新します。

---

## 使い方

テンプレートをコピーして、すぐに使い始められます。

```bash
git clone https://github.com/raiich/jumpstarter.git

# 既存プロジェクトに .claude/ と .github/ をコピー（日本語版）
cp -R jumpstarter/.claude "${YOUR_PROJECT}/"
cp -R jumpstarter/.github "${YOUR_PROJECT}/"

## 英語版
# cp -R jumpstarter/en/.claude "${YOUR_PROJECT}/"
# cp -R jumpstarter/en/.github "${YOUR_PROJECT}/"
```

Claude Code を起動して、スラッシュコマンドでワークフローを起動して開発を進めます。

```bash
claude
```

### 新機能を設計・実装する

機能追加時に頑張って完璧なプロンプトを書く必要はありません。
最小限の指示をもとに、適切に設計・実装するためのワークフローを用意しています。

`/design-feature`スキルが、コードベースを調査した上で的確に要件をヒアリングしてDesign Docを作成します:

```
/design-feature CLIにverboseオプションを追加したい
```

実装のズレを減らすためにあらかじめテストケースを設計したい場合は、`/design-feature-tests`スキルが便利です:

```
/design-feature-tests .local/docs/features/verbose-option/design-doc.md
```

`/implement-feature`スキルにより、Design Docやテストケースにもとづいて実装します:

```
/implement-feature .local/docs/features/verbose-option/design-doc.md
```

設計ドキュメントやテストケースは `.local/docs/features/[名前]/` に整理されます。
各ステップの成果物をチェックポイントとして利用することで、実装がずれても手戻りを小さくできます。

### Claude Code の設定を継続的に改善する

開発体験が良くならないのは、日々のフィードバックが設定に反映されていないことが多いためです。
設定の改善も、ユーザーが頑張って行動する必要はありません。
下記のような仕組みが、継続的な改善の手助けになります。

`/kaizen` コマンドにより、Hooks で記録した対話ログから改善点を抽出できます:

```
/kaizen   対話ログから改善の提案をして
```

`/import-best-practices` コマンドにより、ネット上のベストプラクティスを簡単に取り込めます:

```
/import-best-practices   https://... の記事のベストプラクティスを取り込みたい
```

これらにより、ルール・スキル・フックなどがプロジェクトに合わせて洗練されていきます。頑張って設定を見直す必要はありません。

---

## その他の設定

`.claude/` ディレクトリの各要素が連携して「頑張らない」を実現します。

### 頑張らないための Rules

ユーザーが毎回指示しなくても、一定の品質基準で動作するためのルール群です (`.claude/rules/`)。

- 冗長な出力を頑張って読み解く必要はありません → **writing-style**
- 修正漏れがないか頑張って確認する必要はありません → **fix-guidelines**
- ワークフローを頑張って指示する必要はありません → **workflow-patterns**

### 頑張らないための Hooks

`/kaizen` による改善提案の入力データとなる対話ログを `.local/claude/conversation.log` に自動記録します。

---

## カスタマイズ

テンプレートをコピーした後、プロジェクトに合わせて調整できます。

- **スキル・ルールの追加**: `.claude/skills/` や `.claude/rules/` にファイルを追加
- **権限設定**: `.claude/settings.local.json` で許可・拒否するコマンドを調整
- **自動改善**: `/kaizen` で対話ログから改善点を抽出、`/import-best-practices` で外部記事のベストプラクティスを取り込み

---
name: review
description: Review code, design docs, or user-facing documentation from multiple independent perspectives. Enumerate perspectives first, then systematically review each one.
allowed-tools: Read, Grep, Glob, Bash, Task, AskUserQuestion, Agent
effort: high
---

# Review

コード・設計ドキュメント・利用者向けドキュメントを独立した視点でレビューするスキルです。

## 前提

- **観点を先に列挙し、各観点で漏れなくレビューする**（後から「この観点もあった」とならないようにする）
- 指摘はすべて初回レビューで出し切る。追加パスで新しい指摘が出るのは失敗

## フロー

### 1. レビュー対象の特定と種別判別

レビュー対象と種別（`code` / `design` / `docs`）を特定する。

**種別の判別:**

| 種別 | 対象例 | 典型的なパス |
|---|---|---|
| `code` | 実装コード・テストコード | `src/**`, `**/*.test.*` など |
| `design` | Design Doc・テストケース設計 | `.local/docs/features/design.md`, `.local/docs/features/tests.md` |
| `docs` | 利用者向けドキュメント | `README.md`, `docs/**/*.md` |

判別ロジック:
- 引数で明示（`code` / `design` / `docs`）されていればそれに従う
- 引数がなくパスから一意に推論できる場合はそれに従う
- 判別不能な場合は `AskUserQuestion` で確認する
- 指定なしの場合は `git diff --name-only` で変更ファイルから推論し、ユーザーに確認する

関連ファイルの読み込み:
- `code` の場合、`.local/docs/features/design.md` / `tests.md` が存在すれば設計との整合性確認のため読み込む
- `design` の場合、関連するコードベースの構造も把握する
- `docs` の場合、ドキュメントが参照するコード（公開 API、設定ファイル等）を把握する

**ツール**: Read, Glob, Grep, Bash

### 2. レビュー観点の列挙

種別に応じて参照する perspectives を決定し、適用する観点を自律的に選定する（ユーザーへの事前確認はしない）。選定結果は最後の報告でユーザーに伝える。

| 種別 | 参照 perspectives |
|---|---|
| `code` | [design-and-coding.md](../../guidelines/perspectives/design-and-coding.md), [testing.md](../../guidelines/perspectives/testing.md) |
| `design` | [design-and-coding.md](../../guidelines/perspectives/design-and-coding.md), [documentation.md](../../guidelines/perspectives/documentation.md), [testing.md](../../guidelines/perspectives/testing.md) |
| `docs` | [documentation.md](../../guidelines/perspectives/documentation.md) |

引数やユーザーの当初指示でスコープが明示されていればそれに従う。例:
- 「アルゴリズムは詳しく見ない」→ アルゴリズム正当性の観点を除外
- 「エッジケースはテストで見る」→ 境界条件の観点をテスト側に限定
- 「官能検査は後で」→ UX/見た目の観点を除外

### 3. 多角的レビューの実施

種別に応じてサブエージェント構成を選択し、並行実行する。

| 種別 | サブエージェント構成 | 各エージェントの担当観点 |
|---|---|---|
| `code` | `critical-analyst` + `security-reviewer` + `codebase-explorer` | 構造・設計の問題／セキュリティ・堅牢性／テストコード品質 |
| `design` | `critical-analyst` + `constructive-analyst` + `security-reviewer` | 弱点・リスクの発見／強みの増幅・機会の発見／セキュリティ観点 |
| `docs` | `critical-analyst`（コードとの突合せ検証を明示指示） | ドキュメントの主張がコードと整合するか |

各サブエージェントに渡す情報:
- レビュー対象ファイルのパス一覧
- 適用する観点リスト（ステップ2で確定したもの）
- 関連ドキュメント（`design.md` / `tests.md` 等、存在する場合）

**種別別の補足:**

- `code`: `security-reviewer` は入力検証・リソースリーク・境界条件・型安全性を担当。`codebase-explorer` はテストコードの false positive、循環ロジック、引数不足、assertion 網羅性を担当
- `design`: 作成者の自己レビューとは独立した視点で実施する
- `docs`: 「書いてある通りに動くか」をコードと突き合わせて検証する。推測で「正しそう」と判断しない。コード例・API 記載・設定値は実際にコードを読んで確認する

**ツール**: Agent, Read, Grep, Glob, Bash

### 4. レビュー結果の統合

サブエージェントの結果を統合し、以下の形式で提示する。**レビューしたスコープ（対象ファイル・適用観点・除外した観点とその理由）** を冒頭に明示し、ユーザーがレビュー範囲を後から把握できるようにする。

```markdown
## レビュー結果

### レビュー対象
- ファイル: [レビューしたファイルのリスト]
- 種別: [code / design / docs]

### レビュー観点
- 適用した観点: [一覧]
- 除外した観点: [一覧と理由（スコープ外・指示によるもの等）]

### 指摘事項

| # | 深刻度 | 観点 | 指摘内容 | 該当箇所 |
|---|--------|------|----------|----------|
| 1 | Critical | ... | ... | file:line |

### 総評
[全体的な評価と推奨アクション]
```

**深刻度の基準**: [../../guidelines/perspectives/review-severity.md](../../guidelines/perspectives/review-severity.md)

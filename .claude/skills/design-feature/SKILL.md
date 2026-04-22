---
name: design-feature
description: プロンプトを起点にヒアリング・設計・テスト設計を行い design.md / tests.md を作成するスキル。実装はしない。design.md が既にあればテスト設計から再開。
allowed-tools: Read, Grep, Glob, Edit, Write, Task, AskUserQuestion
effort: high
---

# Design Feature

ユーザーのプロンプトを起点に、調査・ヒアリング・設計・テスト設計までを承認ゲートなしで完走するスキル。
実装は行わない（実装まで含むなら `develop-feature` スキルを使う）。

## 前提

- 承認ゲートなし。ヒアリングで必要情報を揃えたら、設計完了まで自律的に進める
- ユーザー確認は「要件・方針の不足情報」に限定（プロセスの途中承認は取らない）
- 仮定で埋めた箇所は ❓ マークで明示（検証は `develop-feature` の実装フェーズに委ねる）
- 成果物は `design.md` と `tests.md` のみ。実装コードは書かない
- 軽量な修正・新規実装（設計ドキュメント不要レベル）は `fix` スキルを使う

## 保存先

- **フィーチャー固有ドキュメント**: `.local/docs/features/[名前]/`
  - `design.md` - Design Doc
  - `tests.md` - テストケース設計

## 起動時の分岐

対象 feature 名から `.local/docs/features/[名前]/` を確認し、開始フェーズを決定:

| 状態 | 開始フェーズ |
|------|-------------|
| `design.md` なし | [1. 調査・ヒアリング・設計] から |
| `design.md` あり、`tests.md` なし | [2. テスト設計] から（既存設計を尊重） |
| 両方あり | 完了済み。更新要望があれば該当フェーズから再実施 |

feature 名が不明な場合はユーザーに確認する。

## フロー

### 1. 調査・ヒアリング・設計

#### 1.1. コードベース調査

既存コードベース・ドキュメントを調査。広範な探索は Task（Explore エージェント）を活用。

**ツール**: Read, Glob, Grep, Task

#### 1.2. 要件ヒアリング

調査結果を踏まえ、不足情報を効率的にヒアリング。必要なら複数ラウンド実施。

ヒアリングの原則と Good/Bad 例は [../../guidelines/processes/hearing.md](../../guidelines/processes/hearing.md) を参照。

**ツール**: AskUserQuestion

#### 1.3. Design Doc 作成

**保存先**: `.local/docs/features/[名前]/design.md`

**内容:**
```markdown
# Design Doc: [名前]

## 背景・目的
- What: 何を作るか
- Why: なぜ必要か

## 要件
- 機能要件
- 非機能要件・制約条件

## スコープ

## 技術的アプローチ（選択理由、代替案）

## 設計（アーキテクチャ、処理フロー）

## 関連コード・参照
- 変更対象ファイル・関数
- 参考にすべき既存パターン
- 関連ドキュメント

## 実装詳細
- インターフェース/シグネチャのみ。メソッドの中身は書かない
- 重要なアルゴリズムやロジックの分岐のみコード例で示す

## 考慮事項（セキュリティなど）
```

**信頼度マーク**: [../../guidelines/perspectives/documentation.md](../../guidelines/perspectives/documentation.md) の信頼度マークセクションを参照。付与対象は**要件**・**設計**・**実装詳細** セクションの各項目。

**⛔ Design Doc はコードを書く場所ではない**: 書いてよい/いけないコードの基準、Bad/Good 例は [../../guidelines/perspectives/documentation.md](../../guidelines/perspectives/documentation.md) の「設計ドキュメントに実装コードを書きすぎない」セクションを参照。

**ツール**: Write, Edit

#### 1.4. Design Doc セルフレビュー

次フェーズに進む前に `design.md` をレビュー。不備があれば 1.3 に戻って修正する。

観点: [../../rules/self-review.instructions.md](../../rules/self-review.instructions.md)（対象: ドキュメント・コード設計）

### 2. テスト設計

#### 2.1. テスト基盤の調査

既存テスト基盤を調査し、流用可能なものを把握。原則は [../../guidelines/perspectives/testing.md](../../guidelines/perspectives/testing.md) の「既存テスト基盤の流用」セクションを参照。

**ツール**: Glob, Read

#### 2.2. テストケース設計

Design Doc に基づいてテストケースを設計。

**保存先**: `.local/docs/features/[名前]/tests.md`

**内容:**
```markdown
# テストケース設計

## テストケース1: [振る舞い]
- Given: [前提条件]
- When: [実行する操作]
- Then: [期待される結果]

## テストケース2: [振る舞い]の異常系
- Given: [前提条件]
- When: [実行する操作]
- Then: [期待されるエラー処理]
```

**ツール**: Write, Edit

#### 2.3. テストケース セルフレビュー

作成した `tests.md` をレビュー。不備があれば 2.2 に戻って修正する。

観点: [../../rules/self-review.instructions.md](../../rules/self-review.instructions.md)（対象: テスト）

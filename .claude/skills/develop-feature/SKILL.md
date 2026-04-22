---
name: develop-feature
description: プロンプトを起点にヒアリング・設計・テスト設計・実装・テスト実行までを承認ゲートなしで一気通貫に進める。design.md / tests.md が既にあれば該当フェーズをスキップして再開。
allowed-tools: Read, Grep, Glob, Edit, Write, Bash, Task, AskUserQuestion, TaskCreate, TaskUpdate, TaskList, TaskGet, Skill
effort: high
---

# Develop Feature

ユーザーのプロンプトを起点に、調査・ヒアリング・設計・テスト設計・実装・テスト実行までを承認ゲートなしで完走するスキル。
設計・テスト設計フェーズは `design-feature` スキルに委譲する。

## 前提

- 承認ゲートなし。ヒアリングで必要情報を揃えたら、最後まで自律的に進める
- ユーザー確認は「要件・方針の不足情報」と「❓マーク付き仮定項目の検証」に限定（プロセスの途中承認は取らない）
- 成果物は `design.md`、`tests.md`、および実装コード・テストコード
- 実装中に設計不整合が見つかれば `design.md` / `tests.md` を更新してから続行
- 設計ドキュメントのみ必要な場合は `design-feature` スキルを使う
- 軽量な修正・新規実装（設計ドキュメント不要レベル）は `fix` スキルを使う

## 保存先

- **フィーチャー固有ドキュメント**: `.local/docs/features/[名前]/`
  - `design.md` - Design Doc
  - `tests.md` - テストケース設計

## 起動時の分岐

対象 feature 名から `.local/docs/features/[名前]/` を確認し、開始フェーズを決定:

| 状態 | 開始フェーズ |
|------|-------------|
| `design.md` / `tests.md` のいずれかが欠けている | [1. 設計] から |
| 両方あり | [2. 実装・テスト実行] から |

feature 名が不明な場合はユーザーに確認する。

## フロー

### 1. 設計

`design-feature` スキルを呼び出して `design.md` / `tests.md` を作成する。
design-feature 内で「`design.md` 無し→調査・設計から」「`tests.md` 無し→テスト設計から」の分岐が処理される。

**ツール**: Skill

### 2. 実装・テスト実行

#### 2.1. 実装タスクリスト作成

Design Doc とテストケースに基づいてタスクを作成。
`design.md` に ❓（仮定）マーク付き項目がある場合、該当タスクの実装前にユーザーへ確認する。

**タスクの粒度**:
- 1タスク = 1つの関数・クラス・メソッド程度
- 1タスク = テストケース1〜5個程度

**ツール**: TaskCreate

#### 2.2. TDDサイクル

各タスクについて以下を実施:

1. **テストコード作成** — `tests.md` のケースをテストコードに落とす
2. **機能実装** — テストが通るように実装
3. **テスト実行** — プロジェクト設定に従ったコマンドで実行
4. **失敗時の原因分析** — [../../guidelines/processes/root-cause-analysis.md](../../guidelines/processes/root-cause-analysis.md)
5. **リファクタリング** — テスト通過後、重複排除・命名改善・構造簡素化。再度テスト実行で回帰確認
6. **セルフレビュー** — [../../rules/self-review.instructions.md](../../rules/self-review.instructions.md) に従って実装・テストコードをレビュー。修正時は再度テスト実行で回帰確認
7. **タスク完了** — TaskUpdate で completed、次タスクへ

**ツール**: TaskUpdate, Bash, Edit, Write

#### 2.3. 実装中に設計不整合を発見した場合

- 該当タスクを一旦 pending に戻し、`design.md` / `tests.md` を更新
- 更新理由を簡潔に記録（信頼度マークも適宜付与）
- 更新後に該当タスクを in_progress に戻して継続

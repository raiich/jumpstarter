---
name: implement-feature
description: Implement features with TDD workflow based on a design doc and test cases.
allowed-tools: Read, Grep, Glob, Edit, Write, Bash(go test:*), Task, AskUserQuestion, TaskCreate, TaskUpdate, TaskList, TaskGet, Skill
---

# Implement Feature

Design Doc とテストケース設計を入力として、TDD 実装を行うスキルです。

## 前提

- 入力は `design-feature-tests` スキルの成果物（`design-doc.md` + `test-cases.md`）またはユーザーからの直接入力
- 小さい単位で TDD サイクルを回す

## 保存先

- **ナレッジベース**: `docs/` 配下にトピックごとのファイル
- **作業要点**: `.local/claude/learnings.md`（追記式）

## タスクの粒度

- 1タスク = 1つの関数・クラス・メソッド程度
- 1タスク = テストケース1〜5個程度

## フロー

### [入力の確認]

#### 1. Design Doc とテストケースの確認

`design-doc.md` と `test-cases.md` がある場合はそれらを読み込む。ない場合はユーザーの入力から要件を把握する。
不明点があればヒアリングする。

**ツール**: Read, AskUserQuestion

### [実装フェーズ] — タスクごとに繰り返し

#### 2. 実装タスクリスト作成

Design Doc とテストケースに基づいてタスクを作成。

**ツール**: TaskCreate

#### 3. TDDサイクル実施

**ツール**: TaskUpdate でタスクを in_progress に設定

各タスクについて以下を実施：

**3.1. テストコード作成（基本パターン）**
- テストケース設計に基づいてテストコードを実装

**3.2. 機能実装（基本パターン）**
- テストが通るように機能を実装

**3.3. テスト実行と失敗時の対応**

**ツール**: Bash (go test)

失敗時の対応：

1. **実装のバグ** → 修正し再テスト
2. **テスト設計のミス** → テストケース設計を見直し、test-cases.md を更新
3. **要件の認識違い** → Design Doc を見直し、design-doc.md を更新
4. **判断不可** → ユーザーにヒアリング

**3.4. リファクタリング**
- テストが通った後、実装コードを整理（重複排除、命名改善、構造の簡素化など）
- リファクタリング後にテストを再実行し、既存テストが通ることを確認

**3.5. タスク完了**

**ツール**: TaskUpdate で completed に設定、TaskList で次タスク確認

**レビュー粒度**:
- 軽微なタスク: まとめてレビュー可
- 重要なタスク: 個別レビュー

### [完了フェーズ]

ナレッジベース更新 → 作業要点の保存 → /kaizen 実行

## 自己レビュー観点

### テストコード
- テストケース設計との整合性
- 正常系・異常系・エッジケースの網羅性
- テストの粒度は適切か

### 実装コード
- 過剰な実装をしていないか
- セキュリティ上の懸念はないか

観点は内容に応じて調整。不明確な場合はユーザーにヒアリング。

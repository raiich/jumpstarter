---
name: design-feature-tests
description: Plan implementation and design test cases based on a design doc.
allowed-tools: Read, Grep, Glob, Edit, Write, Task, AskUserQuestion, EnterPlanMode, ExitPlanMode, Skill
---

# Design Feature Tests

Design Doc を入力として、実装計画の策定とテストケース設計を行うスキルです。

## 前提

- 入力は `design-feature` スキルの成果物（`design-doc.md`）またはユーザーからの直接入力
- 成果物は `test-cases.md` として保存し、`implement-feature` スキルへの入力となる
- **承認ゲート**: テストケースはユーザーの明示的な承認なしに完了としない

## 保存先

- **ナレッジベース**: `docs/` 配下にトピックごとのファイル
- **フィーチャー固有ドキュメント**: `docs/features/[名前]/` 配下
  - `test-cases.md` - テストケース設計（本スキルの主成果物）
- **作業要点**: `.local/claude/learnings.md`（追記式）

## フロー

### [入力の確認]

#### 1. Design Doc の確認

`design-doc.md` がある場合はそれを読み込む。ない場合はユーザーの入力から要件を把握する。
不明点があればヒアリングする。

**ツール**: Read, AskUserQuestion

### [実装計画フェーズ] — Plan mode

#### 2. Plan mode に入る

**ツール**: EnterPlanMode

#### 3. コードベース調査と実装計画

Design Doc を踏まえて、既存コードベースを調査し、実装方針をプランファイルに記録。
広範な探索が必要な場合は Task（Explore エージェント）を活用。

**ツール**: Read, Glob, Grep, Task

**プランファイルに記録する内容:**
```markdown
# 実装計画

## Design Doc サマリ
- design-doc.md からの要約

## 影響範囲
- 変更対象ファイル
- 影響を受ける既存機能

## 実装ステップ概要
- ステップ1: ...
- ステップ2: ...
```

#### 4. 自己レビューと Plan mode 終了

プランファイルを自己レビューし、ExitPlanMode でユーザー承認を求める。

**ツール**: ExitPlanMode

### [テストケース設計フェーズ] — 通常 mode（承認ゲート）

#### 5. テストケース設計

テストケースを設計。作成後は**基本パターン**（自己レビュー → ユーザーレビュー → 修正）に従い、ユーザー承認を得る。

**保存先**: `docs/features/[名前]/test-cases.md`

**内容:**
```markdown
# テストケース設計

## テストケース1: [機能名]の正常系
- Given: [前提条件]
- When: [実行する操作]
- Then: [期待される結果]

## テストケース2: [機能名]の異常系
- Given: [前提条件]
- When: [実行する操作]
- Then: [期待されるエラー処理]
```

**ツール**: Write, Edit, AskUserQuestion

**⛔ ユーザーの承認なしに次へ進まない**

### [完了フェーズ]

ナレッジベース更新 → 作業要点の保存 → /kaizen 実行

## 自己レビュー観点

### プラン（実装計画フェーズ）
- Design Doc の解釈は正しいか
- 影響範囲を見落としていないか

### テストケース
- 機能に対する網羅性
- 正常系・異常系・エッジケースの網羅性
- テストの粒度は適切か
- `.claude/rules/writing-style.instructions.md` の簡潔さの原則に従っているか

観点は内容に応じて調整。不明確な場合はユーザーにヒアリング。

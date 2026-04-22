---
name: fix
description: ユーザー入力を起点に、ヒアリングしてから修正または軽量な新規実装を承認なしで進めるスキル。設計ドキュメント（design.md / tests.md）は作らない。
allowed-tools: Read, Grep, Glob, Edit, Write, Bash, AskUserQuestion, TaskCreate, TaskUpdate, TaskList
---

# Fix

軽量な修正・新規実装を、ヒアリングしてから承認ゲートなしで進めるスキル。
設計ドキュメント（design.md / tests.md）は作らない。

## 使い分け

- **fix**: 既存コードの修正、局所的な追加、軽量な新規実装（設計ドキュメント不要レベル）
- **design-feature**: 設計ドキュメント（`design.md` / `tests.md`）のみを作成したい場合。実装はしない
- **develop-feature**: 要件整理から必要な機能追加、複数コンポーネントに跨る設計と実装まで一気通貫で進める場合

## 原則

- 場当たり的・段階的な修正は避け、方針を先に決めてから一括で変更する
- 原則（修正方針の事前確定、影響範囲の確認、パターン検索による網羅）は [../../guidelines/processes/fix-in-one-pass.md](../../guidelines/processes/fix-in-one-pass.md) を参照

## フロー

### 1. 入力理解とヒアリング

- ユーザーの指示を理解し、対象と影響範囲を推測
- 不足情報があれば効率的にヒアリング（[../../guidelines/processes/hearing.md](../../guidelines/processes/hearing.md)）

**ツール**: Read, Glob, Grep, AskUserQuestion

### 2. 包括的検索（修正系の場合）

Grep で同じ問題が他にも無いか網羅（fix-in-one-pass.md の「パターン検索」手順）。

**ツール**: Grep, Glob

### 3. 実装・修正

- 修正系: 見つかった全ての箇所を一度に修正
- 新規実装系: 影響範囲を踏まえて必要最小限のコードを実装

**ツール**: Edit, Write

### 4. テスト実行（テストがある場合）

プロジェクト設定に従ってテストを実行。失敗時の原因分析は [../../guidelines/processes/root-cause-analysis.md](../../guidelines/processes/root-cause-analysis.md) を参照。

**ツール**: Bash

### 5. セルフレビュー

- [ ] 該当箇所をすべて修正／実装したか
- [ ] 変更内容は正しいか
- [ ] 記述ルール（`.claude/rules/writing-style.instructions.md`）と セルフレビュー観点（`.claude/rules/self-review.instructions.md`）に従っているか
- [ ] 関連するドキュメント・コードも更新が必要か

### 6. 報告

簡潔に報告: 変更ファイル数、概要、テスト結果（実行した場合）。

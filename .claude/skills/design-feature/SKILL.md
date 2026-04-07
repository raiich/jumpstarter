---
name: design-feature
description: Design features through codebase investigation and user hearing, producing a design doc.
allowed-tools: Read, Grep, Glob, Edit, Write, Task, AskUserQuestion, EnterPlanMode, ExitPlanMode, Skill
effort: high
---

# Design Feature

コードベース調査とユーザーヒアリングを通じて、Design Doc を作成するスキルです。

## 前提

- ユーザーは初期段階で全要件を伝えるとは限らないため、能動的にヒアリングする
- コードベースとドキュメントを先に把握し、その理解に基づいて効率的に質問する
- 成果物は `design.md` として保存し、`implement-feature` スキルへの入力となる
- 各要件・設計項目には**信頼度マーク**を付与し、情報源の確かさを可視化する

## 保存先

- **フィーチャー固有ドキュメント**: `.local/docs/features/[名前]/`
  - `design.md` - Design Doc（本スキルの主成果物）

## フロー

### [調査・ヒアリングフェーズ] — Plan mode

#### 1. ユーザー: 機能の概要や方針の入力

#### 2. Plan mode に入る

**ツール**: EnterPlanMode

#### 3. コードベース調査

既存コードベースとドキュメントを調査し、プランファイルに記録。
広範な探索が必要な場合は Task（Explore エージェント）を活用。

**ツール**: Read, Glob, Grep, Task

#### 4. 要件のヒアリング

調査結果を踏まえて、不足情報をユーザーにヒアリング。
ヒアリング内容をプランファイルに反映。必要に応じて複数回ヒアリングする。

**効率的なヒアリングの原則:**
- コードを読めば分かることは聞かない
- 調査で把握した制約・パターンを踏まえ、具体的な選択肢を提示して聞く
- ユーザーが言及していない観点（エッジケース、既存機能との整合性等）も能動的に確認する

**例: 効率的なヒアリング**

ユーザー入力: 「通知機能を追加したい」

❌ 悪い質問: 「どのような通知機能ですか？」（広すぎる）

✅ 良い質問（調査結果を踏まえた選択肢付き）:
- 「既存の EventBus（events/bus.go）を使ったイベント駆動と、直接呼び出しのどちらを想定していますか？」
- 「通知先は Slack webhook（config に slack_url が既存）以外にもありますか？」

**ツール**: AskUserQuestion

**プランファイルに記録する内容:**

⚠️ **重要**: プランファイルの冒頭に必ず「成果物」セクションを記載すること。
プランモード終了後にコンテキストが失われた場合でも、成果物の種類を正しく判断できるようにするため。

```markdown
# 調査・ヒアリング結果

## 成果物
- 種類: Design Doc（マークダウンドキュメント）
- 保存先: .local/docs/features/[名前]/design.md
- ⚠️ コード実装は行わない（implement-feature の責務）

## 関連する既存機能
- 機能A: path/to/file

## 重要なアーキテクチャパターン
- パターン1: 説明

## ヒアリング結果
- 確認した要件・制約

## 技術的アプローチの候補
- アプローチA: 概要と利点・欠点
- アプローチB: 概要と利点・欠点
```

#### 5. 自己レビューと Plan mode 終了

プランファイルを自己レビューし、ExitPlanMode でユーザー承認を求める。

**ツール**: ExitPlanMode

### [Design Doc 作成フェーズ] — 通常 mode（承認ゲート）

#### 6. Design Doc 作成

プランファイルの調査・ヒアリング結果を基に Design Doc を作成。
作成後は**基本パターン**（自己レビュー → ユーザーレビュー → 修正）に従い、ユーザー承認を得る。

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

**信頼度マーク:**

要件・設計の各項目に、情報源の確かさを示すマークを付与する。

- ✅ **確認済み** — ユーザー発言・コード・ドキュメントから直接確認した事実
- ⚠️ **推測** — 確認済み情報からの妥当な推測
- ❓ **仮定** — 情報源がない仮定。実装前にユーザー確認が必要

```markdown
## 要件（例）
- ✅ 認証は既存の JWT ミドルウェアを使用（auth.go:L42 で確認）
- ⚠️ トークンの有効期限は24時間（現在の設定から推測）
- ❓ リフレッシュトークンの要否は未確認
```

マークの付与対象: **要件**、**設計**、**実装詳細** セクションの各項目。
背景・目的、スコープ、関連コードには不要。

**⛔ Design Doc はコードを書く場所ではない:**

Design Doc の目的は「何をどう作るか」の設計判断を伝えること。実装コードを書くことではない。
コードは設計判断を補足する最小限の断片のみ許可する。

- **書いてよいコード**: 型定義・インターフェース、データ構造、非自明な分岐ロジック（数行）
- **書いてはいけないコード**: クラスの完全実装、関数の本体、初期化処理、ボイラープレート、ユーティリティ関数

**判断基準**: そのコードブロックを削除しても設計意図が伝わるなら、そのコードは不要。

**定量目安**: コードブロックの合計行数がドキュメント全体の 1/3 を超えたら書きすぎ。

**Bad（設計ドキュメントに不適切）**:

```typescript
// ❌ クラスの完全実装を書いている
class MovingState implements State<GameData> {
  name() { return "Moving"; }
  entry(machine: EntryMachine<GameData>, event: object): void {
    const data = machine.value();
    if (event instanceof MoveEvent) {
      data.moveDirection = event.direction;
      data.facing = event.direction;
    }
    const tick = (m: AfterFuncMachine<GameData>): void => {
      const d = m.value();
      const dx = d.moveDirection === "right" ? SPEED : -SPEED;
      d.playerX = Math.max(0, Math.min(d.playerX + dx, d.sceneWidthPx - W));
      updateCamera(d);
      checkSceneTransition(m, d);
      m.afterFunc(d.dispatcher, TICK_MS, tick);
    };
    machine.afterFunc(data.dispatcher, TICK_MS, tick);
  }
}
```

**Good（設計判断のみ伝える）**:

> Moving state は afterFunc チェーン（16ms 間隔）でピクセル単位移動を実現する。
> self-transition で前回チェーンが自動キャンセルされるため、方向変更・停止が安全。

```typescript
// 移動 tick 内でシーン遷移ゾーンを検出 → StopMoveEvent + EnterTownEvent
machine.afterFunc(dispatcher, TICK_MS, tick);
```

**その他の注意:**
- 1つのコード例 = 1つの設計判断（複数の関心事を混ぜない）
- 自明な処理は `// ...validation...` のように省略コメントで飛ばす
- コードよりテキスト・表・Mermaid 図が適切な場合はそちらを使う

**ツール**: Write, Edit, AskUserQuestion

**⛔ ユーザーの承認なしに次へ進まない**

## 自己レビュー観点

### プラン（調査フェーズ）
- ヒアリングすべき不明点を見逃していないか

### Design Doc
- What/Why が明確か
- 要件に曖昧さが残っていないか
- 機能要件はテスト可能な形式か（「〜したら〜になる」のように、操作と結果が明確か）
- 関連コードの調査は十分か
- スコープは明確か
- 代替案の検討は十分か
- **⛔ コード量チェック**: コードブロック合計行数がドキュメント全体の 1/3 を超えていないか
- **⛔ コード必要性チェック**: 各コードブロックを削除しても設計意図が伝わるか？伝わるなら削除する
- **⛔ 実装混入チェック**: クラスの完全実装、関数本体、初期化処理が含まれていないか
- 設計を全部コードで表現しようとしていないか（テキストや図表で十分な箇所はないか）
- 各コード例が1つの設計判断に焦点を当てているか（複数の関心事が混在していないか）
- `.claude/rules/writing-style.instructions.md` の簡潔さの原則に従っているか
- 信頼度マーク: ❓（仮定）が多い項目はヒアリングで解消できないか検討

観点は内容に応じて調整。不明確な場合はユーザーにヒアリング。

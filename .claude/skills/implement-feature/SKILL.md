---
name: implement-feature
description: >
  スタブ / 未実装マーカーを、ユニット単位の TDD で実装に変える。
disable-model-invocation: true
---

# Implement Feature

## 1. 概要

未実装マーカー（`panic("implement me ...")` / `t.Fatal("not implemented ...")` 等の言語慣用の中断）が残るスタブを、ユニット単位の TDD で実装に変える。マーカーが進捗そのものであり、全て消えてテストが通り、`design.md` の検証方法が満たされた状態が完了。別の進捗管理ドキュメントは持たない。

入力は設計スタブ・テストスタブと、あれば薄い `design.md`。`sketch-feature` の成果物が代表的な入力だが、未実装マーカー付きのスタブであれば手書きのものも対象。

動作環境別の手段は [environment.md](environment.md) を参照。参照する観点は §5。

## 2. 原則

- **TDD ループ**: ユニット（型・関数・メソッド）ごとに「テストスタブを具体化（赤）→ 本体を実装（緑）」を 1 ユニットずつ回す。複数を同時に着手しない
- **ビルド維持**: 各ステップ後にビルド・型検査が通る状態を保つ。未着手ユニットのマーカーは残してよい（部分実装の可視化）
- **マーカーが完了判定の根拠**: 未実装マーカーが残る = 未実装。実装時に必ず消す。「通るはず」で報告せず、テストランナーを実行し出力を確認する
- **承認ゲート**: `design.md` の ❓ 仮定は実装前に解消する。実装中に設計の穴が判明したら、設計（`design.md`・設計スタブ）に戻して起こし直す
- **スコープ維持**: スタブにない機能を足さない（設計スタブが委ねる private な内部実装の追加はスコープ内）。実装中に必要が判明した変更は依頼外作業として切り分け、ユーザーに諮る
- **コメントの整理**: 実装したら、そのユニットの未実装ヒントの `// TODO` を消す（将来の改善メモは残す）。コードと重複するコメントは削除し、コードだけでは追いにくい処理（非自明な前提・トレードオフ・所有権）だけコメントで補足する。公開境界の doc コメント（契約・設計意図）は残す
- **完了報告の率直さ**: 残った未実装・前提崩れ・既知の不具合を隠さず列挙する

## 3. 起動と再開

### 3.1 起動条件

明示起動。未実装マーカー付きのスタブが対象。マーカーが見当たらない場合は、先にスタブ・テストスタブを起こす（`/sketch-feature` 等）よう案内する。受け入れ条件はテストスタブのケース名・Given/When/Then と doc コメントの契約から読み取る。`design.md` があれば目的・スコープ・検証方法の根拠として併用し、不足するなら実装前にユーザーへ確認する。

### 3.2 再開

ユニット単位で進むため、残った未実装マーカーから再開できる。どのユニットから再開するかをユーザーに確認してから進める。

## 4. フェーズ

### 4.1 準備

- **インプット**: 設計スタブ・テストスタブ・`design.md`・既存テスト基盤（[testing.md](../sketch-feature/references/testing.md) の「既存テスト基盤の流用」）の把握。実装対象コードの現状の挙動・隣接コードからの依存を確認する
- **アウトプット**: 未実装マーカーを全件洗い出した実装ユニットの一覧と、依存の末端から積み上げる実装順。一覧は [environment.md](environment.md) の TODO 管理で追う
- **確認**: `design.md` の ❓ をユーザー確認で解消する（[environment.md](environment.md) の構造化ヒアリング）

### 4.2 ユニット実装ループ

各ユニットで次を回す。

1. **赤**: テストスタブの `t.Fatal` を外し、Given/When/Then を実コードに具体化する。ケースは `design.md` の目的・スコープと設計スタブの責務（doc コメントの契約）から導く。テストを実行し、未実装（本体の `panic`）ゆえに失敗することを確認する
2. **緑**: 本体を実装して `panic` を外す。テストを実行し、通ることを出力で確認する
3. **確認**: ビルド・型検査・lint が通ること。実装した範囲のコメントを整理する（原則の「コメントの整理」）。テスト具体化は [testing.md](../sketch-feature/references/testing.md)、横断的な盲点は §5 の観点でセルフレビューする

失敗時は [root-cause-analysis.md](../sketch-feature/references/root-cause-analysis.md) で原因を分類する。同じ修正で 2 回失敗したら別アプローチ、3 回でユーザーに相談する（[environment.md](environment.md) の巻き戻し・コンテキスト clear を併用）。設計の問題なら 4.1 へ戻す。

### 4.3 完了検証

- 未実装マーカーが全件消えたことをコード検索で確認する
- 全テストが green であることを実行して確認する
- `design.md` の「検証方法」節を実行する（lint・型・コマンド出力・計測・手動）。`design.md` が無い場合は、ビルド・lint・テストに加え §3.1 で確認した要件の充足を確かめる
- 残課題・前提崩れを率直に報告して終了する

## 5. 参照する観点

セルフレビュー時の観点は固定リストを読み上げず、[derive-perspectives.md](../sketch-feature/references/derive-perspectives.md) の手順でその場で導出し、[missable-checklist.md](../sketch-feature/references/missable-checklist.md)（自力で出にくい取りこぼし項目）と突き合わせる。

導出を補助する手順ファイル:

- テスト具体化・テスト品質 → [testing.md](../sketch-feature/references/testing.md)
- 完了の定義・検証手段の確認 → [verification-strategy.md](../sketch-feature/references/verification-strategy.md)
- 失敗時の分析・エスカレーション → [root-cause-analysis.md](../sketch-feature/references/root-cause-analysis.md)
- エージェント特有の盲点（テスト実行の実証・完了報告バイアス） → [agent-pitfalls.md](../sketch-feature/references/agent-pitfalls.md)

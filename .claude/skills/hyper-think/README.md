# Hyper Think

着眼点が直交する複数のサブエージェントを並列実行し、1 つの視点では見落とす論点を出典タグ付きで統合する深い思考スキル。明示起動のみで、引数のキーワードから `review` / `ideate` / `opine` のモードを判定する。

## 実行例

### review モード（既存物から問題抽出）

```
❯ /hyper-think この PR をレビューして

● review モードで実行します。複数の観点で並列分析した結果を統合します。

  | # | 深刻度 | 観点 | 指摘 | 該当箇所 | 出典 |
  |---|---|---|---|---|---|
  | 1 | Critical | 認証 | トークン検証が省略 | auth.ts:42 | critical-analyst |
  | 2 | Major | エラー処理 | リトライ時の重複登録 | api.ts:88 | stakeholder-sim |
```

### ideate モード（選択肢の発散）

```
❯ /hyper-think 認証フローのアーキテクチャ案を出して

● ideate モードで実行します。複数視点から発散と収束を行います。

  1. **JWT + Refresh Token** (from: framework-analyst)
     - 利点: ステートレス、標準実装が豊富、スケール容易
     - 欠点: 失効が困難、refresh フロー追加で複雑化
  2. **Opaque Token + Session Store** (from: critical-analyst)
     - 利点: 失効即時、サーバ側で完全制御
     - 欠点: セッションストアの可用性に依存

  ## 推奨: 案 ..（理由: ..）
```

### opine モード（所感・論点・前提）

```
❯ /hyper-think この認証設計どう思う?

● opine モードで実行します。多視点で所感・論点・前提を出します。

  - 重要論点: トークン有効期限 30 日は UX 優先、セッション固定脅威モデルに対して長すぎる (from: critical-analyst)
  - 着目した点: 既存の強力な MFA を活かせば短期トークン + 自動更新で UX を保てる (from: constructive-analyst)
  - 検証すべき前提: ログアウト操作と整合する revocation 経路が未設計（前提が暗黙化している） (from: framework-analyst)
```

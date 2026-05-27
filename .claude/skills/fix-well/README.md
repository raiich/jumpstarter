# Fix Well

修正系の指示を受けたとき、指摘箇所だけを場当たり的に直さず、影響範囲・同類問題・整合性を先に洗い出して 1 回でまとめて修正するスキル。

## 実行例

```
❯ /fix-well user_id を account_id にリネームして

● fix-well を実行します。影響範囲を先に洗い出して 1 回でまとめて修正します。

  ## 修正完了: user_id → account_id

  - 定義: src/models/user.ts:42
  - 参照: src/api/auth.ts, src/api/users.ts（5 箇所）
  - 呼び出し元: src/components/*.tsx（8 箇所）
  - ドキュメント: docs/api.md, README.md

  ## 調査したが該当しなかった範囲

  - src/legacy/*: 古い別概念
  - migrations/*: DB カラム名は別タスク
```

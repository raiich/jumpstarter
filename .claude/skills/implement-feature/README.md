# Implement Feature

未実装マーカーが残るスタブを、ユニット単位の TDD で実装に変え、テスト通過と検証方法の充足まで仕上げるスキル。

## 実行例

```
❯ /implement-feature registry のスタブを実装して

● implement-feature を実行します。未実装マーカーを洗い出し、ユニットごとに 赤 → 緑 で埋めます。

  [準備] 未実装マーカー 5 件を洗い出し → 実装ユニット一覧（依存の末端 Add → Close の順）
         design.md の ❓ 1 件をヒアリングで解消
  [ループ] Registry.Add: テスト 3 ケースを具体化（赤・panic で失敗）→ 本体実装（緑）→ // TODO 削除
           Registry.Close: 同様に 赤 → 緑
  [完了検証] 未実装マーカー 0 件・全テスト green・design.md の検証方法（go test / go vet）を実行
  → 残課題なしを報告して終了

  [変更ファイル]
  registry/
  ├── registry.go         # panic を実装に置換、未実装ヒントの TODO を削除
  └── registry_test.go    # t.Fatal を Given/When/Then の実コードに置換
```

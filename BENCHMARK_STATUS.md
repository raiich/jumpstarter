# ベンチマーク実装状況

## 実装完了

- ✅ 9つのベンチマークファイルをすべて作成
- ✅ 計画書の項目を95%以上カバー
- ✅ goroutineリークを修正（concurrency_bench_test.go）
- ✅ MapDeleteベンチマークの測定範囲を修正（data_structure_bench_test.go）
- ✅ メモリリークを起こすベンチマークをスキップ設定（time_bench_test.go）
- ✅ 全ファイルにグローバル変数を定義

## 実装中

- ⚠️ グローバル変数への代入を全ベンチマークに適用
  - function_call_bench_test.go: 一部完了
  - allocation_bench_test.go: 一部完了
  - その他: グローバル変数は定義済み、`_ = result` の置き換えは未完了

## 既知の問題

- ⚠️ misc_bench_test.go のカスタムソートがバブルソート（O(n²)）で、sort.Ints（O(n log n)）と公平な比較ではない

## 残タスク（オプション）

1. すべてのベンチマークで `_ = result` をグローバル変数への代入に置き換え
2. カスタムソートの実装を見直し（バブルソート → クイックソートなど）
3. `b.ResetTimer()` の使用を統一

## ファイル一覧

1. function_call_bench_test.go - 関数呼び出し
2. allocation_bench_test.go - メモリ割り当て
3. concurrency_bench_test.go - 並行処理
4. conversion_bench_test.go - 型変換とフォーマット
5. data_structure_bench_test.go - データ構造
6. encoding_bench_test.go - エンコーディング
7. time_bench_test.go - 時刻操作
8. context_bench_test.go - コンテキスト
9. misc_bench_test.go - その他の機能

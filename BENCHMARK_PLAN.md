# Go言語 オーバーヘッド計測ベンチマーク計画

## 目的
Go言語の各種機能がどれだけのオーバーヘッドを持つかを計測し、パフォーマンスの最適化に役立てる。

## 計測対象の機能

### 1. 関数呼び出し (function_call_bench_test.go)
- ベースライン（何もしない処理）
- 単純な関数呼び出し
- 引数ありの関数呼び出し
- インターフェース経由の関数呼び出し
- 直接メソッド呼び出し
- defer付き関数 vs deferなし関数
- クロージャ呼び出し
- 再帰関数
- 可変長引数の関数

### 2. メモリ割り当て (allocation_bench_test.go)
- スタック割り当て vs ヒープ割り当て
- 小さい構造体の割り当て
- 大きい構造体の割り当て
- スライスの作成（make vs リテラル）
- スライスの append
- マップの作成と操作
- 文字列の連結（+ vs strings.Builder vs bytes.Buffer）
- インターフェース変換
- ポインタ vs 値のコピー

### 3. 並行処理 (concurrency_bench_test.go)
- goroutineの起動オーバーヘッド
- channelの送受信
- unbuffered channel vs buffered channel
- mutexのロック/アンロック
- RWMutexの読み込みロック vs 書き込みロック
- sync.Once
- sync.Pool
- atomic操作 vs mutex
- select文のオーバーヘッド

### 4. データ構造 (data_structure_bench_test.go)
- 配列 vs スライス
- スライスのアクセス vs マップのアクセス
- リンクリストの操作
- range によるイテレーション（インデックス vs 値）
- for文 vs range文

### 5. その他の機能 (misc_bench_test.go)
- 型アサーション
- 型switch
- panic/recover
- reflectionの使用
- エラーハンドリング（error返却 vs panic）
- 空インターフェースの使用
- ジェネリクス vs インターフェース（Go 1.18+）

## ベンチマークの実行方法

```bash
# すべてのベンチマークを実行
go test -bench=. -benchmem

# 特定のベンチマークを実行
go test -bench=BenchmarkFunctionCall -benchmem

# 詳細な統計情報を表示
go test -bench=. -benchmem -benchtime=10s

# CPUプロファイルを取得
go test -bench=. -cpuprofile=cpu.prof

# メモリプロファイルを取得
go test -bench=. -memprofile=mem.prof
```

## 期待される出力
各ベンチマークで以下の情報を取得:
- 操作あたりの実行時間 (ns/op)
- 操作あたりのメモリ割り当て量 (B/op)
- 操作あたりの割り当て回数 (allocs/op)

## ファイル構成
```
jumpstarter/
├── BENCHMARK_PLAN.md              # この計画書
├── BENCHMARK_RESULTS.md           # ベンチマーク結果（実行後に作成）
├── function_call_bench_test.go    # 関数呼び出しのベンチマーク
├── allocation_bench_test.go       # メモリ割り当てのベンチマーク
├── concurrency_bench_test.go      # 並行処理のベンチマーク
├── data_structure_bench_test.go   # データ構造のベンチマーク
└── misc_bench_test.go             # その他の機能のベンチマーク
```

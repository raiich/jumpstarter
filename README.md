# Jumpstarter - Go Performance Benchmarks

Go言語の各種機能のパフォーマンスオーバーヘッドを計測するベンチマークスイート。

## 概要

このプロジェクトは、Go言語の様々な機能がどれだけのオーバーヘッドを持つかを体系的に計測します。関数呼び出し、メモリ割り当て、並行処理、型変換など、160以上のベンチマークを提供します。

## クイックスタート

### すべてのベンチマークを実行

```bash
go test -bench=. -benchmem
```

### 特定のカテゴリのみ実行

```bash
# 関数呼び出しのベンチマークのみ
go test -bench=BenchmarkFunctionCall -benchmem

# メモリ割り当てのベンチマークのみ
go test -bench=BenchmarkAllocation -benchmem
```

### 結果をファイルに保存

```bash
go test -bench=. -benchmem -benchtime=1s > results.txt
```

## ベンチマークカテゴリ

- **関数呼び出し** - 関数、メソッド、インターフェース、defer、クロージャなど
- **メモリ割り当て** - スタック/ヒープ、構造体、スライス、マップなど
- **並行処理** - goroutine、channel、mutex、atomic操作など
- **型変換** - 数値・文字列変換、フォーマット操作など
- **データ構造** - 配列、スライス、マップのアクセスとイテレーション
- **エンコーディング** - JSON、Gob、Base64、Hexなど
- **時刻操作** - time.Now()、タイマー、ティッカーなど
- **コンテキスト** - context作成、値の取得、キャンセル伝播など
- **その他** - 型アサーション、リフレクション、エラーハンドリングなど

## ドキュメント

- **[ベンチマークカタログ](docs/CATALOG.md)** - 実装済みベンチマークの全一覧
- **[ベンチマークの書き方](docs/WRITING_BENCHMARKS.md)** - 新しいベンチマークを追加する方法
- **[実行ガイド](docs/RUNNING_BENCHMARKS.md)** - 詳細な実行方法とプロファイリング

## ベンチマーク結果の見方

```
BenchmarkSimpleFunction-8    100000000    10.5 ns/op    0 B/op    0 allocs/op
```

- `BenchmarkSimpleFunction`: ベンチマーク名
- `-8`: 使用したCPU数（GOMAXPROCS）
- `100000000`: 実行された反復回数
- `10.5 ns/op`: 操作あたりの平均実行時間（ナノ秒）
- `0 B/op`: 操作あたりの平均メモリ割り当て量（バイト）
- `0 allocs/op`: 操作あたりの平均割り当て回数

## 動作環境

- Go 1.20以降

## ライセンス

MIT License

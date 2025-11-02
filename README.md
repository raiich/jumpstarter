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
- **[実行ガイド](docs/RUNNING_BENCHMARKS.md)** - 実行方法と結果の見方
- **[プロファイリング](docs/PROFILING.md)** - 詳細な性能分析

## 動作環境

- Go 1.20以降

## ライセンス

MIT License

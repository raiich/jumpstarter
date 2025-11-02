# ベンチマーク実行ガイド

このドキュメントでは、ベンチマークの実行方法、結果の分析、プロファイリング手法について説明します。

## 基本的な実行方法

### すべてのベンチマークを実行

```bash
go test -bench=. -benchmem
```

- `-bench=.`: すべてのベンチマークを実行（正規表現で指定）
- `-benchmem`: メモリ割り当て統計を表示

### 特定のベンチマークを実行

```bash
# パターンマッチで実行
go test -bench=BenchmarkFunctionCall -benchmem

# より詳細な指定
go test -bench=BenchmarkFunctionCall/NoArgs -benchmem

# 複数カテゴリを実行
go test -bench='Benchmark(FunctionCall|Allocation)' -benchmem
```

### 特定のファイルのみ実行

```bash
# 1つのファイル
go test -bench=. -benchmem function_call_bench_test.go

# 複数のファイル
go test -bench=. -benchmem function_call_bench_test.go allocation_bench_test.go
```

## 実行時間とイテレーション数の調整

### 実行時間を指定

```bash
# 各ベンチマークを10秒間実行
go test -bench=. -benchmem -benchtime=10s

# 短時間で確認（100ミリ秒）
go test -bench=. -benchmem -benchtime=100ms
```

### 固定回数のイテレーション

```bash
# 各ベンチマークを100回実行
go test -bench=. -benchmem -benchtime=100x

# 1000回実行
go test -bench=. -benchmem -benchtime=1000x
```

## 結果の保存と分析

### 結果をファイルに保存

```bash
# 基本的な保存
go test -bench=. -benchmem > results.txt

# タイムスタンプ付きで保存
go test -bench=. -benchmem > "results_$(date +%Y%m%d_%H%M%S).txt"

# 標準エラーも含めて保存
go test -bench=. -benchmem 2>&1 | tee results.txt
```

### 複数回実行して統計的信頼性を高める

```bash
# 10回実行して結果を保存
go test -bench=. -benchmem -count=10 | tee bench.txt
```

## benchstat による統計分析

### インストール

```bash
go install golang.org/x/perf/cmd/benchstat@latest
```

### 単一結果の分析

```bash
# 結果の統計情報を表示
benchstat bench.txt
```

出力例：
```
name                time/op
FunctionCall/NoArgs 1.23ns ± 2%
FunctionCall/Args1  2.45ns ± 1%
```

- `±` の後の数値は変動係数（標準偏差/平均）

### 改善前後の比較

```bash
# 改善前のベンチマーク
go test -bench=. -benchmem -count=10 | tee old.txt

# コードを変更

# 改善後のベンチマーク
go test -bench=. -benchmem -count=10 | tee new.txt

# 比較分析
benchstat old.txt new.txt
```

出力例：
```
name                old time/op  new time/op  delta
FunctionCall/NoArgs  1.23ns ± 2%  0.98ns ± 1%  -20.33%  (p=0.000 n=10+10)
```

- `delta`: 変化率（負の値は改善）
- `p`: 統計的有意性（p<0.05で有意）
- `n`: サンプル数

## ベンチマーク結果の見方

### 出力例

```
BenchmarkSimpleFunction-8    100000000    10.5 ns/op    0 B/op    0 allocs/op
```

### 各項目の意味

| 項目 | 説明 |
|------|------|
| `BenchmarkSimpleFunction` | ベンチマーク名 |
| `-8` | 使用したCPU数（GOMAXPROCS） |
| `100000000` | 実行された反復回数（N） |
| `10.5 ns/op` | 操作あたりの平均実行時間 |
| `0 B/op` | 操作あたりの平均メモリ割り当て量（バイト） |
| `0 allocs/op` | 操作あたりの平均割り当て回数 |

### 時間の単位

- `ns/op`: ナノ秒（10^-9秒）
- `µs/op`: マイクロ秒（10^-6秒）
- `ms/op`: ミリ秒（10^-3秒）
- `s/op`: 秒

### 0 allocs/op の重要性

メモリ割り当てが0の場合：
- スタックのみで処理が完結
- ガベージコレクタの負荷なし
- 最も高速な実装

## 並列実行

### GOMAXPROCS の変更

```bash
# CPU数を指定して実行
go test -bench=. -benchmem -cpu=1,2,4,8

# 結果
BenchmarkExample-1    ...
BenchmarkExample-2    ...
BenchmarkExample-4    ...
BenchmarkExample-8    ...
```

### 並列ベンチマークの実行

並列ベンチマーク（`b.RunParallel`を使用）は自動的に複数goroutineで実行されます。

## トラブルシューティング

### ベンチマークが不安定な場合

```bash
# より多くのイテレーションで安定させる
go test -bench=. -benchmem -benchtime=10s

# 複数回実行して平均を取る
go test -bench=. -benchmem -count=20 | benchstat

# タイムアウトを延長
go test -bench=. -benchmem -timeout=30m
```

**環境要因の最小化:**
- ブラウザや重いアプリケーションを閉じる
- バックグラウンドプロセスを停止
- 電源に接続（バッテリー駆動だとCPUクロックが下がる）
- Linux: CPUガバナーをperformanceモードに設定

```bash
# Linux: CPUクロックを固定
sudo cpupower frequency-set --governor performance
```

### メモリ割り当てが予想より多い場合

```bash
# エスケープ解析を確認
go build -gcflags='-m' . 2>&1 | grep "escapes to heap"

# より詳細な分析
go build -gcflags='-m -m' .
```

### ベンチマークが遅すぎる場合

```bash
# 短時間で確認
go test -bench=. -benchmem -benchtime=10ms

# 特定のベンチマークのみ実行
go test -bench=BenchmarkSpecific -benchmem
```

## 高度な使い方

### ベンチマーク間の比較

```bash
# ベースラインとの比較
go test -bench=BenchmarkBaseline -benchmem -count=5 > baseline.txt
go test -bench=BenchmarkOptimized -benchmem -count=5 > optimized.txt
benchstat baseline.txt optimized.txt
```

### CI/CDでの継続的計測

```bash
# 結果を保存してリグレッション検出
go test -bench=. -benchmem -count=10 > current.txt
benchstat previous.txt current.txt

# 閾値チェック（パフォーマンス低下を検出）
benchstat -alpha=0.05 previous.txt current.txt
```

詳細な性能分析については [プロファイリング](profiling.md) を参照してください。

## 参考資料

- [Go公式: Benchmarks](https://pkg.go.dev/testing#hdr-Benchmarks)
- [benchstat ドキュメント](https://pkg.go.dev/golang.org/x/perf/cmd/benchstat)

# ベンチマーク実行ガイド

## 基本的な実行方法

### すべてのベンチマークを実行
```bash
go test -bench=. -benchmem
```

### 特定のベンチマークを実行
```bash
# パターンマッチで実行
go test -bench=BenchmarkFunctionCall -benchmem

# 特定のファイルのみ実行
go test -bench=. -benchmem function_call_bench_test.go
```

### 実行時間を調整
```bash
# 10秒間実行
go test -bench=. -benchmem -benchtime=10s

# 固定回数実行（100回）
go test -bench=. -benchmem -benchtime=100x
```

## 統計的信頼性を高める方法

### 複数回実行して結果を保存
```bash
go test -bench=. -benchmem -count=10 | tee bench.txt
```

### benchstat による結果分析

#### インストール
```bash
go install golang.org/x/perf/cmd/benchstat@latest
```

#### 単一結果の分析
```bash
benchstat bench.txt
```

#### 改善前後の比較
```bash
# 改善前
go test -bench=. -benchmem -count=10 | tee old.txt

# コード変更

# 改善後
go test -bench=. -benchmem -count=10 | tee new.txt

# 比較
benchstat old.txt new.txt
```

## プロファイリング

### CPUプロファイル
```bash
# プロファイル取得
go test -bench=. -cpuprofile=cpu.prof

# 可視化
go tool pprof -http=:8080 cpu.prof
```

### メモリプロファイル
```bash
# プロファイル取得
go test -bench=. -memprofile=mem.prof

# 可視化
go tool pprof -http=:8080 mem.prof
```

### ブロックプロファイル（並行処理）
```bash
go test -bench=. -blockprofile=block.prof
go tool pprof -http=:8080 block.prof
```

## ベンチマーク結果の見方

### 出力例
```
BenchmarkSimpleFunction-8    100000000    10.5 ns/op    0 B/op    0 allocs/op
```

### 各項目の意味
- `BenchmarkSimpleFunction`: ベンチマーク名
- `-8`: 使用したCPU数（GOMAXPROCS）
- `100000000`: 実行された反復回数
- `10.5 ns/op`: 操作あたりの平均実行時間（ナノ秒）
- `0 B/op`: 操作あたりの平均メモリ割り当て量（バイト）
- `0 allocs/op`: 操作あたりの平均割り当て回数

## ベンチマークのベストプラクティス

### 1. コンパイラ最適化の回避

```go
// デッドコード除去を防ぐ
var result int

func BenchmarkExample(b *testing.B) {
    var r int
    for i := 0; i < b.N; i++ {
        r = someFunction()
    }
    result = r  // グローバル変数に代入して最適化を防ぐ
}
```

```go
// インライン化を防ぐ
//go:noinline
func someFunction() int {
    return 42
}
```

### 2. セットアップコストの除外

```go
func BenchmarkExample(b *testing.B) {
    // セットアップ
    data := make([]int, 1000)
    for i := range data {
        data[i] = i
    }

    b.ResetTimer()  // ここまでの時間を除外

    for i := 0; i < b.N; i++ {
        // 計測対象の処理
    }
}
```

### 3. 測定したくない部分の除外

```go
func BenchmarkExample(b *testing.B) {
    for i := 0; i < b.N; i++ {
        b.StopTimer()
        // 測定から除外する処理
        data := setupData()
        b.StartTimer()

        // 計測対象の処理
        process(data)
    }
}
```

### 4. メモリ割り当ての測定

```bash
# -benchmem フラグを必ず使用
go test -bench=. -benchmem
```

### 5. 並列ベンチマーク

```go
func BenchmarkParallel(b *testing.B) {
    b.RunParallel(func(pb *testing.PB) {
        for pb.Next() {
            // 並列実行される処理
        }
    })
}
```

## トラブルシューティング

### ベンチマークが不安定な場合

```bash
# CPUクロックの固定（Linux）
sudo cpupower frequency-set --governor performance

# 他のプロセスの影響を最小化
# - ブラウザなどを閉じる
# - count を増やして統計的に安定させる
go test -bench=. -benchmem -count=20
```

### メモリ割り当てが0にならない場合

```bash
# エスケープ解析を確認
go build -gcflags='-m' .
```

## 参考リンク

- [Go公式: Writing Benchmarks](https://pkg.go.dev/testing#hdr-Benchmarks)
- [benchstat ドキュメント](https://pkg.go.dev/golang.org/x/perf/cmd/benchstat)
- [pprof ドキュメント](https://pkg.go.dev/net/http/pprof)

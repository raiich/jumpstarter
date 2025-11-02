# ベンチマークの書き方

このドキュメントでは、正確で信頼性の高いベンチマークを書くためのベストプラクティスを説明します。

## 基本構造

```go
package jumpstarter

import "testing"

// グローバル変数（コンパイラ最適化を防ぐため）
var GlobalResultInt int

func BenchmarkExample(b *testing.B) {
    var result int
    for i := 0; i < b.N; i++ {
        result = someFunction()
    }
    GlobalResultInt = result
}
```

## 必須: コンパイラ最適化の防止

### グローバル変数への代入

ベンチマーク結果がコンパイラによってデッドコード除去されないよう、**必ず**グローバル変数に結果を代入します。

```go
// ❌ 悪い例: コンパイラが最適化で削除する可能性がある
func BenchmarkBad(b *testing.B) {
    for i := 0; i < b.N; i++ {
        _ = someFunction()  // 結果を捨てている
    }
}

// ✅ 良い例: グローバル変数に代入
var GlobalResultInt int

func BenchmarkGood(b *testing.B) {
    var result int
    for i := 0; i < b.N; i++ {
        result = someFunction()
    }
    GlobalResultInt = result  // 最適化を防ぐ
}
```

### グローバル変数の命名規則

このプロジェクトでは、以下の命名規則を採用しています：

```
Global[BenchmarkName][Type]
```

例：
- `GlobalBasicCallsInt`
- `GlobalJSONBytes`
- `GlobalStringConversionString`

### グローバル変数のスコープ

**重要**: グローバル変数は各ベンチマーク関数の直前に配置し、エクスポート（大文字始まり）します。

```go
// ✅ 正しい配置とスコープ
var GlobalExampleInt int  // エクスポート（大文字）

func BenchmarkExample(b *testing.B) {
    var result int
    for i := 0; i < b.N; i++ {
        result = someFunction()
    }
    GlobalExampleInt = result
}
```

エクスポートする理由：
- パッケージ内プライベート変数は、コンパイラが最適化する可能性がある
- エクスポートされた変数は、パッケージ外から参照可能なため最適化されにくい

### インライン化の制御

関数呼び出しのオーバーヘッドを正確に測定するため、`//go:noinline` ディレクティブを使用します。

```go
//go:noinline
func someFunction() int {
    return 42
}

func BenchmarkFunctionCall(b *testing.B) {
    var result int
    for i := 0; i < b.N; i++ {
        result = someFunction()
    }
    GlobalResultInt = result
}
```

## セットアップとクリーンアップ

### b.ResetTimer() の使用

セットアップ処理の時間を測定から除外します。

```go
func BenchmarkWithSetup(b *testing.B) {
    // セットアップ処理（測定対象外）
    data := make([]int, 1000)
    for i := range data {
        data[i] = i
    }

    b.ResetTimer()  // ここまでの時間を除外

    // 計測対象の処理
    var sum int
    for i := 0; i < b.N; i++ {
        sum = 0
        for _, v := range data {
            sum += v
        }
    }
    GlobalSum = sum
}
```

### b.StopTimer() / b.StartTimer() の使用

ループ内で測定したくない部分を除外します。

```go
func BenchmarkWithExclusion(b *testing.B) {
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

**注意**: StopTimer/StartTimer はオーバーヘッドがあるため、可能であれば b.ResetTimer() を使用してください。

## 並行処理のベストプラクティス

### goroutineリークの防止

並行処理ベンチマークでは、goroutineを正しく終了させます。

```go
func BenchmarkConcurrent(b *testing.B) {
    ch := make(chan int)
    done := make(chan struct{})

    // ワーカーgoroutine
    go func() {
        for {
            select {
            case <-ch:
                // 処理
            case <-done:
                return  // 正しく終了
            }
        }
    }()

    b.ResetTimer()

    for i := 0; i < b.N; i++ {
        ch <- i
    }

    b.StopTimer()
    close(done)  // goroutineを終了
}
```

### メモリリークの防止

`time.After` と `time.Tick` は停止できずメモリリークの原因となるため、ベンチマークでは使用を避けます。

```go
// ❌ 悪い例: メモリリークを起こす
func BenchmarkBadTimer(b *testing.B) {
    for i := 0; i < b.N; i++ {
        <-time.After(time.Millisecond)  // タイマーが溜まり続ける
    }
}

// ✅ 良い例: 停止可能なタイマーを使用
func BenchmarkGoodTimer(b *testing.B) {
    for i := 0; i < b.N; i++ {
        timer := time.NewTimer(time.Millisecond)
        <-timer.C
        timer.Stop()  // 明示的に停止
    }
}
```

やむを得ず計測が必要な場合は `b.Skip()` でスキップします：

```go
func BenchmarkAfter(b *testing.B) {
    b.Skip("time.After creates timers that cannot be stopped, causing memory leaks in benchmarks")
    // ...
}
```

## サブベンチマーク

関連するベンチマークをグループ化します。

```go
func BenchmarkStringConcat(b *testing.B) {
    parts := []string{"Hello", "World", "From", "Go"}

    b.Run("Plus", func(b *testing.B) {
        var result string
        for i := 0; i < b.N; i++ {
            result = parts[0] + " " + parts[1] + " " + parts[2] + " " + parts[3]
        }
        GlobalString = result
    })

    b.Run("Join", func(b *testing.B) {
        var result string
        for i := 0; i < b.N; i++ {
            result = strings.Join(parts, " ")
        }
        GlobalString = result
    })

    b.Run("Builder", func(b *testing.B) {
        var result string
        for i := 0; i < b.N; i++ {
            var sb strings.Builder
            for j, part := range parts {
                if j > 0 {
                    sb.WriteString(" ")
                }
                sb.WriteString(part)
            }
            result = sb.String()
        }
        GlobalString = result
    })
}
```

## チェックリスト

新しいベンチマークを書く際は、以下を確認してください：

- [ ] グローバル変数を定義し、結果を代入している
- [ ] グローバル変数名が `Global[BenchmarkName][Type]` パターンに従っている
- [ ] グローバル変数がエクスポート（大文字始まり）されている
- [ ] 関数呼び出しベンチマークでは `//go:noinline` を使用している
- [ ] セットアップ処理がある場合は `b.ResetTimer()` を使用している
- [ ] 並行処理ベンチマークでgoroutineを正しく終了している
- [ ] `time.After` や `time.Tick` を使用していない（またはスキップしている）
- [ ] `-benchmem` フラグで実行してメモリ割り当てを確認している

## 参考資料

- [Go公式: Writing Benchmarks](https://pkg.go.dev/testing#hdr-Benchmarks)
- [Dave Cheney: How to write benchmarks in Go](https://dave.cheney.net/2013/06/30/how-to-write-benchmarks-in-go)

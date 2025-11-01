# ベンチマークのベストプラクティス

## コンパイラ最適化の防止

ベンチマーク結果がコンパイラによってデッドコード除去されないよう、グローバル変数に結果を代入します。

```go
// ファイル冒頭でグローバル変数を定義
var globalResultInt int

// ベンチマーク内で結果をグローバル変数に代入
func BenchmarkExample(b *testing.B) {
    var result int
    for i := 0; i < b.N; i++ {
        result = someFunction()
    }
    globalResultInt = result  // _ = result ではなくグローバル変数に代入
}
```

## goroutineリークの防止

並行処理ベンチマークでは、`done` チャネルで goroutine を正しく終了させます。

```go
done := make(chan struct{})
go func() {
    for {
        select {
        case <-ch:
        case <-done:
            return  // 正しく終了
        }
    }
}()
b.ResetTimer()
// ... ベンチマーク実行 ...
b.StopTimer()
close(done)
```

## メモリリークの防止

`time.After` と `time.Tick` は停止できずメモリリークの原因となるため、これらのベンチマークは `b.Skip()` で回避します。

```go
b.Run("After", func(b *testing.B) {
    b.Skip("time.After creates timers that cannot be stopped, causing memory leaks in benchmarks")
    // ...
})
```

## インライン化の制御

関数呼び出しのオーバーヘッドを正確に測定するため、`//go:noinline` ディレクティブを使用します。

```go
//go:noinline
func someFunction() int {
    return 42
}
```

## セットアップコストの除外

セットアップ処理の時間を測定から除外するため、`b.ResetTimer()` を使用します。

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

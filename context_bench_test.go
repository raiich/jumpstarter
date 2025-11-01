package jumpstarter

import (
	"context"
	"testing"
	"time"
)

// グローバル変数（コンパイラ最適化を防ぐため）
var (
	globalCtx       context.Context
	globalCtxValue  interface{}
	globalCtxCancel context.CancelFunc
)

// ============================================================================
// コンテキストの基本操作
// ============================================================================

func BenchmarkContextCreation(b *testing.B) {
	b.Run("Background", func(b *testing.B) {
		var result context.Context
		for i := 0; i < b.N; i++ {
			result = context.Background()
		}
		globalCtx = result
	})

	b.Run("TODO", func(b *testing.B) {
		var result context.Context
		for i := 0; i < b.N; i++ {
			result = context.TODO()
		}
		globalCtx = result
	})
}

// ============================================================================
// context.WithValue
// ============================================================================

type contextKey string

func BenchmarkContextWithValue(b *testing.B) {
	key := contextKey("key")
	value := "value"

	b.Run("WithValue", func(b *testing.B) {
		ctx := context.Background()
		var result context.Context
		for i := 0; i < b.N; i++ {
			result = context.WithValue(ctx, key, value)
		}
		globalCtx = result
	})

	ctx := context.WithValue(context.Background(), key, value)

	b.Run("Value", func(b *testing.B) {
		var result interface{}
		for i := 0; i < b.N; i++ {
			result = ctx.Value(key)
		}
		globalCtxValue = result
	})
}

// ============================================================================
// context.WithCancel / WithTimeout / WithDeadline
// ============================================================================

func BenchmarkContextWithCancel(b *testing.B) {
	b.Run("WithCancel", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			ctx, cancel := context.WithCancel(context.Background())
			cancel()
			globalCtx = ctx
		}
	})

	b.Run("WithTimeout", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			ctx, cancel := context.WithTimeout(context.Background(), time.Hour)
			cancel()
			globalCtx = ctx
		}
	})

	b.Run("WithDeadline", func(b *testing.B) {
		deadline := time.Now().Add(time.Hour)
		for i := 0; i < b.N; i++ {
			ctx, cancel := context.WithDeadline(context.Background(), deadline)
			cancel()
			globalCtx = ctx
		}
	})
}

// ============================================================================
// ネストしたコンテキストからの値取得
// ============================================================================

func BenchmarkNestedContextValue(b *testing.B) {
	key := contextKey("key")
	value := "value"

	b.Run("Depth1", func(b *testing.B) {
		ctx := context.WithValue(context.Background(), key, value)
		var result interface{}
		for i := 0; i < b.N; i++ {
			result = ctx.Value(key)
		}
		globalCtxValue = result
	})

	b.Run("Depth5", func(b *testing.B) {
		ctx := context.Background()
		for i := 0; i < 5; i++ {
			ctx = context.WithValue(ctx, contextKey("dummy"), "dummy")
		}
		ctx = context.WithValue(ctx, key, value)
		b.ResetTimer()
		var result interface{}
		for i := 0; i < b.N; i++ {
			result = ctx.Value(key)
		}
		globalCtxValue = result
	})

	b.Run("Depth10", func(b *testing.B) {
		ctx := context.Background()
		for i := 0; i < 10; i++ {
			ctx = context.WithValue(ctx, contextKey("dummy"), "dummy")
		}
		ctx = context.WithValue(ctx, key, value)
		b.ResetTimer()
		var result interface{}
		for i := 0; i < b.N; i++ {
			result = ctx.Value(key)
		}
		globalCtxValue = result
	})
}

// ============================================================================
// コンテキストキャンセルの検出
// ============================================================================

func BenchmarkContextDone(b *testing.B) {
	b.Run("DoneCheck", func(b *testing.B) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		for i := 0; i < b.N; i++ {
			select {
			case <-ctx.Done():
			default:
			}
		}
	})

	b.Run("CancelledDoneCheck", func(b *testing.B) {
		ctx, cancel := context.WithCancel(context.Background())
		cancel()

		for i := 0; i < b.N; i++ {
			select {
			case <-ctx.Done():
			default:
			}
		}
	})
}

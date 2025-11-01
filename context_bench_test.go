package jumpstarter

import (
	"context"
	"testing"
	"time"
)

// ============================================================================
// コンテキストの基本操作
// ============================================================================

func BenchmarkContextCreation(b *testing.B) {
	b.Run("Background", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = context.Background()
		}
	})

	b.Run("TODO", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = context.TODO()
		}
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
		for i := 0; i < b.N; i++ {
			_ = context.WithValue(ctx, key, value)
		}
	})

	ctx := context.WithValue(context.Background(), key, value)

	b.Run("Value", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = ctx.Value(key)
		}
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
			_ = ctx
		}
	})

	b.Run("WithTimeout", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			ctx, cancel := context.WithTimeout(context.Background(), time.Hour)
			cancel()
			_ = ctx
		}
	})

	b.Run("WithDeadline", func(b *testing.B) {
		deadline := time.Now().Add(time.Hour)
		for i := 0; i < b.N; i++ {
			ctx, cancel := context.WithDeadline(context.Background(), deadline)
			cancel()
			_ = ctx
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
		for i := 0; i < b.N; i++ {
			_ = ctx.Value(key)
		}
	})

	b.Run("Depth5", func(b *testing.B) {
		ctx := context.Background()
		for i := 0; i < 5; i++ {
			ctx = context.WithValue(ctx, contextKey("dummy"), "dummy")
		}
		ctx = context.WithValue(ctx, key, value)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = ctx.Value(key)
		}
	})

	b.Run("Depth10", func(b *testing.B) {
		ctx := context.Background()
		for i := 0; i < 10; i++ {
			ctx = context.WithValue(ctx, contextKey("dummy"), "dummy")
		}
		ctx = context.WithValue(ctx, key, value)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = ctx.Value(key)
		}
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

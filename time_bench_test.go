package jumpstarter

import (
	"testing"
	"time"
)

// ============================================================================
// 時刻操作
// ============================================================================

func BenchmarkTimeOperations(b *testing.B) {
	b.Run("Now", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = time.Now()
		}
	})

	start := time.Now()

	b.Run("Since", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = time.Since(start)
		}
	})

	b.Run("NowSub", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = time.Now().Sub(start)
		}
	})
}

// ============================================================================
// タイムゾーン変換
// ============================================================================

func BenchmarkTimezoneConversion(b *testing.B) {
	t := time.Now()
	loc, _ := time.LoadLocation("America/New_York")

	b.Run("In", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = t.In(loc)
		}
	})
}

// ============================================================================
// フォーマット
// ============================================================================

func BenchmarkTimeFormatting(b *testing.B) {
	t := time.Now()

	b.Run("Format", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = t.Format(time.RFC3339)
		}
	})

	b.Run("String", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = t.String()
		}
	})
}

// ============================================================================
// タイマー操作
// ============================================================================

func BenchmarkTimer(b *testing.B) {
	b.Run("NewTimer", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			timer := time.NewTimer(time.Hour)
			timer.Stop()
		}
	})

	b.Run("After", func(b *testing.B) {
		b.Skip("time.After creates timers that cannot be stopped, causing memory leaks in benchmarks")
		for i := 0; i < b.N; i++ {
			ch := time.After(time.Hour)
			_ = ch
		}
	})

	b.Run("AfterFunc", func(b *testing.B) {
		done := make(chan bool, 1)
		for i := 0; i < b.N; i++ {
			timer := time.AfterFunc(time.Nanosecond, func() {
				done <- true
			})
			<-done
			timer.Stop()
		}
	})
}

// ============================================================================
// ティッカー操作
// ============================================================================

func BenchmarkTicker(b *testing.B) {
	b.Run("NewTicker", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			ticker := time.NewTicker(time.Hour)
			ticker.Stop()
		}
	})

	b.Run("Tick", func(b *testing.B) {
		b.Skip("time.Tick creates tickers that cannot be stopped, causing memory leaks in benchmarks")
		for i := 0; i < b.N; i++ {
			ch := time.Tick(time.Hour)
			_ = ch
		}
	})
}

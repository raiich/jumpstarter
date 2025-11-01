package jumpstarter

import (
	"testing"
	"time"
)

// グローバル変数（コンパイラ最適化を防ぐため）
var (
	globalTime     time.Time
	globalDuration time.Duration
	globalTimeStr  string
	globalTimeChan <-chan time.Time
)

// ============================================================================
// 時刻操作
// ============================================================================

func BenchmarkTimeOperations(b *testing.B) {
	b.Run("Now", func(b *testing.B) {
		var result time.Time
		for i := 0; i < b.N; i++ {
			result = time.Now()
		}
		globalTime = result
	})

	start := time.Now()

	b.Run("Since", func(b *testing.B) {
		var result time.Duration
		for i := 0; i < b.N; i++ {
			result = time.Since(start)
		}
		globalDuration = result
	})

	b.Run("NowSub", func(b *testing.B) {
		var result time.Duration
		for i := 0; i < b.N; i++ {
			result = time.Now().Sub(start)
		}
		globalDuration = result
	})
}

// ============================================================================
// タイムゾーン変換
// ============================================================================

func BenchmarkTimezoneConversion(b *testing.B) {
	t := time.Now()
	loc, err := time.LoadLocation("America/New_York")
	if err != nil {
		b.Fatal(err)
	}

	b.Run("In", func(b *testing.B) {
		var result time.Time
		for i := 0; i < b.N; i++ {
			result = t.In(loc)
		}
		globalTime = result
	})
}

// ============================================================================
// フォーマット
// ============================================================================

func BenchmarkTimeFormatting(b *testing.B) {
	t := time.Now()

	b.Run("Format", func(b *testing.B) {
		var result string
		for i := 0; i < b.N; i++ {
			result = t.Format(time.RFC3339)
		}
		globalTimeStr = result
	})

	b.Run("String", func(b *testing.B) {
		var result string
		for i := 0; i < b.N; i++ {
			result = t.String()
		}
		globalTimeStr = result
	})
}

// ============================================================================
// タイマー操作
// ============================================================================

func BenchmarkTimer(b *testing.B) {
	b.Run("NewTimer", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			timer := time.NewTimer(0)
			<-timer.C
			timer.Stop()
		}
	})

	b.Run("After", func(b *testing.B) {
		b.Skip("time.After creates timers that cannot be stopped, causing memory leaks in benchmarks")
		for i := 0; i < b.N; i++ {
			globalTimeChan = time.After(time.Hour)
		}
	})

	b.Run("AfterFunc", func(b *testing.B) {
		done := make(chan bool, 1)
		for i := 0; i < b.N; i++ {
			timer := time.AfterFunc(0, func() {
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
		ticker := time.NewTicker(time.Nanosecond)
		defer ticker.Stop()
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			<-ticker.C
		}
	})

	b.Run("Tick", func(b *testing.B) {
		b.Skip("time.Tick creates tickers that cannot be stopped, causing memory leaks in benchmarks")
		for i := 0; i < b.N; i++ {
			globalTimeChan = time.Tick(time.Hour)
		}
	})
}

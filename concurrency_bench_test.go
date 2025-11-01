package jumpstarter

import (
	"context"
	"sync"
	"sync/atomic"
	"testing"
)

// グローバル変数（コンパイラ最適化を防ぐため）
var (
	globalCounter int64
	globalBool    bool
)

// ============================================================================
// goroutineの起動と終了
// ============================================================================

func BenchmarkGoroutine(b *testing.B) {
	b.Run("StartAndFinish", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			done := make(chan bool)
			go func() {
				done <- true
			}()
			<-done
		}
	})
}

// ============================================================================
// channelの送受信
// ============================================================================

func BenchmarkChannel(b *testing.B) {
	b.Run("Unbuffered", func(b *testing.B) {
		ch := make(chan int)
		done := make(chan struct{})
		go func() {
			for {
				select {
				case <-ch:
				case <-done:
					return
				}
			}
		}()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			ch <- i
		}
		b.StopTimer()
		close(done)
	})

	b.Run("BufferSize/1", func(b *testing.B) {
		ch := make(chan int, 1)
		done := make(chan struct{})
		go func() {
			for {
				select {
				case <-ch:
				case <-done:
					return
				}
			}
		}()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			ch <- i
		}
		b.StopTimer()
		close(done)
	})

	b.Run("BufferSize/10", func(b *testing.B) {
		ch := make(chan int, 10)
		done := make(chan struct{})
		go func() {
			for {
				select {
				case <-ch:
				case <-done:
					return
				}
			}
		}()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			ch <- i
		}
		b.StopTimer()
		close(done)
	})

	b.Run("BufferSize/100", func(b *testing.B) {
		ch := make(chan int, 100)
		done := make(chan struct{})
		go func() {
			for {
				select {
				case <-ch:
				case <-done:
					return
				}
			}
		}()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			ch <- i
		}
		b.StopTimer()
		close(done)
	})

	b.Run("BufferSize/1000", func(b *testing.B) {
		ch := make(chan int, 1000)
		done := make(chan struct{})
		go func() {
			for {
				select {
				case <-ch:
				case <-done:
					return
				}
			}
		}()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			ch <- i
		}
		b.StopTimer()
		close(done)
	})
}

// ============================================================================
// Mutexのロック/アンロック
// ============================================================================

func BenchmarkMutex(b *testing.B) {
	b.Run("NoContention", func(b *testing.B) {
		var mu sync.Mutex
		for i := 0; i < b.N; i++ {
			mu.Lock()
			mu.Unlock()
		}
	})

	b.Run("WithContention", func(b *testing.B) {
		var mu sync.Mutex
		var counter int
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				mu.Lock()
				counter++
				mu.Unlock()
			}
		})
	})
}

// ============================================================================
// RWMutexの読み込みロック vs 書き込みロック
// ============================================================================

func BenchmarkRWMutex(b *testing.B) {
	b.Run("RLock", func(b *testing.B) {
		var mu sync.RWMutex
		for i := 0; i < b.N; i++ {
			mu.RLock()
			mu.RUnlock()
		}
	})

	b.Run("Lock", func(b *testing.B) {
		var mu sync.RWMutex
		for i := 0; i < b.N; i++ {
			mu.Lock()
			mu.Unlock()
		}
	})
}

// ============================================================================
// sync.Onceの初回実行 vs 2回目以降の呼び出し
// ============================================================================

func BenchmarkSyncOnce(b *testing.B) {
	b.Run("FirstCall", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			var once sync.Once
			once.Do(func() {})
		}
	})

	b.Run("SubsequentCalls", func(b *testing.B) {
		var once sync.Once
		once.Do(func() {})
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			once.Do(func() {})
		}
	})
}

// ============================================================================
// sync.PoolのGet/Put操作
// ============================================================================

func BenchmarkSyncPool(b *testing.B) {
	pool := &sync.Pool{
		New: func() interface{} {
			return new(int)
		},
	}

	b.Run("Get", func(b *testing.B) {
		var result interface{}
		for i := 0; i < b.N; i++ {
			result = pool.Get()
		}
		globalInterface = result
	})

	b.Run("Put", func(b *testing.B) {
		obj := new(int)
		for i := 0; i < b.N; i++ {
			pool.Put(obj)
		}
	})

	b.Run("GetPut", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			obj := pool.Get()
			pool.Put(obj)
		}
	})
}

// ============================================================================
// sync.WaitGroupのAdd/Done/Wait操作
// ============================================================================

func BenchmarkWaitGroup(b *testing.B) {
	b.Run("AddDone", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			var wg sync.WaitGroup
			wg.Add(1)
			wg.Done()
		}
	})

	b.Run("Wait", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			var wg sync.WaitGroup
			wg.Add(1)
			go func() {
				wg.Done()
			}()
			wg.Wait()
		}
	})
}

// ============================================================================
// atomic操作 vs mutex
// ============================================================================

func BenchmarkAtomicVsMutex(b *testing.B) {
	b.Run("AtomicLoad", func(b *testing.B) {
		var counter int64
		var result int64
		for i := 0; i < b.N; i++ {
			result = atomic.LoadInt64(&counter)
		}
		globalCounter = result
	})

	b.Run("AtomicStore", func(b *testing.B) {
		var counter int64
		for i := 0; i < b.N; i++ {
			atomic.StoreInt64(&counter, int64(i))
		}
	})

	b.Run("AtomicAdd", func(b *testing.B) {
		var counter int64
		for i := 0; i < b.N; i++ {
			atomic.AddInt64(&counter, 1)
		}
	})

	b.Run("AtomicCompareAndSwap", func(b *testing.B) {
		var counter int64
		for i := 0; i < b.N; i++ {
			atomic.CompareAndSwapInt64(&counter, 0, 1)
			atomic.CompareAndSwapInt64(&counter, 1, 0)
		}
	})

	b.Run("MutexLoad", func(b *testing.B) {
		var mu sync.Mutex
		var counter int64
		for i := 0; i < b.N; i++ {
			mu.Lock()
			globalCounter = counter
			mu.Unlock()
		}
	})

	b.Run("MutexStore", func(b *testing.B) {
		var mu sync.Mutex
		var counter int64
		for i := 0; i < b.N; i++ {
			mu.Lock()
			counter = int64(i)
			mu.Unlock()
		}
		globalCounter = counter
	})
}

// ============================================================================
// select文のオーバーヘッド
// ============================================================================

func BenchmarkSelect(b *testing.B) {
	b.Run("2Way", func(b *testing.B) {
		ch1 := make(chan int, 1)
		ch2 := make(chan int, 1)
		ch1 <- 1

		for i := 0; i < b.N; i++ {
			select {
			case v := <-ch1:
				ch1 <- v
			case <-ch2:
			}
		}
	})

	b.Run("4Way", func(b *testing.B) {
		ch1 := make(chan int, 1)
		ch2 := make(chan int, 1)
		ch3 := make(chan int, 1)
		ch4 := make(chan int, 1)
		ch1 <- 1

		for i := 0; i < b.N; i++ {
			select {
			case v := <-ch1:
				ch1 <- v
			case <-ch2:
			case <-ch3:
			case <-ch4:
			}
		}
	})

	b.Run("8Way", func(b *testing.B) {
		ch1 := make(chan int, 1)
		ch2 := make(chan int, 1)
		ch3 := make(chan int, 1)
		ch4 := make(chan int, 1)
		ch5 := make(chan int, 1)
		ch6 := make(chan int, 1)
		ch7 := make(chan int, 1)
		ch8 := make(chan int, 1)
		ch1 <- 1

		for i := 0; i < b.N; i++ {
			select {
			case v := <-ch1:
				ch1 <- v
			case <-ch2:
			case <-ch3:
			case <-ch4:
			case <-ch5:
			case <-ch6:
			case <-ch7:
			case <-ch8:
			}
		}
	})
}

// ============================================================================
// contextによるキャンセル伝播
// ============================================================================

func BenchmarkContextCancellation(b *testing.B) {
	b.Run("Propagation/1Goroutine", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			ctx, cancel := context.WithCancel(context.Background())
			done := make(chan bool)
			go func() {
				<-ctx.Done()
				done <- true
			}()
			cancel()
			<-done
		}
	})

	b.Run("Propagation/10Goroutines", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			ctx, cancel := context.WithCancel(context.Background())
			var wg sync.WaitGroup
			wg.Add(10)
			for j := 0; j < 10; j++ {
				go func() {
					<-ctx.Done()
					wg.Done()
				}()
			}
			cancel()
			wg.Wait()
		}
	})

	b.Run("Propagation/100Goroutines", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			ctx, cancel := context.WithCancel(context.Background())
			var wg sync.WaitGroup
			wg.Add(100)
			for j := 0; j < 100; j++ {
				go func() {
					<-ctx.Done()
					wg.Done()
				}()
			}
			cancel()
			wg.Wait()
		}
	})
}

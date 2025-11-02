package jumpstarter

import "testing"

// ============================================================================
// 基本的な関数呼び出し
// ============================================================================

//go:noinline
func simpleFunction() int {
	return 42
}

//go:noinline
func functionWith1Arg(a int) int {
	return a
}

//go:noinline
func functionWith3Args(a, b, c int) int {
	return a + b + c
}

//go:noinline
func functionWith10Args(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10 int) int {
	return a1 + a2 + a3 + a4 + a5 + a6 + a7 + a8 + a9 + a10
}

// グローバル変数（コンパイラ最適化を防ぐため）
var GlobalBasicCallsInt int

func BenchmarkBasicCalls(b *testing.B) {
	b.Run("Baseline", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			GlobalBasicCallsInt = i
		}
	})

	b.Run("Simple", func(b *testing.B) {
		var result int
		for i := 0; i < b.N; i++ {
			result = simpleFunction()
		}
		GlobalBasicCallsInt = result
	})

	b.Run("With1Arg", func(b *testing.B) {
		var result int
		for i := 0; i < b.N; i++ {
			result = functionWith1Arg(42)
		}
		GlobalBasicCallsInt = result
	})

	b.Run("With3Args", func(b *testing.B) {
		var result int
		for i := 0; i < b.N; i++ {
			result = functionWith3Args(1, 2, 3)
		}
		GlobalBasicCallsInt = result
	})

	b.Run("With10Args", func(b *testing.B) {
		var result int
		for i := 0; i < b.N; i++ {
			result = functionWith10Args(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
		}
		GlobalBasicCallsInt = result
	})
}

// ============================================================================
// インターフェースとメソッド呼び出し
// ============================================================================

type Calculator interface {
	Calculate() int
}

type concreteCalculator struct {
	value int
}

func (c *concreteCalculator) Calculate() int {
	return c.value
}

// グローバル変数（コンパイラ最適化を防ぐため）
var GlobalInterfaceAndMethodInt int

func BenchmarkInterfaceAndMethod(b *testing.B) {
	b.Run("InterfaceCall", func(b *testing.B) {
		var calc Calculator = &concreteCalculator{value: 42}
		var result int
		for i := 0; i < b.N; i++ {
			result = calc.Calculate()
		}
		GlobalInterfaceAndMethodInt = result
	})

	b.Run("DirectMethodCall", func(b *testing.B) {
		calc := &concreteCalculator{value: 42}
		var result int
		for i := 0; i < b.N; i++ {
			result = calc.Calculate()
		}
		GlobalInterfaceAndMethodInt = result
	})
}

// ============================================================================
// 値レシーバ vs ポインタレシーバ
// ============================================================================

type SmallStruct struct {
	a int64 // 8バイト
}

type MediumStruct struct {
	a, b, c, d int64 // 32バイト
}

type LargeStruct struct {
	a, b, c, d, e, f, g, h         int64
	i, j, k, l, m, n, o, p         int64
	q, r, s, t, u, v, w, x         int64
	y, z, aa, bb, cc, dd, ee, ff   int64 // 256バイト
}

func (s SmallStruct) ValueReceiverSmall() int64 {
	return s.a
}

func (s *SmallStruct) PointerReceiverSmall() int64 {
	return s.a
}

func (s MediumStruct) ValueReceiverMedium() int64 {
	return s.a + s.b + s.c + s.d
}

func (s *MediumStruct) PointerReceiverMedium() int64 {
	return s.a + s.b + s.c + s.d
}

func (s LargeStruct) ValueReceiverLarge() int64 {
	return s.a + s.b + s.c + s.d
}

func (s *LargeStruct) PointerReceiverLarge() int64 {
	return s.a + s.b + s.c + s.d
}

// グローバル変数（コンパイラ最適化を防ぐため）
var GlobalReceiverTypesInt64 int64

func BenchmarkReceiverTypes(b *testing.B) {
	b.Run("ValueReceiver/Small", func(b *testing.B) {
		s := SmallStruct{a: 1}
		var result int64
		for i := 0; i < b.N; i++ {
			result = s.ValueReceiverSmall()
		}
		GlobalReceiverTypesInt64 = result
	})

	b.Run("PointerReceiver/Small", func(b *testing.B) {
		s := &SmallStruct{a: 1}
		var result int64
		for i := 0; i < b.N; i++ {
			result = s.PointerReceiverSmall()
		}
		GlobalReceiverTypesInt64 = result
	})

	b.Run("ValueReceiver/Medium", func(b *testing.B) {
		s := MediumStruct{a: 1, b: 2, c: 3, d: 4}
		var result int64
		for i := 0; i < b.N; i++ {
			result = s.ValueReceiverMedium()
		}
		GlobalReceiverTypesInt64 = result
	})

	b.Run("PointerReceiver/Medium", func(b *testing.B) {
		s := &MediumStruct{a: 1, b: 2, c: 3, d: 4}
		var result int64
		for i := 0; i < b.N; i++ {
			result = s.PointerReceiverMedium()
		}
		GlobalReceiverTypesInt64 = result
	})

	b.Run("ValueReceiver/Large", func(b *testing.B) {
		s := LargeStruct{a: 1, b: 2, c: 3, d: 4}
		var result int64
		for i := 0; i < b.N; i++ {
			result = s.ValueReceiverLarge()
		}
		GlobalReceiverTypesInt64 = result
	})

	b.Run("PointerReceiver/Large", func(b *testing.B) {
		s := &LargeStruct{a: 1, b: 2, c: 3, d: 4}
		var result int64
		for i := 0; i < b.N; i++ {
			result = s.PointerReceiverLarge()
		}
		GlobalReceiverTypesInt64 = result
	})
}

// ============================================================================
// 型埋め込み（embedding）
// ============================================================================

type Base struct {
	value int
}

func (b *Base) BaseMethod() int {
	return b.value
}

type Embedded struct {
	Base
}

// グローバル変数（コンパイラ最適化を防ぐため）
var GlobalEmbeddingInt int

func BenchmarkEmbedding(b *testing.B) {
	b.Run("DirectFieldAccess", func(b *testing.B) {
		e := Embedded{Base: Base{value: 42}}
		var result int
		for i := 0; i < b.N; i++ {
			result = e.value
		}
		GlobalEmbeddingInt = result
	})

	b.Run("PromotedMethodCall", func(b *testing.B) {
		e := Embedded{Base: Base{value: 42}}
		var result int
		for i := 0; i < b.N; i++ {
			result = e.BaseMethod()
		}
		GlobalEmbeddingInt = result
	})
}

// ============================================================================
// defer付き関数 vs deferなし関数
// ============================================================================

//go:noinline
func functionWithDefer() int {
	defer func() {}()
	return 42
}

//go:noinline
func functionWithoutDefer() int {
	return 42
}

// グローバル変数（コンパイラ最適化を防ぐため）
var GlobalDeferInt int

func BenchmarkDefer(b *testing.B) {
	b.Run("WithDefer", func(b *testing.B) {
		var result int
		for i := 0; i < b.N; i++ {
			result = functionWithDefer()
		}
		GlobalDeferInt = result
	})

	b.Run("WithoutDefer", func(b *testing.B) {
		var result int
		for i := 0; i < b.N; i++ {
			result = functionWithoutDefer()
		}
		GlobalDeferInt = result
	})
}

// ============================================================================
// 特殊な呼び出しパターン
// ============================================================================

//go:noinline
func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

//go:noinline
func variadicFunction(args ...int) int {
	sum := 0
	for _, arg := range args {
		sum += arg
	}
	return sum
}

// グローバル変数（コンパイラ最適化を防ぐため）
var GlobalSpecialCallsInt int

func BenchmarkSpecialCalls(b *testing.B) {
	b.Run("Closure", func(b *testing.B) {
		value := 42
		fn := func() int {
			return value
		}
		var result int
		for i := 0; i < b.N; i++ {
			result = fn()
		}
		GlobalSpecialCallsInt = result
	})

	b.Run("Recursion", func(b *testing.B) {
		var result int
		for i := 0; i < b.N; i++ {
			result = fibonacci(10)
		}
		GlobalSpecialCallsInt = result
	})

	b.Run("Variadic", func(b *testing.B) {
		var result int
		for i := 0; i < b.N; i++ {
			result = variadicFunction(1, 2, 3, 4, 5)
		}
		GlobalSpecialCallsInt = result
	})
}

// ============================================================================
// インライン化
// ============================================================================

func inlinableFunction() int {
	return 42
}

//go:noinline
func nonInlinableFunction() int {
	return 42
}

// グローバル変数（コンパイラ最適化を防ぐため）
var GlobalInliningInt int

func BenchmarkInlining(b *testing.B) {
	b.Run("Inlinable", func(b *testing.B) {
		var result int
		for i := 0; i < b.N; i++ {
			result = inlinableFunction()
		}
		GlobalInliningInt = result
	})

	b.Run("NonInlinable", func(b *testing.B) {
		var result int
		for i := 0; i < b.N; i++ {
			result = nonInlinableFunction()
		}
		GlobalInliningInt = result
	})
}

// ============================================================================
// 間接呼び出し
// ============================================================================

// グローバル変数（コンパイラ最適化を防ぐため）
var GlobalIndirectCallsInt int

func BenchmarkIndirectCalls(b *testing.B) {
	b.Run("FunctionPointer", func(b *testing.B) {
		fn := simpleFunction
		var result int
		for i := 0; i < b.N; i++ {
			result = fn()
		}
		GlobalIndirectCallsInt = result
	})

	b.Run("MethodValue", func(b *testing.B) {
		calc := &concreteCalculator{value: 42}
		fn := calc.Calculate
		var result int
		for i := 0; i < b.N; i++ {
			result = fn()
		}
		GlobalIndirectCallsInt = result
	})
}

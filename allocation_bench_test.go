package jumpstarter

import (
	"bytes"
	"strings"
	"testing"
)

// ============================================================================
// スタック vs ヒープ割り当て
// ============================================================================

//go:noinline
func allocateOnStack() int {
	x := 42
	return x
}

//go:noinline
func allocateOnHeap() *int {
	x := 42
	return &x
}

// グローバル変数（コンパイラ最適化を防ぐため）
var (
	globalStackVsHeapInt    int
	globalStackVsHeapIntPtr *int
)

func BenchmarkStackVsHeap(b *testing.B) {
	b.Run("Stack", func(b *testing.B) {
		var result int
		for i := 0; i < b.N; i++ {
			result = allocateOnStack()
		}
		globalStackVsHeapInt = result
	})

	b.Run("Heap", func(b *testing.B) {
		var result *int
		for i := 0; i < b.N; i++ {
			result = allocateOnHeap()
		}
		globalStackVsHeapIntPtr = result
	})
}

// ============================================================================
// 構造体の割り当て（サイズ別）
// ============================================================================

type SmallStructAlloc struct {
	a int64 // 8バイト
}

type MediumStructAlloc struct {
	a, b, c, d int64 // 32バイト
}

type LargeStructAlloc struct {
	a, b, c, d, e, f, g, h         int64
	i, j, k, l, m, n, o, p         int64
	q, r, s, t, u, v, w, x         int64
	y, z, aa, bb, cc, dd, ee, ff   int64 // 256バイト
}

// グローバル変数（コンパイラ最適化を防ぐため）
var (
	globalStructAllocationSmallPtr  *SmallStructAlloc
	globalStructAllocationMediumPtr *MediumStructAlloc
	globalStructAllocationLargePtr  *LargeStructAlloc
)

func BenchmarkStructAllocation(b *testing.B) {
	b.Run("Small", func(b *testing.B) {
		var result *SmallStructAlloc
		for i := 0; i < b.N; i++ {
			result = &SmallStructAlloc{a: 1}
		}
		globalStructAllocationSmallPtr = result
	})

	b.Run("Medium", func(b *testing.B) {
		var result *MediumStructAlloc
		for i := 0; i < b.N; i++ {
			result = &MediumStructAlloc{a: 1, b: 2, c: 3, d: 4}
		}
		globalStructAllocationMediumPtr = result
	})

	b.Run("Large", func(b *testing.B) {
		var result *LargeStructAlloc
		for i := 0; i < b.N; i++ {
			result = &LargeStructAlloc{a: 1, b: 2, c: 3, d: 4}
		}
		globalStructAllocationLargePtr = result
	})
}

// ============================================================================
// スライスの事前確保 vs 動的拡張
// ============================================================================

// グローバル変数（コンパイラ最適化を防ぐため）
var globalSlicePreallocationSlice []int

func BenchmarkSlicePreallocation(b *testing.B) {
	b.Run("Preallocated", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			s := make([]int, 0, 1000)
			for j := 0; j < 1000; j++ {
				s = append(s, j)
			}
			globalSlicePreallocationSlice = s
		}
	})

	b.Run("Dynamic", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			s := make([]int, 0)
			for j := 0; j < 1000; j++ {
				s = append(s, j)
			}
			globalSlicePreallocationSlice = s
		}
	})
}

// ============================================================================
// スライスの作成（make vs リテラル）
// ============================================================================

// グローバル変数（コンパイラ最適化を防ぐため）
var globalSliceCreationSlice []int

func BenchmarkSliceCreation(b *testing.B) {
	b.Run("MakeWithCapacity", func(b *testing.B) {
		var result []int
		for i := 0; i < b.N; i++ {
			result = make([]int, 0, 10)
		}
		globalSliceCreationSlice = result
	})

	b.Run("Literal", func(b *testing.B) {
		var result []int
		for i := 0; i < b.N; i++ {
			result = []int{}
		}
		globalSliceCreationSlice = result
	})

	b.Run("LiteralWithValues", func(b *testing.B) {
		var result []int
		for i := 0; i < b.N; i++ {
			result = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		}
		globalSliceCreationSlice = result
	})
}

// ============================================================================
// マップの事前容量指定 vs 指定なし
// ============================================================================

// グローバル変数（コンパイラ最適化を防ぐため）
var globalMapPreallocationMap map[int]int

func BenchmarkMapPreallocation(b *testing.B) {
	b.Run("Preallocated", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			m := make(map[int]int, 1000)
			for j := 0; j < 1000; j++ {
				m[j] = j
			}
			globalMapPreallocationMap = m
		}
	})

	b.Run("Dynamic", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			m := make(map[int]int)
			for j := 0; j < 1000; j++ {
				m[j] = j
			}
			globalMapPreallocationMap = m
		}
	})
}

// ============================================================================
// マップの操作（追加、検索、削除）
// ============================================================================

// グローバル変数（コンパイラ最適化を防ぐため）
var (
	globalMapOperationsMap map[int]int
	globalMapOperationsInt int
)

func BenchmarkMapOperations(b *testing.B) {
	b.Run("Insert", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			m := make(map[int]int)
			m[42] = 100
			globalMapOperationsMap = m
		}
	})

	b.Run("Lookup", func(b *testing.B) {
		m := make(map[int]int)
		m[42] = 100
		var result int
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			result = m[42]
		}
		globalMapOperationsInt = result
	})

	b.Run("Delete", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			m := make(map[int]int)
			m[42] = 100
			b.StartTimer()
			delete(m, 42)
			b.StopTimer()
		}
	})
}

// ============================================================================
// 文字列の連結
// ============================================================================

// グローバル変数（コンパイラ最適化を防ぐため）
var globalStringConcatenationString string

func BenchmarkStringConcatenation(b *testing.B) {
	b.Run("PlusOperator", func(b *testing.B) {
		var result string
		for i := 0; i < b.N; i++ {
			result = "Hello" + " " + "World"
		}
		globalStringConcatenationString = result
	})

	b.Run("StringsBuilder", func(b *testing.B) {
		var result string
		for i := 0; i < b.N; i++ {
			var sb strings.Builder
			if _, err := sb.WriteString("Hello"); err != nil {
				b.Fatal(err)
			}
			if _, err := sb.WriteString(" "); err != nil {
				b.Fatal(err)
			}
			if _, err := sb.WriteString("World"); err != nil {
				b.Fatal(err)
			}
			result = sb.String()
		}
		globalStringConcatenationString = result
	})

	b.Run("BytesBuffer", func(b *testing.B) {
		var result string
		for i := 0; i < b.N; i++ {
			var buf bytes.Buffer
			if _, err := buf.WriteString("Hello"); err != nil {
				b.Fatal(err)
			}
			if _, err := buf.WriteString(" "); err != nil {
				b.Fatal(err)
			}
			if _, err := buf.WriteString("World"); err != nil {
				b.Fatal(err)
			}
			result = buf.String()
		}
		globalStringConcatenationString = result
	})
}

// ============================================================================
// 文字列とバイトスライスの変換
// ============================================================================

// グローバル変数（コンパイラ最適化を防ぐため）
var (
	globalStringByteConversionBytes  []byte
	globalStringByteConversionString string
)

func BenchmarkStringByteConversion(b *testing.B) {
	s := "Hello, World!"
	bs := []byte(s)

	b.Run("StringToBytes", func(b *testing.B) {
		var result []byte
		for i := 0; i < b.N; i++ {
			result = []byte(s)
		}
		globalStringByteConversionBytes = result
	})

	b.Run("BytesToString", func(b *testing.B) {
		var result string
		for i := 0; i < b.N; i++ {
			result = string(bs)
		}
		globalStringByteConversionString = result
	})
}

// ============================================================================
// interface{}への代入と型アサーション
// ============================================================================

// グローバル変数（コンパイラ最適化を防ぐため）
var (
	globalInterfaceConversionInterface interface{}
	globalInterfaceConversionInt       int
)

func BenchmarkInterfaceConversion(b *testing.B) {
	b.Run("IntToInterface", func(b *testing.B) {
		var result interface{}
		for i := 0; i < b.N; i++ {
			result = 42
		}
		globalInterfaceConversionInterface = result
	})

	b.Run("StructToInterface", func(b *testing.B) {
		s := SmallStructAlloc{a: 42}
		var result interface{}
		for i := 0; i < b.N; i++ {
			result = s
		}
		globalInterfaceConversionInterface = result
	})

	b.Run("PointerToInterface", func(b *testing.B) {
		s := &SmallStructAlloc{a: 42}
		var result interface{}
		for i := 0; i < b.N; i++ {
			result = s
		}
		globalInterfaceConversionInterface = result
	})

	b.Run("TypeAssertion", func(b *testing.B) {
		var iface interface{} = 42
		var result int
		for i := 0; i < b.N; i++ {
			result = iface.(int)
		}
		globalInterfaceConversionInt = result
	})
}

// ============================================================================
// ポインタ vs 値のコピー
// ============================================================================

// グローバル変数（コンパイラ最適化を防ぐため）
var (
	globalPointerVsValueSmallVal   SmallStructAlloc
	globalPointerVsValueSmallPtr   *SmallStructAlloc
	globalPointerVsValueMediumVal  MediumStructAlloc
	globalPointerVsValueMediumPtr  *MediumStructAlloc
	globalPointerVsValueLargeVal   LargeStructAlloc
	globalPointerVsValueLargePtr   *LargeStructAlloc
)

func BenchmarkPointerVsValue(b *testing.B) {
	b.Run("ValueCopy/Small", func(b *testing.B) {
		s := SmallStructAlloc{a: 42}
		var result SmallStructAlloc
		for i := 0; i < b.N; i++ {
			result = s
		}
		globalPointerVsValueSmallVal = result
	})

	b.Run("PointerCopy/Small", func(b *testing.B) {
		s := &SmallStructAlloc{a: 42}
		var result *SmallStructAlloc
		for i := 0; i < b.N; i++ {
			result = s
		}
		globalPointerVsValueSmallPtr = result
	})

	b.Run("ValueCopy/Medium", func(b *testing.B) {
		s := MediumStructAlloc{a: 1, b: 2, c: 3, d: 4}
		var result MediumStructAlloc
		for i := 0; i < b.N; i++ {
			result = s
		}
		globalPointerVsValueMediumVal = result
	})

	b.Run("PointerCopy/Medium", func(b *testing.B) {
		s := &MediumStructAlloc{a: 1, b: 2, c: 3, d: 4}
		var result *MediumStructAlloc
		for i := 0; i < b.N; i++ {
			result = s
		}
		globalPointerVsValueMediumPtr = result
	})

	b.Run("ValueCopy/Large", func(b *testing.B) {
		s := LargeStructAlloc{a: 1, b: 2, c: 3, d: 4}
		var result LargeStructAlloc
		for i := 0; i < b.N; i++ {
			result = s
		}
		globalPointerVsValueLargeVal = result
	})

	b.Run("PointerCopy/Large", func(b *testing.B) {
		s := &LargeStructAlloc{a: 1, b: 2, c: 3, d: 4}
		var result *LargeStructAlloc
		for i := 0; i < b.N; i++ {
			result = s
		}
		globalPointerVsValueLargePtr = result
	})
}

// ============================================================================
// 構造体のゼロ値初期化 vs フィールド指定初期化
// ============================================================================

// グローバル変数（コンパイラ最適化を防ぐため）
var globalStructInitializationSmallVal SmallStructAlloc

func BenchmarkStructInitialization(b *testing.B) {
	b.Run("ZeroValue", func(b *testing.B) {
		var result SmallStructAlloc
		for i := 0; i < b.N; i++ {
			result = SmallStructAlloc{}
		}
		globalStructInitializationSmallVal = result
	})

	b.Run("FieldSpecified", func(b *testing.B) {
		var result SmallStructAlloc
		for i := 0; i < b.N; i++ {
			result = SmallStructAlloc{a: 0}
		}
		globalStructInitializationSmallVal = result
	})
}

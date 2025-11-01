package jumpstarter

import (
	"reflect"
	"testing"
)

// ============================================================================
// 配列 vs スライスのインデックスアクセス
// ============================================================================

// グローバル変数（コンパイラ最適化を防ぐため）
var globalArrayVsSliceAccessInt int

func BenchmarkArrayVsSliceAccess(b *testing.B) {
	arr := [100]int{}
	slice := make([]int, 100)
	for i := 0; i < 100; i++ {
		arr[i] = i
		slice[i] = i
	}

	b.Run("ArrayAccess", func(b *testing.B) {
		var result int
		for i := 0; i < b.N; i++ {
			result = arr[i%100]
		}
		globalArrayVsSliceAccessInt = result
	})

	b.Run("SliceAccess", func(b *testing.B) {
		var result int
		for i := 0; i < b.N; i++ {
			result = slice[i%100]
		}
		globalArrayVsSliceAccessInt = result
	})
}

// ============================================================================
// スライスのインデックスアクセス vs マップのキーアクセス
// ============================================================================

// グローバル変数（コンパイラ最適化を防ぐため）
var globalSliceVsMapAccessInt int

func BenchmarkSliceVsMapAccess(b *testing.B) {
	slice := make([]int, 100)
	m := make(map[int]int, 100)
	for i := 0; i < 100; i++ {
		slice[i] = i
		m[i] = i
	}

	b.Run("SliceAccess", func(b *testing.B) {
		var result int
		for i := 0; i < b.N; i++ {
			result = slice[i%100]
		}
		globalSliceVsMapAccessInt = result
	})

	b.Run("MapAccess", func(b *testing.B) {
		var result int
		for i := 0; i < b.N; i++ {
			result = m[i%100]
		}
		globalSliceVsMapAccessInt = result
	})
}

// ============================================================================
// スライスのコピー
// ============================================================================

// グローバル変数（コンパイラ最適化を防ぐため）
var globalSliceCopySlice []int

func BenchmarkSliceCopy(b *testing.B) {
	src := make([]int, 1000)
	for i := range src {
		src[i] = i
	}

	b.Run("BuiltinCopy", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			dst := make([]int, len(src))
			copy(dst, src)
			globalSliceCopySlice = dst
		}
	})

	b.Run("ForLoop", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			dst := make([]int, len(src))
			for j := range src {
				dst[j] = src[j]
			}
			globalSliceCopySlice = dst
		}
	})

	b.Run("AppendVariadic", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			dst := append([]int{}, src...)
			globalSliceCopySlice = dst
		}
	})
}

// ============================================================================
// マップの削除
// ============================================================================

func BenchmarkMapDelete(b *testing.B) {
	b.Run("Delete", func(b *testing.B) {
		m := make(map[int]int, 100)
		for i := 0; i < 100; i++ {
			m[i] = i
		}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			key := i % 100
			delete(m, key)
			m[key] = key // 次の測定のために再度追加
		}
	})
}

// ============================================================================
// スライスのイテレーション
// ============================================================================

// グローバル変数（コンパイラ最適化を防ぐため）
var globalSliceIterationSum int

func BenchmarkSliceIteration(b *testing.B) {
	slice := make([]int, 1000)
	for i := range slice {
		slice[i] = i
	}

	b.Run("IndexLoop", func(b *testing.B) {
		var sum int
		for i := 0; i < b.N; i++ {
			sum = 0
			for j := 0; j < len(slice); j++ {
				sum += slice[j]
			}
		}
		globalSliceIterationSum = sum
	})

	b.Run("RangeIndex", func(b *testing.B) {
		var sum int
		for i := 0; i < b.N; i++ {
			sum = 0
			for j := range slice {
				sum += slice[j]
			}
		}
		globalSliceIterationSum = sum
	})

	b.Run("RangeValue", func(b *testing.B) {
		var sum int
		for i := 0; i < b.N; i++ {
			sum = 0
			for _, v := range slice {
				sum += v
			}
		}
		globalSliceIterationSum = sum
	})

	b.Run("RangeBoth", func(b *testing.B) {
		var sum int
		for i := 0; i < b.N; i++ {
			sum = 0
			for j, v := range slice {
				sum += j + v
			}
		}
		globalSliceIterationSum = sum
	})
}

// ============================================================================
// 構造体の比較
// ============================================================================

type CompareStruct struct {
	A int
	B string
	C float64
}

// グローバル変数（コンパイラ最適化を防ぐため）
var globalStructComparisonBool bool

func BenchmarkStructComparison(b *testing.B) {
	s1 := CompareStruct{A: 1, B: "test", C: 3.14}
	s2 := CompareStruct{A: 1, B: "test", C: 3.14}

	b.Run("EqualOperator", func(b *testing.B) {
		var result bool
		for i := 0; i < b.N; i++ {
			result = s1 == s2
		}
		globalStructComparisonBool = result
	})

	b.Run("DeepEqual", func(b *testing.B) {
		var result bool
		for i := 0; i < b.N; i++ {
			result = reflect.DeepEqual(s1, s2)
		}
		globalStructComparisonBool = result
	})
}

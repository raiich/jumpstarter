package jumpstarter

import (
	"reflect"
	"testing"
)

// グローバル変数（コンパイラ最適化を防ぐため）
var (
	globalDsInt    int
	globalDsBool   bool
	globalDsSum    int
	globalDsSlice  []int
)

// ============================================================================
// 配列 vs スライスのインデックスアクセス
// ============================================================================

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
		globalDsInt = result
	})

	b.Run("SliceAccess", func(b *testing.B) {
		var result int
		for i := 0; i < b.N; i++ {
			result = slice[i%100]
		}
		globalDsInt = result
	})
}

// ============================================================================
// スライスのインデックスアクセス vs マップのキーアクセス
// ============================================================================

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
		globalDsInt = result
	})

	b.Run("MapAccess", func(b *testing.B) {
		var result int
		for i := 0; i < b.N; i++ {
			result = m[i%100]
		}
		globalDsInt = result
	})
}

// ============================================================================
// スライスのコピー
// ============================================================================

func BenchmarkSliceCopy(b *testing.B) {
	src := make([]int, 1000)
	for i := range src {
		src[i] = i
	}

	b.Run("BuiltinCopy", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			dst := make([]int, len(src))
			copy(dst, src)
			globalDsSlice = dst
		}
	})

	b.Run("ForLoop", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			dst := make([]int, len(src))
			for j := range src {
				dst[j] = src[j]
			}
			globalDsSlice = dst
		}
	})

	b.Run("AppendVariadic", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			dst := append([]int{}, src...)
			globalDsSlice = dst
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
		globalDsSum = sum
	})

	b.Run("RangeIndex", func(b *testing.B) {
		var sum int
		for i := 0; i < b.N; i++ {
			sum = 0
			for j := range slice {
				sum += slice[j]
			}
		}
		globalDsSum = sum
	})

	b.Run("RangeValue", func(b *testing.B) {
		var sum int
		for i := 0; i < b.N; i++ {
			sum = 0
			for _, v := range slice {
				sum += v
			}
		}
		globalDsSum = sum
	})

	b.Run("RangeBoth", func(b *testing.B) {
		var sum int
		for i := 0; i < b.N; i++ {
			sum = 0
			for j, v := range slice {
				sum += j + v
			}
		}
		globalDsSum = sum
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

func BenchmarkStructComparison(b *testing.B) {
	s1 := CompareStruct{A: 1, B: "test", C: 3.14}
	s2 := CompareStruct{A: 1, B: "test", C: 3.14}

	b.Run("EqualOperator", func(b *testing.B) {
		var result bool
		for i := 0; i < b.N; i++ {
			result = s1 == s2
		}
		globalDsBool = result
	})

	b.Run("DeepEqual", func(b *testing.B) {
		var result bool
		for i := 0; i < b.N; i++ {
			result = reflect.DeepEqual(s1, s2)
		}
		globalDsBool = result
	})
}

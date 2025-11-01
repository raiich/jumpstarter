package jumpstarter

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

// ============================================================================
// 数値と文字列の変換
// ============================================================================

// グローバル変数（コンパイラ最適化を防ぐため）
var (
	globalNumberStringConversionString string
	globalNumberStringConversionInt    int
	globalNumberStringConversionInt64  int64
)

func BenchmarkNumberStringConversion(b *testing.B) {
	n := 12345

	b.Run("Itoa", func(b *testing.B) {
		var result string
		for i := 0; i < b.N; i++ {
			result = strconv.Itoa(n)
		}
		globalNumberStringConversionString = result
	})

	b.Run("Sprintf", func(b *testing.B) {
		var result string
		for i := 0; i < b.N; i++ {
			result = fmt.Sprintf("%d", n)
		}
		globalNumberStringConversionString = result
	})

	s := "12345"

	b.Run("ParseInt", func(b *testing.B) {
		var result int64
		var err error
		for i := 0; i < b.N; i++ {
			result, err = strconv.ParseInt(s, 10, 64)
			if err != nil {
				b.Fatal(err)
			}
		}
		globalNumberStringConversionInt64 = result
	})

	b.Run("Sscanf", func(b *testing.B) {
		var result int
		for i := 0; i < b.N; i++ {
			_, err := fmt.Sscanf(s, "%d", &result)
			if err != nil {
				b.Fatal(err)
			}
		}
		globalNumberStringConversionInt = result
	})

	b.Run("FormatInt/Base2", func(b *testing.B) {
		var result string
		for i := 0; i < b.N; i++ {
			result = strconv.FormatInt(int64(n), 2)
		}
		globalNumberStringConversionString = result
	})

	b.Run("FormatInt/Base10", func(b *testing.B) {
		var result string
		for i := 0; i < b.N; i++ {
			result = strconv.FormatInt(int64(n), 10)
		}
		globalNumberStringConversionString = result
	})

	b.Run("FormatInt/Base16", func(b *testing.B) {
		var result string
		for i := 0; i < b.N; i++ {
			result = strconv.FormatInt(int64(n), 16)
		}
		globalNumberStringConversionString = result
	})
}

// ============================================================================
// 数値型間の変換
// ============================================================================

// グローバル変数（コンパイラ最適化を防ぐため）
var (
	globalNumericConversionInt     int
	globalNumericConversionInt64   int64
	globalNumericConversionFloat64 float64
)

func BenchmarkNumericConversion(b *testing.B) {
	var i int = 12345
	var i64 int64 = 12345
	var f64 float64 = 12345.67

	b.Run("IntToInt64", func(b *testing.B) {
		var result int64
		for j := 0; j < b.N; j++ {
			result = int64(i)
		}
		globalNumericConversionInt64 = result
	})

	b.Run("Int64ToFloat64", func(b *testing.B) {
		var result float64
		for j := 0; j < b.N; j++ {
			result = float64(i64)
		}
		globalNumericConversionFloat64 = result
	})

	b.Run("Float64ToInt", func(b *testing.B) {
		var result int
		for j := 0; j < b.N; j++ {
			result = int(f64)
		}
		globalNumericConversionInt = result
	})
}

// ============================================================================
// 文字列とバイトスライスの相互変換
// ============================================================================

// グローバル変数（コンパイラ最適化を防ぐため）
var (
	globalStringByteSliceConversionBytes  []byte
	globalStringByteSliceConversionString string
)

func BenchmarkStringByteSliceConversion(b *testing.B) {
	s := "Hello, World! This is a test string."
	bs := []byte(s)

	b.Run("StringToBytes", func(b *testing.B) {
		var result []byte
		for i := 0; i < b.N; i++ {
			result = []byte(s)
		}
		globalStringByteSliceConversionBytes = result
	})

	b.Run("BytesToString", func(b *testing.B) {
		var result string
		for i := 0; i < b.N; i++ {
			result = string(bs)
		}
		globalStringByteSliceConversionString = result
	})
}

// ============================================================================
// フォーマット操作
// ============================================================================

// グローバル変数（コンパイラ最適化を防ぐため）
var globalFormattingString string

func BenchmarkFormatting(b *testing.B) {
	n := 42
	s := "world"

	b.Run("Sprintf", func(b *testing.B) {
		var result string
		for i := 0; i < b.N; i++ {
			result = fmt.Sprintf("Hello %s %d", s, n)
		}
		globalFormattingString = result
	})

	b.Run("Sprint", func(b *testing.B) {
		var result string
		for i := 0; i < b.N; i++ {
			result = fmt.Sprint("Hello ", s, " ", n)
		}
		globalFormattingString = result
	})

	parts := []string{"Hello", "World", "From", "Go"}

	b.Run("Join", func(b *testing.B) {
		var result string
		for i := 0; i < b.N; i++ {
			result = strings.Join(parts, " ")
		}
		globalFormattingString = result
	})

	b.Run("PlusOperator", func(b *testing.B) {
		var result string
		for i := 0; i < b.N; i++ {
			result = parts[0] + " " + parts[1] + " " + parts[2] + " " + parts[3]
		}
		globalFormattingString = result
	})

	b.Run("Builder", func(b *testing.B) {
		var result string
		for i := 0; i < b.N; i++ {
			var sb strings.Builder
			for j, part := range parts {
				if j > 0 {
					if _, err := sb.WriteString(" "); err != nil {
						b.Fatal(err)
					}
				}
				if _, err := sb.WriteString(part); err != nil {
					b.Fatal(err)
				}
			}
			result = sb.String()
		}
		globalFormattingString = result
	})
}

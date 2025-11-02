package jumpstarter

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"sort"
	"strings"
	"testing"
)

// ============================================================================
// 型アサーション
// ============================================================================

// グローバル変数（コンパイラ最適化を防ぐため）
var (
	GlobalTypeAssertionInt    int
	GlobalTypeAssertionString string
)

func BenchmarkTypeAssertion(b *testing.B) {
	var iface interface{} = 42

	b.Run("Success", func(b *testing.B) {
		var result int
		for i := 0; i < b.N; i++ {
			result = iface.(int)
		}
		GlobalTypeAssertionInt = result
	})

	b.Run("SuccessWithCheck", func(b *testing.B) {
		var result int
		for i := 0; i < b.N; i++ {
			var ok bool
			result, ok = iface.(int)
			if !ok {
				b.Fatal("type assertion failed")
			}
		}
		GlobalTypeAssertionInt = result
	})

	b.Run("FailureWithCheck", func(b *testing.B) {
		var result string
		for i := 0; i < b.N; i++ {
			var ok bool
			result, ok = iface.(string)
			if ok {
				b.Fatal("type assertion should have failed")
			}
		}
		GlobalTypeAssertionString = result
	})
}

// ============================================================================
// 型switch
// ============================================================================

// グローバル変数（コンパイラ最適化を防ぐため）
var GlobalTypeSwitchInt int

func BenchmarkTypeSwitch(b *testing.B) {
	var iface interface{} = 42

	b.Run("2Way", func(b *testing.B) {
		var result int
		for i := 0; i < b.N; i++ {
			switch v := iface.(type) {
			case int:
				result = v
			case string:
				result = len(v)
			}
		}
		GlobalTypeSwitchInt = result
	})

	b.Run("5Way", func(b *testing.B) {
		var result int
		for i := 0; i < b.N; i++ {
			switch v := iface.(type) {
			case int:
				result = v
			case string:
				result = len(v)
			case bool:
				result = 0
			case float64:
				result = int(v)
			case []byte:
				result = len(v)
			}
		}
		GlobalTypeSwitchInt = result
	})

	b.Run("10Way", func(b *testing.B) {
		var result int
		for i := 0; i < b.N; i++ {
			switch v := iface.(type) {
			case int:
				result = v
			case string:
				result = len(v)
			case bool:
				result = 0
			case float64:
				result = int(v)
			case []byte:
				result = len(v)
			case int64:
				result = int(v)
			case uint:
				result = int(v)
			case []int:
				result = len(v)
			case map[string]int:
				result = len(v)
			case *int:
				result = *v
			}
		}
		GlobalTypeSwitchInt = result
	})
}

// ============================================================================
// panic/recover
// ============================================================================

func panicFunction() {
	panic("error")
}

func recoverFunction() (recovered bool) {
	defer func() {
		if r := recover(); r != nil {
			recovered = true
		}
	}()
	panicFunction()
	return false
}

// グローバル変数（コンパイラ最適化を防ぐため）
var GlobalPanicRecoverBool bool

func BenchmarkPanicRecover(b *testing.B) {
	var result bool
	for i := 0; i < b.N; i++ {
		result = recoverFunction()
	}
	GlobalPanicRecoverBool = result
}

// ============================================================================
// reflection操作
// ============================================================================

type ReflectStruct struct {
	Field1 int
	Field2 string
}

func (r *ReflectStruct) Method() int {
	return r.Field1
}

// グローバル変数（コンパイラ最適化を防ぐため）
var (
	GlobalReflectionType   reflect.Type
	GlobalReflectionValue  reflect.Value
	GlobalReflectionInt    int
	GlobalReflectionValues []reflect.Value
)

func BenchmarkReflection(b *testing.B) {
	s := &ReflectStruct{Field1: 42, Field2: "test"}

	b.Run("TypeOf", func(b *testing.B) {
		var result reflect.Type
		for i := 0; i < b.N; i++ {
			result = reflect.TypeOf(s)
		}
		GlobalReflectionType = result
	})

	b.Run("ValueOf", func(b *testing.B) {
		var result reflect.Value
		for i := 0; i < b.N; i++ {
			result = reflect.ValueOf(s)
		}
		GlobalReflectionValue = result
	})

	b.Run("FieldAccess/Direct", func(b *testing.B) {
		var result int
		for i := 0; i < b.N; i++ {
			result = s.Field1
		}
		GlobalReflectionInt = result
	})

	b.Run("FieldAccess/Reflection", func(b *testing.B) {
		v := reflect.ValueOf(s).Elem()
		var result int
		for i := 0; i < b.N; i++ {
			result = int(v.Field(0).Int())
		}
		GlobalReflectionInt = result
	})

	b.Run("MethodCall/Direct", func(b *testing.B) {
		var result int
		for i := 0; i < b.N; i++ {
			result = s.Method()
		}
		GlobalReflectionInt = result
	})

	b.Run("MethodCall/Reflection", func(b *testing.B) {
		v := reflect.ValueOf(s)
		method := v.MethodByName("Method")
		var result []reflect.Value
		for i := 0; i < b.N; i++ {
			result = method.Call(nil)
		}
		GlobalReflectionValues = result
	})
}

// ============================================================================
// エラーハンドリング
// ============================================================================

func returnError() error {
	return errors.New("error")
}

// グローバル変数（コンパイラ最適化を防ぐため）
var (
	GlobalErrorHandlingErr       error
	GlobalErrorHandlingInterface interface{}
)

func BenchmarkErrorHandling(b *testing.B) {
	b.Run("ErrorReturn", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			err := returnError()
			if err != nil {
				GlobalErrorHandlingErr = err
			}
		}
	})

	b.Run("PanicRecover", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			func() {
				defer func() {
					if r := recover(); r != nil {
						GlobalErrorHandlingInterface = r
					}
				}()
				panic("error")
			}()
		}
	})
}

// ============================================================================
// interface{}の操作
// ============================================================================

// グローバル変数（コンパイラ最適化を防ぐため）
var (
	GlobalInterfaceOperationsInterface interface{}
	GlobalInterfaceOperationsInt       int
)

func BenchmarkInterfaceOperations(b *testing.B) {
	b.Run("AssignInt", func(b *testing.B) {
		var iface interface{}
		for i := 0; i < b.N; i++ {
			iface = 42
		}
		GlobalInterfaceOperationsInterface = iface
	})

	b.Run("AssignStruct", func(b *testing.B) {
		s := ReflectStruct{Field1: 42, Field2: "test"}
		var iface interface{}
		for i := 0; i < b.N; i++ {
			iface = s
		}
		GlobalInterfaceOperationsInterface = iface
	})

	b.Run("AssignPointer", func(b *testing.B) {
		s := &ReflectStruct{Field1: 42, Field2: "test"}
		var iface interface{}
		for i := 0; i < b.N; i++ {
			iface = s
		}
		GlobalInterfaceOperationsInterface = iface
	})

	b.Run("ExtractInt", func(b *testing.B) {
		var iface interface{} = 42
		var result int
		for i := 0; i < b.N; i++ {
			result = iface.(int)
		}
		GlobalInterfaceOperationsInt = result
	})
}

// ============================================================================
// ジェネリクス vs interface{}
// ============================================================================

func genericMax[T int | float64](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func interfaceMax(a, b interface{}) interface{} {
	switch av := a.(type) {
	case int:
		bv, ok := b.(int)
		if !ok {
			return nil
		}
		if av > bv {
			return av
		}
		return bv
	case float64:
		bv, ok := b.(float64)
		if !ok {
			return nil
		}
		if av > bv {
			return av
		}
		return bv
	}
	return nil
}

// グローバル変数（コンパイラ最適化を防ぐため）
var (
	GlobalGenericsVsInterfaceInt       int
	GlobalGenericsVsInterfaceInterface interface{}
)

func BenchmarkGenericsVsInterface(b *testing.B) {
	b.Run("Generics", func(b *testing.B) {
		var result int
		for i := 0; i < b.N; i++ {
			result = genericMax(10, 20)
		}
		GlobalGenericsVsInterfaceInt = result
	})

	b.Run("Interface", func(b *testing.B) {
		var result interface{}
		for i := 0; i < b.N; i++ {
			result = interfaceMax(10, 20)
		}
		GlobalGenericsVsInterfaceInterface = result
	})
}

// ============================================================================
// 正規表現
// ============================================================================

var precompiledRegex = regexp.MustCompile(`test\d+`)

// グローバル変数（コンパイラ最適化を防ぐため）
var GlobalRegexBool bool

func BenchmarkRegex(b *testing.B) {
	input := "This is test123 string"

	b.Run("CompileEachTime", func(b *testing.B) {
		var matched bool
		var err error
		for i := 0; i < b.N; i++ {
			matched, err = regexp.MatchString(`test\d+`, input)
			if err != nil {
				b.Fatal(err)
			}
		}
		GlobalRegexBool = matched
	})

	b.Run("Precompiled", func(b *testing.B) {
		var result bool
		for i := 0; i < b.N; i++ {
			result = precompiledRegex.MatchString(input)
		}
		GlobalRegexBool = result
	})

	b.Run("StringContains", func(b *testing.B) {
		var result bool
		for i := 0; i < b.N; i++ {
			result = strings.Contains(input, "test")
		}
		GlobalRegexBool = result
	})

	b.Run("StringHasPrefix", func(b *testing.B) {
		var result bool
		for i := 0; i < b.N; i++ {
			result = strings.HasPrefix(input, "This")
		}
		GlobalRegexBool = result
	})
}

// ============================================================================
// エラーラッピング
// ============================================================================

// グローバル変数（コンパイラ最適化を防ぐため）
var GlobalErrorWrappingErr error

func BenchmarkErrorWrapping(b *testing.B) {
	baseErr := errors.New("base error")

	b.Run("Errorf", func(b *testing.B) {
		var result error
		for i := 0; i < b.N; i++ {
			result = fmt.Errorf("wrapped: %w", baseErr)
		}
		GlobalErrorWrappingErr = result
	})

	b.Run("Join", func(b *testing.B) {
		err2 := errors.New("second error")
		var result error
		for i := 0; i < b.N; i++ {
			result = errors.Join(baseErr, err2)
		}
		GlobalErrorWrappingErr = result
	})
}

// ============================================================================
// ソート
// ============================================================================

func BenchmarkSort(b *testing.B) {
	data := make([]int, 1000)
	for i := range data {
		data[i] = 1000 - i
	}

	b.Run("Ints", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			temp := make([]int, len(data))
			copy(temp, data)
			sort.Ints(temp)
		}
	})

	b.Run("Slice", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			temp := make([]int, len(data))
			copy(temp, data)
			sort.Slice(temp, func(i, j int) bool {
				return temp[i] < temp[j]
			})
		}
	})

	b.Run("CustomSort", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			temp := make([]int, len(data))
			copy(temp, data)
			// 簡易的なバブルソート
			for j := 0; j < len(temp); j++ {
				for k := j + 1; k < len(temp); k++ {
					if temp[j] > temp[k] {
						temp[j], temp[k] = temp[k], temp[j]
					}
				}
			}
		}
	})
}

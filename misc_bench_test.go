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

// グローバル変数（コンパイラ最適化を防ぐため）
var (
	globalMiscInt       int
	globalMiscString    string
	globalMiscBool      bool
	globalMiscInterface interface{}
	globalMiscType      reflect.Type
	globalMiscValue     reflect.Value
	globalMiscValues    []reflect.Value
	globalMiscErr       error
)

// ============================================================================
// 型アサーション
// ============================================================================

func BenchmarkTypeAssertion(b *testing.B) {
	var iface interface{} = 42

	b.Run("Success", func(b *testing.B) {
		var result int
		for i := 0; i < b.N; i++ {
			result = iface.(int)
		}
		globalMiscInt = result
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
		globalMiscInt = result
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
		globalMiscString = result
	})
}

// ============================================================================
// 型switch
// ============================================================================

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
		globalMiscInt = result
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
		globalMiscInt = result
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
		globalMiscInt = result
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

func BenchmarkPanicRecover(b *testing.B) {
	var result bool
	for i := 0; i < b.N; i++ {
		result = recoverFunction()
	}
	globalMiscBool = result
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

func BenchmarkReflection(b *testing.B) {
	s := &ReflectStruct{Field1: 42, Field2: "test"}

	b.Run("TypeOf", func(b *testing.B) {
		var result reflect.Type
		for i := 0; i < b.N; i++ {
			result = reflect.TypeOf(s)
		}
		globalMiscType = result
	})

	b.Run("ValueOf", func(b *testing.B) {
		var result reflect.Value
		for i := 0; i < b.N; i++ {
			result = reflect.ValueOf(s)
		}
		globalMiscValue = result
	})

	b.Run("FieldAccess/Direct", func(b *testing.B) {
		var result int
		for i := 0; i < b.N; i++ {
			result = s.Field1
		}
		globalMiscInt = result
	})

	b.Run("FieldAccess/Reflection", func(b *testing.B) {
		v := reflect.ValueOf(s).Elem()
		var result int
		for i := 0; i < b.N; i++ {
			result = int(v.Field(0).Int())
		}
		globalMiscInt = result
	})

	b.Run("MethodCall/Direct", func(b *testing.B) {
		var result int
		for i := 0; i < b.N; i++ {
			result = s.Method()
		}
		globalMiscInt = result
	})

	b.Run("MethodCall/Reflection", func(b *testing.B) {
		v := reflect.ValueOf(s)
		method := v.MethodByName("Method")
		var result []reflect.Value
		for i := 0; i < b.N; i++ {
			result = method.Call(nil)
		}
		globalMiscValues = result
	})
}

// ============================================================================
// エラーハンドリング
// ============================================================================

func returnError() error {
	return errors.New("error")
}

func BenchmarkErrorHandling(b *testing.B) {
	b.Run("ErrorReturn", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			err := returnError()
			if err != nil {
				globalMiscErr = err
			}
		}
	})

	b.Run("PanicRecover", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			func() {
				defer func() {
					if r := recover(); r != nil {
						globalMiscInterface = r
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

func BenchmarkInterfaceOperations(b *testing.B) {
	b.Run("AssignInt", func(b *testing.B) {
		var iface interface{}
		for i := 0; i < b.N; i++ {
			iface = 42
		}
		globalMiscInterface = iface
	})

	b.Run("AssignStruct", func(b *testing.B) {
		s := ReflectStruct{Field1: 42, Field2: "test"}
		var iface interface{}
		for i := 0; i < b.N; i++ {
			iface = s
		}
		globalMiscInterface = iface
	})

	b.Run("AssignPointer", func(b *testing.B) {
		s := &ReflectStruct{Field1: 42, Field2: "test"}
		var iface interface{}
		for i := 0; i < b.N; i++ {
			iface = s
		}
		globalMiscInterface = iface
	})

	b.Run("ExtractInt", func(b *testing.B) {
		var iface interface{} = 42
		var result int
		for i := 0; i < b.N; i++ {
			result = iface.(int)
		}
		globalMiscInt = result
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

func BenchmarkGenericsVsInterface(b *testing.B) {
	b.Run("Generics", func(b *testing.B) {
		var result int
		for i := 0; i < b.N; i++ {
			result = genericMax(10, 20)
		}
		globalMiscInt = result
	})

	b.Run("Interface", func(b *testing.B) {
		var result interface{}
		for i := 0; i < b.N; i++ {
			result = interfaceMax(10, 20)
		}
		globalMiscInterface = result
	})
}

// ============================================================================
// 正規表現
// ============================================================================

var precompiledRegex = regexp.MustCompile(`test\d+`)

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
		globalMiscBool = matched
	})

	b.Run("Precompiled", func(b *testing.B) {
		var result bool
		for i := 0; i < b.N; i++ {
			result = precompiledRegex.MatchString(input)
		}
		globalMiscBool = result
	})

	b.Run("StringContains", func(b *testing.B) {
		var result bool
		for i := 0; i < b.N; i++ {
			result = strings.Contains(input, "test")
		}
		globalMiscBool = result
	})

	b.Run("StringHasPrefix", func(b *testing.B) {
		var result bool
		for i := 0; i < b.N; i++ {
			result = strings.HasPrefix(input, "This")
		}
		globalMiscBool = result
	})
}

// ============================================================================
// エラーラッピング
// ============================================================================

func BenchmarkErrorWrapping(b *testing.B) {
	baseErr := errors.New("base error")

	b.Run("Errorf", func(b *testing.B) {
		var result error
		for i := 0; i < b.N; i++ {
			result = fmt.Errorf("wrapped: %w", baseErr)
		}
		globalMiscErr = result
	})

	b.Run("Join", func(b *testing.B) {
		err2 := errors.New("second error")
		var result error
		for i := 0; i < b.N; i++ {
			result = errors.Join(baseErr, err2)
		}
		globalMiscErr = result
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

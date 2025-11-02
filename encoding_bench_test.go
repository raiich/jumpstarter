package jumpstarter

import (
	"bytes"
	"encoding/base64"
	"encoding/gob"
	"encoding/hex"
	"encoding/json"
	"testing"
)

// ============================================================================
// JSON エンコーディング
// ============================================================================

type SmallJSON struct {
	Name  string
	Age   int
	Email string
}

type MediumJSON struct {
	Field1  string
	Field2  int
	Field3  float64
	Field4  bool
	Field5  string
	Field6  int
	Field7  float64
	Field8  bool
	Field9  string
	Field10 int
}

type LargeJSON struct {
	Field1, Field2, Field3, Field4, Field5     string
	Field6, Field7, Field8, Field9, Field10    int
	Field11, Field12, Field13, Field14, Field15 float64
	Field16, Field17, Field18, Field19, Field20 bool
	Field21, Field22, Field23, Field24, Field25 string
	Field26, Field27, Field28, Field29, Field30 int
	Field31, Field32, Field33, Field34, Field35 float64
	Field36, Field37, Field38, Field39, Field40 bool
	Field41, Field42, Field43, Field44, Field45 string
	Field46, Field47, Field48, Field49, Field50 int
}

// グローバル変数（コンパイラ最適化を防ぐため）
var GlobalJSONBytes []byte

func BenchmarkJSON(b *testing.B) {
	small := SmallJSON{Name: "Alice", Age: 30, Email: "alice@example.com"}
	medium := MediumJSON{
		Field1: "test", Field2: 42, Field3: 3.14, Field4: true, Field5: "hello",
		Field6: 100, Field7: 2.71, Field8: false, Field9: "world", Field10: 200,
	}
	large := LargeJSON{
		Field1: "a", Field2: "b", Field3: "c", Field4: "d", Field5: "e",
	}

	b.Run("Marshal/Small", func(b *testing.B) {
		var result []byte
		var err error
		for i := 0; i < b.N; i++ {
			result, err = json.Marshal(small)
			if err != nil {
				b.Fatal(err)
			}
		}
		GlobalJSONBytes = result
	})

	b.Run("Marshal/Medium", func(b *testing.B) {
		var result []byte
		var err error
		for i := 0; i < b.N; i++ {
			result, err = json.Marshal(medium)
			if err != nil {
				b.Fatal(err)
			}
		}
		GlobalJSONBytes = result
	})

	b.Run("Marshal/Large", func(b *testing.B) {
		var result []byte
		var err error
		for i := 0; i < b.N; i++ {
			result, err = json.Marshal(large)
			if err != nil {
				b.Fatal(err)
			}
		}
		GlobalJSONBytes = result
	})

	smallJSON, err := json.Marshal(small)
	if err != nil {
		b.Fatal(err)
	}
	b.Run("Unmarshal/Small", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			var result SmallJSON
			err := json.Unmarshal(smallJSON, &result)
			if err != nil {
				b.Fatal(err)
			}
		}
	})

	b.Run("Encoder/Small", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			var buf bytes.Buffer
			enc := json.NewEncoder(&buf)
			err := enc.Encode(small)
			if err != nil {
				b.Fatal(err)
			}
		}
	})

	b.Run("Decoder/Small", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			buf := bytes.NewBuffer(smallJSON)
			dec := json.NewDecoder(buf)
			var result SmallJSON
			err := dec.Decode(&result)
			if err != nil {
				b.Fatal(err)
			}
		}
	})
}

// 構造体タグの影響
type NoTags struct {
	Name  string
	Age   int
	Email string
}

type WithTags struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

type WithOmitEmpty struct {
	Name  string `json:"name,omitempty"`
	Age   int    `json:"age,omitempty"`
	Email string `json:"email,omitempty"`
}

// グローバル変数（コンパイラ最適化を防ぐため）
var GlobalJSONTagsBytes []byte

func BenchmarkJSONTags(b *testing.B) {
	noTags := NoTags{Name: "Alice", Age: 30, Email: "alice@example.com"}
	withTags := WithTags{Name: "Alice", Age: 30, Email: "alice@example.com"}
	withOmitEmpty := WithOmitEmpty{Name: "Alice", Age: 30, Email: "alice@example.com"}

	b.Run("NoTags", func(b *testing.B) {
		var result []byte
		var err error
		for i := 0; i < b.N; i++ {
			result, err = json.Marshal(noTags)
			if err != nil {
				b.Fatal(err)
			}
		}
		GlobalJSONTagsBytes = result
	})

	b.Run("WithTags", func(b *testing.B) {
		var result []byte
		var err error
		for i := 0; i < b.N; i++ {
			result, err = json.Marshal(withTags)
			if err != nil {
				b.Fatal(err)
			}
		}
		GlobalJSONTagsBytes = result
	})

	b.Run("WithOmitEmpty", func(b *testing.B) {
		var result []byte
		var err error
		for i := 0; i < b.N; i++ {
			result, err = json.Marshal(withOmitEmpty)
			if err != nil {
				b.Fatal(err)
			}
		}
		GlobalJSONTagsBytes = result
	})
}

// ============================================================================
// その他のエンコーディング
// ============================================================================

// グローバル変数（コンパイラ最適化を防ぐため）
var (
	GlobalOtherEncodingsBytes []byte
	GlobalOtherEncodingsStr   string
)

func BenchmarkOtherEncodings(b *testing.B) {
	data := SmallJSON{Name: "Alice", Age: 30, Email: "alice@example.com"}

	b.Run("Gob/Encode", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			var buf bytes.Buffer
			enc := gob.NewEncoder(&buf)
			err := enc.Encode(data)
			if err != nil {
				b.Fatal(err)
			}
		}
	})

	var gobBuf bytes.Buffer
	enc := gob.NewEncoder(&gobBuf)
	err := enc.Encode(data)
	if err != nil {
		b.Fatal(err)
	}
	gobData := gobBuf.Bytes()

	b.Run("Gob/Decode", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			buf := bytes.NewBuffer(gobData)
			dec := gob.NewDecoder(buf)
			var result SmallJSON
			err := dec.Decode(&result)
			if err != nil {
				b.Fatal(err)
			}
		}
	})

	input := []byte("Hello, World! This is a test string for encoding.")

	b.Run("Base64/Encode", func(b *testing.B) {
		var result string
		for i := 0; i < b.N; i++ {
			result = base64.StdEncoding.EncodeToString(input)
		}
		GlobalOtherEncodingsStr = result
	})

	encoded := base64.StdEncoding.EncodeToString(input)

	b.Run("Base64/Decode", func(b *testing.B) {
		var result []byte
		var err error
		for i := 0; i < b.N; i++ {
			result, err = base64.StdEncoding.DecodeString(encoded)
			if err != nil {
				b.Fatal(err)
			}
		}
		GlobalOtherEncodingsBytes = result
	})

	b.Run("Hex/Encode", func(b *testing.B) {
		var result string
		for i := 0; i < b.N; i++ {
			result = hex.EncodeToString(input)
		}
		GlobalOtherEncodingsStr = result
	})

	hexEncoded := hex.EncodeToString(input)

	b.Run("Hex/Decode", func(b *testing.B) {
		var result []byte
		var err error
		for i := 0; i < b.N; i++ {
			result, err = hex.DecodeString(hexEncoded)
			if err != nil {
				b.Fatal(err)
			}
		}
		GlobalOtherEncodingsBytes = result
	})
}

package jumpstarter

import (
	"bytes"
	"encoding/base64"
	"encoding/gob"
	"encoding/hex"
	"encoding/json"
	"testing"
)

// グローバル変数（コンパイラ最適化を防ぐため）
var (
	globalEncBytes []byte
	globalEncStr   string
	globalEncErr   error
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
		for i := 0; i < b.N; i++ {
			globalEncBytes, globalEncErr = json.Marshal(small)
		}
	})

	b.Run("Marshal/Medium", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			globalEncBytes, globalEncErr = json.Marshal(medium)
		}
	})

	b.Run("Marshal/Large", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			globalEncBytes, globalEncErr = json.Marshal(large)
		}
	})

	smallJSON, _ := json.Marshal(small)
	b.Run("Unmarshal/Small", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			var result SmallJSON
			globalEncErr = json.Unmarshal(smallJSON, &result)
		}
	})

	b.Run("Encoder/Small", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			var buf bytes.Buffer
			enc := json.NewEncoder(&buf)
			globalEncErr = enc.Encode(small)
		}
	})

	b.Run("Decoder/Small", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			buf := bytes.NewBuffer(smallJSON)
			dec := json.NewDecoder(buf)
			var result SmallJSON
			globalEncErr = dec.Decode(&result)
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

func BenchmarkJSONTags(b *testing.B) {
	noTags := NoTags{Name: "Alice", Age: 30, Email: "alice@example.com"}
	withTags := WithTags{Name: "Alice", Age: 30, Email: "alice@example.com"}
	withOmitEmpty := WithOmitEmpty{Name: "Alice", Age: 30, Email: "alice@example.com"}

	b.Run("NoTags", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			globalEncBytes, globalEncErr = json.Marshal(noTags)
		}
	})

	b.Run("WithTags", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			globalEncBytes, globalEncErr = json.Marshal(withTags)
		}
	})

	b.Run("WithOmitEmpty", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			globalEncBytes, globalEncErr = json.Marshal(withOmitEmpty)
		}
	})
}

// ============================================================================
// その他のエンコーディング
// ============================================================================

func BenchmarkOtherEncodings(b *testing.B) {
	data := SmallJSON{Name: "Alice", Age: 30, Email: "alice@example.com"}

	b.Run("Gob/Encode", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			var buf bytes.Buffer
			enc := gob.NewEncoder(&buf)
			globalEncErr = enc.Encode(data)
		}
	})

	var gobBuf bytes.Buffer
	enc := gob.NewEncoder(&gobBuf)
	globalEncErr = enc.Encode(data)
	gobData := gobBuf.Bytes()

	b.Run("Gob/Decode", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			buf := bytes.NewBuffer(gobData)
			dec := gob.NewDecoder(buf)
			var result SmallJSON
			globalEncErr = dec.Decode(&result)
		}
	})

	input := []byte("Hello, World! This is a test string for encoding.")

	b.Run("Base64/Encode", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			globalEncStr = base64.StdEncoding.EncodeToString(input)
		}
	})

	encoded := base64.StdEncoding.EncodeToString(input)

	b.Run("Base64/Decode", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			globalEncBytes, globalEncErr = base64.StdEncoding.DecodeString(encoded)
		}
	})

	b.Run("Hex/Encode", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			globalEncStr = hex.EncodeToString(input)
		}
	})

	hexEncoded := hex.EncodeToString(input)

	b.Run("Hex/Decode", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			globalEncBytes, globalEncErr = hex.DecodeString(hexEncoded)
		}
	})
}

// go test -bench=.
// goos: darwin
// goarch: arm64
// pkg: github.com/triumphpc/go-study/tricks/serialize/gob
// cpu: Apple M1 Pro
// BenchmarkJSONMarshal-10        	 1408658	       883.9 ns/op
// BenchmarkJSONUnmarshal-10      	  352075	      3383 ns/op
// BenchmarkGobEncode-10          	  269751	      4212 ns/op
// BenchmarkGobDecode-10          	   81046	     15419 ns/op
// BenchmarkGobReuseEncoder-10    	 2006162	       600.2 ns/op
// BenchmarkCompareSize-10        	1000000000	         0.0000291 ns/op
// --- BENCH: BenchmarkCompareSize-10
//
//	   main_test.go:143: JSON size: 221 bytes
//	   main_test.go:144: GOB size: 349 bytes
//	   main_test.go:143: JSON size: 221 bytes
//	   main_test.go:144: GOB size: 349 bytes
//	   main_test.go:143: JSON size: 221 bytes
//	   main_test.go:144: GOB size: 349 bytes
//	   main_test.go:143: JSON size: 221 bytes
//	   main_test.go:144: GOB size: 349 bytes
//	   main_test.go:143: JSON size: 221 bytes
//	   main_test.go:144: GOB size: 349 bytes
//		... [output truncated]
//
// PASS
package main

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"testing"
)

// Тестовая структура с разными типами данных
type TestData struct {
	ID       int64             `json:"id"`
	Name     string            `json:"name"`
	Price    float64           `json:"price"`
	IsActive bool              `json:"is_active"`
	Tags     []string          `json:"tags"`
	Metadata map[string]string `json:"metadata"`
	Numbers  []int             `json:"numbers"`
	Nested   *NestedData       `json:"nested"`
}

type NestedData struct {
	Code  string  `json:"code"`
	Value float64 `json:"value"`
}

// Создаем тестовые данные
func createTestData() *TestData {
	return &TestData{
		ID:       12345,
		Name:     "Test Product",
		Price:    99.99,
		IsActive: true,
		Tags:     []string{"tag1", "tag2", "tag3"},
		Metadata: map[string]string{
			"color": "red",
			"size":  "large",
			"type":  "electronics",
		},
		Numbers: []int{1, 2, 3, 4, 5},
		Nested: &NestedData{
			Code:  "ABC123",
			Value: 42.42,
		},
	}
}

// Бенчмарк для JSON маршалинга
func BenchmarkJSONMarshal(b *testing.B) {
	data := createTestData()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := json.Marshal(data)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// Бенчмарк для JSON анмаршалинга
func BenchmarkJSONUnmarshal(b *testing.B) {
	data := createTestData()
	jsonData, _ := json.Marshal(data)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var result TestData
		err := json.Unmarshal(jsonData, &result)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// Бенчмарк для GOB кодирования
func BenchmarkGobEncode(b *testing.B) {
	data := createTestData()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var buf bytes.Buffer
		enc := gob.NewEncoder(&buf)
		err := enc.Encode(data)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// Бенчмарк для GOB декодирования
func BenchmarkGobDecode(b *testing.B) {
	data := createTestData()
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(data)
	if err != nil {
		b.Fatal(err)
	}
	gobData := buf.Bytes()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var result TestData
		dec := gob.NewDecoder(bytes.NewBuffer(gobData))
		err := dec.Decode(&result)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// Бенчмарк для повторного использования энкодера/декодера GOB
func BenchmarkGobReuseEncoder(b *testing.B) {
	data := createTestData()
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		buf.Reset()
		err := enc.Encode(data)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// Сравнение размеров сериализованных данных
func BenchmarkCompareSize(b *testing.B) {
	data := createTestData()

	// JSON размер
	jsonData, _ := json.Marshal(data)
	jsonSize := len(jsonData)

	// GOB размер
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	enc.Encode(data)
	gobSize := buf.Len()

	b.Logf("JSON size: %d bytes", jsonSize)
	b.Logf("GOB size: %d bytes", gobSize)
}

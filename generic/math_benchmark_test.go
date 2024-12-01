package generic

import (
	"math/rand"
	"testing"
)

const benchmarkSize = 10000 // Number of elements for benchmarks

func BenchmarkMinInt(b *testing.B) {
	a, bVal := 100, 200
	for i := 0; i < b.N; i++ {
		Min(a, bVal)
	}
}

func BenchmarkMaxInt(b *testing.B) {
	a, bVal := 100, 200
	for i := 0; i < b.N; i++ {
		Max(a, bVal)
	}
}

func BenchmarkMinString(b *testing.B) {
	a, bVal := "apple", "banana"
	for i := 0; i < b.N; i++ {
		Min(a, bVal)
	}
}

func BenchmarkMaxString(b *testing.B) {
	a, bVal := "apple", "banana"
	for i := 0; i < b.N; i++ {
		Max(a, bVal)
	}
}

func BenchmarkContainsSliceInt(b *testing.B) {
	// Create a large slice of integers
	slice := make([]int, benchmarkSize)
	for i := 0; i < benchmarkSize; i++ {
		slice[i] = i
	}
	searchValue := rand.Intn(benchmarkSize)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ContainsSlice(slice, searchValue)
	}
}

func BenchmarkContainsSliceString(b *testing.B) {
	// Create a large slice of strings
	slice := make([]string, benchmarkSize)
	for i := 0; i < benchmarkSize; i++ {
		slice[i] = "value" + string(rune(i))
	}
	searchValue := "value" + string(rune(rand.Intn(benchmarkSize)))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ContainsSlice(slice, searchValue)
	}
}

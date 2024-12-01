package set

import (
	"math/rand"
	"testing"
)

const benchmarkSize = 10000 // Number of elements for benchmarks

func BenchmarkSetAdd(b *testing.B) {
	s := NewSet[int]()
	for i := 0; i < b.N; i++ {
		s.Add(i)
	}
}

func BenchmarkSetAddMultiple(b *testing.B) {
	s := NewSet[int]()
	elements := make([]int, benchmarkSize)
	for i := 0; i < benchmarkSize; i++ {
		elements[i] = i
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.AddMultiple(elements...)
	}
}

func BenchmarkSetAddSlice(b *testing.B) {
	s := NewSet[int]()
	elements := make([]int, benchmarkSize)
	for i := 0; i < benchmarkSize; i++ {
		elements[i] = i
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.AddSlice(elements)
	}
}

func BenchmarkSetContains(b *testing.B) {
	s := NewSet[int]()
	for i := 0; i < benchmarkSize; i++ {
		s.Add(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Contains(rand.Intn(benchmarkSize))
	}
}

func BenchmarkSetRemove(b *testing.B) {
	s := NewSet[int]()
	for i := 0; i < benchmarkSize; i++ {
		s.Add(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Remove(rand.Intn(benchmarkSize))
	}
}

func BenchmarkSetUnion(b *testing.B) {
	s1 := NewSet[int]()
	s2 := NewSet[int]()
	for i := 0; i < benchmarkSize; i++ {
		s1.Add(i)
		s2.Add(i + benchmarkSize)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = s1.Union(s2)
	}
}

func BenchmarkSetIntersection(b *testing.B) {
	s1 := NewSet[int]()
	s2 := NewSet[int]()
	for i := 0; i < benchmarkSize; i++ {
		s1.Add(i)
		if i%2 == 0 {
			s2.Add(i)
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = s1.Intersection(s2)
	}
}

func BenchmarkSetDifference(b *testing.B) {
	s1 := NewSet[int]()
	s2 := NewSet[int]()
	for i := 0; i < benchmarkSize; i++ {
		s1.Add(i)
		if i%2 == 0 {
			s2.Add(i)
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = s1.Difference(s2)
	}
}

func BenchmarkSetToSlice(b *testing.B) {
	s := NewSet[int]()
	for i := 0; i < benchmarkSize; i++ {
		s.Add(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = s.ToSlice()
	}
}

func BenchmarkSetForEach(b *testing.B) {
	s := NewSet[int]()
	for i := 0; i < benchmarkSize; i++ {
		s.Add(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.ForEach(func(element int) {
			_ = element // Simulate a read operation
		})
	}
}

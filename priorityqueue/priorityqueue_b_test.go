package priorityqueue

import (
	"math/rand"
	"testing"
)

func BenchmarkPriorityQueueEnqueue(b *testing.B) {
	// Comparator for a min-heap
	minHeapComparator := func(a, b int) bool {
		return a < b
	}

	h := NewPriorityQueue(minHeapComparator)

	// Generate a large number of random integers for testing
	data := make([]int, b.N)
	for i := 0; i < b.N; i++ {
		data[i] = rand.Intn(b.N)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		h.Enqueue(data[i])
	}
}

func BenchmarkPriorityQueueDequeue(b *testing.B) {
	// Comparator for a min-heap
	minHeapComparator := func(a, b int) bool {
		return a < b
	}

	h := NewPriorityQueue(minHeapComparator)

	// Pre-fill the queue with random integers
	for i := 0; i < b.N; i++ {
		h.Enqueue(rand.Intn(b.N))
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		h.Dequeue()
	}
}

func BenchmarkPriorityQueueMixedOperations(b *testing.B) {
	// Comparator for a min-heap
	minHeapComparator := func(a, b int) bool {
		return a < b
	}

	h := NewPriorityQueue(minHeapComparator)

	// Generate random data
	data := make([]int, b.N)
	for i := 0; i < b.N; i++ {
		data[i] = rand.Intn(b.N)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		h.Enqueue(data[i])
		if i%2 == 0 { // Dequeue every second operation
			h.Dequeue()
		}
	}
}

func BenchmarkPriorityQueuePeek(b *testing.B) {
	// Comparator for a min-heap
	minHeapComparator := func(a, b int) bool {
		return a < b
	}

	h := NewPriorityQueue(minHeapComparator)

	// Pre-fill the queue with random integers
	for i := 0; i < 1000; i++ {
		h.Enqueue(rand.Intn(1000))
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		h.Peek()
	}
}

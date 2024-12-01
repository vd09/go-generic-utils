package queue

import "testing"

func BenchmarkQueueEnqueue(b *testing.B) {
	queue := NewQueue[int]()
	for i := 0; i < b.N; i++ {
		queue.Enqueue(i)
	}
}

func BenchmarkQueueDequeue(b *testing.B) {
	queue := NewQueue[int]()
	// Pre-fill the queue with `b.N` elements
	for i := 0; i < b.N; i++ {
		queue.Enqueue(i)
	}

	b.ResetTimer() // Reset the timer to exclude setup time

	for i := 0; i < b.N; i++ {
		queue.Dequeue()
	}
}

func BenchmarkQueuePeek(b *testing.B) {
	queue := NewQueue[int]()
	// Push one element to ensure Peek can work
	queue.Enqueue(1)

	b.ResetTimer() // Reset the timer to exclude setup time

	for i := 0; i < b.N; i++ {
		queue.Peek()
	}
}

func BenchmarkQueueEnqueueAndDequeue(b *testing.B) {
	queue := NewQueue[int]()

	b.ResetTimer() // Reset the timer to exclude setup time

	for i := 0; i < b.N; i++ {
		queue.Enqueue(i)
		queue.Dequeue()
	}
}

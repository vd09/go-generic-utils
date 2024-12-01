package stack

import (
	"testing"
)

func BenchmarkStackPush(b *testing.B) {
	stack := NewStack[int]()
	for i := 0; i < b.N; i++ {
		stack.Push(i)
	}
}

func BenchmarkStackPop(b *testing.B) {
	stack := NewStack[int]()
	// Pre-fill the stack with `b.N` elements
	for i := 0; i < b.N; i++ {
		stack.Push(i)
	}

	b.ResetTimer() // Reset the timer to exclude setup time

	for i := 0; i < b.N; i++ {
		stack.Pop()
	}
}

func BenchmarkStackPeek(b *testing.B) {
	stack := NewStack[int]()
	// Push one element to ensure Peek can work
	stack.Push(1)

	b.ResetTimer() // Reset the timer to exclude setup time

	for i := 0; i < b.N; i++ {
		stack.Peek()
	}
}

func BenchmarkStackPushAndPop(b *testing.B) {
	stack := NewStack[int]()

	b.ResetTimer() // Reset the timer to exclude setup time

	for i := 0; i < b.N; i++ {
		stack.Push(i)
		stack.Pop()
	}
}

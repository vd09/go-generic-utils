package priorityqueue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPriorityQueueBasicOperations(t *testing.T) {
	// Comparator for a min-heap
	minHeapComparator := func(a, b int) bool {
		return a < b
	}

	h := NewPriorityQueue(minHeapComparator)

	// Test Enqueue
	h.Enqueue(3)
	h.Enqueue(1)
	h.Enqueue(2)

	assert.Equal(t, 3, h.Len(), "PriorityQueue length should be 3")

	// Test Peek
	peekValue, ok := h.Peek()
	assert.True(t, ok, "Peek should return true for a non-empty queue")
	assert.Equal(t, 1, peekValue, "Peek value should be the smallest element")

	// Test Dequeue
	poppedValue, ok := h.Dequeue()
	assert.True(t, ok, "Dequeue should return true for a non-empty queue")
	assert.Equal(t, 1, poppedValue, "Dequeue should return the smallest element")
	assert.Equal(t, 2, h.Len(), "PriorityQueue length should be 2 after dequeuing an element")

	// Test Peek again
	peekValue, _ = h.Peek()
	assert.Equal(t, 2, peekValue, "Next smallest element should be 2")
}

func TestPriorityQueueDescendingOrder(t *testing.T) {
	// Comparator for a max-heap
	maxHeapComparator := func(a, b int) bool {
		return a > b
	}

	h := NewPriorityQueue(maxHeapComparator)

	// Enqueue elements in random order
	h.Enqueue(10)
	h.Enqueue(20)
	h.Enqueue(15)

	// Verify Peek returns the largest element
	peekValue, ok := h.Peek()
	assert.True(t, ok, "Peek should return true for a non-empty queue")
	assert.Equal(t, 20, peekValue, "Peek value should be the largest element")

	// Dequeue and check order
	expectedOrder := []int{20, 15, 10}
	for _, expected := range expectedOrder {
		value, ok := h.Dequeue()
		assert.True(t, ok, "Dequeue should return true for a non-empty queue")
		assert.Equal(t, expected, value, "Dequeued value should match expected")
	}

	// Verify the queue is empty
	assert.Equal(t, 0, h.Len(), "PriorityQueue should be empty after all elements are dequeued")
}

func TestPriorityQueueEdgeCases(t *testing.T) {
	// Comparator for a min-heap
	minHeapComparator := func(a, b int) bool {
		return a < b
	}

	h := NewPriorityQueue(minHeapComparator)

	// Test Dequeue on an empty queue
	value, ok := h.Dequeue()
	assert.False(t, ok, "Dequeue should return false for an empty queue")
	assert.Equal(t, 0, value, "Dequeue should return zero value for an empty queue")

	// Test Peek on an empty queue
	peekValue, ok := h.Peek()
	assert.False(t, ok, "Peek should return false for an empty queue")
	assert.Equal(t, 0, peekValue, "Peek should return zero value for an empty queue")

	// Test Enqueue and Dequeue with a single element
	h.Enqueue(42)
	assert.Equal(t, 1, h.Len(), "PriorityQueue length should be 1 after one Enqueue")

	peekValue, ok = h.Peek()
	assert.True(t, ok, "Peek should return true for a non-empty queue")
	assert.Equal(t, 42, peekValue, "Peek should return the single element in the queue")

	value, ok = h.Dequeue()
	assert.True(t, ok, "Dequeue should return true for a non-empty queue")
	assert.Equal(t, 42, value, "Dequeue should return the single element in the queue")

	assert.Equal(t, 0, h.Len(), "PriorityQueue should be empty after dequeuing the single element")
}

func TestPriorityQueueCustomTypes(t *testing.T) {
	// Comparator for a max-heap of strings
	stringLengthComparator := func(a, b string) bool {
		return len(a) > len(b)
	}

	h := NewPriorityQueue(stringLengthComparator)

	// Enqueue strings
	h.Enqueue("short")
	h.Enqueue("medium")
	h.Enqueue("verylongstring")

	// Check Peek and Dequeue order
	expectedOrder := []string{"verylongstring", "medium", "short"}
	for _, expected := range expectedOrder {
		value, ok := h.Dequeue()
		assert.True(t, ok, "Dequeue should return true for a non-empty queue")
		assert.Equal(t, expected, value, "Dequeued value should match expected")
	}

	// Verify the queue is empty
	assert.Equal(t, 0, h.Len(), "PriorityQueue should be empty after all elements are dequeued")
}

func TestPriorityQueueDuplicateElements(t *testing.T) {
	// Comparator for a min-heap
	minHeapComparator := func(a, b int) bool {
		return a < b
	}

	h := NewPriorityQueue(minHeapComparator)

	// Enqueue duplicate elements
	h.Enqueue(5)
	h.Enqueue(5)
	h.Enqueue(5)

	// Verify Peek and Dequeue order
	for i := 0; i < 3; i++ {
		value, ok := h.Dequeue()
		assert.True(t, ok, "Dequeue should return true for a non-empty queue")
		assert.Equal(t, 5, value, "Dequeued value should be 5")
	}

	// Verify the queue is empty
	assert.Equal(t, 0, h.Len(), "PriorityQueue should be empty after dequeuing all elements")
}

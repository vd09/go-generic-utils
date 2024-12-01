package priorityqueue

import (
	"container/heap" // Use the standard import without alias
)

// PriorityQueue implements a priority queue using a priorityqueue
type PriorityQueue[T any] struct {
	items      []T               // The items in the priorityqueue
	comparator func(a, b T) bool // Custom comparator function
}

// NewPriorityQueue initializes a new priority queue with a custom comparator
func NewPriorityQueue[T any](comparator func(a, b T) bool) *PriorityQueue[T] {
	h := &PriorityQueue[T]{comparator: comparator}
	heap.Init(h)
	return h
}

// Len returns the number of elements in the priorityqueue
func (h *PriorityQueue[T]) Len() int {
	return len(h.items)
}

// Less uses the custom comparator function to determine the order
func (h *PriorityQueue[T]) Less(i, j int) bool {
	return h.comparator(h.items[i], h.items[j])
}

// Swap swaps the elements at indices i and j
func (h *PriorityQueue[T]) Swap(i, j int) {
	h.items[i], h.items[j] = h.items[j], h.items[i]
}

// Push adds an item to the priorityqueue
func (h *PriorityQueue[T]) Push(x any) {
	h.items = append(h.items, x.(T))
}

// Pop removes and returns the item with the highest priority
func (h *PriorityQueue[T]) Pop() any {
	old := h.items
	n := len(old)
	item := old[n-1]
	h.items = old[:n-1]
	return item
}

// Peek returns the top element of the priorityqueue without removing it
func (h *PriorityQueue[T]) Peek() (T, bool) {
	if h.Len() == 0 {
		var zero T
		return zero, false
	}
	return h.items[0], true
}

// Enqueue adds a new item to the priority queue
func (h *PriorityQueue[T]) Enqueue(value T) {
	heap.Push(h, value)
}

// Dequeue removes and returns the item with the highest priority
func (h *PriorityQueue[T]) Dequeue() (T, bool) {
	if h.Len() == 0 {
		var zero T
		return zero, false
	}
	item := heap.Pop(h).(T)
	return item, true
}

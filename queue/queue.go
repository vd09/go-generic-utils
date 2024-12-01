package queue

import (
	"container/list"
)

// Queue represents a generic queue data structure
type Queue[T any] struct {
	list *list.List
}

// NewQueue creates a new Queue
func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{list: list.New()}
}

// Enqueue adds an element to the end of the queue
func (q *Queue[T]) Enqueue(value T) {
	q.list.PushBack(value)
}

// Dequeue removes and returns the front element of the queue
// It returns the zero value of T and false if the queue is empty
func (q *Queue[T]) Dequeue() (T, bool) {
	element := q.list.Front()
	if element != nil {
		q.list.Remove(element)
		return element.Value.(T), true
	}
	var zero T
	return zero, false
}

// Peek returns the front element without removing it
// It returns the zero value of T and false if the queue is empty
func (q *Queue[T]) Peek() (T, bool) {
	element := q.list.Front()
	if element != nil {
		return element.Value.(T), true
	}
	var zero T
	return zero, false
}

// IsEmpty checks if the queue is empty
func (q *Queue[T]) IsEmpty() bool {
	return q.list.Len() == 0
}

// Size returns the number of elements in the queue
func (q *Queue[T]) Size() int {
	return q.list.Len()
}

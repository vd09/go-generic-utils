package stack

import (
	"container/list"
)

// Stack represents a generic stack data structure
type Stack[T any] struct {
	list *list.List
}

// NewStack creates a new Stack
func NewStack[T any]() *Stack[T] {
	return &Stack[T]{list: list.New()}
}

// Push adds an element to the top of the stack
func (s *Stack[T]) Push(value T) {
	s.list.PushBack(value)
}

// Pop removes and returns the top element of the stack
// It returns the zero value of T and false if the stack is empty
func (s *Stack[T]) Pop() (T, bool) {
	element := s.list.Back()
	if element != nil {
		s.list.Remove(element)
		return element.Value.(T), true
	}
	var zero T
	return zero, false
}

// Peek returns the top element without removing it
// It returns the zero value of T and false if the stack is empty
func (s *Stack[T]) Peek() (T, bool) {
	element := s.list.Back()
	if element != nil {
		return element.Value.(T), true
	}
	var zero T
	return zero, false
}

// IsEmpty checks if the stack is empty
func (s *Stack[T]) IsEmpty() bool {
	return s.list.Len() == 0
}

// Size returns the number of elements in the stack
func (s *Stack[T]) Size() int {
	return s.list.Len()
}

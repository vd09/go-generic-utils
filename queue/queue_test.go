package queue

import "testing"

func TestQueue(t *testing.T) {
	// Test with integers
	intQueue := NewQueue[int]()

	if !intQueue.IsEmpty() {
		t.Errorf("Expected queue to be empty")
	}

	// Enqueue elements
	intQueue.Enqueue(10)
	intQueue.Enqueue(20)
	intQueue.Enqueue(30)

	if intQueue.Size() != 3 {
		t.Errorf("Expected queue size to be 3, got %d", intQueue.Size())
	}

	// Peek the front element
	if front, ok := intQueue.Peek(); !ok || front != 10 {
		t.Errorf("Expected front element to be 10, got %v", front)
	}

	// Dequeue elements
	if value, ok := intQueue.Dequeue(); !ok || value != 10 {
		t.Errorf("Expected dequeued element to be 10, got %v", value)
	}

	if value, ok := intQueue.Dequeue(); !ok || value != 20 {
		t.Errorf("Expected dequeued element to be 20, got %v", value)
	}

	if value, ok := intQueue.Dequeue(); !ok || value != 30 {
		t.Errorf("Expected dequeued element to be 30, got %v", value)
	}

	// Check if queue is empty
	if !intQueue.IsEmpty() {
		t.Errorf("Expected queue to be empty")
	}

	// Dequeue from an empty queue
	if value, ok := intQueue.Dequeue(); ok {
		t.Errorf("Expected Dequeue to return false, but got value %v", value)
	}

	// Test with strings
	stringQueue := NewQueue[string]()

	// Enqueue elements
	stringQueue.Enqueue("hello")
	stringQueue.Enqueue("world")

	if stringQueue.Size() != 2 {
		t.Errorf("Expected queue size to be 2, got %d", stringQueue.Size())
	}

	// Peek the front element
	if front, ok := stringQueue.Peek(); !ok || front != "hello" {
		t.Errorf("Expected front element to be 'hello', got %v", front)
	}

	// Dequeue elements
	if value, ok := stringQueue.Dequeue(); !ok || value != "hello" {
		t.Errorf("Expected dequeued element to be 'hello', got %v", value)
	}

	if value, ok := stringQueue.Dequeue(); !ok || value != "world" {
		t.Errorf("Expected dequeued element to be 'world', got %v", value)
	}

	// Check if queue is empty
	if !stringQueue.IsEmpty() {
		t.Errorf("Expected queue to be empty")
	}
}

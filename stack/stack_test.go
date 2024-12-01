package stack

import (
	"testing"
)

func TestStack(t *testing.T) {
	// Test with integers
	intStack := NewStack[int]()

	if !intStack.IsEmpty() {
		t.Errorf("Expected stack to be empty")
	}

	// Push elements
	intStack.Push(10)
	intStack.Push(20)
	intStack.Push(30)

	if intStack.Size() != 3 {
		t.Errorf("Expected stack size to be 3, got %d", intStack.Size())
	}

	// Peek the top element
	if top, ok := intStack.Peek(); !ok || top != 30 {
		t.Errorf("Expected top element to be 30, got %v", top)
	}

	// Pop elements
	if value, ok := intStack.Pop(); !ok || value != 30 {
		t.Errorf("Expected popped element to be 30, got %v", value)
	}

	if value, ok := intStack.Pop(); !ok || value != 20 {
		t.Errorf("Expected popped element to be 20, got %v", value)
	}

	if value, ok := intStack.Pop(); !ok || value != 10 {
		t.Errorf("Expected popped element to be 10, got %v", value)
	}

	// Check if stack is empty
	if !intStack.IsEmpty() {
		t.Errorf("Expected stack to be empty")
	}

	// Pop from an empty stack
	if value, ok := intStack.Pop(); ok {
		t.Errorf("Expected Pop to return false, but got value %v", value)
	}

	// Test with strings
	stringStack := NewStack[string]()

	// Push elements
	stringStack.Push("hello")
	stringStack.Push("world")

	if stringStack.Size() != 2 {
		t.Errorf("Expected stack size to be 2, got %d", stringStack.Size())
	}

	// Peek the top element
	if top, ok := stringStack.Peek(); !ok || top != "world" {
		t.Errorf("Expected top element to be 'world', got %v", top)
	}

	// Pop elements
	if value, ok := stringStack.Pop(); !ok || value != "world" {
		t.Errorf("Expected popped element to be 'world', got %v", value)
	}

	if value, ok := stringStack.Pop(); !ok || value != "hello" {
		t.Errorf("Expected popped element to be 'hello', got %v", value)
	}

	// Check if stack is empty
	if !stringStack.IsEmpty() {
		t.Errorf("Expected stack to be empty")
	}
}

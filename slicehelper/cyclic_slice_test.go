package slicehelper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCyclicSlice_AddAndGet(t *testing.T) {
	cs := NewCyclicSlice[int](3)

	// Add elements and check GetAt
	cs.Add(1)
	cs.Add(2)
	cs.Add(3)

	assert.Equal(t, 1, cs.GetAt(0))
	assert.Equal(t, 2, cs.GetAt(1))
	assert.Equal(t, 3, cs.GetAt(2))
}

func TestCyclicSlice_AddAndWrapAround(t *testing.T) {
	cs := NewCyclicSlice[int](3)

	// Add more elements than size to check wrap around
	cs.Add(1)
	cs.Add(2)
	cs.Add(3)
	cs.Add(4)

	assert.Equal(t, 2, cs.GetAt(0)) // Oldest element should be overwritten
	assert.Equal(t, 3, cs.GetAt(1))
	assert.Equal(t, 4, cs.GetAt(2))
}

func TestCyclicSlice_SetAt(t *testing.T) {
	cs := NewCyclicSlice[int](3)

	cs.Add(1)
	cs.Add(2)
	cs.Add(3)

	cs.SetAt(1, 10) // Set element at index 1

	assert.Equal(t, 1, cs.GetAt(0))
	assert.Equal(t, 10, cs.GetAt(1))
	assert.Equal(t, 3, cs.GetAt(2))
}

func TestCyclicSlice_GetLast(t *testing.T) {
	cs := NewCyclicSlice[int](3)

	cs.Add(1)
	cs.Add(2)
	cs.Add(3)

	assert.Equal(t, 3, cs.GetLast())

	cs.Add(4) // Wrap around
	assert.Equal(t, 4, cs.GetLast())
}

func TestCyclicSlice_GetData(t *testing.T) {
	cs := NewCyclicSlice[int](3)

	assert.Equal(t, []int{}, cs.GetData())

	cs.Add(1)
	cs.Add(2)
	assert.Equal(t, []int{1, 2}, cs.GetData())

	cs.Add(3)
	assert.Equal(t, []int{1, 2, 3}, cs.GetData())

	cs.Add(4) // Wrap around
	assert.Equal(t, []int{2, 3, 4}, cs.GetData())

	cs.Add(5)
	cs.Add(6)
	cs.Add(7)
	cs.Add(8)
	assert.Equal(t, []int{6, 7, 8}, cs.GetData())

}

func TestCyclicSlice_GetPrevious(t *testing.T) {
	cs := NewCyclicSlice[int](3)

	cs.Add(1)
	cs.Add(2)
	cs.Add(3)

	assert.Equal(t, 2, cs.GetPrevious())

	cs.Add(4) // Wrap around
	assert.Equal(t, 3, cs.GetPrevious())
}

func TestCyclicSlice_IsFull(t *testing.T) {
	cs := NewCyclicSlice[int](3)

	assert.False(t, cs.IsFull())

	cs.Add(1)
	cs.Add(2)
	cs.Add(3)

	assert.True(t, cs.IsFull())

	cs.Add(4) // Add more to check it stays full
	assert.True(t, cs.IsFull())
}

func TestCyclicSlice_GetCurrentSize(t *testing.T) {
	cs := NewCyclicSlice[int](3)

	assert.Equal(t, 0, cs.GetCurrentSize())

	cs.Add(1)
	cs.Add(2)

	assert.Equal(t, 2, cs.GetCurrentSize())

	cs.Add(3)
	cs.Add(4) // Wrap around

	assert.Equal(t, 3, cs.GetCurrentSize())
}

func TestCyclicSlice_PanicOnOutOfBounds(t *testing.T) {
	cs := NewCyclicSlice[int](3)

	cs.Add(1)
	cs.Add(2)

	assert.Panics(t, func() { cs.GetAt(-1) })
	assert.Panics(t, func() { cs.GetAt(3) })

	assert.Panics(t, func() { cs.SetAt(-1, 100) })
	assert.Panics(t, func() { cs.SetAt(3, 100) })
}

func TestCyclicSlice_PanicOnEmptyGetLast(t *testing.T) {
	cs := NewCyclicSlice[int](3)

	assert.Panics(t, func() { cs.GetLast() })
}

func TestCyclicSlice_PanicOnEmptyGetPrevious(t *testing.T) {
	cs := NewCyclicSlice[int](3)

	assert.Panics(t, func() { cs.GetPrevious() })
}

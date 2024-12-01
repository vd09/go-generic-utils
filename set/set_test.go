package set

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewSet(t *testing.T) {
	s := NewSet[int]()
	assert.NotNil(t, s, "NewSet should return a non-nil set")
	assert.True(t, s.IsEmpty(), "New set should be empty")
}

func TestAddAndContains(t *testing.T) {
	s := NewSet[int]()
	s.Add(1)
	s.Add(2)

	assert.True(t, s.Contains(1), "Set should contain element 1")
	assert.True(t, s.Contains(2), "Set should contain element 2")
	assert.False(t, s.Contains(3), "Set should not contain element 3")
	assert.Equal(t, 2, s.Size(), "Set size should be 2")
}

func TestAddMultiple(t *testing.T) {
	s := NewSet[int]()
	s.AddMultiple(1, 2, 3)

	assert.True(t, s.Contains(1), "Set should contain element 1")
	assert.True(t, s.Contains(2), "Set should contain element 2")
	assert.True(t, s.Contains(3), "Set should contain element 3")
	assert.Equal(t, 3, s.Size(), "Set size should be 3")
}

func TestAddSlice(t *testing.T) {
	s := NewSet[int]()
	s.AddSlice([]int{1, 2, 3})

	assert.True(t, s.Contains(1), "Set should contain element 1")
	assert.True(t, s.Contains(2), "Set should contain element 2")
	assert.True(t, s.Contains(3), "Set should contain element 3")
	assert.Equal(t, 3, s.Size(), "Set size should be 3")
}

func TestRemove(t *testing.T) {
	s := NewSet[int]()
	s.AddMultiple(1, 2, 3)
	s.Remove(2)

	assert.False(t, s.Contains(2), "Set should not contain element 2 after removal")
	assert.Equal(t, 2, s.Size(), "Set size should be 2 after removal")
}

func TestToSlice(t *testing.T) {
	s := NewSet[int]()
	s.AddMultiple(1, 2, 3)

	slice := s.ToSlice()
	assert.ElementsMatch(t, []int{1, 2, 3}, slice, "ToSlice should return all elements in the set")
}

func TestUnion(t *testing.T) {
	s1 := NewSet[int]()
	s1.AddMultiple(1, 2, 3)

	s2 := NewSet[int]()
	s2.AddMultiple(3, 4, 5)

	union := s1.Union(s2)
	assert.ElementsMatch(t, []int{1, 2, 3, 4, 5}, union.ToSlice(), "Union should contain all unique elements from both sets")
}

func TestIntersection(t *testing.T) {
	s1 := NewSet[int]()
	s1.AddMultiple(1, 2, 3)

	s2 := NewSet[int]()
	s2.AddMultiple(3, 4, 5)

	intersection := s1.Intersection(s2)
	assert.ElementsMatch(t, []int{3}, intersection.ToSlice(), "Intersection should contain only common elements")
}

func TestDifference(t *testing.T) {
	s1 := NewSet[int]()
	s1.AddMultiple(1, 2, 3)

	s2 := NewSet[int]()
	s2.AddMultiple(3, 4, 5)

	difference := s1.Difference(s2)
	assert.ElementsMatch(t, []int{1, 2}, difference.ToSlice(), "Difference should contain elements in s1 but not in s2")
}

func TestIsEmpty(t *testing.T) {
	s := NewSet[int]()
	assert.True(t, s.IsEmpty(), "New set should be empty")

	s.Add(1)
	assert.False(t, s.IsEmpty(), "Set should not be empty after adding an element")

	s.Remove(1)
	assert.True(t, s.IsEmpty(), "Set should be empty after removing all elements")
}

func TestForEach(t *testing.T) {
	s := NewSet[int]()
	s.AddMultiple(1, 2, 3)

	sum := 0
	s.ForEach(func(element int) {
		sum += element
	})
	assert.Equal(t, 6, sum, "ForEach should iterate over all elements in the set")
}

func TestNewSetWithElements(t *testing.T) {
	s := NewSetWithElements(1, 2, 3)
	assert.ElementsMatch(t, []int{1, 2, 3}, s.ToSlice(), "NewSetWithElements should initialize the set with provided elements")
}

func TestNewSetFromSlice(t *testing.T) {
	s := NewSetFromSlice([]int{1, 2, 3})
	assert.ElementsMatch(t, []int{1, 2, 3}, s.ToSlice(), "NewSetFromSlice should initialize the set with elements from the slice")
}

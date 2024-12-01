package generic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMin(t *testing.T) {
	assert.Equal(t, 3, Min(3, 5), "Min should return the smaller integer")
	assert.Equal(t, -5, Min(-5, 0), "Min should return the smaller negative integer")
	assert.Equal(t, 3.5, Min(3.5, 5.2), "Min should return the smaller float")
	assert.Equal(t, "apple", Min("apple", "banana"), "Min should return the lexicographically smaller string")
}

func TestMax(t *testing.T) {
	assert.Equal(t, 5, Max(3, 5), "Max should return the larger integer")
	assert.Equal(t, 0, Max(-5, 0), "Max should return the larger negative integer")
	assert.Equal(t, 5.2, Max(3.5, 5.2), "Max should return the larger float")
	assert.Equal(t, "banana", Max("apple", "banana"), "Max should return the lexicographically larger string")
}

func TestContainsSlice(t *testing.T) {
	// Test with integers
	intSlice := []int{1, 2, 3, 4, 5}
	assert.True(t, ContainsSlice(intSlice, 3), "ContainsSlice should return true for an existing integer")
	assert.False(t, ContainsSlice(intSlice, 6), "ContainsSlice should return false for a non-existing integer")

	// Test with floats
	floatSlice := []float64{1.1, 2.2, 3.3, 4.4, 5.5}
	assert.True(t, ContainsSlice(floatSlice, 3.3), "ContainsSlice should return true for an existing float")
	assert.False(t, ContainsSlice(floatSlice, 6.6), "ContainsSlice should return false for a non-existing float")

	// Test with strings
	stringSlice := []string{"apple", "banana", "cherry"}
	assert.True(t, ContainsSlice(stringSlice, "banana"), "ContainsSlice should return true for an existing string")
	assert.False(t, ContainsSlice(stringSlice, "grape"), "ContainsSlice should return false for a non-existing string")
}

func TestMinMaxEdgeCases(t *testing.T) {
	// Test Min and Max with equal values
	assert.Equal(t, 5, Min(5, 5), "Min should return the value when both inputs are equal")
	assert.Equal(t, 5, Max(5, 5), "Max should return the value when both inputs are equal")

	// Test Min and Max with large values
	assert.Equal(t, -1000000000, Min(-1000000000, 1000000000), "Min should handle large negative numbers")
	assert.Equal(t, 1000000000, Max(-1000000000, 1000000000), "Max should handle large positive numbers")
}

func TestContainsSliceEdgeCases(t *testing.T) {
	// Test ContainsSlice with an empty slice
	assert.False(t, ContainsSlice([]int{}, 1), "ContainsSlice should return false for an empty slice")

	// Test ContainsSlice with single-element slice
	assert.True(t, ContainsSlice([]int{1}, 1), "ContainsSlice should return true for a single-element slice with a match")
	assert.False(t, ContainsSlice([]int{1}, 2), "ContainsSlice should return false for a single-element slice without a match")
}

package set

// Set is a generic set data structure interface that holds elements of type T.
// The type T must be comparable to be used as a map key.
type Set[T comparable] interface {
	Add(element T)
	AddMultiple(elements ...T)
	AddSlice(elements []T)
	Remove(element T)
	Contains(element T) bool
	Size() int
	IsEmpty() bool
	ToSlice() []T
	Union(other Set[T]) Set[T]
	Intersection(other Set[T]) Set[T]
	Difference(other Set[T]) Set[T]
	ForEach(f func(T))
}

// setImp is the concrete implementation of the Set interface.
type setImp[T comparable] struct {
	elements map[T]struct{}
}

// NewSet creates and returns a new empty Set.
func NewSet[T comparable]() Set[T] {
	return &setImp[T]{elements: make(map[T]struct{})}
}

// NewSetWithElements creates a new set initialized with the provided elements.
func NewSetWithElements[T comparable](elements ...T) Set[T] {
	set := NewSet[T]()
	set.AddMultiple(elements...)
	return set
}

// NewSetFromSlice creates a new set initialized with elements from the provided slice.
func NewSetFromSlice[T comparable](elements []T) Set[T] {
	set := NewSet[T]()
	set.AddSlice(elements)
	return set
}

// Add inserts an element into the set.
func (s *setImp[T]) Add(element T) {
	s.elements[element] = struct{}{}
}

// AddMultiple inserts multiple elements into the set using variadic arguments.
func (s *setImp[T]) AddMultiple(elements ...T) {
	for _, element := range elements {
		s.Add(element)
	}
}

// AddSlice inserts multiple elements from a slice into the set.
func (s *setImp[T]) AddSlice(elements []T) {
	for _, element := range elements {
		s.Add(element)
	}
}

// Remove deletes an element from the set.
func (s *setImp[T]) Remove(element T) {
	delete(s.elements, element)
}

// Contains checks if the set contains the given element.
func (s *setImp[T]) Contains(element T) bool {
	_, exists := s.elements[element]
	return exists
}

// Size returns the number of elements in the set.
func (s *setImp[T]) Size() int {
	return len(s.elements)
}

// IsEmpty returns the number of elements in the set is Zero.
func (s *setImp[T]) IsEmpty() bool {
	return s.Size() == 0
}

// ToSlice returns a slice containing all elements of the set.
func (s *setImp[T]) ToSlice() []T {
	slice := make([]T, 0, len(s.elements))
	for key := range s.elements {
		slice = append(slice, key)
	}
	return slice
}

// Union returns a new set that is the union of s and other.
func (s *setImp[T]) Union(other Set[T]) Set[T] {
	result := NewSet[T]()
	s.ForEach(func(element T) {
		result.Add(element)
	})
	other.ForEach(func(element T) {
		result.Add(element)
	})
	return result
}

// Intersection returns a new set that is the intersection of s and other.
func (s *setImp[T]) Intersection(other Set[T]) Set[T] {
	result := NewSet[T]()
	if s.Size() < other.Size() {
		s.ForEach(func(element T) {
			if other.Contains(element) {
				result.Add(element)
			}
		})
	} else {
		other.ForEach(func(element T) {
			if s.Contains(element) {
				result.Add(element)
			}
		})
	}
	return result
}

// Difference returns a new set that contains elements in s but not in other.
func (s *setImp[T]) Difference(other Set[T]) Set[T] {
	result := NewSet[T]()
	s.ForEach(func(element T) {
		if !other.Contains(element) {
			result.Add(element)
		}
	})
	return result
}

// ForEach iterates over each element in the set and applies the provided function.
func (s *setImp[T]) ForEach(f func(T)) {
	for key := range s.elements {
		f(key)
	}
}

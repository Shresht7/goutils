package set

import (
	"github.com/Shresht7/sliceutils/slice"
)

// Set is a set of unique values
type Set[T comparable] slice.Slice[T]

// New creates a new set from a slice
func New[T comparable](slice []T) Set[T] {

	// Instantiate a new set
	result := make([]T, 0, len(slice))

	// Create a map to store unique values
	uniqueValues := make(map[T]bool)

	// Loop through the slice and add unique values to the set
	for _, v := range slice {
		_, ok := uniqueValues[v]
		if !ok {
			result = append(result, v)
			uniqueValues[v] = true
		}
	}

	// Return the set
	return result
}

// Returns the length of the Set
func (s *Set[T]) Len() int {
	return len(*s)
}

// Returns the capacity of the Set
func (s *Set[T]) Cap() int {
	return cap(*s)
}

// Returns the Set as a normal slice
func (s *Set[T]) AsSlice() []T {
	return *s
}

// Returns the size of the Set
func (s *Set[T]) Size() int {
	return s.Len()
}

// Checks if the set has the given value
func (s *Set[T]) Has(value T) bool {
	return slice.Includes(*s, value)
}

// Add another element to the Set
func (s *Set[T]) Add(value T) *Set[T] {
	if !s.Has(value) {
		s.Add(value)
	}
	return s
}

// Delete an element from the Set
func (s *Set[T]) Delete(value T) *Set[T] {
	*s = slice.Filter(*s, func(v T, _ int) bool { return v != value })
	return s
}

// Clears the entire Set
func (s *Set[T]) Clear() {
	*s = make([]T, 0)
}

// Executes a callback function for each element in the Set
func (s *Set[T]) ForEach(f slice.Callback[T]) {
	slice.ForEach(*s, f)
}

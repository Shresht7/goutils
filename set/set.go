package set

import (
	"github.com/Shresht7/sliceutils/slice"
)

type Set[T comparable] []T

func GetSet[T comparable](slice []T) Set[T] {
	result := make([]T, 0, len(slice))
	uniqueValues := make(map[T]bool)
	for _, v := range slice {
		_, ok := uniqueValues[v]
		if !ok {
			result = append(result, v)
			uniqueValues[v] = true
		}
	}
	return result
}

//	Returns the size of the set
func (s *Set[T]) Size() int {
	return len(*s)
}

//	Set has value
func (s *Set[T]) Has(value T) bool {
	return slice.Includes(*s, value)
}

//	Add another element to the set
func (s *Set[T]) Add(value T) *Set[T] {
	if !s.Has(value) {
		*s = append(*s, value)
	}
	return s
}

//	TODO: Needs Splice
// func (s *Set[T]) Delete(value T) *Set[T] {

// }

//	Clears the entire set
func (s *Set[T]) Clear() {
	*s = make([]T, 0, s.Size())
}

//	Returns a boolean indicating whether this set is a subset of the given set
func (s *Set[T]) IsSubSet(x *Set[T]) bool {
	return slice.Every(*s, func(v T, _ int) bool { return slice.Includes(*x, v) })
}

//	Returns a boolean indicating whether this set is a superset of the given set
func (s *Set[T]) IsSuperSet(x *Set[T]) bool {
	return slice.Every(*x, func(v T, _ int) bool { return slice.Includes(*s, v) })
}

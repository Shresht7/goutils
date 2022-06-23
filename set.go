package sliceutils

import "golang.org/x/exp/constraints"

type Set[T constraints.Ordered] []T

func GetSet[T constraints.Ordered](slice []T) Set[T] {
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
	return Includes(*s, value)
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

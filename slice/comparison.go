package slice

import "golang.org/x/exp/constraints"

//	Check if two slices are equal
func Equal[T constraints.Ordered](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}

	for i, value := range a {
		if value != b[i] {
			return false
		}
	}

	return true
}

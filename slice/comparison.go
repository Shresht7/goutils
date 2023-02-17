package slice

// * EQUAL *//

// Check if two slices are equal
func Equal[T comparable](a, b []T) bool {

	// If the length of the slices are not equal, they are not equal
	if len(a) != len(b) {
		return false
	}

	// Check if each element is equal to the corresponding element in the other slice
	for i, value := range a {
		if value != b[i] {
			return false
		}
	}

	// If all elements are equal, the slices are equal
	return true

}

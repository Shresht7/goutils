package test

// equal returns true if the two given slices are equal, false otherwise
func Equal[T comparable](a, b []T) bool {
	// Check if the slices are the same length
	if len(a) != len(b) {
		return false
	}

	// Check if the slices have the same elements
	for i, v := range a {
		if b[i] != v {
			return false
		}

	}

	// The slices are equal
	return true
}

// Inequality is a helper function that returns true if the actual value is not equal to the expected value
func Inequality[T comparable](actual, expected T) bool {
	return actual != expected
}

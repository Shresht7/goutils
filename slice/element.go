package slice

// * FIRST * //

// Returns the first element of the slice
func First[T any](slice []T) T {
	return slice[0]
}

// Returns the first element of the slice
func (slice *Slice[T]) First() T {
	return First(*slice)
}

// * LAST * //

// Returns the last element of the slice
func Last[T any](slice []T) T {
	return slice[len(slice)-1]
}

// Returns the last element of the slice
func (slice *Slice[T]) Last() T {
	return Last(*slice)
}

// * AT * //

// Returns the element at the given index
func At[T any](slice []T, index int) T {

	// If the index is negative, we need to count from the end of the slice
	if index < 0 {
		index = len(slice) + index
	}

	// If the index is greater than the length of the slice, we need to wrap around
	if index > len(slice)-1 {
		index = index % len(slice)
	}

	// Return the element at the given index
	return slice[index]
}

// Returns the element at the given index
func (slice *Slice[T]) At(index int) T {
	return At(*slice, index)
}

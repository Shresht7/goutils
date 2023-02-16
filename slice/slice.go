package slice

//	Custom Implementation of Slice of type T
type Slice[T any] []T

// Instantiate a new Slice[T]
func New[T any](slice []T) Slice[T] {
	return Slice[T](slice)
}

// Returns the length of the Slice
func (slice *Slice[T]) Len() int {
	return len(*slice)
}

// Returns the capacity of the Slice
func (slice *Slice[T]) Cap() int {
	return cap(*slice)
}

// Returns the Slice as a normal slice
func (slice *Slice[T]) AsSlice() []T {
	return *slice
}

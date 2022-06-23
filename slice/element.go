package slice

//	Returns the first element of the slice
func (slice *Slice[T]) First() T {
	return (*slice)[0]
}

//	Returns the last element of the slice
func (slice *Slice[T]) Last() T {
	return (*slice)[len(*slice)-1]
}

//	Returns the nth element of the slice
func (slice *Slice[T]) Nth(n int) T {
	return (*slice)[n-1]
}

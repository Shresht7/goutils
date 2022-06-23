package slice

//	Returns the first element of the slice
func First[T any](slice []T) T {
	return slice[0]
}

//	Returns the last element of the slice
func Last[T any](slice []T) T {
	return slice[len(slice)-1]
}

//	Returns the nth element of the slice
func Nth[T any](slice []T, n int) T {
	return slice[n-1]
}

package slice

type Slice[T any] []T

//	Turn a slice into a Slice.
func From[T any](slice []T) Slice[T] {
	return slice
}

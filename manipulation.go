package sliceutils

//	Reverse a slice
func Reverse[T any](slice []T) []T {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
	return slice
}

//	Add elements on to the end of the slice. Same as the builtin append().
func Push[T any](slice []T, elements ...T) []T {
	return append(slice, elements...)
}

//	Removes elements from the end of the slice.
func Pop[T any](slice *[]T) T {
	last := len(*slice) - 1
	elem := (*slice)[last:]
	*slice = (*slice)[:last]
	return elem[0]
}

//	Add elements on to the start of the slice.
func Unshift[T any](slice []T, elements ...T) []T {
	return append(elements, slice...)
}

//	Remove elements from the start of the slice.
func Shift[T any](slice *[]T) T {
	elem := (*slice)[:1]
	(*slice) = (*slice)[1:]
	return elem[0]
}

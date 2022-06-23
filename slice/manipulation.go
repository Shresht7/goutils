package slice

import (
	"fmt"
	"strings"
)

//	Concatenate several slices
func Concat[T any](slices ...[]T) []T {
	ret := make([]T, 0, len(slices[0])) //	TODO: Maybe calculate the length ahead of time?
	for _, v := range slices {
		ret = append(ret, v...)
	}
	return ret
}

//	Reverse a slice
func Reverse[T any](slice []T) []T {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
	return slice
}

//	Assemble elements of a slice into an string using fmt.Sprint(element)
func Join[T any](slice []T, separator string) string {
	result := make([]string, len(slice))
	for i, v := range slice {
		result[i] = fmt.Sprint(v)
	}
	return strings.Join(result, separator)
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
	return append(elements, slice...) //	?	Does this actually modify the slice?
}

//	Remove elements from the start of the slice.
func Shift[T any](slice *[]T) T {
	elem := (*slice)[:1]
	(*slice) = (*slice)[1:]
	return elem[0]
}

//	TODO: #8 Implement Splice

//	Distribute elements of a slice into chunks of (at most) given size.
func Chunk[T any](slice []T, size int) [][]T {
	chunks := make([][]T, 0, (len(slice)/size)+1)

	for len(slice)/size > 0 {
		chunks = append(chunks, slice[:size])
		slice = slice[size:]
	}
	chunks = append(chunks, slice)

	return chunks
}

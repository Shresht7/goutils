package slice

import (
	"fmt"
	"strings"
)

// * CONCAT *

// Concatenate several slices
func Concat[T any](slices ...[]T) []T {
	ret := make([]T, 0, len(slices[0])) // ? Maybe calculate the total capacity ahead of time?
	for _, v := range slices {
		ret = append(ret, v...)
	}
	return ret
}

func (slice *Slice[T]) Concat(slices ...[]T) *Slice[T] {
	*slice = append(*slice, Concat(slices...)...)
	return slice
}

// * REVERSE *

// Reverse a slice
func Reverse[T any](slice []T) []T {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
	return slice
}

// Reverse a slice in place
func (slice *Slice[T]) Reverse() *Slice[T] {
	*slice = Reverse(*slice)
	return slice
}

// * JOIN *

// Assemble elements of a slice into an string using fmt.Sprint(element)
func Join[T any](slice []T, separator string) string {
	result := make([]string, len(slice))
	for i, v := range slice {
		result[i] = fmt.Sprint(v)
	}
	return strings.Join(result, separator)
}

// Assemble elements of a slice into an string using fmt.Sprint(element)
func (slice *Slice[T]) Join(separator string) string {
	return Join(*slice, separator)
}

// * PUSH *

// Add elements on to the end of the slice. Same as the builtin append().
func Push[T any](slice []T, elements ...T) []T {
	return append(slice, elements...)
}

// Add elements on to the end of the slice. Same as the builtin append().
func (slice *Slice[T]) Push(elements ...T) *Slice[T] {
	*slice = append(*slice, elements...)
	return slice
}

// * POP *

// Removes elements from the end of the slice.
func Pop[T any](slice *[]T) T {
	last := len(*slice) - 1
	elem := (*slice)[last:]
	*slice = (*slice)[:last]
	return elem[0]
}

// Removes elements from the end of the slice.
func (slice *Slice[T]) Pop() T {
	last := len(*slice) - 1
	elem := (*slice)[last:]
	*slice = (*slice)[:last]
	return elem[0]
}

// * UNSHIFT *

// Add elements on to the start of the slice.
func Unshift[T any](slice []T, elements ...T) []T {
	return append(elements, slice...) //	?	Does this actually modify the slice?
}

// Add elements on to the start of the slice.
func (slice *Slice[T]) Unshift(elements ...T) *Slice[T] {
	*slice = append(elements, *slice...)
	return slice
}

// * SHIFT *

// Remove elements from the start of the slice.
func Shift[T any](slice *[]T) T {
	elem := (*slice)[:1]
	(*slice) = (*slice)[1:]
	return elem[0]
}

// Remove elements from the start of the slice.
func (slice *Slice[T]) Shift() T {
	elem := (*slice)[:1]
	*slice = (*slice)[1:]
	return elem[0]
}

//	TODO: #8 Implement Splice

// * CHUNK *

// Distribute elements of a slice into chunks of (at most) given size.
func Chunk[T any](slice []T, size int) [][]T {
	chunks := make([][]T, 0, (len(slice)/size)+1)

	for len(slice)/size > 0 {
		chunks = append(chunks, slice[:size])
		slice = slice[size:]
	}
	chunks = append(chunks, slice)

	return chunks
}

// Distribute elements of a slice into chunks of (at most) given size.
func (slice *Slice[T]) Chunk(size int) [][]T {
	return Chunk(*slice, size)
}

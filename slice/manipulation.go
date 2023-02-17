package slice

import (
	"fmt"
	"strings"
)

// * CONCAT * //

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

// * REVERSE * //

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

// * JOIN * //

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

// * PUSH * //

// Add elements on to the end of the slice. Same as the builtin append().
func Push[T any](slice []T, elements ...T) []T {
	return append(slice, elements...)
}

// Add elements on to the end of the slice. Same as the builtin append().
func (slice *Slice[T]) Push(elements ...T) *Slice[T] {
	*slice = append(*slice, elements...)
	return slice
}

// * POP * //

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

// * UNSHIFT * //

// Add elements on to the start of the slice.
func Unshift[T any](slice []T, elements ...T) []T {
	return append(elements, slice...) //	?	Does this actually modify the slice?
}

// Add elements on to the start of the slice.
func (slice *Slice[T]) Unshift(elements ...T) *Slice[T] {
	*slice = append(elements, *slice...)
	return slice
}

// * SHIFT * //

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

// * SPLICE * //

// Remove elements from the slice and replace them with new elements.
func Splice[T any](slice *[]T, start, deleteCount int, elements ...T) []T {
	// If the start is negative, start from the end of the slice
	if start < 0 {
		start = len(*slice) + start
	}

	// If the start is greater than the length of the slice, start from the end of the slice
	if start > len(*slice) {
		start = len(*slice)
	}

	// If the deleteCount is negative, set it to 0
	if deleteCount < 0 {
		deleteCount = 0
	}

	// If the deleteCount is greater than the length of the slice, set it to the length of the slice
	if deleteCount > len(*slice) {
		deleteCount = len(*slice)
	}

	// If the deleteCount is greater than the length of the slice minus the start, set it to the length of the slice minus the start
	if deleteCount > len(*slice)-start {
		deleteCount = len(*slice) - start
	}

	// Remove the elements
	removed := (*slice)[start : start+deleteCount]
	*slice = append((*slice)[:start], (*slice)[start+deleteCount:]...)

	// Add the new elements
	*slice = append((*slice)[:start], append(elements, (*slice)[start:]...)...)

	// Return the removed elements
	return removed
}

// * CHUNK * //

// Distribute elements of a slice into chunks of (at most) given size
func Chunk[T any](slice []T, size int) [][]T {

	// If the size is 0 or negative, return the slice as a single chunk
	if size <= 0 {
		return [][]T{slice}
	}

	// Instantiate a slice of slices to hold the chunks
	chunks := make([][]T, 0, (len(slice)/size)+1)

	// While elements can be distributed into chunks of given size
	// append the chunk to the chunks slice and remove the chunk from the slice
	for len(slice)/size > 0 {
		chunks = append(chunks, slice[:size])
		slice = slice[size:]
	}
	// Append the remaining elements to the chunks slice
	chunks = append(chunks, slice)

	// Return the chunks slice
	return chunks
}

// Distribute elements of a slice into chunks of (at most) given size.
func (slice *Slice[T]) Chunk(size int) [][]T {
	return Chunk(*slice, size)
}

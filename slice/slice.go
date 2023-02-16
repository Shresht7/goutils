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

//	Append slices to this slice
func (slice *Slice[T]) Concat(slices ...[]T) Slice[T] {
	for _, v := range slices {
		*slice = append(*slice, v...)
	}
	return *slice
}

//	Reverse this slice
func (slice *Slice[T]) Reverse() Slice[T] {
	*slice = Reverse(*slice)
	return *slice
}

//	Join elements of this slice into a string using fmt.Sprint(element)
func (slice *Slice[T]) Join(separator string) string {
	return Join(*slice, separator)
}

//	Add elements on to the end of the slice. Same as the builtin append.
func (slice *Slice[T]) Push(elements ...T) []T {
	*slice = append(*slice, elements...)
	return *slice
}

//	Removes element from the end of the slice
func (slice *Slice[T]) Pop() T {
	last := len(*slice) - 1
	elem := (*slice)[last:]
	*slice = (*slice)[:last]
	return elem[0]
}

//	Add elements to the start of the slice
func (slice *Slice[T]) Unshift(elements ...T) []T {
	*slice = append(elements, *slice...)
	return *slice
}

//	Remove the element from the start of the slice
func (slice *Slice[T]) Shift() T {
	elem := (*slice)[:1]
	(*slice) = (*slice)[1:]
	return elem[0]
}

//	Distribute elements of a slice into chunks of (at most) given size
func (slice *Slice[T]) Chunk(size int) [][]T {
	return Chunk(*slice, size)
}

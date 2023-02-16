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

//	Executes the callback function for each element of the slice.
//	The callback receives the value and the index of the element.
func (slice *Slice[T]) ForEach(cb Callback[T]) {
	ForEach(*slice, cb)
}

//	Filter the slice based on the callback criteria. Elements are filtered out
//	if the callback returns false. The callback receives the element's value and index.
func (slice *Slice[T]) Filter(cb ReturnCallback[T, bool]) []T {
	return Filter(*slice, cb)
}

//	Map returns a new slice by applying the callback to each element of the given slice.
//	The callback receives the value and the index and returns the mapped value.
//	Note that the Map method can only return a slice of the same type.
//	For more powerful transforms use the Map function.
func (slice *Slice[T]) Map(cb ReturnCallback[T, T]) []T {
	return Map(*slice, cb)
}

//	Reduces a slice into a single value using the given callback function.
//	The callback receives the accumulated value, current value, index and the slice.
//	Note that the Reduce method can only reduce the slice to a value of the constituting type.
//	For more powerful transforms use the Reduce function.
func (slice *Slice[T]) Reduce(cb ReducerCallback[T, T], initializer T) T {
	return Reduce(*slice, cb, initializer)
}

//	Returns true if all elements satisfy the callback function
func (slice *Slice[T]) Every(cb ReturnCallback[T, bool]) bool {
	return Every(*slice, cb)
}

//	Returns true if any one element satisfies the callback function
func (slice *Slice[T]) Some(cb ReturnCallback[T, bool]) bool {
	return Some(*slice, cb)
}

//	Returns the first element in the slice that satisfies the callback criteria.
//	Returns a boolean indicating the success of the find operation, followed by the value and it's index.
//	If the criteria fails, the boolean is false - i.e. the function failed to find the value.
func (slice *Slice[T]) Find(cb ReturnCallback[T, bool]) (bool, T, int) {
	return Find(*slice, cb)
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

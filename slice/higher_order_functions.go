package slice

// * FOREACH * //

// Execute the callback function for each element of the slice.
// The callback receives the element's value and index.
func ForEach[T any](slice []T, cb Callback[T]) {
	for i, v := range slice {
		cb(v, i)
	}
}

// Execute the callback function for each element of the slice.
// The callback receives the element's value and index.
func (slice *Slice[T]) ForEach(cb Callback[T]) {
	ForEach(*slice, cb)
}

// * FILTER * //

// Filter the slice based on the callback criteria. Elements are filtered out
// if the callback returns false. The callback receives the element's value and index.
func Filter[T any](slice []T, cb ReturnCallback[T, bool]) []T {
	ret := make([]T, 0, len(slice))

	for i, v := range slice {
		if cb(v, i) {
			ret = append(ret, v)
		}
	}

	return ret
}

// Filter the slice based on the callback criteria. Elements are filtered out
// if the callback returns false. The callback receives the element's value and index.
func (slice *Slice[T]) Filter(cb ReturnCallback[T, bool]) []T {
	return Filter(*slice, cb)
}

// * MAP * //

// Map returns a new slice by applying the callback to each element of the given slice.
// The callback receives the value and the index and returns the mapped value.
func Map[From, To any](slice []From, cb ReturnCallback[From, To]) []To {
	ret := make([]To, len(slice))

	for i, v := range slice {
		ret[i] = cb(v, i)
	}

	return ret
}

// Map returns a new slice by applying the callback to each element of the given slice.
// The callback receives the value and the index and returns the mapped value.
func (slice *Slice[T]) Map(cb ReturnCallback[T, T]) []T {
	return Map(*slice, cb)
}

// * REDUCE * //

// Reduces a slice into a single value using the given callback function.
// The callback receives the accumulated value, current value, index and the slice.
func Reduce[From, To any](slice []From, cb ReducerCallback[From, To], initializer To) To {
	val := initializer

	for i, v := range slice {
		val = cb(val, v, i, slice)
	}

	return val
}

// Reduces a slice into a single value using the given callback function.
// The callback receives the accumulated value, current value, index and the slice.
func (slice *Slice[T]) Reduce(cb ReducerCallback[T, T], initializer T) T {
	return Reduce(*slice, cb, initializer)
}

// * EVERY *//

// Returns true if all elements satisfy the callback function
func Every[T any](slice []T, cb ReturnCallback[T, bool]) bool {

	for i, v := range slice {
		result := cb(v, i)
		if !result {
			return false
		}
	}

	return true
}

// Returns true if all elements satisfy the callback function
func (slice *Slice[T]) Every(cb ReturnCallback[T, bool]) bool {
	return Every(*slice, cb)
}

// * SOME * //

// Returns true if any one element satisfies the callback function
func Some[T any](slice []T, cb ReturnCallback[T, bool]) bool {

	for i, v := range slice {
		if cb(v, i) {
			return true
		}
	}

	return false
}

// Returns true if any one element satisfies the callback function
func (slice *Slice[T]) Some(cb ReturnCallback[T, bool]) bool {
	return Some(*slice, cb)
}

//	? Use All and Any instead of Every and Some

// * FIND * //

// Returns the first element in the slice that satisfies the callback criteria.
// Returns a boolean indicating the success of the find operation, followed by the value and it's index.
// If the criteria fails, the boolean is false - i.e. the function failed to find the value.
func Find[T any](slice []T, cb ReturnCallback[T, bool]) (bool, T, int) {

	for i, v := range slice {
		if cb(v, i) {
			return true, v, i
		}
	}

	return false, slice[0], -1

}

// Returns the first element in the slice that satisfies the callback criteria.
// Returns a boolean indicating the success of the find operation, followed by the value and it's index.
// If the criteria fails, the boolean is false - i.e. the function failed to find the value.
func (slice *Slice[T]) Find(cb ReturnCallback[T, bool]) (bool, T, int) {
	return Find(*slice, cb)
}

// * INCLUDES * //

// Returns true when the slice contains the given element.
func Includes[T comparable](slice []T, element T) bool {
	ok, _, _ := Find(slice, func(v T, _ int) bool { return v == element })
	return ok
}

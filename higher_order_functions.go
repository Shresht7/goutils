package sliceutils

//	Execute the callback function for each element of the slice.
//	The callback receives the element's value and index.
func ForEach[T any](slice []T, cb Callback[T]) {
	for i, v := range slice {
		cb(v, i)
	}
}

//	Filter the slice based on the callback criteria. Elements are filtered out
//	if the callback returns false. The callback receives the element's value and index.
func Filter[T any](slice []T, cb ReturnCallback[T, bool]) []T {
	ret := make([]T, 0, len(slice))

	for i, v := range slice {
		if cb(v, i) {
			ret = append(ret, v)
		}
	}

	return ret
}

//	Map returns a new slice by applying the callback to each element of the given slice.
//	The callback receives the value and the index and returns the mapped value.
func Map[From, To any](slice []From, cb ReturnCallback[From, To]) []To {
	ret := make([]To, len(slice))

	for i, v := range slice {
		ret[i] = cb(v, i)
	}

	return ret
}

//	Reduces a slice into a single value using the given callback function.
//	The callback receives the accumulated value, current value, index and the slice.
func Reduce[From, To any](slice []From, cb ReducerCallback[From, To], initializer To) To {
	val := initializer

	for i, v := range slice {
		val = cb(val, v, i, slice)
	}

	return val
}

//	Returns true if all elements satisfy the callback function
func Every[T any](slice []T, cb ReturnCallback[T, bool]) bool {

	for i, v := range slice {
		result := cb(v, i)
		if !result {
			return false
		}
	}

	return true
}

//	Returns true if any one element satisfies the callback function
func Some[T any](slice []T, cb ReturnCallback[T, bool]) bool {

	for i, v := range slice {
		if cb(v, i) {
			return true
		}
	}

	return false
}

//	? Use All and Any instead of Every and Some

//	Returns the first element in the slice that satisfies the callback criteria.
//	Returns a boolean indicating the success of the find operation, followed by the value and it's index.
//	If the criteria fails, the boolean is false - i.e. the function failed to find the value.
func Find[T any](slice []T, cb ReturnCallback[T, bool]) (bool, T, int) {

	for i, v := range slice {
		if cb(v, i) {
			return true, v, i
		}
	}

	return false, slice[0], -1

}

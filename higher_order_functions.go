package sliceutils

func ForEach[T any](slice []T, cb Callback[T]) {
	for i, v := range slice {
		cb(v, i)
	}
}

func Filter[T any](slice []T, cb ReturnCallback[T, bool]) []T {
	ret := make([]T, 0, len(slice))

	for i, v := range slice {
		if cb(v, i) {
			ret = append(ret, v)
		}
	}

	return ret
}

func Map[From, To any](slice []From, cb ReturnCallback[From, To]) []To {
	ret := make([]To, len(slice))

	for i, v := range slice {
		ret[i] = cb(v, i)
	}

	return ret
}

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

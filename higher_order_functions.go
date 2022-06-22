package sliceutils

func ForEach[T any](slice []T, cb func(value T, index int)) {
	for i, v := range slice {
		cb(v, i)
	}
}

func Filter[T any](slice []T, cb func(value T, index int) bool) []T {
	ret := make([]T, 0, len(slice))

	for i, v := range slice {
		if cb(v, i) {
			ret = append(ret, v)
		}
	}

	return ret
}

func Map[T any](slice []T, cb func(value T, index int) T) []T {
	ret := make([]T, len(slice))

	for i, v := range slice {
		ret[i] = cb(v, i)
	}

	return ret
}

func Reduce[From any, To any](slice []From, cb func(accumulator To, current From, index int, slice []From) To, initializer To) To {
	val := initializer

	for i, v := range slice {
		val = cb(val, v, i, slice)
	}

	return val
}

//	Returns true if all elements satisfy the callback function
func Every[T any](slice []T, cb func(value T, index int) bool) bool {

	for i, v := range slice {
		result := cb(v, i)
		if !result {
			return false
		}
	}

	return true
}

//	Returns true if any one element satisfies the callback function
func Some[T any](slice []T, cb func(value T, index int) bool) bool {

	for i, v := range slice {
		if cb(v, i) {
			return true
		}
	}

	return false
}

//	? Use All and Any instead of Every and Some

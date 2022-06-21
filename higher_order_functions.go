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

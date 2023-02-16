package slice

//	Callback function
type Callback[T any] func(value T, index int)

//	Callback function that returns a value
type ReturnCallback[T, R any] func(value T, index int) R

//	Reducer callback function
type ReducerCallback[T, R any] func(accumulator R, current T, index int, slice []T) R

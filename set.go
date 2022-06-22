package sliceutils

import "golang.org/x/exp/constraints"

type Set[T any] []T

func GetSet[T constraints.Ordered](slice []T) Set[T] {
	result := make([]T, 0, len(slice))
	uniqueValues := make(map[T]bool)
	for _, v := range slice {
		_, ok := uniqueValues[v]
		if !ok {
			result = append(result, v)
			uniqueValues[v] = true
		}
	}
	return result
}

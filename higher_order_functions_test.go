package sliceutils

import (
	"testing"
)

//	TODO: Improve these tests

var slice []int = []int{0, 1, 2, 3, 4, 5, 6}

func TestForEach(t *testing.T) {

	length := 0
	ForEach(slice, func(_, _ int) { length++ })

	if length != len(slice) {
		t.Error("Failed to forEach over a slice")
	}

}

func TestFilter(t *testing.T) {

	filteredSlice := Filter(slice, func(v, _ int) bool { return v%2 == 0 })

	if !Equal(filteredSlice, []int{0, 2, 4, 6}) {
		t.Error("Failed to filter slice")
	}

}

func TestMap(t *testing.T) {

	mappedSlice := Map(slice, func(v, _ int) int { return v * 2 })

	if !Equal(mappedSlice, []int{0, 2, 4, 6, 8, 10, 12}) {
		t.Error("Failed to map over slice")
	}

}

func TestReduce(t *testing.T) {

	sum := Reduce(slice, func(accumulator, current int, _ int, _ []int) int { return accumulator + current }, 0)

	if sum != 21 {
		t.Error("Failed to reduce slice")
	}

}

func TestEvery(t *testing.T) {

	slice := []int{2, 4, 6, 8, 10}

	divisibleBy2 := Every(slice, func(value, idx int) bool { return value%2 == 0 })

	if !divisibleBy2 {
		t.Error("Every element should be divisible by 2")
	}

}

func TestSome(t *testing.T) {

	slice := []int{1, 3, 5, 7, 9, 10}

	divisibleBy2 := Some(slice, func(value, idx int) bool { return value%2 == 0 })

	if !divisibleBy2 {
		t.Error("At least one element should be divisible by 2")
	}

}

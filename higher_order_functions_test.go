package sliceutils

import (
	"testing"
)

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

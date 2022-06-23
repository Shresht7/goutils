package slice

import (
	"testing"
)

var originalSlice []int = []int{0, 1, 2, 3}

//	Test First
func TestFirstNthAndLast(t *testing.T) {

	slice := From(originalSlice)

	first := slice.First()
	last := slice.Last()

	if first != 0 {
		t.Error("Failed to get the first element from the slice")
	}

	if last != 3 {
		t.Error("Failed to get the last element from the slice")
	}

	n := slice.Nth(2)

	if n != 1 {
		t.Error("Failed to get 2nd element")
	}

}

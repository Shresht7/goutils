package sliceutils

import (
	"fmt"
	"testing"
)

var originalSlice []int = []int{0, 1, 2, 3}

//	Test First
func TestFirstNthAndLast(t *testing.T) {

	var slice []int = originalSlice[:]

	first := First(slice)
	last := Last(slice)

	if first != 0 {
		t.Error("Failed to get the first element from the slice")
	}

	if last != 3 {
		t.Error("Failed to get the last element from the slice")
	}

	n := Nth(slice, 2)

	if n != 1 {
		t.Error("Failed to get 2nd element")
	}

}

//	Test Pop
func TestPop(t *testing.T) {

	var slice []int = originalSlice[:]

	fmt.Println(len(slice))

	//	Pop an element off the slice
	elem := Pop(&slice)

	fmt.Println(len(slice))

	if elem != Last(originalSlice) || len(slice) != len(originalSlice)-1 {
		t.Error("Failed to pop an element off the slice")
	}

}

//	Test Unshift
func TestUnshift(t *testing.T) {

	var slice []int = originalSlice[:]

	//	Add element to the start
	slice = Unshift(slice, -1)

	if First(slice) != -1 || len(slice) != len(originalSlice)+1 {
		t.Error("Failed to unshift an element onto the slice")
	}

	//	Add multiple elements to the start
	slice = Unshift(originalSlice, -3, -2, -1)

	if First(slice) != -3 || len(slice) != len(originalSlice)+3 {
		t.Error("Failed to add multiple elements onto the slice")
	}

}

//	Test Shift
func TestShift(t *testing.T) {

	var slice []int = originalSlice[:]

	//	Remove an element off the slice
	elem := Shift(&slice)

	if elem != First(originalSlice) || len(slice) != len(originalSlice)-1 {
		t.Error("Failed to pop an element off the slice")
	}

}

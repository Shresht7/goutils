package set

import (
	"fmt"
	"testing"

	"github.com/Shresht7/sliceutils/slice"
)

// * SET *//

func TestNew(t *testing.T) {

	// Create a slice
	slice := []int{1, 2, 2, 3, 3, 3, 4, 4, 4, 4}

	// Create a set from the slice
	set := New(slice)

	// Check if the set has the correct length
	if set.Len() != len(slice) {
		t.Errorf("Set has incorrect length. Expected %v, got %v", len(slice), set.Len())
	}

	// Check if the set has the correct capacity
	if set.Cap() != cap(slice) {
		t.Errorf("Set has incorrect capacity. Expected %v, got %v", cap(slice), set.Cap())
	}

	// Check if the set has the correct size
	if set.Size() != len(slice) {
		t.Errorf("Set has incorrect size. Expected %v, got %v", len(slice), set.Size())
	}

	// Check if the set has the correct values
	for _, v := range slice {
		if !set.Has(v) {
			t.Errorf("Set does not contain value %v", v)
		}
	}

}

func ExampleNew() {
	// Create a slice
	slice := []int{1, 2, 2, 3, 3, 3, 4, 4, 4, 4}

	// Create a set from the slice
	set := New(slice)

	// Print the set
	fmt.Println(set)

	// Output: [1 2 3 4]
}

func ExampleSet_sliceMethods() {
	// Create a Set
	set := New([]int{1, 2, 2, 3, 3, 3, 4, 4, 4, 4})

	// Cast the set to a slice
	slice := slice.Slice[int](set)

	slice.ForEach(func(v, i int) {
		fmt.Println(v)
	})

	mapped := slice.Map(func(v, i int) int {
		return v * 2
	})

	fmt.Println(mapped)

	// Output:
	// 1
	// 2
	// 3
	// 4
	// [2 4 6 8]
}

// * HAS *//

func TestHas(t *testing.T) {

	// Create a slice
	slice := []int{1, 2, 2, 3, 3, 3, 4, 4, 4, 4}
	// Create a set from the slice
	set := New(slice)

	// Test Cases
	testCases := []struct {
		value    int
		expected bool
	}{
		{1, true},
		{2, true},
		{3, true},
		{4, true},
		{5, false},
		{6, false},
		{7, false},
		{8, false},
	}

	// Run Test Cases
	for _, tc := range testCases {
		if set.Has(tc.value) != tc.expected {
			t.Errorf("Set.Has(%v) returned %v, expected %v", tc.value, set.Has(tc.value), tc.expected)
		}
	}

}

func ExampleSet_Has() {
	// Create a slice
	slice := []int{1, 2, 2, 3, 3, 3, 4, 4, 4, 4}

	// Create a set from the slice
	set := New(slice)

	// Check if the set has the value 2
	fmt.Println(set.Has(2))

	// Output: true
}

// * ADD *//

func TestAdd(t *testing.T) {

	// Create a Set
	set := New([]int{1, 2, 2, 3, 3, 3, 4, 4, 4, 4})

	set.Add(5)
	set.Add(7)

	// Test Cases
	testCases := []struct {
		value    int
		expected bool
	}{
		{1, true},
		{2, true},
		{3, true},
		{4, true},
		{5, true},
		{6, false},
		{7, true},
		{8, false},
	}

	// Run Test Cases
	for _, tc := range testCases {
		if set.Has(tc.value) != tc.expected {
			t.Errorf("Set.Has(%v) returned %v, expected %v", tc.value, set.Has(tc.value), tc.expected)
		}
	}

}

func ExampleSet_Add() {
	// Create a Set
	set := New([]int{1, 2, 2, 3, 3, 3, 4, 4, 4, 4})

	set.Add(5)
	set.Add(7)

	// Print the set
	fmt.Println(set)

	// Output: [1 2 3 4 5 7]
}

// * CLEAR *//

func TestClear(t *testing.T) {

	// Create a Set
	set := New([]int{1, 2, 2, 3, 3, 3, 4, 4, 4, 4})

	// Clear the set
	set.Clear()

	// Check if the set is empty
	if len(set) != 0 {
		t.Errorf("Set is not empty")
	}

}

func ExampleSet_Clear() {
	// Create a Set
	set := New([]int{1, 2, 2, 3, 3, 3, 4, 4, 4, 4})

	// Clear the set
	set.Clear()

	// Print the set
	fmt.Println(set)

	// Output: []
}

// * DELETE *//

func TestDelete(t *testing.T) {

	// Create a Set
	set := New([]int{1, 2, 2, 3, 3, 3, 4, 4, 4, 4})

	set.Delete(2)
	set.Delete(4)

	// Test Cases
	testCases := []struct {
		value    int
		expected bool
	}{
		{1, true},
		{2, false},
		{3, true},
		{4, false},
		{5, false},
		{6, false},
		{7, false},
		{8, false},
	}

	// Run Test Cases
	for _, tc := range testCases {
		if set.Has(tc.value) != tc.expected {
			t.Errorf("Set.Has(%v) returned %v, expected %v", tc.value, set.Has(tc.value), tc.expected)
		}
	}

}

func ExampleSet_Delete() {
	// Create a Set
	set := New([]int{1, 2, 2, 3, 3, 3, 4, 4, 4, 4})

	set.Delete(2)
	set.Delete(4)

	// Print the set
	fmt.Println(set)

	// Output: [1 3]
}

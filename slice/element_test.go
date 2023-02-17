package slice

import (
	"fmt"
	"testing"
)

// * FIRST * //

func TestFirst(t *testing.T) {

	// Test Data
	slice := []int{0, 1, 2, 3}

	// Test Cases
	testCases := []struct {
		desc     string
		actual   int
		expected int
	}{
		{
			desc:     "Should return the first element of the slice",
			actual:   First(slice),
			expected: 0,
		},
	}

	// Run Tests
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if tC.actual != tC.expected {
				t.Errorf("%s\nwant %d, got %d", tC.desc, tC.expected, tC.actual)
			}
		})
	}

}

func ExampleFirst() {
	// Test Data
	slice := []int{0, 1, 2, 3}

	first := First(slice)

	fmt.Println(first)

	// Output: 0
}

func ExampleFirst_method() {
	// Test Data
	slice := New([]int{0, 1, 2, 3})

	first := slice.First()

	fmt.Println(first)

	// Output: 0
}

// * LAST * //

func TestLast(t *testing.T) {

	// Test Data
	slice := []int{0, 1, 2, 3}

	// Test Cases
	testCases := []struct {
		desc     string
		actual   int
		expected int
	}{
		{
			desc:     "Should return the last element of the slice",
			actual:   Last(slice),
			expected: 3,
		},
	}

	// Run Tests
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if tC.actual != tC.expected {
				t.Errorf("%s\nwant %d, got %d", tC.desc, tC.expected, tC.actual)
			}
		})
	}

}

func ExampleLast() {
	// Test Data
	slice := []int{0, 1, 2, 3}

	last := Last(slice)

	fmt.Println(last)

	// Output: 3
}

func ExampleLast_method() {
	// Test Data
	slice := New([]int{0, 1, 2, 3})

	last := slice.Last()

	fmt.Println(last)

	// Output: 3
}

// * NTH * //

func TestNth(t *testing.T) {

	// Test Data
	slice := []int{0, 1, 2, 3}

	// Test Cases
	testCases := []struct {
		desc     string
		actual   int
		expected int
	}{
		{
			desc:     "Should return the nth element of the slice",
			actual:   Nth(slice, 2),
			expected: 1,
		},
	}

	// Run Tests
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if tC.actual != tC.expected {
				t.Errorf("%s\nwant %d, got %d", tC.desc, tC.expected, tC.actual)
			}
		})
	}

}

func ExampleNth() {
	// Test Data
	slice := []int{0, 1, 2, 3}

	nth := Nth(slice, 2)

	fmt.Println(nth)

	// Output: 1
}

func ExampleNth_method() {
	// Test Data
	slice := New([]int{0, 1, 2, 3})

	nth := slice.Nth(2)

	fmt.Println(nth)

	// Output: 1
}

// * AT *//

func TestAt(t *testing.T) {

	// Test Data
	slice := []int{0, 1, 2, 3}

	// Test Cases
	testCases := []struct {
		desc     string
		actual   int
		expected int
	}{
		{
			desc:     "Should return the first element of the slice",
			actual:   At(slice, 0),
			expected: 0,
		},
		{
			desc:     "Should return the last element of the slice",
			actual:   At(slice, -1),
			expected: 3,
		},
		{
			desc:     "Should return the nth element of the slice",
			actual:   At(slice, 2),
			expected: 2,
		},
	}

	// Run Tests
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if tC.actual != tC.expected {
				t.Errorf("%s\nwant %d, got %d", tC.desc, tC.expected, tC.actual)
			}
		})
	}

}

func ExampleAt() {
	a := []int{1, 2, 3, 4, 5}

	fmt.Println(At(a, 0))
	fmt.Println(At(a, -1))
	fmt.Println(At(a, 2))

	// Output:
	// 1
	// 5
	// 3
}

func ExampleAt_method() {
	a := New([]int{1, 2, 3, 4, 5})

	fmt.Println(a.At(0))
	fmt.Println(a.At(-1))
	fmt.Println(a.At(2))

	// Output:
	// 1
	// 5
	// 3
}

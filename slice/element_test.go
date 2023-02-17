package slice

import (
	"fmt"
	"testing"
)

func TestElement(t *testing.T) {

	// Test Cases
	testCases := []struct {
		desc     string
		slice    []int
		fn       func(slice []int) int
		expected int
	}{
		{
			desc:     "Should return the first element of the slice",
			slice:    []int{0, 1, 2, 3},
			fn:       First[int],
			expected: 0,
		},
		{
			desc:     "Should return the last element of the slice",
			slice:    []int{0, 1, 2, 3},
			fn:       Last[int],
			expected: 3,
		},
		{
			desc:     "Should return the nth element of the slice",
			slice:    []int{0, 1, 2, 3},
			fn:       func(slice []int) int { return Nth(slice, 2) },
			expected: 1,
		},
	}

	// Run Tests
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			actual := tC.fn(tC.slice)
			if actual != tC.expected {
				t.Errorf("want %d, got %d", tC.expected, actual)
			}
		})
	}

}

func ExampleFirst() {
	a := []int{1, 2, 3, 4, 5}

	fmt.Println(First(a))

	// Output:
	// 1
}

func ExampleFirst_method() {
	a := New([]int{1, 2, 3, 4, 5})

	fmt.Println(a.First())

	// Output:
	// 1
}

func ExampleLast() {
	a := []int{1, 2, 3, 4, 5}

	fmt.Println(Last(a))

	// Output:
	// 5
}

func ExampleLast_method() {
	a := New([]int{1, 2, 3, 4, 5})

	fmt.Println(a.Last())

	// Output:
	// 5
}

func ExampleNth() {
	a := []int{1, 2, 3, 4, 5}

	fmt.Println(Nth(a, 2))

	// Output:
	// 2
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

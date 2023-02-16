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

	// Run test cases
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

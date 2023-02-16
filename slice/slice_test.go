package slice

import (
	"testing"
)

func TestSlice(t *testing.T) {

	originalSlice := []int{1, 2, 3, 4, 5}

	// Test Cases
	testCases := []struct {
		desc     string
		fn       func() int
		expected int
	}{
		{
			desc: "Should return the length of the slice",
			fn: func() int {
				slice := New(originalSlice)
				return slice.Len()
			},
			expected: len(originalSlice),
		},
		{
			desc: "Should return the capacity of the slice",
			fn: func() int {
				slice := New(originalSlice)
				return slice.Cap()
			},
			expected: cap(originalSlice),
		},
		{
			desc: "Should return the slice as a normal slice",
			fn: func() int {
				slice := New(originalSlice)
				return len(slice.AsSlice())
			},
			expected: len(originalSlice),
		},
	}

	// Run Tests
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			actual := tc.fn()
			if actual != tc.expected {
				t.Errorf("Expected: %v, Actual: %v", tc.expected, actual)
			}
		})
	}

}

func TestSliceInteroperability(t *testing.T) {

	originalSlice := []int{1, 2, 3, 4, 5}

	// Test Cases
	testCases := []struct {
		desc     string
		fn       func() int
		expected int
	}{
		{
			desc: "Should be able to append to the slice",
			fn: func() int {
				slice := New(originalSlice)
				slice = append(slice, 6)
				return slice.Len()
			},
			expected: len(originalSlice) + 1,
		},
		{
			desc: "Should be interoperable with normal slices",
			fn: func() int {
				slice := New(originalSlice)
				slice = append(slice, 6)
				return len(slice)
			},
			expected: len(originalSlice) + 1,
		},
		{
			desc: "Should be able to use the slice in functions that accept normal slices",
			fn: func() int {
				slice := New(originalSlice)
				return CalculateTotal(slice)
			},
			expected: CalculateTotal(originalSlice),
		},
	}

	// Run Tests
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			actual := tc.fn()
			if actual != tc.expected {
				t.Errorf("Expected: %v, Actual: %v", tc.expected, actual)
			}
		})
	}

}

// -------
// HELPERS
// -------

// CalculateTotal calculates the total of a slice of numbers
func CalculateTotal(numbers []int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

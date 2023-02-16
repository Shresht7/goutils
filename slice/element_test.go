package slice

import (
	"testing"
)

var slice = []int{0, 1, 2, 3}

func TestElement(t *testing.T) {

	// Test Cases
	testCases := []struct {
		desc     string
		fn       func(slice []int) int
		expected int
	}{
		{
			desc:     "Should return the first element of the slice",
			fn:       First[int],
			expected: 0,
		},
		{
			desc:     "Should return the last element of the slice",
			fn:       Last[int],
			expected: 3,
		},
		{
			desc:     "Should return the nth element of the slice",
			fn:       func(slice []int) int { return Nth(slice, 2) },
			expected: 1,
		},
	}

	// Run test cases
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			actual := tC.fn(slice)
			if actual != tC.expected {
				t.Errorf("want %d, got %d", tC.expected, actual)
			}
		})
	}

}

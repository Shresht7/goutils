package slice

import "testing"

func TestEqual(t *testing.T) {

	// Test Cases
	testCases := []struct {
		desc     string
		a, b     []int
		expected bool
	}{
		{
			desc:     "Should return true if both slices are equal",
			a:        []int{0, 1, 2, 3},
			b:        []int{0, 1, 2, 3},
			expected: true,
		},
		{
			desc:     "Should return false if both slices are not equal",
			a:        []int{0, 1, 2, 3},
			b:        []int{9, 8, 7, 6},
			expected: false,
		},
		{
			desc:     "Should return false if both slices do not have same length",
			a:        []int{0, 0, 0},
			b:        []int{0, 0, 0, 0},
			expected: false,
		},
		{
			desc:     "Should return true if both slices are empty",
			a:        []int{},
			b:        []int{},
			expected: true,
		},
		{
			desc:     "Should return false if the order of elements is different",
			a:        []int{0, 1, 2, 3},
			b:        []int{3, 2, 1, 0},
			expected: false,
		},
	}

	// Run test cases
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if Equal(tC.a, tC.b) != tC.expected {
				t.Error("Failed to check equality of the two slices")
			}
		})
	}

}

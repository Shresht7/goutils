package sliceutils

import "testing"

func TestConcat(t *testing.T) {

	testCases := []struct {
		desc     string
		inputs   [][]int
		expected []int
	}{
		{
			desc: "Concatenate two slices",
			inputs: [][]int{
				{0, 1, 2},
				{3, 4},
			},
			expected: []int{0, 1, 2, 3, 4},
		},
		{
			desc: "Concatenate three slices",
			inputs: [][]int{
				{0, 1, 2},
				{3, 4},
				{5, 6, 7},
			},
			expected: []int{0, 1, 2, 3, 4, 5, 6, 7},
		},
		{
			desc: "Concatenate four slices",
			inputs: [][]int{
				{0, 1, 2},
				{3, 4},
				{5, 6, 7},
				{8, 9},
			},
			expected: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			actual := Concat(tC.inputs...)
			if !Equal(actual, tC.expected) {
				t.Error("Failed to concatenate slices")
			}
		})
	}

}

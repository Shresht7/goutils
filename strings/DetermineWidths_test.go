package helpers

import "testing"

func TestDetermineWidths(t *testing.T) {

	// Test Cases
	testCases := []struct {
		Array  [][]string
		Widths []int
	}{
		{
			Array: [][]string{
				{"Alice", "20"},
				{"Bob", "30"},
			},
			Widths: []int{5, 2},
		},
		{
			Array: [][]string{
				{"Alice", "20"},
				{"Bob", "30"},
				{"Charlie", "40"},
			},
			Widths: []int{7, 2},
		},
		{
			Array: [][]string{
				{"Alice", "20", "521"},
				{"Bob", "30", "3781"},
				{"Charlie", "40", "123"},
			},
			Widths: []int{7, 2, 4},
		},
	}

	// Run Test Cases
	for _, testCase := range testCases {
		widths := DetermineWidths(testCase.Array)
		if !equal(widths, testCase.Widths) {
			t.Errorf("want: %v got: %v", testCase.Widths, widths)
		}

	}
}

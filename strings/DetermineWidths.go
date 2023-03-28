package helpers

// Determines the width of each column in a 2D array of strings
func DetermineWidths(array [][]string) []int {
	// Variable to store the widths of each column
	widths := make([]int, len(array[0]))

	// Iterate over the rows and determine the column widths
	for _, row := range array {
		for i, cell := range row {

			// If the cell is longer than the current width, update the width
			if len(cell) > widths[i] {
				widths[i] = len(cell)
			}

		}
	}

	// Return the widths
	return widths
}

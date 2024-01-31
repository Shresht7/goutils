package helpers

import "strings"

// Determines the minimum and maximum level of indentation for the given lines.
// The lines can be passed as a slice of strings or as a single multi-line string.
func DetermineIndentation(lines ...string) (min, max int) {

	// If only one argument was passed, split it into lines, assuming it to be a multi-line string
	if len(lines) == 1 {
		lines = strings.Split(lines[0], "\n")
	}

	// Iterate over the lines and determine the minimum and maximum indentation levels
	for _, line := range lines {
		current := 0 // number of spaces at the beginning of the line

		// Determine indentation level
		for _, r := range line {
			if IsSpace(r) {
				current++ // While we encounter spaces, increment the indentation level
			} else {
				break // Break on encountering a non-space character
			}
		}

		// Update minimum indentation level
		if current < min || min == 0 {
			min = current
		}

		// Update maximum indentation level
		if current > max {
			max = current
		}
	}

	// Return minimum and maximum indentation levels
	return min, max

}

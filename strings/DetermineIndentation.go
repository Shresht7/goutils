package helpers

// Determines the minimum and maximum level of indentation for the given lines
func DetermineIndentation(lines []string) (min, max int) {

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

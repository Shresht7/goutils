package helpers

// Determines the minimum and maximum level of indentation for the given lines
func DetermineIndentation(lines []string) (min, max int) {
	// Iterate over the lines and determine the minimum and maximum indentation levels
	for _, line := range lines {
		indentationLevel := 0 // number of spaces at the beginning of the line
		// Determine indentation level
		for _, r := range line {
			if IsSpace(r) {
				indentationLevel++
			} else {
				break
			}
		}

		// Update minimum indentation level
		if indentationLevel < min || min == 0 {
			min = indentationLevel
		}

		// Update maximum indentation level
		if indentationLevel > max {
			max = indentationLevel
		}
	}

	// Return minimum and maximum indentation levels
	return min, max
}

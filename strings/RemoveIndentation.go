package helpers

// Removes the given number of spaces from the beginning of each line
func RemoveIndentation(n int, lines []string) []string {
	for i, line := range lines {
		if len(line) > n {
			lines[i] = line[n:]
		}
	}
	return lines
}

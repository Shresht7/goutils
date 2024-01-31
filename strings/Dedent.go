package helpers

import "bytes"

// Dedent removes leading spaces and tabs from each line of the given string up to the specified level.
func Dedent(s string, level int) string {

	var buf bytes.Buffer // Buffer to store the resulting string

	var omitted int = 0   // Number of characters omitted from this line
	var done bool = false // Whether the line has been processed

	// Iterate over string characters ...
	for i := 0; i < len(s); i++ {
		switch s[i] {

		// On encountering a newline, reset the state
		case '\n':
			omitted = 0
			done = false
			_ = buf.WriteByte(s[i]) // Write the character to the buffer

		// Remove leading spaces and tabs up to the given level
		case ' ', '\t':
			if !done { // If the line is not done processing ...

				if omitted < level {
					// If the number of omitted characters is less than the specified level,
					// increment the count of tracked omitted characters and skip adding the character to the buffer
					omitted++
					continue
				} else {
					// ... otherwise, if the number of omitted characters has reached the specified level,
					// mark the line as done
					done = true
				}

			}
			_ = buf.WriteByte(s[i]) // Write the character to the buffer

		// For any other character, mark as done and write it to the buffer directly
		default:
			done = true
			_ = buf.WriteByte(s[i])
		}

	}

	// Return the resulting string
	return buf.String()
}

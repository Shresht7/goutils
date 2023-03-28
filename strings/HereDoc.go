package helpers

import (
	"fmt"
	"strings"
)

// HereDoc returns a here-document representation of the given string.
// See https://en.m.wikipedia.org/wiki/Here_document for more information.
func HereDoc(s string) string {

	// Remove leading newline
	if len(s) > 0 && s[0] == '\n' {
		s = s[1:]
	}

	// Remove trailing whitespace
	s = strings.TrimRight(s, "\n\t\r")

	// Split string into lines
	lines := strings.Split(s, "\n")

	// Determine minimum indentation level
	indentationLevel, _ := DetermineIndentation(lines)

	// Remove indentation from lines
	lines = RemoveIndentation(indentationLevel, lines)

	// Return here-document representation
	return strings.Join(lines, "\n")

}

// HereDocf returns a here-document (https://en.m.wikipedia.org/wiki/Here_document)
// representation of the given string. The can be formatted using
// the given arguments as in fmt.Sprintf.
func HereDocf(format string, args ...any) string {
	return fmt.Sprintf(HereDoc(format), args...)
}

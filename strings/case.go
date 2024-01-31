package strutils

import "strings"

// Capitalize the string (make the first letter uppercase)
func Capitalize(s string) string {
	if len(s) == 0 {
		return s
	}
	return strings.ToUpper(s[:1]) + s[1:]
}

package strutils

import "strings"

// Capitalize the string (make the first letter uppercase)
func Capitalize(s string) string {
	if len(s) == 0 {
		return s
	}
	return strings.ToUpper(s[:1]) + s[1:]
}

// Uncapitalize the string (make the first letter lowercase)
func Uncapitalize(s string) string {
	if len(s) == 0 {
		return s
	}
	return strings.ToLower(s[:1]) + s[1:]
}


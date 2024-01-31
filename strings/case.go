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

// HELPER FUNCTIONS
// ----------------

var wordRegex = regexp.MustCompile(`([A-Z])[^A-Z]*?`)

func splitIntoWords(s string) []string {
	r := wordRegex.ReplaceAllString(s, "_$1")
	r = strings.ReplaceAll(r, "-", "_")
	w := strings.Split(r, "_")
	w = slice.Filter(w, func(s string, _ int) bool { return s != "" })
	return w
}

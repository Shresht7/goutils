package strutils

import (
	"regexp"
	"strings"

	"github.com/Shresht7/goutils/slice"
)

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

type Case int

const (
	CamelCase Case = iota
	SnakeCase
	KebabCase
	TitleCase
)

var CaseRegex = map[Case]*regexp.Regexp{
	CamelCase: regexp.MustCompile(`^[a-z][a-zA-Z0-9]+(?:[A-Z][a-z0-9]+)*$`),
	SnakeCase: regexp.MustCompile(`^[a-z][a-zA-Z0-9]+(?:[_][a-z][a-z0-9]+)*$`),
	KebabCase: regexp.MustCompile(`^[a-z][a-zA-Z0-9]+(?:[-][a-z][a-z0-9]+)*$`),
	TitleCase: regexp.MustCompile(`^[A-Z][a-zA-Z0-9]+(?:[A-Z][a-zA-Z0-9]+)*$`),
}

// IsCase returns true if the string is in the specified case
func IsCase(s string, c Case) bool {
	return CaseRegex[c].MatchString(s)
}

// DetermineCase returns the case of the string or -1 if it is not a known case
func DetermineCase(s string) Case {
	for c, r := range CaseRegex {
		if r.MatchString(s) {
			return c
		}
	}
	return -1
}

// Convert the string to camelCase
func ToCamelCase(s string) string {
	words := splitIntoWords(s)

	if len(words) == 0 {
		return ""
	}

	first, rest := strings.ToLower(words[0]), words[1:]
	for i, w := range rest {
		rest[i] = Capitalize(w)
	}

	return first + strings.Join(rest, "")
}

// Convert the string to snake_case
func ToSnakeCase(s string) string {
	words := splitIntoWords(s)

	if len(words) == 0 {
		return ""
	}

	for i, w := range words {
		words[i] = strings.ToLower(w)
	}

	return strings.Join(words, "_")
}

// Convert the string to kebab-case
func ToKebabCase(s string) string {
	words := splitIntoWords(s)

	if len(words) == 0 {
		return ""
	}

	for i, w := range words {
		words[i] = strings.ToLower(w)
	}

	return strings.Join(words, "-")
}

// Convert the string to TitleCase
func ToTitleCase(s string) string {
	words := splitIntoWords(s)

	if len(words) == 0 {
		return ""
	}

	for i, w := range words {
		words[i] = Capitalize(w)
	}

	return strings.Join(words, "")
}

// Convert the string to the specified case
func ToCase(s string, c Case) string {
	switch c {
	case CamelCase:
		return ToCamelCase(s)
	case SnakeCase:
		return ToSnakeCase(s)
	case KebabCase:
		return ToKebabCase(s)
	case TitleCase:
		return ToTitleCase(s)
	default:
		return s
	}
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

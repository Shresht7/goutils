package helpers

import "strings"

//	Apply padding around the string
func Pad(s, char string, count int) string {
	return strings.Repeat(char, count) + s + strings.Repeat(char, count)
}

//	Apply padding to the left of the string
func PadLeft(s, char string, count int) string {
	return strings.Repeat(char, count) + s
}

//	Apply padding to the right of the string
func PadRight(s, char string, count int) string {
	return s + strings.Repeat(char, count)
}

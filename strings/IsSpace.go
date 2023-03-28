package helpers

// Checks if the given rune is a space character.
func IsSpace(r rune) bool {
	return r == ' ' || r == '\t'
}

package helpers

import (
	"math"
	"strings"
)

type AlignOptions struct {
	Mode  string
	Split string
	Pad   string
	Width int
}

// AssignValues assigns default values to the options
func (opts *AlignOptions) AssignValues() {

	// Split on newlines by default
	if opts.Split == "" {
		opts.Split = "\n"
	}

	// Pad with spaces by default
	if opts.Pad == "" {
		opts.Pad = " "
	}

	// Default to left alignment if no mode is specified
	switch strings.ToLower(opts.Mode) {
	case "left", "right", "center":
	default:
		opts.Mode = "center"
	}

}

type StringAndWidthTuple struct {
	str   string
	width int
}

func Align(str string, opts *AlignOptions) string {

	opts.AssignValues()

	// Maximum width of the string to display
	maxWidth := opts.Width

	// Split the string into lines
	lines := strings.Split(str, opts.Split)

	tupleStr := []StringAndWidthTuple{}
	// Find the maximum width of the string
	for _, line := range lines {
		// ! This might not be entirely accurate as unicode characters can be more than one byte in length
		// TODO: Find a better way to calculate the width of a string in bytes
		width := len(line)

		// Update the maximum width if the current line is longer
		if width > maxWidth {
			maxWidth = len(line)
		}

		// Append the string and its width to the slice
		tupleStr = append(tupleStr, StringAndWidthTuple{line, width})
	}

	// Create a new slice to store the aligned strings
	alignedLines := []string{}
	for _, tuple := range tupleStr {
		// Calculate the amount of padding needed
		padding := maxWidth - tuple.width
		// Apply the padding
		alignedLines = append(alignedLines, applyPadding(opts.Mode, tuple.str, opts.Pad, padding))
	}

	// Join the aligned lines and return the result
	return strings.Join(alignedLines, opts.Split)

}

// applyPadding applies the padding to the string based on the alignment options
func applyPadding(mode, str, pad string, space int) string {
	switch strings.ToLower(mode) {

	case "left":
		// Left align the string by adding padding to the right
		return str + strings.Repeat(pad, space)

	case "right":
		// Right align the string by adding padding to the left
		return strings.Repeat(pad, space) + str

	case "center":
		// Center align the string by adding padding to both sides

		// Calculate the amount of padding to add to each side
		half := int(math.Floor(float64(space) / 2.0))

		// If the space is even, we can split the padding evenly
		if space%2 == 0 {
			// Return the string with the padding on both sides
			return Pad(str, pad, half)
		} else {
			// If the space is odd, we need to add one more padding to the right side
			return strings.Repeat(pad, half) + str + strings.Repeat(pad, half+1)
		}

	default:
		// Return the string by default
		return str

	}
}

// Center aligns the string
func AlignCenter(str string) string {
	return Align(str, &AlignOptions{Mode: "center"})
}

// Left aligns the string
func AlignLeft(str string) string {
	return Align(str, &AlignOptions{Mode: "left"})
}

// Right aligns the string
func AlignRight(str string) string {
	return Align(str, &AlignOptions{Mode: "right"})
}

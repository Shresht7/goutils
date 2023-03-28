package helpers

import (
	"fmt"
	"testing"
)

func TestHereDoc(t *testing.T) {

	// Setup test cases
	cases := []struct {
		in, want string
	}{
		{"", ""},
		{"Hello, World!", "Hello, World!"},
		{`Hello, World!`, `Hello, World!`},
		{
			`
			Hello, World!
			`,
			"Hello, World!",
		},
		{
			`
			One
			Two
			Three
			`,
			"One\nTwo\nThree",
		},
		{
			`
			Hello,
				World!
			`,
			"Hello,\n\tWorld!",
		},
	}

	// Run test cases
	for _, c := range cases {
		got := HereDoc(c.in)
		if got != c.want {
			t.Errorf("Doc(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestHereDocf(t *testing.T) {

	// Setup test cases
	cases := []struct {
		in, want string
	}{
		{"", ""},
		{HereDocf("%s", "DO IT"), "DO IT"},
		{HereDocf("%d", 30), "30"},
		{
			HereDocf("%s - %d", "ABC", 30),
			"ABC - 30",
		},
		{
			HereDocf(`
			%s - %d
			`, "ABC", 30),
			"ABC - 30",
		},
	}

	// Run test cases
	for _, c := range cases {
		if c.in != c.want {
			t.Errorf("Docf(%q) == %q, want %q", c.in, c.in, c.want)
		}
	}

}

// Simple example
func ExampleHereDoc_simple() {
	// Print a here-document
	fmt.Println(HereDoc(`
		Hello, World!
	`))
	// Output:
	// Hello, World!
}

// Multiline example
func ExampleHereDoc_multiline() {
	// Print a here-document
	fmt.Println(HereDoc(`
		Hello,
			World!
	`))
	// Output:
	// Hello,
	// 	World!
}

// Simple example
func ExampleHereDocf_simple() {
	// Print a here-document
	fmt.Println(HereDocf(`
		%s
	`, "Hello, World!"))
	// Output:
	// Hello, World!
}

// Multiline example
func ExampleHereDocf_multiline() {
	// Print a here-document
	fmt.Println(HereDocf(`
		%s,
			%s!
	`, "Hello", "World"))
	// Output:
	// Hello,
	// 	World!
}

package helpers

import (
	"testing"
)

func TestRemoveIndentation(t *testing.T) {

	// Test Cases
	cases := []struct {
		in, want []string
		n        int
	}{
		{
			in:   []string{},
			want: []string{},
			n:    0,
		},
		{
			in:   []string{"Hello, World!"},
			want: []string{"Hello, World!"},
			n:    0,
		},
		{
			in:   []string{" Hello, World!"},
			want: []string{"Hello, World!"},
			n:    1,
		},
		{
			in:   []string{"  Hello, World!"},
			want: []string{"Hello, World!"},
			n:    2,
		},
		{
			in:   []string{"  Hello, World!"},
			want: []string{" Hello, World!"},
			n:    1,
		},
		{
			in: []string{
				"\tHello, World!",
				"\t\tGO!",
			},
			want: []string{
				"Hello, World!",
				"\tGO!",
			},
			n: 1,
		},
		{
			in: []string{
				"\t\tHello, World!",
				"\tGO!",
			},
			want: []string{
				"\tHello, World!",
				"GO!",
			},
			n: 1,
		},
	}

	// Run test cases
	for _, c := range cases {
		got := RemoveIndentation(c.n, c.in)
		if !equal(got, c.want) {
			t.Errorf("removeIndentation(%q) == %q, want %q", c.in, got, c.want)
		}
	}

}

// equal compares two slices of comparable types and returns true if they are equal
func equal[T comparable](a, b []T) bool {
	// Return false if slices have different length
	if len(a) != len(b) {
		return false
	}
	// Return false if slices have different elements
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	// Return true if slices are equal
	return true
}

package helpers

import (
	"testing"
)

func TestDetermineIndentation(t *testing.T) {

	// Test Cases
	cases := []struct {
		lines []string
		level [2]int
	}{
		{
			lines: []string{},
			level: [2]int{0, 0},
		},
		{
			lines: []string{"Test String..."},
			level: [2]int{0, 0},
		},
		{
			lines: []string{" Test String..."},
			level: [2]int{1, 1},
		},
		{
			lines: []string{"  Test String..."},
			level: [2]int{2, 2},
		},
		{
			lines: []string{"\tTest String...", "\t\tGO!"},
			level: [2]int{1, 2},
		},
		{
			lines: []string{"\t\tTest String...", "\tGO!"},
			level: [2]int{1, 2},
		},
		{
			lines: []string{"\t\tTest String...", "GO!"},
			level: [2]int{0, 2},
		},
		{
			lines: []string{"\t\t\tTest String...", "   GO!"},
			level: [2]int{3, 3},
		},
	}

	// Run test cases
	for _, c := range cases {
		min, max := DetermineIndentation(c.lines)
		if min != c.level[0] || max != c.level[1] {
			t.Errorf("DetermineIndentation(%q) == (%d, %d), want (%d, %d)", c.lines, min, max, c.level[0], c.level[1])
		}
	}

}

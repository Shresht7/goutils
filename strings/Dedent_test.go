package helpers

import (
	"strings"
	"testing"
)

func TestDedent(t *testing.T) {

	testCases := []struct {
		input string
		level int
		want  string
	}{
		{
			input: "    Hello, world!\n    This is a test string.\n    This is another test string.\n",
			level: 4,
			want:  "Hello, world!\nThis is a test string.\nThis is another test string.\n",
		},
		{
			input: "    Hello, world!\n    This is a test string.\n    This is another test string.\n",
			level: 2,
			want:  "  Hello, world!\n  This is a test string.\n  This is another test string.\n",
		},
		{
			input: "    Hello, world!\n    This is a test string.\n    This is another test string.\n",
			level: 0,
			want:  "    Hello, world!\n    This is a test string.\n    This is another test string.\n",
		},
		{
			input: "    Hello, world!\n    This is a test string.\n    This is another test string.\n",
			level: 6,
			want:  "Hello, world!\nThis is a test string.\nThis is another test string.\n",
		},
	}

	for _, testCase := range testCases {
		got := Dedent(testCase.input, testCase.level)
		if strings.Compare(got, testCase.want) != 0 {
			t.Errorf("Dedent(%q, %d) = %q, want %q", testCase.input, testCase.level, got, testCase.want)
		}
	}

}

package helpers

import (
	"strings"
	"testing"
)

func TestAlign(t *testing.T) {

	testCases := []struct {
		desc     string
		actual   string
		expected string
	}{
		{
			desc: "should align to the center by default",
			actual: Align(join(
				"text-align",
				"center",
			), &AlignOptions{}),
			expected: join(
				"text-align",
				"  center  ",
			),
		},
		{
			desc: "should handle an odd difference",
			actual: Align(join(
				"txt-align",
				"center",
			), &AlignOptions{}),
			expected: join(
				"txt-align",
				" center  ",
			),
		},
		{
			desc: "should align to the left",
			actual: Align(join(
				"text-align",
				"left"), &AlignOptions{Mode: "Left"}),
			expected: join(
				"text-align",
				"left      ",
			),
		},
		{
			desc: "should align to the right",
			actual: Align(join(
				"text-align",
				"right"), &AlignOptions{Mode: "Right"}),
			expected: join(
				"text-align",
				"     right",
			),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.desc, func(t *testing.T) {
			if testCase.actual != testCase.expected {
				t.Errorf("want:\n%s\n\ngot:\n%s", testCase.expected, testCase.actual)
			}
		})
	}

}

func join(str ...string) string {
	return strings.Join(str, "\n")
}

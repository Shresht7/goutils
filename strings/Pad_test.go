package helpers

import (
	"fmt"
	"testing"
)

func TestPad(t *testing.T) {

	testCases := []struct {
		desc     string
		actual   string
		expected string
	}{
		{
			desc:     "should apply padding around the string",
			actual:   Pad("Go", " ", 1),
			expected: " Go ",
		},
		{
			desc:     "should apply padding around the string using the given character",
			actual:   Pad("Go", "-", 1),
			expected: "-Go-",
		},
		{
			desc:     "should apply the correct amount of padding around the string",
			actual:   Pad("Go", "-", 3),
			expected: "---Go---",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.desc, func(t *testing.T) {
			if testCase.actual != testCase.expected {
				t.Errorf("want:\t%s\ngot:\t%s\n", testCase.expected, testCase.actual)
			}
		})
	}

}

func ExamplePad() {
	str := "Golang"
	newStr := Pad(str, "|", 2)
	fmt.Println(newStr)
	//	Output: ||Golang||
}

func TestPadLeft(t *testing.T) {

	testCases := []struct {
		desc     string
		actual   string
		expected string
	}{
		{
			desc:     "should apply padding to the left of the string",
			actual:   PadLeft("Go", " ", 1),
			expected: " Go",
		},
		{
			desc:     "should apply padding to the left of the string using the given character",
			actual:   PadLeft("Go", "-", 1),
			expected: "-Go",
		},
		{
			desc:     "should apply the correct amount of padding to the left of the string",
			actual:   PadLeft("Go", "-", 3),
			expected: "---Go",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.desc, func(t *testing.T) {
			if testCase.actual != testCase.expected {
				t.Errorf("want:\t%s\ngot:\t%s\n", testCase.expected, testCase.actual)
			}
		})
	}

}

func ExamplePadLeft() {
	str := "Golang"
	newStr := PadLeft(str, ">>", 1)
	fmt.Println(newStr)
	//	Output: >>Golang
}

func TestPadRight(t *testing.T) {

	testCases := []struct {
		desc     string
		actual   string
		expected string
	}{
		{
			desc:     "should apply padding to the right of the string",
			actual:   PadRight("Go", " ", 1),
			expected: "Go ",
		},
		{
			desc:     "should apply padding to the right of the string using the given character",
			actual:   PadRight("Go", "-", 1),
			expected: "Go-",
		},
		{
			desc:     "should apply the correct amount of padding to the right of the string",
			actual:   PadRight("Go", "-", 3),
			expected: "Go---",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.desc, func(t *testing.T) {
			if testCase.actual != testCase.expected {
				t.Errorf("want:\t%s\ngot:\t%s\n", testCase.expected, testCase.actual)
			}
		})
	}

}

func ExamplePadRight() {
	str := "Golang"
	newStr := PadRight(str, ">", 3)
	fmt.Println(newStr)
	//	Output: Golang>>>
}

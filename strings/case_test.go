package strutils

import "testing"

func TestCapitalize(t *testing.T) {

	testCases := []struct {
		input    string
		expected string
	}{
		{"", ""},
		{"a", "A"},
		{"A", "A"},
		{"abc", "Abc"},
		{"Abc", "Abc"},
		{"ABC", "ABC"},
		{"123", "123"},
		{"123abc", "123abc"},
		{"123Abc", "123Abc"},
		{"123ABC", "123ABC"},
		{"abc123", "Abc123"},
		{"Abc123", "Abc123"},
	}

	for _, testCase := range testCases {
		actual := Capitalize(testCase.input)
		if actual != testCase.expected {
			t.Errorf("Capitalize(%q) expected %q, got %q", testCase.input, testCase.expected, actual)
		}
	}

}

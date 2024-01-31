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

func TestUncapitalize(t *testing.T) {

	testCases := []struct {
		input    string
		expected string
	}{
		{"", ""},
		{"a", "a"},
		{"A", "a"},
		{"abc", "abc"},
		{"Abc", "abc"},
		{"ABC", "aBC"},
		{"123", "123"},
		{"123abc", "123abc"},
		{"123Abc", "123Abc"},
		{"123ABC", "123ABC"},
		{"abc123", "abc123"},
		{"Abc123", "abc123"},
		{"A sentence", "a sentence"},
		{"A Sentence", "a Sentence"},
	}

	for _, testCase := range testCases {
		actual := Uncapitalize(testCase.input)
		if actual != testCase.expected {
			t.Errorf("Uncapitalize(%q) expected %q, got %q", testCase.input, testCase.expected, actual)
		}
	}

}

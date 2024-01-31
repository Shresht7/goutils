package strutils

import (
	"testing"

	"github.com/Shresht7/goutils/internal/test"
)

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
		{"a sentence", "A sentence"},
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

func TestSplitIntoWords(t *testing.T) {

	testCases := []struct {
		desc     string
		actual   []string
		expected []string
	}{
		{"empty string", splitIntoWords(""), []string{}},
		{"single word", splitIntoWords("abc"), []string{"abc"}},
		{"camelCase", splitIntoWords("camelCase"), []string{"camel", "Case"}},
		{"camelCaseString", splitIntoWords("camelCaseString"), []string{"camel", "Case", "String"}},
		{"snake_case", splitIntoWords("snake_case"), []string{"snake", "case"}},
		{"snake_case_string", splitIntoWords("snake_case_string"), []string{"snake", "case", "string"}},
		{"kebab-case", splitIntoWords("kebab-case"), []string{"kebab", "case"}},
		{"kebab-case-string", splitIntoWords("kebab-case-string"), []string{"kebab", "case", "string"}},
		{"TitleCase", splitIntoWords("TitleCase"), []string{"Title", "Case"}},
		{"TitleCaseString", splitIntoWords("TitleCaseString"), []string{"Title", "Case", "String"}},
	}

	for _, testCase := range testCases {
		if !test.Equal(testCase.actual, testCase.expected) {
			t.Errorf("splitIntoWords(%q) expected %q, got %q", testCase.desc, testCase.expected, testCase.actual)
		}
	}
}

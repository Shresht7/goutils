package helpers

import (
	"testing"
)

func TestWrap(t *testing.T) {
	// Test Cases
	testCases := []struct {
		s, str, expected string
	}{
		{"", "", ""},
		{"", "a", "a"},
		{"*", "", "**"},
		{"*", "a", "*a*"},
		{"*", "ab", "*ab*"},
		{"*", "abc", "*abc*"},
		{"==", "abc", "==abc=="},
	}

	// Run Test Cases
	for _, tc := range testCases {
		actual := Wrap(tc.s, tc.str)
		if actual != tc.expected {
			t.Errorf("Wrap(%v, %v): expected %v, actual %v", tc.s, tc.str, tc.expected, actual)
		}
	}

}

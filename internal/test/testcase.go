package test

import (
	"testing"
)

// ---------
// TEST CASE
// ---------

type TestCase[T any] struct {
	Desc     string
	Fn       func() T
	Expected T
	Fail     func(actual, expected T) bool
}

// --------------
// RUN TEST CASES
// --------------

func RunTestCases[T any](testCases *[]TestCase[T], t *testing.T) {

	for _, testCase := range *testCases {
		t.Run(testCase.Desc, func(t *testing.T) {
			actual := testCase.Fn()
			if testCase.Fail(actual, testCase.Expected) {
				t.Errorf("Failed: %v\nwant:\t%v\ngot:\t%v", testCase.Desc, testCase.Expected, actual)
			}
		})
	}

}

// -------
// HELPERS
// -------

func Inequality[T comparable](actual, expected T) bool {
	return actual != expected
}

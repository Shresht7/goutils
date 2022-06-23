package test

import (
	"testing"

	"golang.org/x/exp/constraints"
)

type TestCase[T, E any] struct {
	Desc     string
	Fn       func() E
	Expected E
	Fail     func(actual, expected E) bool
}

func RunTestCases[T, E any](testCases *[]TestCase[T, E], t *testing.T) {

	for _, testCase := range *testCases {
		t.Run(testCase.Desc, func(t *testing.T) {
			actual := testCase.Fn()
			if testCase.Fail(actual, testCase.Expected) {
				t.Errorf("Failed: %v\nwant:\t%v\ngot:\t%v", testCase.Desc, testCase.Expected, actual)
			}
		})
	}

}

func Inequality[T constraints.Ordered | bool](actual, expected T) bool {
	return actual != expected
}

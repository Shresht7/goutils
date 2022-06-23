package sliceutils

import (
	"testing"

	. "github.com/Shresht7/sliceutils/internal"
	"golang.org/x/exp/constraints"
)

//	TODO: Improve these tests

var sliceA []int = []int{0, 1, 2, 3, 4, 5, 6}
var sliceB []string = []string{"A", "B", "C"}

func TestForEach(t *testing.T) {

	integerSliceTestCases := &[]TestCase[int, int]{
		{
			Desc: "Should iterate over each entry in the slice",
			Fn: func() int {
				actual := 0
				ForEach(sliceA, func(_, _ int) { actual++ })
				return actual
			},
			Expected: len(sliceA),
			Fail:     Inequality[int],
		},
		{
			Desc: "Should not iterate over an empty slice",
			Fn: func() int {
				actual := 0
				ForEach(sliceA, func(_, _ int) { actual++ })
				return actual
			},
			Expected: len(sliceA),
			Fail:     Inequality[int],
		},
	}

	stringSliceTestCases := &[]TestCase[string, string]{
		{
			Desc: "Should iterate over each entry in a string slice",
			Fn: func() string {
				actual := ""
				ForEach(sliceB, func(v string, _ int) { actual += v + "->" })
				return actual
			},
			Expected: "A->B->C->",
			Fail:     Inequality[string],
		},
	}

	RunTestCases(integerSliceTestCases, t)
	RunTestCases(stringSliceTestCases, t)

}

func TestFilter(t *testing.T) {

	testCases := &[]TestCase[int, []int]{
		{
			Desc: "Should return nothing",
			Fn: func() []int {
				return Filter(sliceA, func(v, _ int) bool { return false })
			},
			Expected: []int{},
			Fail:     SliceInequality[int],
		},
		{
			Desc: "Should return everything",
			Fn: func() []int {
				return Filter(sliceA, func(v, _ int) bool { return true })
			},
			Expected: sliceA,
			Fail:     SliceInequality[int],
		},
		{
			Desc: "Should return even numbers",
			Fn: func() []int {
				return Filter(sliceA, func(v, _ int) bool { return v%2 == 0 })
			},
			Expected: []int{0, 2, 4, 6},
			Fail:     SliceInequality[int],
		},
		{
			Desc: "Should return odd numbers",
			Fn: func() []int {
				return Filter(sliceA, func(v, _ int) bool { return v%2 != 0 })
			},
			Expected: []int{1, 3, 5},
			Fail:     SliceInequality[int],
		},
		{
			Desc: "Should return numbers greater than 2",
			Fn: func() []int {
				return Filter(sliceA, func(v, _ int) bool { return v > 2 })
			},
			Expected: []int{3, 4, 5, 6},
			Fail:     SliceInequality[int],
		},
		{
			Desc: "Should return numbers equal to 5",
			Fn: func() []int {
				return Filter(sliceA, func(v, _ int) bool { return v == 5 })
			},
			Expected: []int{5},
			Fail:     SliceInequality[int],
		},
	}

	RunTestCases(testCases, t)
}

func TestMap(t *testing.T) {

	testCases := &[]TestCase[int, []int]{
		{
			Desc: "Should return a slice with each value doubled",
			Fn: func() []int {
				return Map(sliceA, func(v, _ int) int { return 2 * v })
			},
			Expected: []int{0, 2, 4, 6, 8, 10, 12},
			Fail:     SliceInequality[int],
		},
		{
			Desc: "Should return a slice with 0 for even numbers and 1 for odd numbers",
			Fn: func() []int {
				return Map(sliceA, func(v, _ int) int { return v % 2 })
			},
			Expected: []int{0, 1, 0, 1, 0, 1, 0},
			Fail:     SliceInequality[int],
		},
		{
			Desc: "Should return an empty slice for an empty slice",
			Fn: func() []int {
				return Map([]int{}, func(v, _ int) int { return 1 })
			},
			Expected: []int{},
			Fail:     SliceInequality[int],
		},
	}

	RunTestCases(testCases, t)

}

func TestReduce(t *testing.T) {

	expectedSum := 0
	for _, v := range sliceA {
		expectedSum += v
	}

	integerTestCases := &[]TestCase[int, int]{
		{
			Desc: "Should return the sum of all elements in the slice",
			Fn: func() int {
				return Reduce(sliceA, func(a, c, _ int, _ []int) int { return a + c }, 0)
			},
			Expected: expectedSum,
			Fail:     Inequality[int],
		},
	}

	expectedConcatenation := ""
	for _, v := range sliceB {
		expectedConcatenation += v
	}

	stringTestCases := &[]TestCase[string, string]{
		{
			Desc: "Should concatenate all the elements in the slice",
			Fn: func() string {
				return Reduce(sliceB, func(a, c string, _ int, _ []string) string { return a + c }, "")
			},
			Expected: expectedConcatenation,
			Fail:     Inequality[string],
		},
	}

	RunTestCases(integerTestCases, t)
	RunTestCases(stringTestCases, t)

}

func TestEvery(t *testing.T) {

	slice := []int{2, 4, 6, 8, 10}

	testCases := &[]TestCase[int, bool]{
		{
			Desc:     "Every element should be divisible by 2",
			Fn:       func() bool { return Every(slice, func(v, _ int) bool { return v%2 == 0 }) },
			Expected: true,
			Fail:     Inequality[bool],
		},
		{
			Desc:     "Every element should not be divisible by 3",
			Fn:       func() bool { return Every(slice, func(v, _ int) bool { return v%3 == 0 }) },
			Expected: false,
			Fail:     Inequality[bool],
		},
	}

	RunTestCases(testCases, t)

}

func TestSome(t *testing.T) {

	slice := []int{1, 3, 5, 7, 9, 10}

	testCases := &[]TestCase[int, bool]{
		{
			Desc:     "Every element should not be divisible by 2",
			Fn:       func() bool { return Every(slice, func(v, _ int) bool { return v%2 == 0 }) },
			Expected: false,
			Fail:     Inequality[bool],
		},
		{
			Desc:     "Every element should be smaller than or equal to 10",
			Fn:       func() bool { return Every(slice, func(v, _ int) bool { return v <= 10 }) },
			Expected: true,
			Fail:     Inequality[bool],
		},
	}

	RunTestCases(testCases, t)

}
func TestFind(t *testing.T) {

	okTestCases := &[]TestCase[int, bool]{
		{
			Desc: "Find 3 in the slice",
			Fn: func() bool {
				ok, _, _ := Find(sliceA, func(v, _ int) bool { return v == 3 })
				return ok
			},
			Expected: true,
			Fail:     Inequality[bool],
		},
		{
			Desc: "Do not find 11 in the slice",
			Fn: func() bool {
				ok, _, _ := Find(sliceA, func(v, _ int) bool { return v == 11 })
				return ok
			},
			Expected: false,
			Fail:     Inequality[bool],
		},
	}

	valueTestCases := &[]TestCase[int, int]{
		{
			Desc: "Find the first multiple of 4",
			Fn: func() int {
				_, v, _ := Find(sliceA, func(v, _ int) bool { return v%4 == 0 && v != 0 })
				return v
			},
			Expected: 4,
			Fail:     Inequality[int],
		},
		{
			Desc: "Find the second multiple of 3",
			Fn: func() int {
				_, v, _ := Find(sliceA, func(v, _ int) bool { return v%3 == 0 && v != 3 && v != 0 })
				return v
			},
			Expected: 6,
			Fail:     Inequality[int],
		},
	}

	RunTestCases(okTestCases, t)
	RunTestCases(valueTestCases, t)

}

// ----------------
// HELPER FUNCTIONS
// ----------------

//	Helper function to check if the two given slices are not equal
func SliceInequality[T constraints.Ordered](actual, expected []T) bool {
	return !Equal(actual, expected)
}

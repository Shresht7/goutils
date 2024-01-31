package slice

import (
	"fmt"
	"testing"

	"github.com/Shresht7/goutils/internal/test"
)

// ----------------
// HELPER FUNCTIONS
// ----------------

// SliceInequality compares two slices and returns true if they are not equal
func SliceInequality[T comparable](actual, expected []T) bool {
	return !Equal(actual, expected)
}

// =====
// TESTS
// =====

// * FOR EACH *//

func TestForEach(t *testing.T) {

	// Test Data
	var sliceA []int = []int{0, 1, 2, 3, 4, 5, 6}
	var sliceB []string = []string{"A", "B", "C"}

	// Test Cases
	testCases1 := []test.TestCase[int]{
		{
			Desc: "Should iterate over each entry in the slice",
			Actual: func() int {
				actual := 0
				ForEach(sliceA, func(_, _ int) { actual++ })
				return actual
			},
			Expected: len(sliceA),
			Fail:     test.Inequality[int],
		},
		{
			Desc: "Should not iterate over an empty slice",
			Actual: func() int {
				actual := 0
				ForEach(sliceA, func(_, _ int) { actual++ })
				return actual
			},
			Expected: len(sliceA),
			Fail:     test.Inequality[int],
		},
	}
	testCases2 := []test.TestCase[string]{
		{
			Desc: "Should iterate over each entry in a string slice",
			Actual: func() string {
				actual := ""
				ForEach(sliceB, func(v string, _ int) { actual += v + "->" })
				return actual
			},
			Expected: "A->B->C->",
			Fail:     test.Inequality[string],
		},
	}

	// Run Tests
	for _, tc := range testCases1 {
		t.Run(tc.Desc, func(t *testing.T) {
			actual := tc.Actual()
			if tc.Fail(actual, tc.Expected) {
				t.Errorf("%s: Expected: %v, Actual: %v", tc.Desc, tc.Expected, actual)
			}
		})
	}
	for _, tc := range testCases2 {
		t.Run(tc.Desc, func(t *testing.T) {
			actual := tc.Actual()
			if tc.Fail(actual, tc.Expected) {
				t.Errorf("%s: Expected: %v, Actual: %v", tc.Desc, tc.Expected, actual)
			}
		})
	}

}

func ExampleForEach() {
	slice := []int{1, 2, 3}

	ForEach(slice, func(v, i int) {
		fmt.Printf("Value: %v, Index: %v\n", v, i)
	})

	// Output:
	// Value: 1, Index: 0
	// Value: 2, Index: 1
	// Value: 3, Index: 2
}

func ExampleForEach_method() {
	slice := New([]int{1, 2, 3})

	slice.ForEach(func(v, i int) {
		fmt.Printf("Value: %v, Index: %v\n", v, i)
	})

	// Output:
	// Value: 1, Index: 0
	// Value: 2, Index: 1
	// Value: 3, Index: 2
}

// * FILTER *//

func TestFilter(t *testing.T) {

	// Test Data
	var sliceA []int = []int{0, 1, 2, 3, 4, 5, 6}

	// Test Cases
	testCases := []test.TestCase[[]int]{
		{
			Desc: "Should return nothing",
			Actual: func() []int {
				return Filter(sliceA, func(v, _ int) bool { return false })
			},
			Expected: []int{},
			Fail:     SliceInequality[int],
		},
		{
			Desc: "Should return everything",
			Actual: func() []int {
				return Filter(sliceA, func(v, _ int) bool { return true })
			},
			Expected: sliceA,
			Fail:     SliceInequality[int],
		},
		{
			Desc: "Should return even numbers",
			Actual: func() []int {
				return Filter(sliceA, func(v, _ int) bool { return v%2 == 0 })
			},
			Expected: []int{0, 2, 4, 6},
			Fail:     SliceInequality[int],
		},
		{
			Desc: "Should return odd numbers",
			Actual: func() []int {
				return Filter(sliceA, func(v, _ int) bool { return v%2 != 0 })
			},
			Expected: []int{1, 3, 5},
			Fail:     SliceInequality[int],
		},
		{
			Desc: "Should return numbers greater than 2",
			Actual: func() []int {
				return Filter(sliceA, func(v, _ int) bool { return v > 2 })
			},
			Expected: []int{3, 4, 5, 6},
			Fail:     SliceInequality[int],
		},
		{
			Desc: "Should return numbers equal to 5",
			Actual: func() []int {
				return Filter(sliceA, func(v, _ int) bool { return v == 5 })
			},
			Expected: []int{5},
			Fail:     SliceInequality[int],
		},
	}

	// Run Tests
	for _, tc := range testCases {
		t.Run(tc.Desc, func(t *testing.T) {
			actual := tc.Actual()
			if tc.Fail(actual, tc.Expected) {
				t.Errorf("%s: Expected: %v, Actual: %v", tc.Desc, tc.Expected, actual)
			}
		})
	}

}

func ExampleFilter() {
	slice := []int{1, 2, 3, 4, 5, 6}

	filtered := Filter(slice, func(v, _ int) bool {
		return v%2 == 0
	})

	fmt.Println(filtered)

	// Output:
	// [2 4 6]
}

func ExampleFilter_method() {
	slice := New([]int{1, 2, 3, 4, 5, 6})

	filtered := slice.Filter(func(v, _ int) bool {
		return v%2 == 0
	})

	fmt.Println(filtered)

	// Output:
	// [2 4 6]
}

// * MAP *//

func TestMap(t *testing.T) {

	// Test Data
	var sliceA []int = []int{0, 1, 2, 3, 4, 5, 6}

	testCases := []test.TestCase[[]int]{
		{
			Desc: "Should return a slice with each value doubled",
			Actual: func() []int {
				return Map(sliceA, func(v, _ int) int { return 2 * v })
			},
			Expected: []int{0, 2, 4, 6, 8, 10, 12},
			Fail:     SliceInequality[int],
		},
		{
			Desc: "Should return a slice with 0 for even numbers and 1 for odd numbers",
			Actual: func() []int {
				return Map(sliceA, func(v, _ int) int { return v % 2 })
			},
			Expected: []int{0, 1, 0, 1, 0, 1, 0},
			Fail:     SliceInequality[int],
		},
		{
			Desc: "Should return an empty slice for an empty slice",
			Actual: func() []int {
				return Map([]int{}, func(v, _ int) int { return 1 })
			},
			Expected: []int{},
			Fail:     SliceInequality[int],
		},
	}

	// Run Tests
	for _, tc := range testCases {
		t.Run(tc.Desc, func(t *testing.T) {
			actual := tc.Actual()
			if tc.Fail(actual, tc.Expected) {
				t.Errorf("%s: Expected: %v, Actual: %v", tc.Desc, tc.Expected, actual)
			}
		})
	}

}

func ExampleMap() {
	slice := []int{1, 2, 3, 4, 5, 6}

	mapped := Map(slice, func(v, _ int) int {
		return 2 * v
	})

	fmt.Println(mapped)

	// Output:
	// [2 4 6 8 10 12]
}

func ExampleMap_method() {
	slice := New([]int{1, 2, 3, 4, 5, 6})

	mapped := slice.Map(func(v, _ int) int {
		return 2 * v
	})

	fmt.Println(mapped)

	// Output:
	// [2 4 6 8 10 12]
}

// * REDUCE *//

func TestReduce(t *testing.T) {

	// Test Data
	var sliceA []int = []int{0, 1, 2, 3, 4, 5, 6}
	var sliceB []string = []string{"A", "B", "C"}

	expectedSum := 0
	for _, v := range sliceA {
		expectedSum += v
	}

	testCases := []test.TestCase[int]{
		{
			Desc: "Should return the sum of all elements in the slice",
			Actual: func() int {
				return Reduce(sliceA, func(a, c, _ int, _ []int) int { return a + c }, 0)
			},
			Expected: expectedSum,
			Fail:     test.Inequality[int],
		},
	}

	// Run Tests
	for _, tc := range testCases {
		t.Run(tc.Desc, func(t *testing.T) {
			actual := tc.Actual()
			if tc.Fail(actual, tc.Expected) {
				t.Errorf("%s: Expected: %v, Actual: %v", tc.Desc, tc.Expected, actual)
			}
		})
	}

	expectedConcatenation := ""
	for _, v := range sliceB {
		expectedConcatenation += v
	}

	testCases2 := []test.TestCase[string]{
		{
			Desc: "Should concatenate all the elements in the slice",
			Actual: func() string {
				return Reduce(sliceB, func(a, c string, _ int, _ []string) string { return a + c }, "")
			},
			Expected: expectedConcatenation,
			Fail:     test.Inequality[string],
		},
	}

	// Run Tests
	for _, tc := range testCases2 {
		t.Run(tc.Desc, func(t *testing.T) {
			actual := tc.Actual()
			if tc.Fail(actual, tc.Expected) {
				t.Errorf("%s: Expected: %v, Actual: %v", tc.Desc, tc.Expected, actual)
			}
		})
	}

}

func ExampleReduce() {
	slice := []int{1, 2, 3, 4, 5, 6}

	sum := Reduce(slice, func(a, c, _ int, _ []int) int {
		return a + c
	}, 0)

	fmt.Println(sum)

	// Output:
	// 21
}

func ExampleReduce_method() {
	slice := New([]int{1, 2, 3, 4, 5, 6})

	sum := slice.Reduce(func(a, c, _ int, _ []int) int {
		return a + c
	}, 0)

	fmt.Println(sum)

	// Output:
	// 21
}

// * EVERY *//

func TestEvery(t *testing.T) {

	slice := []int{2, 4, 6, 8, 10}

	testCases := []test.TestCase[bool]{
		{
			Desc:     "Every element should be divisible by 2",
			Actual:   func() bool { return Every(slice, func(v, _ int) bool { return v%2 == 0 }) },
			Expected: true,
			Fail:     test.Inequality[bool],
		},
		{
			Desc:     "Every element should not be divisible by 3",
			Actual:   func() bool { return Every(slice, func(v, _ int) bool { return v%3 == 0 }) },
			Expected: false,
			Fail:     test.Inequality[bool],
		},
	}

	// Run Tests
	for _, tc := range testCases {
		t.Run(tc.Desc, func(t *testing.T) {
			actual := tc.Actual()
			if tc.Fail(actual, tc.Expected) {
				t.Errorf("%s: Expected: %v, Actual: %v", tc.Desc, tc.Expected, actual)
			}
		})
	}

}

func ExampleEvery() {
	slice := []int{2, 4, 6, 8, 10}

	allEven := Every(slice, func(v, _ int) bool {
		return v%2 == 0
	})

	fmt.Println(allEven)

	// Output:
	// true
}

func ExampleEvery_method() {
	slice := New([]int{2, 4, 6, 8, 10})

	allEven := slice.Every(func(v, _ int) bool {
		return v%2 == 0
	})

	fmt.Println(allEven)

	// Output:
	// true
}

// * SOME * //

func TestSome(t *testing.T) {

	slice := []int{1, 3, 5, 7, 9, 10}

	testCases := []test.TestCase[bool]{
		{
			Desc:     "Every element should not be divisible by 2",
			Actual:   func() bool { return Every(slice, func(v, _ int) bool { return v%2 == 0 }) },
			Expected: false,
			Fail:     test.Inequality[bool],
		},
		{
			Desc:     "Every element should be smaller than or equal to 10",
			Actual:   func() bool { return Every(slice, func(v, _ int) bool { return v <= 10 }) },
			Expected: true,
			Fail:     test.Inequality[bool],
		},
	}

	// Run Tests
	for _, tc := range testCases {
		t.Run(tc.Desc, func(t *testing.T) {
			actual := tc.Actual()
			if tc.Fail(actual, tc.Expected) {
				t.Errorf("%s: Expected: %v, Actual: %v", tc.Desc, tc.Expected, actual)
			}
		})
	}

}

func ExampleSome() {
	slice := []int{1, 3, 5, 7, 9, 10}

	anyEven := Some(slice, func(v, _ int) bool {
		return v%2 == 0
	})

	fmt.Println(anyEven)

	// Output:
	// true
}

func ExampleSome_method() {
	slice := New([]int{1, 3, 5, 7, 9, 10})

	anyEven := slice.Some(func(v, _ int) bool {
		return v%2 == 0
	})

	fmt.Println(anyEven)

	// Output:
	// true
}

// * FIND * //

func TestFind(t *testing.T) {

	// Test Data
	var sliceA []int = []int{0, 1, 2, 3, 4, 5, 6}

	// Test Cases
	testCases1 := []test.TestCase[bool]{
		{
			Desc: "Find 3 in the slice",
			Actual: func() bool {
				ok, _, _ := Find(sliceA, func(v, _ int) bool { return v == 3 })
				return ok
			},
			Expected: true,
			Fail:     test.Inequality[bool],
		},
		{
			Desc: "Do not find 11 in the slice",
			Actual: func() bool {
				ok, _, _ := Find(sliceA, func(v, _ int) bool { return v == 11 })
				return ok
			},
			Expected: false,
			Fail:     test.Inequality[bool],
		},
	}
	testCases2 := []test.TestCase[int]{
		{
			Desc: "Find the first multiple of 4",
			Actual: func() int {
				_, v, _ := Find(sliceA, func(v, _ int) bool { return v%4 == 0 && v != 0 })
				return v
			},
			Expected: 4,
			Fail:     test.Inequality[int],
		},
		{
			Desc: "Find the second multiple of 3",
			Actual: func() int {
				_, v, _ := Find(sliceA, func(v, _ int) bool { return v%3 == 0 && v != 3 && v != 0 })
				return v
			},
			Expected: 6,
			Fail:     test.Inequality[int],
		},
	}

	// Run Tests
	for _, tc := range testCases1 {
		t.Run(tc.Desc, func(t *testing.T) {
			actual := tc.Actual()
			if tc.Fail(actual, tc.Expected) {
				t.Errorf("%s: Expected: %v, Actual: %v", tc.Desc, tc.Expected, actual)
			}
		})
	}
	for _, tc := range testCases2 {
		t.Run(tc.Desc, func(t *testing.T) {
			actual := tc.Actual()
			if tc.Fail(actual, tc.Expected) {
				t.Errorf("%s: Expected: %v, Actual: %v", tc.Desc, tc.Expected, actual)
			}
		})
	}

}

func ExampleFind() {
	slice := []int{1, 3, 5, 7, 9, 10}

	ok, v, _ := Find(slice, func(v, _ int) bool {
		return v%2 == 0
	})

	fmt.Println(ok, v)

	// Output:
	// true 10
}

func ExampleFind_method() {
	slice := New([]int{1, 3, 5, 7, 9, 10})

	ok, v, _ := slice.Find(func(v, _ int) bool {
		return v%2 == 0
	})

	fmt.Println(ok, v)

	// Output:
	// true 10
}

// * FIND LAST * //

func TestFindLast(t *testing.T) {

	// Test Data
	var sliceA []int = []int{0, 1, 2, 3, 4, 5, 6}

	// Test Cases
	testCases := []test.TestCase[bool]{
		{
			Desc: "Find 3 in the slice",
			Actual: func() bool {
				ok, _, _ := FindLast(sliceA, func(v, _ int) bool { return v == 3 })
				return ok
			},
			Expected: true,
			Fail:     test.Inequality[bool],
		},
		{
			Desc: "Do not find 11 in the slice",
			Actual: func() bool {
				ok, _, _ := FindLast(sliceA, func(v, _ int) bool { return v == 11 })
				return ok
			},
			Expected: false,
			Fail:     test.Inequality[bool],
		},
		{
			Desc: "Find the last multiple of 4",
			Actual: func() bool {
				ok, element, _ := FindLast(sliceA, func(v, _ int) bool { return v%4 == 0 && v != 0 })
				return ok && element == 4
			},
			Expected: true,
			Fail:     test.Inequality[bool],
		},
	}

	// Run Tests
	for _, tc := range testCases {
		t.Run(tc.Desc, func(t *testing.T) {
			actual := tc.Actual()
			if tc.Fail(actual, tc.Expected) {
				t.Errorf("%s: Expected: %v, Actual: %v", tc.Desc, tc.Expected, actual)
			}
		})
	}

}

func ExampleFindLast() {
	slice := []int{1, 3, 5, 7, 9, 10}

	ok, v, _ := FindLast(slice, func(v, _ int) bool {
		return v%2 == 0
	})

	fmt.Println(ok, v)

	// Output:
	// true 10
}

func ExampleFindLast_method() {
	slice := New([]int{1, 3, 5, 7, 9, 10})

	ok, v, _ := slice.FindLast(func(v, _ int) bool {
		return v%2 == 0
	})

	fmt.Println(ok, v)

	// Output:
	// true 10
}

// * INCLUDES * //

func TestIncludes(t *testing.T) {

	// Test Data
	var sliceA []int = []int{0, 1, 2, 3, 4, 5, 6}

	// Test Cases
	testCases := []test.TestCase[bool]{
		{
			Desc:     "Should include 3",
			Actual:   func() bool { return Includes(sliceA, 3) },
			Expected: true,
			Fail:     test.Inequality[bool],
		},
		{
			Desc:     "Should not include 99",
			Actual:   func() bool { return Includes(sliceA, 99) },
			Expected: false,
			Fail:     test.Inequality[bool],
		},
	}

	// Run Tests
	for _, tc := range testCases {
		t.Run(tc.Desc, func(t *testing.T) {
			actual := tc.Actual()
			if tc.Fail(actual, tc.Expected) {
				t.Errorf("%s\nExpected: %v\nActual: %v", tc.Desc, tc.Expected, actual)
			}
		})
	}

}

func ExampleIncludes() {
	slice := []int{1, 3, 5, 7, 9, 10}

	ok := Includes(slice, 10)

	fmt.Println(ok)

	// Output:
	// true
}

package slice

import (
	"fmt"
	"testing"
)

// * CONCAT * //

func TestConcat(t *testing.T) {

	// Test Cases
	testCases := []struct {
		desc     string
		inputs   [][]int
		expected []int
	}{
		{
			desc: "Concatenate two slices",
			inputs: [][]int{
				{0, 1, 2},
				{3, 4},
			},
			expected: []int{0, 1, 2, 3, 4},
		},
		{
			desc: "Concatenate three slices",
			inputs: [][]int{
				{0, 1, 2},
				{3, 4},
				{5, 6, 7},
			},
			expected: []int{0, 1, 2, 3, 4, 5, 6, 7},
		},
		{
			desc: "Concatenate four slices",
			inputs: [][]int{
				{0, 1, 2},
				{3, 4},
				{5, 6, 7},
				{8, 9},
			},
			expected: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
	}

	// Run Tests
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			actual := Concat(tC.inputs...)
			if !Equal(actual, tC.expected) {
				t.Errorf("%s\nwant: %v\ngot: %v", tC.desc, tC.expected, actual)
			}
		})
	}

}

func ExampleConcat() {
	a := []int{1, 2, 3}
	b := []int{4, 5, 6}

	res := Concat(a, b)
	fmt.Println(res)

	// Output:
	// [1 2 3 4 5 6]
}

func ExampleConcat_method() {
	a := []int{1, 2, 3}
	b := []int{4, 5, 6}
	c := Slice[int]([]int{7, 8, 9})

	res := New(a)
	res.Concat(b, c)
	fmt.Println(res)

	// Output:
	// [1 2 3 4 5 6 7 8 9]
}

func ExampleConcat_mixed() {
	a := []int{1, 2, 3}
	b := Slice[int]([]int{4, 5, 6})
	c := New([]int{7, 8, 9})
	d := []int{10, 11, 12}

	res := Concat(a, b, c, d)
	fmt.Println(res)

	// Output:
	// [1 2 3 4 5 6 7 8 9 10 11 12]
}

// * REVERSE * //

func TestReverse(t *testing.T) {

	// Test Cases
	testCases := []struct {
		desc     string
		slice    []int
		expected []int
	}{
		{
			desc:     "Should reverse a slice with two elements",
			slice:    []int{0, 1},
			expected: []int{1, 0},
		},
		{
			desc:     "Should reverse a slice with many elements",
			slice:    []int{0, 1, 2, 3},
			expected: []int{3, 2, 1, 0},
		},
		{
			desc:     "Should reverse a slice with one element",
			slice:    []int{0},
			expected: []int{0},
		},
		{
			desc:     "Should reverse a slice with no element",
			slice:    []int{},
			expected: []int{},
		},
	}

	// Run Tests
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			actual := Reverse(tC.slice)
			if !Equal(actual, tC.expected) {
				t.Errorf("%s\nwant: %v\ngot: %v", tC.desc, tC.expected, actual)
			}
		})
	}

}

func ExampleReverse() {
	a := []int{1, 2, 3, 4, 5}

	fmt.Println(Reverse(a))

	// Output:
	// [5 4 3 2 1]
}

func ExampleReverse_method() {
	a := New([]int{1, 2, 3, 4, 5})

	fmt.Println(*a.Reverse())

	// Output:
	// [5 4 3 2 1]
}

// * JOIN * //

func TestJoin(t *testing.T) {

	// Test Cases
	testCases := []struct {
		desc     string
		input    []any
		sep      string
		expected string
	}{
		{
			desc:     "Should convert an integer slice into a string",
			input:    []any{0, 1, 2, 3},
			sep:      " + ",
			expected: "0 + 1 + 2 + 3",
		},
		{
			desc:     "Should convert an string slice into a string",
			input:    []any{"A", "B", "C", "D", "E"},
			sep:      "-->",
			expected: "A-->B-->C-->D-->E",
		},
		{
			desc:     "Should convert an string slice into a string",
			input:    []any{[]int{0, 1}, []int{2, 3}},
			sep:      "::",
			expected: "[0 1]::[2 3]",
		},
	}

	// Run Tests
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			actual := Join(tC.input, tC.sep)
			if actual != tC.expected {
				t.Errorf("%s\nwant: %v\ngot: %v", tC.desc, tC.expected, actual)
			}
		})
	}

}

func ExampleJoin() {
	a := []int{1, 2, 3, 4, 5}

	fmt.Println(Join(a, ", "))

	// Output:
	// 1, 2, 3, 4, 5
}

func ExampleJoin_method() {
	a := New([]int{1, 2, 3, 4, 5})

	fmt.Println(a.Join(", "))

	// Output:
	// 1, 2, 3, 4, 5
}

// * PUSH * //

func TestPush(t *testing.T) {

	// Test Cases
	testCases := []struct {
		desc     string
		input    []int
		toPush   []int
		expected []int
	}{
		{
			desc:     "Should push an element onto the slice",
			input:    []int{0, 1, 2},
			toPush:   []int{3},
			expected: []int{0, 1, 2, 3},
		},
		{
			desc:     "Should push multiple elements onto the slice",
			input:    []int{0, 1, 2},
			toPush:   []int{3, 4, 5, 6, 7},
			expected: []int{0, 1, 2, 3, 4, 5, 6, 7},
		},
	}

	// Run Tests
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			actual := Push(tC.input, tC.toPush...)
			if !Equal(actual, tC.expected) {
				t.Errorf("%s\nwant: %v\ngot: %v", tC.desc, tC.expected, actual)
			}
		})
	}

}

func ExamplePush() {
	a := []int{1, 2, 3, 4, 5}

	fmt.Println(Push(a, 6, 7, 8, 9, 10))

	// Output:
	// [1 2 3 4 5 6 7 8 9 10]
}

func ExamplePush_method() {
	a := New([]int{1, 2, 3, 4, 5})

	fmt.Println(*a.Push(6, 7, 8, 9, 10))

	// Output:
	// [1 2 3 4 5 6 7 8 9 10]
}

// * POP * //

func TestPop(t *testing.T) {

	// Test Cases
	testCases := []struct {
		desc     string
		input    []int
		expected int
	}{
		{
			desc:     "Should pop off the last element from the slice",
			input:    []int{0, 1, 2, 3},
			expected: 3,
		},
	}

	// Run Tests
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			actual := Pop(&tC.input)
			if actual != tC.expected {
				t.Errorf("%s\nwant: %v\n got: %v", tC.desc, tC.expected, actual)
			}
		})
	}

}

func ExamplePop() {
	a := []int{1, 2, 3, 4, 5}

	fmt.Println(Pop(&a))

	// Output:
	// 5
}

func ExamplePop_method() {
	a := New([]int{1, 2, 3, 4, 5})

	fmt.Println(a.Pop())

	// Output:
	// 5
}

// * UNSHIFT * //

func TestUnshift(t *testing.T) {

	// Test Cases
	testCases := []struct {
		desc      string
		input     []int
		toUnshift []int
		expected  []int
	}{
		{
			desc:      "Should unshift an element onto the slice",
			input:     []int{0, 1, 2},
			toUnshift: []int{-3, -2, -1},
			expected:  []int{-3, -2, -1, 0, 1, 2},
		},
	}

	// Run Tests
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			actual := Unshift(tC.input, tC.toUnshift...)
			if !Equal(actual, tC.expected) {
				t.Errorf("%s\nwant: %v\ngot: %v", tC.desc, tC.expected, actual)
			}
		})
	}

}

func ExampleUnshift() {
	a := []int{1, 2, 3, 4, 5}

	fmt.Println(Unshift(a, 0))

	// Output:
	// [0 1 2 3 4 5]
}

func ExampleUnshift_method() {
	a := New([]int{1, 2, 3, 4, 5})

	fmt.Println(*a.Unshift(0))

	// Output:
	// [0 1 2 3 4 5]
}

// * SHIFT * //

func TestShift(t *testing.T) {

	// Test Cases
	testCases := []struct {
		desc     string
		input    []int
		expected int
	}{
		{
			desc:     "Should remove the first element from the slice",
			input:    []int{0, 1, 2},
			expected: 0,
		},
	}

	// Run Tests
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			actual := Shift(&tC.input)
			if actual != tC.expected {
				t.Errorf("%s\nwant: %v\ngot: %v", tC.desc, tC.expected, actual)
			}
		})
	}

}

func ExampleShift() {
	a := []int{1, 2, 3, 4, 5}

	fmt.Println(Shift(&a))

	// Output:
	// 1
}

func ExampleShift_method() {
	a := New([]int{1, 2, 3, 4, 5})

	fmt.Println(a.Shift())

	// Output:
	// 1
}

// * SPLICE * //

func TestSplice(t *testing.T) {

	// Test Cases
	testCases := []struct {
		desc     string
		input    []int
		start    int
		delete   int
		toAdd    []int
		expected []int
	}{
		{
			desc:     "Should remove 2 elements from the slice and add 2 elements",
			input:    []int{0, 1, 2, 3, 4, 5},
			start:    2,
			delete:   2,
			toAdd:    []int{6, 7},
			expected: []int{0, 1, 6, 7, 4, 5},
		},
		{
			desc:     "Should remove 2 elements from the slice and add 0 elements",
			input:    []int{0, 1, 2, 3, 4, 5},
			start:    2,
			delete:   2,
			toAdd:    []int{},
			expected: []int{0, 1, 4, 5},
		},
		{
			desc:     "Should remove 0 elements from the slice and add 2 elements",
			input:    []int{0, 1, 2, 3, 4, 5},
			start:    2,
			delete:   0,
			toAdd:    []int{6, 7},
			expected: []int{0, 1, 6, 7, 2, 3, 4, 5},
		},
		{
			desc:     "Should remove 0 elements from the slice and add 0 elements",
			input:    []int{0, 1, 2, 3, 4, 5},
			start:    2,
			delete:   0,
			toAdd:    []int{},
			expected: []int{0, 1, 2, 3, 4, 5},
		},
	}

	// Run Tests
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			actual := Splice(&tC.input, tC.start, tC.delete, tC.toAdd...)
			if !Equal(tC.input, tC.expected) {
				t.Errorf("%s\nwant: %v\ngot: %v", tC.desc, tC.expected, actual)
			}
		})
	}

}

func ExampleSplice() {
	a := []int{1, 2, 3, 4, 5}

	Splice(&a, 2, 2, 6, 7)

	fmt.Println(a)

	// Output:
	// [1 2 6 7 5]
}

// * CHUNK * //

func TestChunks(t *testing.T) {

	// Test Cases
	testCases := []struct {
		desc     string
		input    []int
		size     int
		expected [][]int
	}{
		{
			desc:     "Chunk the slice with 2 elements each",
			input:    []int{0, 1, 2, 3, 4, 5},
			size:     2,
			expected: [][]int{{0, 1}, {2, 3}, {4, 5}},
		},
		{
			desc:     "Chunk the slice with 3 elements each",
			input:    []int{0, 1, 2, 3, 4, 5},
			size:     3,
			expected: [][]int{{0, 1, 2}, {3, 4, 5}},
		},
		{
			desc:     "Chunk an uneven slice with 2 elements",
			input:    []int{0, 1, 2, 3, 4},
			size:     2,
			expected: [][]int{{0, 1}, {2, 3}, {4}},
		},
	}

	// Run Tests
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			actual := Chunk(tC.input, tC.size)
			for i := 0; i < len(tC.expected); i++ {
				if !Equal(actual[i], tC.expected[i]) {
					t.Errorf("%s\nwant: %v\ngot: %v", tC.desc, tC.expected[i], actual[i])
				}
			}
		})
	}

}

func ExampleChunk() {
	a := []int{1, 2, 3, 4, 5}

	fmt.Println(Chunk(a, 2))

	// Output:
	// [[1 2] [3 4] [5]]
}

func ExampleChunk_method() {
	a := New([]int{1, 2, 3, 4, 5})

	fmt.Println(a.Chunk(2))

	// Output:
	// [[1 2] [3 4] [5]]
}

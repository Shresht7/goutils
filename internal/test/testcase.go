package test

// ---------
// TEST CASE
// ---------

// TestCase is a struct that contains the information needed to describe a test case.
type TestCase[T any] struct {
	// Description of the test case
	Desc string
	// Function to generate the actual value
	Fn func() T
	// Expected value to compare the actual value with
	Expected T
	// Function to compare the actual and expected values. If true, the test case has failed.
	Fail func(actual, expected T) bool
}

// -------
// HELPERS
// -------

// Inequality is a helper function that returns true if the actual value is not equal to the expected value.
func Inequality[T comparable](actual, expected T) bool {
	return actual != expected
}

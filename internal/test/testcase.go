package test

// ---------
// TEST CASE
// ---------

// TestCase is a struct that contains the information needed to describe a test case
type TestCase[T any] struct {
	Desc     string                        // Description of the test case
	Actual   func() T                      // Function to generate the actual value
	Expected T                             // Expected value to compare the actual value with
	Fail     func(actual, expected T) bool // Function to compare the actual and expected values. If true, the test case has failed
}

// -------
// HELPERS
// -------

// Inequality is a helper function that returns true if the actual value is not equal to the expected value
func Inequality[T comparable](actual, expected T) bool {
	return actual != expected
}

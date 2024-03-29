package generic

import (
	"testing"
)

// https://go.dev/doc/tutorial/generics#add_generic_function
func TestGenericBasic(t *testing.T) {
	// Initialize a map for the integer values
	ints := map[string]int64{
		"first":  34,
		"second": 12,
	}

	// Initialize a map for the float values
	floats := map[string]float64{
		"first":  35.98,
		"second": 26.99,
	}

	t.Logf("Generic Sums: %v and %v\n", Sum(ints), Sum(floats))
}

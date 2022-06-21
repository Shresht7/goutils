package sliceutils

import (
	"fmt"
	"testing"
)

func TestForEach(t *testing.T) {

	slice := []string{"A", "B", "C"}

	ForEach(slice, func(char string, idx int) {
		fmt.Println(idx, char)
	})

}

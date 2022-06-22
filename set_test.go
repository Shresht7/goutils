package sliceutils

import "testing"

func TestGetSet(t *testing.T) {

	slice := []int{0, 1, 1, 2, 3, 3, 3, 4, 5}

	set := GetSet(slice)

	if !Equal(set, []int{0, 1, 2, 3, 4, 5}) {
		t.Error("Failed to create set")
	}

}

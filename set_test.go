package sliceutils

import "testing"

func TestSet(t *testing.T) {

	slice := []int{0, 1, 1, 1, 2, 3, 3, 3, 3, 3, 4, 5, 5}

	set := GetSet(slice)

	if !Equal(set, []int{0, 1, 2, 3, 4, 5}) {
		t.Error("Failed to create set")
	}

	set.Add(6)
	set.Add(7)
	set.Add(7)

	if set.Size() != 8 {
		t.Error("Failed to add to set")
	}

	set.Clear()

	if set.Size() != 0 {
		t.Error("Failed to clear set")
	}

}

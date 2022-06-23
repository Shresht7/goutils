package set

import (
	"testing"

	"github.com/Shresht7/sliceutils"
)

func TestSet(t *testing.T) {

	slice := []int{0, 1, 1, 1, 2, 3, 3, 3, 3, 3, 4, 5, 5}

	set := GetSet(slice)

	if !sliceutils.Equal(set, []int{0, 1, 2, 3, 4, 5}) {
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

func TestSubSetAndSuperSet(t *testing.T) {

	setA := GetSet([]int{0, 1, 2, 3, 4, 5, 6, 7})
	setB := GetSet([]int{3, 4})
	setC := GetSet([]int{3, 11})

	if !setB.IsSubSet(&setA) {
		t.Errorf("%v should be a subset of %v", setB, setA)
	}

	if !setA.IsSuperSet(&setB) {
		t.Errorf("%v should be a superset of %v", setA, setB)
	}

	if setC.IsSubSet(&setA) || setC.IsSuperSet(&setB) {
		t.Errorf("%v should not be a subset of %v or superset of %v", setC, setA, setB)
	}

}

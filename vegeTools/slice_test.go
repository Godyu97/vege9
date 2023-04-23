package vegeTools

import (
	"testing"
)

func TestItemIsInSlice(t *testing.T) {
	strslice := []string{"1", "2", "3"}
	if ItemIsInSlice("2", strslice) == false {
		t.Error("ItemIsInSlice failed")
	}
	intslice := []int{1, 2, 3}
	if ItemIsInSlice(2, intslice) == false {
		t.Error("ItemIsInSlice failed")
	}
}

func TestSliceIsInSlice(t *testing.T) {
	strslice := []string{"1", "2", "3"}
	if SliceIsInSlice([]string{"2", "3"}, strslice) == false {
		t.Error("SliceIsInSlice failed")
	}
	intslice := []int{1, 2, 3}
	if SliceIsInSlice([]int{2, 3}, intslice) == false {
		t.Error("SliceIsInSlice failed")
	}
}

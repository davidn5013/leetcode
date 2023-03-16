package grind75

import (
	"reflect"
	"testing"
)

func TestMergeIntervals(t *testing.T) {
	gut := merge([][]int{{1, 3}, {2, 6}, {8, 19}, {9, 18}})
	want := [][]int{{1, 6}, {8, 19}}
	if !reflect.DeepEqual(gut, want) {
		t.Errorf("gut %v; wanted %v", gut, want)
	}

	gut = merge([][]int{{1, 4}, {0, 4}})
	want = [][]int{{0, 4}}
	if !reflect.DeepEqual(gut, want) {
		t.Errorf("gut %v; wanted %v", gut, want)
	}

}

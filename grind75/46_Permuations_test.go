package grind75

import (
	"reflect"
	"testing"
)

func TestPermute(t *testing.T) {
	input := []int{1, 2, 3}
	want := [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 2, 1}, {3, 1, 2}}
	gut := permute(input)

	if !reflect.DeepEqual(gut, want) {
		t.Errorf("gut %v,\n wanted %v", gut, want)
	}
}

package grind75

import (
	"reflect"
	"testing"
)

func TestDistansToZeroInMatrix(t *testing.T) {
	inp := [][]int{
		{0, 0, 0},
		{0, 1, 0},
		{1, 1, 1},
	}

	got := DistansToZeroInMatrix(inp)

	want := [][]int{
		{0, 0, 0},
		{0, 1, 0},
		{1, 2, 1},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v; wanted %v", got, want)
	}
}

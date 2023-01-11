package grind75

import (
	"reflect"
	"testing"
)

func TestDistThreeSum(t *testing.T) {
	want := [][]int{
		{-1, 0, 1},
		{-1, -1, 2},
	}
	wantOr := [][]int{
		{-1, -1, 2},
		{-1, 0, 1},
	}
	got := DistThreeSum([]int{-1, 0, 1, 2, -1, -4})
	if reflect.DeepEqual(got, want) == false &&
		reflect.DeepEqual(got, wantOr) == false {
		t.Errorf("got %v; wanted %v", got, want)
	}
}

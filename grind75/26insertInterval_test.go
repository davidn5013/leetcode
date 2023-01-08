package grind75

import (
	"reflect"
	"testing"
)

func TestInsertInterval(t *testing.T) {
	got := InsertInterval([][]int{
		{1, 2},
		{3, 5},
		{6, 7},
		{8, 10},
		{12, 16},
	}, []int{4, 8})
	want := [][]int{{1, 2}, {3, 10}, {12, 16}}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v;wanted %v", got, want)
	}
}

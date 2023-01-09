package grind75

import (
	"reflect"
	"testing"

	"github.com/davidn5013/leetcode/tools"
)

func TestKclosest(t *testing.T) {
	points := [][]int{{1, 3}, {-2, 2}}
	k := 1

	got := Kclosest(points, k)

	want := [][]int{{-2, 2}}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("got %v; wanted %v", got, want)
	}
}

func TestCopy2dInt(t *testing.T) {
	src := [][]int{{3, 3}, {5, -1}, {-2, 4}, {0, 0}, {1, 0}}
	dst := tools.Copy2dInt(src)

	if !reflect.DeepEqual(src, dst) {
		t.Errorf("got %v; wanted %v", dst, src)
	}
}

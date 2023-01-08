package grind75

import (
	"reflect"
	"testing"
)

func TestKclosest(t *testing.T) {
	// Input: points = [[3,3],[5,-1],[-2,4]], k = 2
	// Output: [[3,3],[-2,4]]
	// Explanation: The answer [[-2,4],[3,3]] would also be accepted.
	points := [][]int{{3, 3}, {5, -1}, {-2, 4}}
	k := 2

	got := Kclosest(points, k)

	want := [][]int{{3, 3}, {-2, 4}}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("got %v; wanted %v", got, want)
	}
}

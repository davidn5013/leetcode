package grind75

import "testing"

func TestCanFinish(t *testing.T) {
	want := true
	got := canFinish(2, [][]int{{1, 0}})
	if want != got {
		t.Errorf("got %v ; wanted %v", got, want)
	}
}

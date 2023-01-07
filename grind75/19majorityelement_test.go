package grind75

import "testing"

func TestMajorityElement(t *testing.T) {
	want := 2
	got := MajorityElement([]int{2, 1, 2, 2, 1, 2, 1})
	if got != want {
		t.Errorf("MajorityElement []int{2,2,1,1,1,2,2} gut %d; want %d", got, want)
	}
}

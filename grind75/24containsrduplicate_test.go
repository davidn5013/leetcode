package grind75

import "testing"

func TestContainDuplicate(t *testing.T) {
	got, want := ContainDuplicate([]int{1, 2, 3, 1}), true
	if got != want {
		t.Errorf("got %t; want %t", got, want)
	}
	got, want = ContainDuplicate([]int{1, 2, 3, 4}), false
	if got != want {
		t.Errorf("got %t; want %t", got, want)
	}
	got, want = ContainDuplicate([]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}), true
	if got != want {
		t.Errorf("got %t; want %t", got, want)
	}
}

func BenchmarkContainDuplicate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ContainDuplicate([]int{1, 2, 3, 1})
	}
}

package grind75

import "testing"

func TestClimbStairs(t *testing.T) {
	ans := ClimbStairs(4)
	want := 5
	if ans != want {
		t.Errorf("got %d, want %d", ans, want)
	}
}

func TestClimbStairsRecursive(t *testing.T) {
	ans := ClimbStairsRecursive(4)
	want := 5
	if ans != want {
		t.Errorf("got %d, want %d", ans, want)
	}
}

func BenchmarkClimbStairsRecursive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ClimbStairsRecursive(9999999)
	}
}

func BenchmarkClimbStairs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ClimbStairs(9999999)
	}
}

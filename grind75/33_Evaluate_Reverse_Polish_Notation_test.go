package grind75

import (
	"testing"
)

func TestEvalRPN(t *testing.T) {
	got := EvalRPN([]string{"2", "1", "+", "3", "*"})
	want := 9
	if got != want {
		t.Errorf("got %d ; wanted %d", got, want)
	}
}

func TestEvalRPN2(t *testing.T) {
	got := EvalRPN([]string{"4", "13", "5", "/", "+"})
	want := 6
	if got != want {
		t.Errorf("got %d ; wanted %d", got, want)
	}
}

func TestEvalRPN3(t *testing.T) {
	got := EvalRPN([]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"})
	want := 22
	if got != want {
		t.Errorf("got %d ; wanted %d", got, want)
	}
}

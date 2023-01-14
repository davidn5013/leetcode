package grind75

import (
	"fmt"
	"testing"
)

func ExampleEvalRPN(t *testing.T) {
	got := EvalRPN([]string{"2", "1", "+", "3", "*"})
	want := 9

	fmt.Println(got, want)

	if got != want {
		t.Errorf("got %d ; wanted %d", got, want)
	}

}

func ExampleEvalRPN2(t *testing.T) {
	got := EvalRPN([]string{"4", "13", "5", "/", "+"})
	want := 6

	fmt.Println(got, want)

	if got != want {
		t.Errorf("got %d ; wanted %d", got, want)
	}

}

func ExampleEvalRPN3(t *testing.T) {
	got := EvalRPN([]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"})
	want := 22

	fmt.Println(got, want)

	if got != want {
		t.Errorf("got %d ; wanted %d", got, want)
	}

}

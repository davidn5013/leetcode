package p1

import (
	"reflect"
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	inp := 5
	want := []string{"1", "2", "Fizz", "4", "Buzz"}
	gut := FizzBuzz(inp)
	if !reflect.DeepEqual(want, gut) {
		t.Errorf("gut %v; wanted %v", gut, want)
	}
}

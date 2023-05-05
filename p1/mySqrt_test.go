package p1

import (
	"math"
	"testing"
)

func Test_usqr4(t *testing.T) {
	const test = 1_000_000_000_000_000
	gut := usqrt4(test)
	want := uint(math.Sqrt(float64(test)))
	if gut != want {
		// fmt.Errorf("wanted %d ;gute %d", want, gut)
	}
}

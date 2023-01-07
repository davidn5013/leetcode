package grind75

import (
	"fmt"
	"testing"
)

// 0020. Test of isValid
func TestIsValid(t *testing.T) {
	var tests = []struct {
		s    string
		want bool
	}{
		{"()", true}, {"()[]{}", true}, {"(]", false},
		{"{[]}", true}, {"{[()]}", true}, {"{[(])", false},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("%s,%v", tt.s, tt.want)
		t.Run(testname, func(t *testing.T) {
			ans := IsValid(tt.s)
			if tt.want != ans {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}

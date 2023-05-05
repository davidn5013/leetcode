package p1

import (
	"fmt"
	"testing"
)

func TestTitleToNumber(t *testing.T) {
	var tests = []struct {
		a    string
		want int
	}{
		{"A", 1},
		{"AB", 28},
		{"ZY", 701},
	}

	for _, tt := range tests {

		testname := fmt.Sprintf(" %s", tt.a)
		t.Run(testname, func(t *testing.T) {
			ans := titleToNumber(tt.a)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}

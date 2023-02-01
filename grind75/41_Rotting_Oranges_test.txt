package grind75

import (
	"testing"
)

func TestOrangesRottning(t *testing.T) {
	var input = [][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}}

	/*
		prt := func() {
			for r := 0; r < len(input); r++ {
				for c := 0; c < len(input[1]); c++ {
					log.Print(input[r][c])
				}
				log.Println()
			}
		}
	*/

	// prt()
	want := 4
	gut := orangesRotting(input)
	// prt()

	if want != gut {
		t.Errorf("gut %d; wanted %d", gut, want)
	}
}

func TestOrangesRottning2(t *testing.T) {
	var input = [][]int{
		{0, 2},
	}
	want := -1
	gut := orangesRotting(input)

	if want != gut {
		t.Errorf("gut %d; wanted %d", gut, want)
	}
}

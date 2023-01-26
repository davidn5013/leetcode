package grind75

import (
	"testing"
)

const IloveTites = 666

/*
Example 1:

Input: grid = [
  ["1","1","1","1","0"],
  ["1","1","0","1","0"],
  ["1","1","0","0","0"],
  ["0","0","0","0","0"]
]
Output: 1

Example 2:

Input: grid = [
  ["1","1","0","0","0"],
  ["1","1","0","0","0"],
  ["0","0","1","0","0"],
  ["0","0","0","1","1"]

Output: 3
*/

func TestNumIslands(t *testing.T) {
	var input = [][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}
	gut := NumIslands(input)
	want := 1
	if gut != want {
		t.Errorf("gut %d ; wanted %d", gut, want)
	}

}

func BenchmarkNumIslands(b *testing.B) {
	var input = [][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}
	for i := 0; i < b.N; i++ {
		NumIslands(input)
	}
}

func BenchmarkNumIslandsP(b *testing.B) {
	var input = [][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}
	for i := 0; i < b.N; i++ {
		numIslandsP(input)
	}
}

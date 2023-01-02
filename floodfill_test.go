// Using exampel to test floodfill
package leetcode

import (
	"fmt"
	"testing"
)

var testGrid = [][]int{
	{1, 1, 1},
	{1, 1, 0},
	{1, 0, 1},
}

func ExampleFloodFill() {
	grid2 := FloodFill(testGrid, 0, 0, 2)
	prtMatrix(grid2)

	// OutPut:
	// 222
	// 220
	// 201
}

func prtMatrix(grid [][]int) {
	for _, m := range grid {
		for _, v := range m {
			fmt.Printf("%d", v)
		}
		fmt.Println()
	}
}

func BenchmarkFloodFill(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FloodFill(testGrid, 0, 0, 2)
	}
}

func BenchmarkFloodFillIt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FloodFillIt(testGrid, 0, 0, 2)
	}
}

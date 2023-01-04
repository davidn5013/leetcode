package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

// failing exampel on leetcood
func TestFloodFill(t *testing.T) {
	image := [][]int{
		{0, 0, 0},
		{0, 0, 0},
	}
	ans := FloodFill(image, 1, 0, 2)
	want := [][]int{
		{2, 2, 2},
		{2, 2, 2},
	}

	if reflect.DeepEqual(ans, want) != true {
		t.Errorf("got %v = ; want %v", ans, want)
	}
}

// Using exampel to test floodfill

var testGrid = [][]int{
	{1, 1, 1},
	{1, 1, 0},
	{1, 0, 1},
}

func ExampleFloodFill() {
	grid2 := FloodFill(testGrid, 0, 0, 2)
	printGrid(grid2)

	// OutPut:
	// 222
	// 220
	// 201
}

func printGrid(grid [][]int) {
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

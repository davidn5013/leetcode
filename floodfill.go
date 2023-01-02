// Package leetcode 733. Flood Fill
package leetcode

import "fmt"

// FloodFill fills the region starting at point (x, y)
// with color c and returns the resulting grid.
// Recursive version.
func FloodFill(grid [][]int, row, col, color int) [][]int {
	const wall = 0
	if row < 0 ||
		col < 0 ||
		row > len(grid)-1 ||
		col > len(grid[0])-1 ||
		grid[row][col] == wall ||
		grid[row][col] == color {
		return grid
	}
	grid[row][col] = color
	FloodFill(grid, row+1, col, color)
	FloodFill(grid, row-1, col, color)
	FloodFill(grid, row, col+1, color)
	FloodFill(grid, row, col-1, color)
	return grid
}

// FloodFill fills the region starting at point (x, y)
// with color c and returns the resulting grid.
// Iteration verson.
func FloodFillIt(grid [][]int, row, col int, target_color int) [][]int {
	oldCol := grid[row][col]
	if oldCol == target_color {
		return grid
	}

	visit := map[string]bool{getPosString(row, col): true}
	queue := [][]int{{row, col}}
	for len(queue) > 0 {
		cRow, cCol := queue[0][0], queue[0][1]
		queue = queue[1:]
		grid[cRow][cCol] = target_color

		directions := [][]int{
			{-1, 0},
			{0, 1},
			{1, 0},
			{0, -1},
		}

		for _, direction := range directions {
			rowOff, colOff := direction[0], direction[1]
			nearRow := cRow + rowOff
			nearCol := cCol + colOff
			pos := getPosString(nearRow, nearCol)
			if isInBounds(grid, nearRow, nearCol) &&
				grid[nearRow][nearCol] == oldCol {
				if _, found := visit[pos]; found {
					continue
				}
				visit[pos] = true
				queue = append(queue, []int{nearRow, nearCol})
			}

		}
	}
	return grid
}

func isInBounds(grid [][]int, row, col int) bool {
	rowInBounds := 0 <= row &&
		row < len(grid)
	colInBounds := 0 <= col &&
		col < len(grid[0])
	return rowInBounds && colInBounds
}

func getPosString(row, col int) string {
	return fmt.Sprintf("%d:%d", row, col)
}

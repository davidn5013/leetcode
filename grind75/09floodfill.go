package grind75

import (
	p "github.com/davidn5013/leetcode/tools/ds"
)

/* My first all made by me flood fill. This gave realy good feeling after loot of
 * work for day to realy understand.
 */

var floodfillDirection = []p.Pos{
	{Row: -1, Col: 0},
	{Row: 0, Col: 1},
	{Row: 1, Col: 0},
	{Row: 0, Col: -1},
}

// FloodFill 0733 fills the region starting at point (sr,sc) with color and returns the
// resulting grid. Iterative version using array of position.
func FloodFill(image [][]int, sr, sc, color int) [][]int {
	oldColor := image[sr][sc]
	if oldColor == color {
		return image
	}

	stack := make([]p.Pos, 1)
	stack[0] = p.Pos{Col: sr, Row: sc}

	for len(stack) > 0 {
		// pop from position stack
		curr := stack[0]
		stack = stack[1:]

		row, col := curr.Row, curr.Col // r=row , c=coll
		image[row][col] = color

		func(direction []p.Pos) {
			for _, dir := range direction {
				// nxtrow=next row , nxtcol=next col
				nxtrow, nxtcol := row+dir.Row, col+dir.Col
				switch {
				// check inbound height and width
				case nxtrow < 0 || nxtrow >= len(image) ||
					nxtcol < 0 || nxtcol >= len(image[0]):
					continue
				case image[nxtrow][nxtcol] == oldColor:
					stack = append(stack, p.Pos{Row: nxtrow, Col: nxtcol})
				}
			}
		}(floodfillDirection)
	}

	return image
}

// FloodFill2 fills the region starting at point (x, y) with color c and returns the
// resulting grid. Recursive version.
/*func FloodFill2(image [][]int, sr, sc, color int) [][]int {
	// I did not find a way to store start position in a recursive version
	// so FloodFill is iterative and stores start position color
	const wall = 0
	if sr < 0 ||
		sc < 0 ||
		sr > len(image)-1 ||
		sc > len(image[0])-1 ||
		image[sr][sc] == wall ||
		image[sr][sc] == color {
		return image
	}
	image[sr][sc] = color
	FloodFill(image, sr+1, sc, color)
	FloodFill(image, sr-1, sc, color)
	FloodFill(image, sr, sc+1, color)
	FloodFill(image, sr, sc-1, color)
	return image
}

// FloodFillIt fills the region starting at point (x, y)  with color c and returns the
// resulting grid. Recursive version. Iteration verson.
func FloodFillIt(grid [][]int, row, col, targetColor int) [][]int {
	oldCol := grid[row][col]
	if oldCol == targetColor {
		return grid
	}

	visit := map[string]bool{getPosString(row, col): true}
	queue := [][]int{{row, col}}
	for len(queue) > 0 {
		cRow, cCol := queue[0][0], queue[0][1]
		queue = queue[1:]
		grid[cRow][cCol] = targetColor

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
*/

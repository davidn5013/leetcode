package grind75

// NumIslands (Numbers of Islands ) in 2d binary grid wherer 1 land and 0 water find the islands
/* 	{1, 1, 1, 1, 0},
{1, 1, 0, 1, 0},
{1, 1, 0, 0, 0},
{0, 0, 0, 0, 0},
has 1 island */

// NumIslands number of islands of 1 in matrix with sea as 0
// Source: https://ljun20160606.github.io/leetcode/algorithms/200/
func NumIslands(grid [][]byte) int {
	var count int
	for x := range grid {
		for y := range grid[x] {
			if grid[x][y] == '1' {
				search(grid, x, y)
				count++
			}
		}
	}
	return count
}

func search(grid [][]byte, x, y int) {
	if x < 0 || x >= len(grid) || y < 0 || y >= len(grid[0]) || grid[x][y] == '0' {
		return
	}
	grid[x][y] = '0'
	search(grid, x-1, y)
	search(grid, x+1, y)
	search(grid, x, y-1)
	search(grid, x, y+1)
}

// Fastest version on LeetCode using prt to grid think thats on gcc. Faster for some reason.
func dfsP(gridptr *[][]byte, i, j int) {
	grid := (*gridptr)
	m, n := len(grid), len(grid[0])
	if i > m-1 || i < 0 || j > n-1 || j < 0 {
		return
	}
	if grid[i][j] == '0' {
		return
	}
	grid[i][j] = '0'
	dfsP(gridptr, i+1, j)
	dfsP(gridptr, i-1, j)
	dfsP(gridptr, i, j+1)
	dfsP(gridptr, i, j-1)
}

func numIslandsP(grid [][]byte) int {
	count := 0
	for i := range grid {
		for j := range grid[0] {
			if grid[i][j] == '1' {
				dfsP(&grid, i, j)
				count++
			}
		}
	}
	return count
}

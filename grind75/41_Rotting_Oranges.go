package grind75

import (
	"github.com/davidn5013/leetcode/tools/ds"
)

// organgesRotting every cycle will change near by cells i 4 dirs
/*You are given an m x n grid where each cell can have one of three values:
    0 representing an empty cell,
    1 representing a fresh orange, or
    2 representing a rotten orange.
Every minute, any fresh orange that is 4-directionally adjacent to a rotten orange becomes rotten.

Return the minimum number of minutes that must elapse until no cell has a fresh orange.
If this is impossible, return -1.*/
func orangesRotting(grid [][]int) int {
	var (
		res int      // hold return value
		q   []ds.Pos // struct with (row int, col int}
		// directions
		dir = [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	)

	const fresh = 1
	const rotten = 2

	// loop throw all cells and find rotten store position of rotten in queue
	// if the queue was 0 then set return -1. res is now = 0
	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid); c++ {
			if grid[r][c] == rotten {
				q = append(q, ds.Pos{Row: r, Col: c})
			}
		}
	}

	// no rotten oranges return -1
	if len(q) == 0 {
		return -1
	}

	for len(q) > 0 {
		// pop front queue
		curr := q[0]
		q = q[1:]

		// loop throw queue updating new rotten and new position of rotten
		for _, d := range dir {
			r, c := curr.Row+d[0], curr.Row+d[1]

			if r >= 0 && c >= 0 &&
				r < len(grid) && c < len(grid[1]) {

				if grid[r][c] == fresh {
					grid[r][c] = rotten
					// push to end of queue
					q = append(q, ds.Pos{Row: r, Col: c})
				}

			}
		}
		res++
	}

	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid); c++ {
			if grid[r][c] == fresh {
				return -1
			}
		}
	}

	return res
}

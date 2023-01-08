package grind75

import "github.com/davidn5013/leetcode/tools"

// UpdateMatrix return distans of nearest o for each cell
// Exampel
// Input  Result
// 0 0 0  0 0 0
// 0 1 0  0 1 0
// 1 1 1  1 2 1
// 2,1 has 2 because that how meen sideway or up that is need to get to 0
func DistansToZeroInMatrix(mat [][]int) (res [][]int) {
	return updateMatrix(mat)
}

// TODO Go back and understand this function better

// updateMatrix return distans of nearest 0 for each cell
// leetcode 542 01 Matrix
func updateMatrix(mat [][]int) (res [][]int) {
	res = make([][]int, len(mat))

	// copy mat to res
	for i := range mat {
		res[i] = make([]int, len(mat[i]))
		copy(res[i], mat[i])
	}

	// Calculate shortest manhattan distans for ever cell not 1
	for row := 0; row < len(res); row++ {
		for col := 0; col < len(res[row]); col++ {
			if res[row][col] != 1 {
				continue
			}
			res[row][col] = minDist(res, row, col)
		}
	}
	return res

}

// minDist travers and return Manhattan distance
// helper for updateMatrix
func minDist(mat [][]int, row, col int) (minDist int) {
	minDist = len(mat) + len(mat[row]) // maximum possible distance

	for r := 0; r < len(mat); r++ {
		for c := 0; c < len(mat[r]); c++ {
			if mat[r][c] != 0 {
				continue
			}

			distManh := tools.AbsInt(r-row) + tools.AbsInt(c-col) // Manhattan distance

			if distManh < minDist {
				minDist = distManh
			}

		}
	}

	return minDist
}

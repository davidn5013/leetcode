package grind75

import "github.com/davidn5013/leetcode/tools"

// Leetcode 973 Grind75 number 28

// Kclosest sort input array after distant to zero (0,0) return k part of array
// Original array is intact BUT not it's order!
func Kclosest(res [][]int, k int) [][]int {
	for i := 1; i < len(res); i++ {
		if i >= 1 {

			// distance from 0,0 to coordination previous and current in the on loop
			distPrev := tools.EucliDistSquared(0, 0, res[i-1][0], res[i-1][1])
			dist := tools.EucliDistSquared(0, 0, res[i][0], res[i][1])

			// Sorting input array points
			if distPrev > dist {
				res[i-1], res[i] = res[i], res[i-1]
				i -= 2
			}
		}
	}

	return res[:k]
}

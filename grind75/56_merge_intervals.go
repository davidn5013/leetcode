package grind75

import (
	"sort"
)

// Leetcode 56. Merge Intervals
// Grind75 number 56

// merge overlapping intervals
func merge(intervals [][]int) (ret [][]int) {
	sort.Ints(intervals[0])
	cur := intervals[0]

	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] > cur[1] {
			ret = append(ret, append(make([]int, 0), cur...))
			cur = intervals[i]
		} else if intervals[i][1] > cur[1] {
			cur[1] = intervals[i][1]
		}
	}

	ret = append(ret, cur)

	return ret
}

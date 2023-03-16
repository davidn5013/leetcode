package grind75

import (
	"sort"
)

// Leetcode 56. Merge Intervals
// Grind75 number 56

// merge overlapping intervals
func merge(intervals [][]int) [][]int {
	const start, end = 0, 1

	var merged [][]int

	// Sort make merge start interval unnecessary
	// No we can not use sort.Ints(intervals[1])
	if len(intervals) > 1 {
		// custom sort on intervals start
		sort.Slice(intervals, func(i, j int) bool {
			return intervals[i][start] < intervals[j][start]
		})
	}

	// Because of sort we do not need to if start on sub array is lower then current array,
	for _, interval := range intervals {
		last := len(merged) - 1
		// first merge or a new high interval then store in merge
		if last < 0 || interval[start] > merged[last][end] {
			merged = append(merged,
				[]int{start: interval[start], end: interval[end]},
			)
			// intersect ting then merge
		} else if interval[end] > merged[last][end] {
			merged[last][end] = interval[end]
		}
	}

	return merged[:len(merged):len(merged)]
}

package grind75

import "github.com/davidn5013/leetcode/tools"

// Leetcode 57 Grind75 26

// InsertInterval add new interval to a array of non-overlapping intervals
// Example
// [][]int{{1,3},{6,9}} adding []int{2,5} only extend 1,3 to 1,5
// an result in [][]int{{1,5},{6,9}}
// [][]int{{1,2},{3,5},{6,7},{8,10},{12,16}} adding {4,8}
// [][]int{{1,2},{3,10},{12,16}}
// Because the new interval [4,8] overlaps with [3,5],[6,7],[8,10].
func InsertInterval(intervals [][]int, newInterval []int) (res [][]int) {
	for _, interval := range intervals {
		if interval[1] < newInterval[0] {
			res = append(res, interval)
			continue
		} else if interval[0] > newInterval[1] {
			res = append(res, newInterval)
			newInterval = interval
		} else {
			newInterval = []int{
				tools.MinInt(interval[0], newInterval[0]),
				tools.MaxInt(interval[1], newInterval[1]),
			}
		}
	}
	res = append(res, newInterval)
	return res
}

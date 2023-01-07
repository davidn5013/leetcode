package grind75

import t "github.com/davidn5013/leetcode/tools"

// Leetcode 53 and grind75 nr 25

// MaxSubArray 0053 biggest sum of sub arrays in array
// Exampel:
// fmt.Println(MaxSubArray([-2,1,-3,4,-1,2,1,-5,4])
// Prints out : 6
func MaxSubArray(nums []int) int {
	curBest, largest := nums[0], nums[0]
	for _, n := range nums[1:] {
		curBest = t.MaxInt(n, curBest+n)
		largest = t.MaxInt(largest, curBest)
	}
	return largest
}

package grind75

// TwoSum 0001. Two Sum
// Given an array of integers nums and an integer target,
// return indices of the two numbers such that they add up to
// target. You may assume that each input would have exactly
// one solution, and you may  not use the same element twice.
// You can return the answer in any order.
func TwoSum(nums []int, target int) []int {

	// Constraints:
	//	Only one valid answer exists.
	//	2 <= nums.length <= 10e4	, -10e9 <= target <= 10e9
	if len(nums) <= 1 || len(nums) >= 10e4 || target <= -10e9 || target >= 10e9 {
		return []int{0, 0}
	}

	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			switch {
			// Constraints:
			// -10e9 <= nums[i] <= 10e9
			case i == j:
				continue
			case j <= -10e9 || j >= 10e9:
				return []int{0, 0}
			case nums[i]+nums[j] == target:
				return []int{i, j}
			}
		}
	}

	return []int{0, 0}
}

// beautiful solution from bharadwaj6 on https://leetcode.com/problems/two-sum/discuss/261590/100-golang
// no constraints check
/*
func twoSum(nums []int, target int) []int {
	count := make(map[int]int)
	for i, num := range nums {
		j, ok := count[num]
		if ok {
			return []int{j, i}
		}
		count[target-num] = i
	}
	return []int{}

}
*/

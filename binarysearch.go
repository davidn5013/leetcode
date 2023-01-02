// Package leetcode 704. Binary Search
package leetcode

// Search Given an array of integers nums which is sorted in ascending order,
// and an integer target, search target in nums.
// If target exists, then return its index. Otherwise, return -1.
func Search(nums []int, target int) int {
	for i, n := range nums {
		if n == target {
			return i
		}
	}
	return -1
}

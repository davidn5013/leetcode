package grind75

import "sort"

// Grind75 Number 43
// leetcode 39. Combination Sum

/*
Given an array of distinct integers candidates and a target integer target, return a list
of all unique combinations of candidates where the chosen numbers sum to target.
You may return the combinations in any order.

The same number may be chosen from candidates an unlimited number of times.
Two combinations are unique if the frequency of at least one of the chosen numbers is different.

The test cases are generated such that the number of unique combinations that sum up to target
is less than 150 combinations for the given input.
*/

// Backtracking Algorithm

// TODO Read more about backtracking alogrightm

// combinationSum return a list of all unique combinations of candidates
func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	var helper func(candidate, share []int, target int)
	var r [][]int
	helper = func(c, share []int, target int) {
		if target == 0 {
			r = append(r, share)
			return
		}
		if len(c) == 0 || target < c[0] {
			return
		}
		share = share[:len(share):len(share)]

		helper(c, append(share, c[0]), target-c[0])
		helper(c[1:], share, target)
	}
	helper(candidates, []int{}, target)
	return r
}

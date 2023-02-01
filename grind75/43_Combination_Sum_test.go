package grind75

import (
	"reflect"
	"testing"
)

/*
Example 1:

Input: candidates = [2,3,6,7], target = 7
Output: [[2,2,3],[7]]
Explanation:
2 and 3 are candidates, and 2 + 2 + 3 = 7. Note that 2 can be used multiple times.
7 is a candidate, and 7 = 7.
These are the only two combinations.

Example 2:

Input: candidates = [2,3,5], target = 8
Output: [[2,2,2,2],[2,3,3],[3,5]]

Example 3:

Input: candidates = [2], target = 1
Output: []

*/

func TestCombinationSum(t *testing.T) {
	input := []int{2, 3, 6, 7}
	target := 7
	want := [][]int{{2, 2, 3}, {7}}
	gut := combinationSum(input, target)
	if !reflect.DeepEqual(want, gut) {
		t.Errorf("gut %v; wanted %v", gut, want)
	}
	input = []int{2}
	target = 1
	want = [][]int{}
	gut = combinationSum(input, target)
	if gut != nil {
		t.Errorf("gut %v; wanted %v", gut, want)
	}
	input = []int{2, 3, 5}
	target = 8
	want = [][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}}
	gut = combinationSum(input, target)
	if !reflect.DeepEqual(want, gut) {
		t.Errorf("gut %v; wanted %v", gut, want)
	}

}

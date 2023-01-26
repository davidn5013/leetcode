package grind75

import "testing"

/*
Example 1:

Input: nums = [4,5,6,7,0,1,2], target = 0
Output: 4

Example 2:

Input: nums = [4,5,6,7,0,1,2], target = 3
Output: -1

Example 3:

Input: nums = [1], target = 0
Output: -1
*/

func TestSearchRotatedSortedArray(t *testing.T) {
	input := []int{4, 5, 6, 7, 0, 1, 2}
	want := 4
	target := 0
	gut := SearchRotatedSortedArray(input, target)
	if want != gut {
		t.Errorf("gut %d; wanted %d", gut, want)
	}
}

func TestSearchRotatedSortedArray2(t *testing.T) {
	input := []int{4, 5, 6, 7, 0, 1, 2}
	want := -1
	target := 3
	gut := SearchRotatedSortedArray(input, target)
	if want != gut {
		t.Errorf("gut %d; wanted %d", gut, want)
	}
}

func TestSearchRotatedSortedArray3(t *testing.T) {
	input := []int{1}
	want := -1
	target := 0
	gut := SearchRotatedSortedArray(input, target)
	if want != gut {
		t.Errorf("gut %d; wanted %d", gut, want)
	}
}

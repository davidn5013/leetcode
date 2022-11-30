package twoSum

import (
	"fmt"
	"reflect"
	"testing"
)

// Example 1:
// Input: nums = [2,7,11,15], target = 9
// Output: [0,1]
// Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].
// Example 2:
// Input: nums = [3,2,4], target = 6
// Output: [1,2]
// Example 3:
// Input: nums = [3,3], target = 6
// Output: [0,1]
func TestTwoSum(t *testing.T) {
	var tests = []struct {
		nums   []int
		target int
		want   []int
	}{
		{
			[]int{2, 7, 11, 15}, 9, []int{0, 1},
		},
		{
			[]int{3, 2, 4}, 6, []int{1, 2},
		},
		{
			[]int{3, 3}, 6, []int{0, 1},
		},
		{
			[]int{2222222, 2222222}, 4444444, []int{0, 1},
		},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("%d,%d", tt.nums, tt.target)
		t.Run(testname, func(t *testing.T) {
			ans := TwoSum(tt.nums, tt.target)
			if reflect.DeepEqual(ans, tt.want) != true {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}

func BenchmarkIntMin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TwoSum([]int{2, 7, 11, 15}, 9)
	}
}

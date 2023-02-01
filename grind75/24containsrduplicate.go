package grind75

// LeetCode 217 and grind75 nr 24

// ContainDuplicate true if array has duplicate values
func ContainDuplicate(nums []int) (res bool) {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				return true
			}
		}
	}
	return false
}

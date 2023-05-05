package p1

// removeDuplicates, only on sorted array. Remove dubblet numbers
func removeDuplicates(nums []int) int {
	// Only on sort array. Swap place for ever unequal to prev and
	// short array to length with only uniq numbers
	start, prev := 1, nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] != prev {
			start, prev, nums[start] = start+1, nums[i], nums[i]
		}
	}
	return start
}

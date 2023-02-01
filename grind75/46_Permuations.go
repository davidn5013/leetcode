package grind75

// permute given an array nums of distinct integers, return
// all the possible permutations in any order.
func permute(nums []int) [][]int {
	ret := make([][]int, 0)

	var backtrace func(int)

	backtrace = func(start int) {
		if start == len(nums)-1 {
			ans := make([]int, len(nums))
			copy(ans, nums)
			ret = append(ret, ans)
			return
		}
		for i := start; i < len(nums); i++ {
			nums[i], nums[start] = nums[start], nums[i]
			backtrace(start + 1)
			nums[i], nums[start] = nums[start], nums[i]
		}
	}

	backtrace(0)

	return ret
}

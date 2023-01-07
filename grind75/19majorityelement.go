package grind75

// MajorityElement what number that is the moste in a list
// Example
// []int{2,2,1,1,1,2,2} hase four number of twos result in : 2
func MajorityElement(nums []int) int {
	x, t := nums[0], 1
	for i := 1; i < len(nums); i++ {
		switch {
		case x == nums[i]:
			t++
		case t > 0:
			t--
		default:
			x = nums[i]
			t = 1
		}
	}
	return x
}

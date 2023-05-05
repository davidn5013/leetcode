package p1

func zeroFilledSubarray(nums []int) int64 {
	var r, c int64
	for i := range nums {
		if nums[i] == 0 {
			c++
			r += c
			continue
		}
		c = 0
	}
	return r
}

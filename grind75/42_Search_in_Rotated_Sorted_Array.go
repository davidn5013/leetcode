package grind75

// SearchRotatedSortedArray give the index of a target in a pivot sorted array
func SearchRotatedSortedArray(nums []int, target int) int {
	left, right := 0, len(nums)

	// modified binary search switch left and right to
	// support rotated but sort data
	for {
		middle := (left + right) / 2
		crnt := nums[middle]

		switch {
		case crnt == target:
			return middle // target found
		case left == right:
			return -1 // no target found
		case crnt > nums[left]:
			if crnt > target && nums[left] <= target {
				right = middle
			} else {
				left = middle
			}
		default: // crnt < nums[i]:
			if crnt < target && nums[right-1] >= target {
				left = middle
			} else {
				right = middle
			}

		} // switch

	} // for

} // func

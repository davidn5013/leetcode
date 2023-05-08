package p1

// 11. Container With Most Water

func maxArea(height []int) int {
	l := 0
	r := len(height) - 1
	area := 0

	for l < r {
		// Calculating the max area
		area = max(area, min(height[l], height[r])*(r-l))

		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return area
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

package p1

// 11. Container With Most Water

func maxArea(height []int) int {
	var maxA int
	if len(height) < 2 {
		return 0
	}

	for i := 0; i < len(height); i++ {
		for j := i + 1; j < len(height); j++ {
			minH := min(height[i], height[j])
			vol := minH * (j - i)
			if maxA < vol {
				maxA = vol
			}

		}
	}
	return maxA
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

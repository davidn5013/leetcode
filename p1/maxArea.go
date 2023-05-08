package p1

import "math"

// 11. Container With Most Water

func maxArea(height []int) int {
	var maxA int
	if len(height) < 2 {
		return 0
	}

	for i := 0; i < len(height); i++ {
		for j := i + 1; j < len(height); j++ {
			minH := int(math.Min(float64(height[i]), float64(height[j])))
			maxA = int(math.Max(float64(maxA), float64(minH*(j-i))))
		}
	}
	return maxA
}

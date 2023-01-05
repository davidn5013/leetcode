package tools

// MaxInt bigest number of x, y in int
func MaxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// AbsInt 1=1 -1=1 (by doing -1 * -1 = 1)
func AbsInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

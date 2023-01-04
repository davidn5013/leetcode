package tools

// Max bigest number of x, y in int
func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// Abs 1=1 -1=1 (by doing -1 * -1 = 1)
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

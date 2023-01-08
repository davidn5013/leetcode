package tools

// MaxInt bigest number of x, y in int
func MaxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// MinInt smallest number of x, y in int
func MinInt(x, y int) int {
	if x < y {
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

// EqualMatrix return two are the same [][]int
func EqualMatrix(a, b [][]int) (res bool) {
	if len(a) != len(b) {
		return false
	}
	res = true

	for idx1, arr1 := range a {
		if len(arr1) != len(b[1]) {
			res = false
			break
		}
		arr2 := b[idx1]
		for idx2, value1 := range arr1 {
			value2 := arr2[idx2]
			if value1 != value2 {
				res = false
				goto endloop
			}
		}
	}
endloop:
	return res
}

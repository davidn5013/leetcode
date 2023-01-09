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

// EucliDistSquared return part of euclidean distance (x1 - x2)2 + (y1 - y2)2
/*
The full is euclidean distance is (sqrt) √(x1 - x2)2 + (y1 - y2)2) but this function is meant to
be used to compare the short distance so if the answer is square 2 too big does not matter.

If the correct euclidean distance is need remember to
int(math.Ceil(float64(EucliDistSquared(x,y)))) the result

Example
fmt.Println(EucliDistSquared(0,0,1,3))
// OutPut : 10

log.Println(math.Ceil(math.Sqrt(float64(EucliDistSquared(0, 0, 1, 3)))))
// This is the real distans
// OutPut : 4
*/
func EucliDistSquared(x1, y1, x2, y2 int) int {
	return Square(x1-x2) + Square(y1-y2)

}

// Square return n*n.
func Square(n int) int {
	return n * n
}

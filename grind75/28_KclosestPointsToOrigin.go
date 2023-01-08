package grind75

// Leetcode 973 Grind75 number 28

// Kclosest return distance between two points on the X-Y plane is the Euclidean distance (i.e., √(x1 - x2)2 + (y1 - y2)2).
// k is the how many of the closes point we want
func Kclosest(points [][]int, k int) (res [][]int) {
	// we can't as for more point then we get
	if len(points) > k {
		return nil
	}

	// Testing EculideanDistInt()
	// log.Println(EculideanDistSqrtInt(0, 0, 1, 3))
	// return 10
	// works

	return [][]int{{0}}
}

// EuclideanDistSqrtInt return euclidean distance  √(x1 - x2)2 + (y1 - y2)2).
// Return value before square root
// Exampel
// fmt.Println(EculideanDistSqrtInt(0,0,1,3))
// 10
func EculideanDistSqrtInt(x1, y1, x2, y2 int) int {
	return Square(x1-x2) + Square(y1-y2)

}

// Square return n*n
func Square(n int) int {
	return n * n
}

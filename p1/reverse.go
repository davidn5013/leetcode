package p1

import (
	"log"
	"math"
	"strconv"
)

func reverse(x int) int {
	var (
		res                int
		neg, xStr, revxStr string
	)

	if x < 0 {
		neg = "-"
	}

	// convert to int to string, reverse, restore minus
	// and convert backup to int
	xStr = strconv.Itoa(int(math.Abs(float64(x))))

	for i := len(xStr) - 1; i >= 0; i-- {
		revxStr = revxStr + string(xStr[i])
	}

	revxStr = neg + revxStr
	res, err := strconv.Atoi(revxStr)
	if err != nil {
		log.Fatal(err)
	}

	return res
}

/* All int version

func reverse(x int) int {
	res := 0

	for x != 0 {
		res = res * 10 + x % 10
		x = x / 10
	}

	// leetcode constrans
	if res < math.MinInt32 || res > math.MaxInt32 {
		return 0
	}
	return res
}

*/

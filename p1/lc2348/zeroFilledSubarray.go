package main

import "fmt"

func main() {
	fmt.Println(zeroFilledSubarray([]int{1, 3, 0, 0, 2, 0, 0, 4}))
}

func zeroFilledSubarray(nums []int) int64 {
	var r, c int64
	for i := range nums {
		if nums[i] == 0 {
			c++
			r += c
			continue
		}
		c = 0
	}
	return r
}

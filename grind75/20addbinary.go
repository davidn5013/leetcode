package grind75

import "github.com/davidn5013/leetcode/tools"

// Leetcode 67 grind75 number 20

// AddBinary add too a string of binaries "11"+"1"="100". Arbitrary length
func AddBinary(a string, b string) (res string) {
	return tools.BigIntToBinStr(tools.BinStrAdd(a, b))
}

// package lc14 Longest Common Prefix
package lc14

import (
	"math"
)

// minStr shortest string in string slice
func minStr(strs []string) int {
	res := math.MaxInt
	for _, str := range strs {
		lenstr := len(str)
		if lenstr < res {
			res = lenstr
		}
	}
	return res
}

// LongestCommonPrefix return prefix used in all string
func LongestCommonPrefix(strs []string) string {

	var prefix string

	nrSlices := len(strs)     // number of string slices in []string
	minStrLen := minStr(strs) // len of shortest string in string slice

	for rIdx := 0; rIdx < minStrLen; rIdx++ {
		prefix = strs[0][:rIdx]

		// compare prefix in ever slice
		for sIdx := 0; sIdx < nrSlices; sIdx++ {
			begStr := string(strs[sIdx][:rIdx])
			if begStr != prefix {
				// loop moves prefix on rune to far
				prefix = prefix[:len(prefix)-1]
				// break out of two loops
				goto endloops
			}
		}
	}

endloops:

	return prefix
}

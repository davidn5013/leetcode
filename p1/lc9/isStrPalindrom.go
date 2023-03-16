// package lc9 is LeetCode 0009 IsPalindrome Check if a integer is a palindrome.
package lc9

import (
	"strconv"
)

// isStrPal true if a string is palindrom "anna" == true
func isStrPal(x []byte) bool {
	for i := 0; i < len(x); i++ {
		if x[i] != x[len(x)-1-i] {
			return false
		}

	}
	return true
}

// IsPlaindromeInt true if a integer is a palindrome 121==true
func IsPlaindromeInt(x int) bool {
	if x < 0 {
		return false
	}
	s := strconv.Itoa(x)
	return isStrPal([]byte(s))
}

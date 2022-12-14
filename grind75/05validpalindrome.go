package grind75

import (
	"fmt"
	"strings"
)

// 125. Valid Palindrome
// https://leetcode.com/problems/valid-palindrome/

// IsPalindrome 0125 check if the alfanumerics is palindrome in lowercase returning true if so
func IsPalindrome(s string) bool {
	// Keep alfanumerics
	var sb strings.Builder
	sb.Grow(len(s))
	for i := 0; i < len(s); i++ {
		b := s[i]
		if ('a' <= b && b <= 'z') ||
			('A' <= b && b <= 'Z') ||
			('0' <= b && b <= '9') {
			sb.WriteByte(b)
		}
	}
	testString := sb.String()
	testString = strings.ToLower(testString)

	var revtestString string
	for _, v := range testString {
		revtestString = string(v) + revtestString
	}

	fmt.Println(revtestString)

	return strings.Compare(testString, revtestString) == 0

}

// Package isValid solution to leetcode number 0020
package isValid

// IsValid 0020. An input string is valid if: Open brackets must be closed by
// the same type of brackets. Open brackets must be closed in the correct
// order. Every close bracket has a corresponding open bracket of the same
// type.
func IsValid(s string) bool {
	// Constraints: 1 <= s.length <= 104s consists of parentheses only '()[]{}'.
	if len(s) <= 1 || len(s) >= 10e4 {
		return false
	}
	add := []rune{'(', '[', '{'}
	sub := []rune{')', ']', '}'}
	count := 0
	for _, c := range s {
		for _, a := range add {
			if c == a {
				count++
			}
		}
		for _, a := range sub {
			if c == a {
				count--
			}
		}
	}
	return count == 0
}

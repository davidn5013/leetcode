package grind75

import (
	"log"

	"github.com/davidn5013/leetcode/tools/ds"
)

// isValid solution to leetcode number 0020

// IsValid 0020. An input string is valid if: Open brackets must be closed by
// the same type of brackets. Open brackets must be closed in the correct
// order. Every close bracket has a corresponding open bracket of the same
// type.
func IsValid(s string) bool {
	// Constraints: 1 <= s.length <= 104s consists of parentheses only '()[]{}'.
	if len(s) <= 1 || len(s) >= 10e4 {
		return false
	}

	type void struct{}

	pre := map[rune]void{
		'(': void{},
		'{': void{},
		'[': void{},
	}

	post := map[rune]void{
		')': void{},
		'}': void{},
		']': void{},
	}

	// New RuneStack
	stack := ds.NewRuneStack()
	for _, r := range s {
		// if r == '(' || r == '[' || r == '{' {
		if _, ok := pre[r]; ok {
			stack.Push(r)
			// } else if r == ')' || r == ']' || r == '}' {
		} else if _, ok := post[r]; ok {
			if stack.IsEmpty() {
				return false
			}
			item, _ := stack.Pop()
			if _, ok := pre[item]; !ok {
				log.Printf("%#v\n", pre[item])
				return false
			}
			// if r == ')' && item != '(' {
			// 	return false
			// }
			// if r == ']' && item != '[' {
			// 	return false
			// }
			// if r == '}' && item != '{' {
			// 	return false
			// }
		}
	}
	return stack.IsEmpty()
}

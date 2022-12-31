// Package leetcode isValid solution to leetcode number 0020
package leetcode

import (
	"fmt"
)

// IsValid 0020. An input string is valid if: Open brackets must be closed by
// the same type of brackets. Open brackets must be closed in the correct
// order. Every close bracket has a corresponding open bracket of the same
// type.
func IsValid(s string) bool {
	// Constraints: 1 <= s.length <= 104s consists of parentheses only '()[]{}'.
	if len(s) <= 1 || len(s) >= 10e4 {
		return false
	}

	stack := NewStack()
	for _, r := range s {
		if r == '(' || r == '[' || r == '{' {
			stack.Push(r)
		} else if r == ')' || r == ']' || r == '}' {
			if stack.IsEmpty() {
				return false
			}
			item, _ := stack.Pop()
			if r == ')' && item != '(' {
				return false
			}
			if r == ']' && item != '[' {
				return false
			}
			if r == '}' && item != '{' {
				return false
			}
		}
	}
	return stack.IsEmpty()
}

/***************
 *
 * Rune Stack
 */

// Stack - Rune stack data structure.
type Stack struct {
	items []rune
}

// NewStack creates a new stack.
func NewStack() *Stack {
	return &Stack{
		items: []rune{},
	}
}

// Push adds an item to the top of the stack.
func (s *Stack) Push(item rune) {
	s.items = append(s.items, item)
}

// Pop removes and returns the top item from the stack.
// If the stack is empty, it returns an error.
func (s *Stack) Pop() (rune, error) {
	if len(s.items) == 0 {
		return 0, fmt.Errorf("stack is empty")
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item, nil
}

// IsEmpty returns true if the stack is empty, and false otherwise.
func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

package ds

import "fmt"

/***************
 *
 * Rune Stack
 */

// RuneStack - Rune stack data structure.
type RuneStack struct {
	items []rune
}

// NewRuneStack creates a new stack.
func NewRuneStack() *RuneStack {
	return &RuneStack{
		items: []rune{},
	}
}

// Push adds an item to the top of the stack.
func (s *RuneStack) Push(item rune) {
	s.items = append(s.items, item)
}

// Pop removes and returns the top item from the stack.
// If the stack is empty, it returns an error.
func (s *RuneStack) Pop() (rune, error) {
	if len(s.items) == 0 {
		return 0, fmt.Errorf("stack is empty")
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item, nil
}

// IsEmpty returns true if the stack is empty, and false otherwise.
func (s *RuneStack) IsEmpty() bool {
	return len(s.items) == 0
}

package ds

import (
	"math"
)

// Leetcode 155. Min Stack
// Grind75 Number 38 Min Stackk

// MinStack is a stack that continuously storing minimum value
// Pop somewhat slower then other models because it need to go trough the stack and update min
// Push fast in this module it only a array update and a check for min value
type MinStack struct {
	minIdx int
	stack  []int
}

// NewMinStack create new MinStack
func NewMinStack() MinStack {
	return MinStack{
		minIdx: -1,
		stack:  []int{},
	}
}

// Push new value on stack and maintain smallest value
func (s *MinStack) Push(val int) {
	s.stack = append(s.stack, val)
	if s.minIdx == -1 {
		s.minIdx = 0
	} else if val < s.stack[s.minIdx] {
		s.minIdx = len(s.stack) - 1
	}
	//this.Print()
}

// Pop remove last value on stack
func (s *MinStack) Pop() {
	if s.minIdx == len(s.stack)-1 {
		s.stack[len(s.stack)-1] = math.MaxInt
		for i := 0; i < len(s.stack); i++ {
			if s.stack[i] < s.stack[s.minIdx] {
				s.minIdx = i
			}
		}
	}
	s.stack = s.stack[:len(s.stack)-1]

	//this.Print()
}

// Top peek top of stack
func (s *MinStack) Top() int {
	n := len(s.stack)
	if n > 0 {
		//this.Print()
		return s.stack[n-1]
	}
	return 0
}

// GetMin min value in stack
func (s *MinStack) GetMin() int {
	if s.minIdx >= 0 {
		return s.stack[s.minIdx]
	}
	return 0
}

package grind75

import (
	"fmt"

	ds "github.com/davidn5013/leetcode/tools/ds"
)

func ExampleHasCycle() {
	// Create test TreeNode
	l := &ds.ListNode{Val: 3}
	l.Next = &ds.ListNode{Val: 2}
	loop := l.Next
	l = l.Next
	l.Next = &ds.ListNode{Val: 0}
	l = l.Next
	l.Next = loop
	fmt.Println(HasCycle(l))

	// Create test TreeNode
	l = &ds.ListNode{Val: 3}
	l.Next = &ds.ListNode{Val: 2}
	l = l.Next
	l.Next = &ds.ListNode{Val: 0}
	l = l.Next

	fmt.Println(HasCycle(l))
	// OutPut: true
	// false
}

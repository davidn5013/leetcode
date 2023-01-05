package grind75

import ds "github.com/davidn5013/leetcode/tools/ds"

// HasCycle check for cycle in list returning true if cycle exist
func HasCycle(head *ds.ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	curr, forw := head, head.Next
	for curr != forw {
		if forw == nil || forw.Next == nil {
			return false
		}
		curr = curr.Next
		forw = forw.Next.Next
	}
	return true
}

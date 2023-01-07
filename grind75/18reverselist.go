package grind75

import ds "github.com/davidn5013/leetcode/tools/ds"

// ReverseList 0206 Reverse a linked list 1,2,3,4,5 -> 5,4,3,2,1
// This leadves a empty head behind. You have to nil your pointer to Head
// after this it, look in example below
// Example :
// listOfNodesReversed := ReverseList(listOfNodes)
// listOfNodes = nil
func ReverseList(head *ds.ListNode) *ds.ListNode {
	var prevList *ds.ListNode
	for head != nil {
		// swap next pointer throw the list
		temp := head.Next
		head.Next = prevList
		prevList = head
		head = temp
	}
	return prevList
}

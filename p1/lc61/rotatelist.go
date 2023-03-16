package lc61

// ListNode definition for singly-linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || k < 1 {
		return head
	}

	// move the last k node first

	// get the lenght of the nodelist
	length, tail := 1, head
	for tail.Next != nil {
		tail = tail.Next
		length++
	}

	// rotating nodelist of 5 6 times is the same a rotating 1
	// and if k same as length then we do note need to routate
	k = k % length
	if k == 0 {
		return head
	}

	// find cute of if the list is 1,2,3,4,5 and k it
	// between 3,4 så lenght-k-1 start from 1 = 2 jumps
	// makes cur = 3
	cur := head
	for i := 1; i <= length-k-1; i++ {
		cur = cur.Next
	}

	// cutting of and moving nodes inte the list node
	newHead := cur.Next
	cur.Next = nil
	tail.Next = head

	return newHead
}

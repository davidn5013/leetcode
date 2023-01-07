package grind75

import "github.com/davidn5013/leetcode/tools/ds"

// MiddleNode return pointer to middle of list
func MiddleNode(head *ds.ListNode) *ds.ListNode {
	mid := head
	end := head
	for end.Next != nil && end.Next.Next != nil {
		end = end.Next.Next
		mid = mid.Next
	}
	if end.Next != nil {
		mid = mid.Next
	}
	return mid
}

/*

 func middleNode(head *ListNode) *ListNode {
    var mid = head
    for head != nil && head.Next != nil {
        head, mid = head.Next.Next, mid.Next
    }
    return mid
}

*/

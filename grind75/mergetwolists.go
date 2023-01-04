package leetcode

import "github.com/davidn5013/leetcode/tools/ds"

// 21. Merge Two Sorted Lists
// https://leetcode.com/problems/merge-two-sorted-lists/

// MergeTwoSortedLists Join two sorted list to one. l1+l2=l2
// Warning corrupt l1 and changes l2 to two the new join list
func MergeTwoSortedLists(l1 *ds.ListNode, l2 *ds.ListNode) *ds.ListNode {
	dummy := new(ds.ListNode)

	for node := dummy; l1 != nil || l2 != nil; node = node.Next {
		if l1 == nil {
			node.Next = l2

			break
		} else if l2 == nil {
			node.Next = l1

			break
		} else if l1.Val < l2.Val {
			node.Next = l1

			l1 = l1.Next
		} else {
			node.Next = l2

			l2 = l2.Next
		}
	}

	return dummy.Next
}

// Package leetcode - 21. Merge Two Sorted Lists
// https://leetcode.com/problems/merge-two-sorted-lists/
package leetcode

import (
	"fmt"
	"sort"
)

// MergeTwoSortedLists Join two sorted list to one. l1+l2=l2
// Warning corrupt l1 and changes l2 to two the new join list
func MergeTwoSortedLists(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := new(ListNode)

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

// ListNode Linked list of for int values
type ListNode struct {
	Val  int
	Next *ListNode
}

// New return pointer too new linked list of int values
func New(val int) *ListNode {
	l := ListNode{
		Val:  val,
		Next: nil,
	}
	return &l
}

// Insert a new node of int value in ListNode
func (l *ListNode) Insert(val int) {
	newnode := &ListNode{
		Val:  val,
		Next: nil,
	}

	if l.Next == nil {
		l.Next = newnode
	} else {
		n := l.Next
		for n.Next != nil {
			n = n.Next
		}
		n.Next = newnode
	}
}

// String Stringer for ListNode usage fmt.Println("%s",newLinkelist)
func (l *ListNode) String() (ret string) {
	var (
		nodeCount int
		nextNode  = l
	)

	for {
		ret += fmt.Sprintf("\nNode: %d Value: %d", nodeCount, nextNode.Val)
		if nextNode = nextNode.Next; nextNode == nil {
			break
		} else {
			nodeCount++
		}
	}

	return ret
}

// SortArray Sorts ListNode by converting to array and returning new ListNode
func (l *ListNode) SortArray() {
	var nodeCount int
	var nextNode = l

	// Count nodes
	for {
		if nextNode = nextNode.Next; nextNode == nil {
			break
		} else {
			nodeCount++
		}
	}
	nodeCount++

	sortArray := make([]int, nodeCount)

	var index int
	nextNode = l
	// fill sortArray
	for {
		sortArray[index] = nextNode.Val
		if nextNode = nextNode.Next; nextNode == nil {
			break
		} else if index++; index >= nodeCount {
			break
		}
	}

	// Sort
	sort.Ints(sortArray)

	index = 0
	nextNode = l
	// fill node with sorted array
	for {
		nextNode.Val = sortArray[index]
		if nextNode = nextNode.Next; nextNode == nil {
			break
		} else if index++; index >= nodeCount {
			break
		}
	}
}

// JoinTwoList Join two lists same order list1+list2=list1
func JoinTwoList(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1.Next == nil {
		list1.Next = list2
	} else {
		p := list1.Next
		for p.Next != nil {
			p = p.Next
		}
		p.Next = list2
	}
	return list1
}

// MergeTwoListsArr Join two unsorted ListNodes list1+list2=list2
// sort new list by converting to array
func MergeTwoListsArr(list1 *ListNode, list2 *ListNode) *ListNode {
	res := JoinTwoList(list1, list2)
	res.SortArray()
	return res
}

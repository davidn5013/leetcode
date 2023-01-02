package leetcode

import (
	"testing"
)

func TestMergeTwoSortedList(t *testing.T) {
	l1 := NewListNode(11)
	l2 := NewListNode(12)

	l1.Insert(21)
	l1.Insert(31)
	l1.Insert(41)

	l2.Insert(22)
	l2.Insert(32)
	l2.Insert(42)

	// Sorted: 11,12,21,22,31,32,41,42

	ll := MergeTwoSortedLists(l1, l2)

	t.Log(ll)
}

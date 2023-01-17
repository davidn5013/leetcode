package grind75

import (
	"testing"

	"github.com/davidn5013/leetcode/tools/ds"
)

func TestIsValidBST1(t *testing.T) {
	// [5,4,6,null,null,3,7] false

	node := new(ds.TreeNode)
	node.InsertArrValInTreeNode([]int{5, 4, 6, 0, 0, 3, 7})
	t.Log(node.LevelOrderArr())

	got := IsValidBST(node)
	want := false

	if got != want {
		t.Errorf("got %t; wanted %t", got, want)
		t.Errorf("Value in node %v", node)
	}
}

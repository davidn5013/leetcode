package grind75

import (
	r "math/rand"
	"testing"

	"github.com/davidn5013/leetcode/tools/ds"
)

func TestMaxDepth(t *testing.T) {
	tree := RandomTreeNode(10)
	got := MaxDepth(tree)
	wanted := 10
	if got != wanted {
		t.Errorf("got %d; wanted %d", got, wanted)
	}
}

func RandomTreeNode(depth int) *ds.TreeNode {
	if depth == 0 {
		return nil
	}

	node := &ds.TreeNode{
		Val: r.Intn(200),
	}

	node.Left = RandomTreeNode(depth - 1)
	node.Right = RandomTreeNode(depth - 1)

	return node
}

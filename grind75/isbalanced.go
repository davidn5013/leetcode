package grind75

import (
	t "github.com/davidn5013/leetcode/tools"
	ds "github.com/davidn5013/leetcode/tools/ds"
)

// leetcode 110. Balanced Binary Tree https://leetcode.com/problems/balanced-binary-tree/

// IsBalanced 0110 test if a simple tree node is balanced
func IsBalanced(root *ds.TreeNode) bool {
	if root == nil {
		return true
	}
	var stack []*ds.TreeNode
	stack = append(stack, root)
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if node == nil {
			continue
		}
		leftHeight := getHeight(node.Left)
		rightHeight := getHeight(node.Right)
		if t.AbsInt(leftHeight-rightHeight) > 1 {
			return false
		}
		stack = append(stack, node.Left, node.Right)
	}
	return true
}

func getHeight(node *ds.TreeNode) int {
	if node == nil {
		return 0
	}
	return 1 + t.MaxInt(getHeight(node.Left), getHeight(node.Right))
}

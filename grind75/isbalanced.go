package leetcode

import "github.com/davidn5013/leetcode/tools/ds"

// leetcode 110. Balanced Binary Tree https://leetcode.com/problems/balanced-binary-tree/

// IsBalanced test if a simple tree node is balanced
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
		if abs(leftHeight-rightHeight) > 1 {
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
	return 1 + max(getHeight(node.Left), getHeight(node.Right))
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

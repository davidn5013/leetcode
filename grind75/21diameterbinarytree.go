package grind75

import (
	t "github.com/davidn5013/leetcode/tools"
	ds "github.com/davidn5013/leetcode/tools/ds"
)

// Leetcode 543 Grind75 nr 21

// DiameterOfBinaryTree return the length of the diameter of the tree.
func DiameterOfBinaryTree(root *ds.TreeNode) int {
	if root == nil {
		return 0
	}

	/* Go down the tree recursive finding node == nil (in height)
	   go down left then right tree and save max diameter until nil
	   and go down height from next level plus 1 we can then take max of height
	   And the big of height or diameter is the max diameter */

	maxDia := t.MaxInt(DiameterOfBinaryTree(root.Left), DiameterOfBinaryTree(root.Right))
	maxHeight := height(root.Left) + height(root.Right)

	return t.MaxInt(maxDia, maxHeight)
}

func height(node *ds.TreeNode) int {
	if node == nil {
		return 0
	}
	return 1 + t.MaxInt(height(node.Left), height(node.Right))
}

/*
func DiameterOfBinaryTreeOld(root *ds.TreeNode) int {
	_, res := nodeLenDialoop(root)
	return res
}

func nodeLenDialoop(root *ds.TreeNode) (length, diameter int) {
	if root == nil {
		return
	}

	leftLen, LeftDia := nodeLenDialoop(root.Left)
	rightLen, rightDia := nodeLenDialoop(root.Right)

	length = t.MaxInt(leftLen, rightLen) + 1
	diameter = t.MaxInt(leftLen+rightLen, t.MaxInt(LeftDia, rightDia))
	return
}
*/

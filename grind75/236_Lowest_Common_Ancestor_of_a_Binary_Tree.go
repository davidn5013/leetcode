package grind75

import "github.com/davidn5013/leetcode/tools/ds"

// 236 Lowest Common Ancestor of a Binary Tree
// Grind75 number 46

// lowestCommonAncestor rGiven a binary tree, find the lowest common ancestor (LCA) of two given nodes in the tree.
func lowestCommonAncestor(root, p, q *ds.TreeNode) *ds.TreeNode {
	if root == nil {
		return nil
	}
	if root == p {
		return root
	}
	if root == q {
		return root
	}
	l := lowestCommonAncestor(root.Left, p, q)
	r := lowestCommonAncestor(root.Right, p, q)
	if l != nil && r != nil {
		return root
	}
	if l != nil {
		return l
	}
	if r != nil {
		return r
	}
	return nil
}

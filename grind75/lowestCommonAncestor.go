package grind75

import t "github.com/davidn5013/leetcode/tools/ds"

// 235. Lowest Common Ancestor of a Binary Search Tree
// https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-search-tree/description/

// Lca Lowest Common Ancestor of a Binary Search Tree
func Lca(root, p, q *t.TreeNode) *t.TreeNode {
	if root == nil || root == p || root == q {
		return root
	}
	left, right := Lca(root.Left, p, q), Lca(root.Right, p, q)
	switch {
	case left != nil && right != nil:
		return root
	case left == nil:
		return right
	}
	return left
}

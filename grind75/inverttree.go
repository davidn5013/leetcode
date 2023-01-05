package grind75

import (
	t "github.com/davidn5013/leetcode/tools/ds"
)

// 226. Invert Binary Tree
// https://leetcode.com/problems/invert-binary-tree/

// InvertTree Resurcive Input: root = [2,1,3] Output: [2,3,1]
func InvertTree(root *t.TreeNode) *t.TreeNode {
	if root == nil {
		return nil
	}

	// Swap the left and right child nodes.
	root.Left, root.Right = root.Right, root.Left

	// Recursively invert the left and right subtrees.
	InvertTree(root.Left)
	InvertTree(root.Right)

	return root
}

// InvertTreeIterative Input: root = [2,1,3] Output: [2,3,1]
// Not realy need for this leetcode
func InvertTreeIterative(root *t.TreeNode) *t.TreeNode {
	if root == nil {
		return nil
	}

	queue := []*t.TreeNode{root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		// Swap the left and right child nodes.
		node.Left, node.Right = node.Right, node.Left

		// Add the child nodes to the queue for processing.
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}

	return root
}

/* chat GPT solutions
Can I invert a binary tree in golang an not use recursive?
Yes, you can use an iterative approach to invert a binary tree in Go without using recursion. Here is one way to do it:

InvertBinaryTree inverts the given binary tree.

This function uses a breadth-first traversal of the binary tree, using a queue to store the nodes that need to be processed. It swaps the left and right child nodes of each node and adds the child nodes to the queue. This continues until all nodes in the tree have been processed.

Note that this approach has a time complexity of O(n), where n is the number of nodes in the tree, and a space complexity of O(n), since the queue can hold up to n nodes at any given time.
*/

/*
func InvertBinaryTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	queue := []*TreeNode{root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		// Swap the left and right child nodes.
		node.Left, node.Right = node.Right, node.Left

		// Add the child nodes to the queue for processing.
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}

	return root
}
*/

/* And recursive?

Yes, you can also use a recursive approach to invert a binary tree in Go. Here is one way to do it:

This function uses a recursive depth-first traversal of the binary tree. It first swaps the left and right child nodes of the current node, and then recursively calls itself on the left and right child nodes to invert the subtrees rooted at those nodes. This continues until all nodes in the tree have been processed.

Note that this approach has a time complexity of O(n), where n is the number of nodes in the tree, and a space complexity of O(n), since the function call stack can hold up to n function calls at any given time.
*/

// InvertBinaryTree inverts the given binary tree.
/*
func InvertBinaryTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	// Swap the left and right child nodes.
	root.Left, root.Right = root.Right, root.Left

	// Recursively invert the left and right subtrees.
	InvertBinaryTree(root.Left)
	InvertBinaryTree(root.Right)

	return root
}
*/

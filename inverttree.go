// Package leetcode - 226. Invert Binary Tree
// https://leetcode.com/problems/invert-binary-tree/
package leetcode

import "fmt"

// InvertTree Resurcive Input: root = [2,1,3] Output: [2,3,1]
func InvertTree(root *TreeNode) *TreeNode {
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
func InvertTreeIterative(root *TreeNode) *TreeNode {
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

/**
 * Definition for a binary tree node.
 */

// TreeNode Binary tree storing int
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// NewTreeNode val int, Left, Right
func NewTreeNode() *TreeNode {
	return &TreeNode{}
}

// PrintInOrder prints the values of the nodes in the binary tree in in-order traversal order.
func (t *TreeNode) PrintInOrder() {
	if t == nil {
		return
	}

	fmt.Print(t.Val, " ")
	t.Right.PrintInOrder()
	t.Left.PrintInOrder()
}

// Set value root node
func (t *TreeNode) Set(val int) {
	t.Val = val
}

// Not realy need for this leetcode but this give antoher view and is a stringer example to keep
func (t *TreeNode) String() string {
	if t == nil {
		return " "
	}
	if t.Left == nil && t.Right == nil {
		return fmt.Sprintf("%d", t.Val)
	}
	return fmt.Sprintf("%d %s %s", t.Val, t.Left.String(), t.Left.String())
}

// PrintBreadthFirst prints the values of the nodes in the binary tree in breadth-first order.
// Not realy need for this leetcode but this give antoher view and is a stringer example to keep
func (t *TreeNode) PrintBreadthFirst() (res string) {
	if t == nil {
		return
	}

	queue := []*TreeNode{t}
	for len(queue) > 0 {
		node := queue[len(queue)-1]
		queue = queue[:len(queue)-1]

		res += fmt.Sprint(node.Val, " ")

		// Add the child nodes to the queue for processing.
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
	}
	return res
}

// From Chat GPT

// InsertLeft adds a new node with the given value as the left child of the current node.
// If the current node does not have a left child, it creates a new node.
// TODO Does not work!
func (t *TreeNode) InsertLeft(val int) {
	if t == nil {
		return
	}
	if t.Left == nil {
		t.Left = &TreeNode{Val: val}
	} else {
		t.Left.Val = val
	}
}

// InsertRight adds a new node with the given value as the right child of the current node.
// If the current node does not have a right child, it creates a new node.
// TODO Does not work!
func (t *TreeNode) InsertRight(val int) {
	if t == nil {
		return
	}
	if t.Right == nil {
		t.Right = &TreeNode{Val: val}
	} else {
		t.Right.Val = val
	}
}

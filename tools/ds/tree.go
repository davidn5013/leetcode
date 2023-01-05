package ds

import "fmt"

// TreeNode Binary tree storing int, left and right
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// NewTreeNode return new pointer to TreeNode
func NewTreeNode() *TreeNode {
	return &TreeNode{}
}

// Set value root node
func (t *TreeNode) Set(val int) {
	t.Val = val
}

// PrintTree print the hollow tree
func PrintTree(root *TreeNode) {
	printTreeRecursive(root)
	fmt.Println()
}

func printTreeRecursive(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Printf("%d ", root.Val)
	printTreeRecursive(root.Left)
	printTreeRecursive(root.Right)
}

// PrintInOrder prints the values of the nodes in the binary tree in in-order traversal order.
func (t *TreeNode) PrintInOrder() {
	if t == nil {
		return
	}
	fmt.Print(t.Val, " ")
	t.Left.PrintInOrder()
	t.Right.PrintInOrder()
}

// PrintBreadthFirst prints the values of the nodes in the binary tree in breadth-first order.
// Not realy need for this leetcode but this give antoher view and is a stringer example to keep
func (t *TreeNode) PrintBreadthFirst(root *TreeNode) (res string) {
	if root == nil {
		return
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		res += fmt.Sprint(node.Val, " ")

		// Add the child nodes to the queue for processing.
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}
	return res
}

// Not realy need for this leetcode but this give antoher view and is a stringer example to keep
func (t *TreeNode) String() string {
	if t == nil {
		return "nil"
	}
	return fmt.Sprintf("%v {%v %v}", t.Val, t.Left.String(), t.Right.String())
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

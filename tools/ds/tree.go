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

/* Thanks to https://blog.devgenius.io/the-main-idea-of-implementing-a-binary-tree-with-golang-556fac53ced4
 */

// PreOrder Print tree Data -> left -> right order
// Example
// t.Val=3; t.Left.Val=9; t.Right.Val=20;
// t.Right.Left.Val=15; t.Right.Right.Val=7
// Print as : 3 9 20 15 7
func (t *TreeNode) PreOrder() {
	if t != nil {
		// print root
		fmt.Print(t.Val, " ")
		// print left tree
		t.Left.PreOrder()
		// print right tree
		t.Right.PreOrder()

	}
}

// PostOrder Print tree Left -> Right -> Data
// Example
// t.Val=3; t.Left.Val=9; t.Right.Val=20;
// t.Right.Left.Val=15; t.Right.Right.Val=7
// Print as : 9 15 7 20 3
func (t *TreeNode) PostOrder() {
	if t != nil {
		// print left tree
		t.Left.PostOrder()
		// print right tree
		t.Right.PostOrder()
		// print root
		fmt.Print(t.Val, " ")
	}
}

// MidOrder Print Tree Left -> Data -> Right
// Example
// t.Val=3; t.Left.Val=9; t.Right.Val=20;
// t.Right.Left.Val=15; t.Right.Right.Val=7
// Print as : 9 3 15 20 7
func (t *TreeNode) MidOrder() {
	if t != nil {
		// print left tree
		t.Left.MidOrder()
		// print root
		fmt.Print(t.Val, " ")
		// print right tree
		t.Right.MidOrder()
	}
}

/* From Chat GPT
 */

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

// LevelOrderArr return TreeNode as [][]int
func (t *TreeNode) LevelOrderArr() (res [][]int) {
	// t.Val == 00 && t.Left .. is only so we not get any value back on a empty
	// if this second part after || (.. is remove we get [][]int{{0}} back on empty
	// This is for compliance to LeetCod
	if t.Val == 0 && t.Left == nil && t.Right == nil {
		return [][]int{}
	}

	var (
		queue = []*TreeNode{t} // hold this levels node pointers
	)

	// Loop until there is no more queue. In start the queue only has the root node
	// but in the inner loop we store the data and queue up the new nodes then we need
	// to drop the part of queue that run in the inner loop and start over
	for len(queue) > 0 {

		/* we need to store how long the queue is now so we not over run in
		   the inner loop while we are append to the queue */

		curLevel := []int{} // Temp store for result data,inner part of array, clear
		sz := len(queue)    // length of part of queue to run

		// inner loop first run root node, next run left,right then wider and wider
		for i := 0; i < sz; i++ { // running store queue so far
			cur := queue[i]
			curLevel = append(curLevel, cur.Val) // adding data to temp then to result
			if cur.Left != nil {
				queue = append(queue, cur.Left) // queue for next run of inner loop
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
		}
		queue = queue[sz:]          // drop part of the queue we just ran throw
		res = append(res, curLevel) // We we are append temp array to the main matrix [][]int{[]int{}} if you want
	}
	return res
}

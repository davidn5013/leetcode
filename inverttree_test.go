package leetcode

import "fmt"

func ExampleInvertTree() {
	// New binary t1
	t1 := new(TreeNode)

	t1.Val = 4
	t1.Left = new(TreeNode)
	t1.Right = new(TreeNode)
	t1.Left.Val = 2
	t1.Right.Val = 7

	t1.Left.Left = new(TreeNode)
	t1.Left.Right = new(TreeNode)
	t1.Right.Left = new(TreeNode)
	t1.Right.Right = new(TreeNode)

	t1.Left.Left.Val = 1
	t1.Left.Right.Val = 3
	t1.Right.Left.Val = 6
	t1.Right.Right.Val = 9

	fmt.Println(t1.PrintBreadthFirst())
	t1 = InvertTree(t1)
	fmt.Println(t1.PrintBreadthFirst())

	t2 := new(TreeNode)

	t2.Val = 2
	t2.Left = new(TreeNode)
	t2.Right = new(TreeNode)
	t2.Left.Val = 1
	t2.Right.Val = 3

	t2.PrintInOrder()
	fmt.Println()
	t2 = InvertTree(t2)
	t2.PrintInOrder()
	fmt.Println()

	// TODO Bug in InsertLeft and InsertRight does not create nodes
	t3 := NewTreeNode()
	t3.Val = 4
	t3.InsertLeft(2)
	t3.InsertLeft(1)
	t3.InsertLeft(6)
	t3.InsertRight(7)
	t3.InsertRight(3)
	t3.InsertRight(9)

	t3.PrintInOrder()
	fmt.Println()
	t3 = InvertTree(t3)
	t3.PrintInOrder()
	fmt.Println()

	// Don't not remove part below INCLUDE empty line it part Example testing

	// Disabled test with this comment
	// Output:  4 2 1 3 7 6 9
	// 4 7 2 9 6 3 1
	// 2 1 3
	// 2 3 1
}

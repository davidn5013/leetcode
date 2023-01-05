package grind75

import (
	"fmt"

	ds "github.com/davidn5013/leetcode/tools/ds"
)

func ExampleIsBalanced() {
	// Create test TreeNode
	root := &ds.TreeNode{Val: 1}
	root.Left = &ds.TreeNode{Val: 2}
	root.Right = &ds.TreeNode{Val: 3}
	root.Left.Left = &ds.TreeNode{Val: 4}
	root.Left.Right = &ds.TreeNode{Val: 5}
	fmt.Println(IsBalanced(root))
	root.Left.Right.Right = &ds.TreeNode{Val: 10}
	fmt.Println(IsBalanced(root))
	// Output:  true
	// false
}

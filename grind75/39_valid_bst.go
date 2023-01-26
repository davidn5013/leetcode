package grind75

import (
	"github.com/davidn5013/leetcode/tools/ds"
)

// Leetcode 98 Validate that a binary three is binary search tree
// Grind75 number 39

// BUG fail on leetcode setting [5,4,6,null,null,3,7] false but not on unity test

// IsValidBST valid binary tree as binary search tree
func IsValidBST(root *ds.TreeNode) bool {
	var (
		parVal  int                       // queue of the node values from last nodes
		nodeS   = make([]*ds.TreeNode, 1) // node queue
		isValid bool                      // start false
	)

	parVal = root.Val // start
	nodeS[0] = root   // start

	// in current node get children node values and check
	// drop node value store children value
	// return new node list, new values that the childre w had (new parVal)
	// and if the node was not valid
	for len(nodeS) > 0 {
		nodeS, parVal, isValid = getChildNodeValues(nodeS, parVal)

		if !isValid {
			break
		}
	}

	return isValid
}

func getChildNodeValues(nodeS []*ds.TreeNode, compVal int) (nextNodeS []*ds.TreeNode, nextCompVal int, isValid bool) {
	var leftVal, rightVal int
	isValid = true

	for len(nodeS) > 0 {
		// poping from queue
		node := nodeS[0]
		nodeS = nodeS[1:]

		if node.Left != nil {
			leftVal = node.Left.Val
			leftnode := node.Left

			nextNodeS = append(nextNodeS, leftnode)

			if leftVal >= compVal {
				isValid = false
				goto endloop
			}
		}

		if node.Right != nil {
			rightVal = node.Right.Val
			rightnode := node.Right

			nextNodeS = append(nextNodeS, rightnode)

			if rightVal <= compVal {
				isValid = false
				goto endloop
			}
			compVal = rightVal
		}

	}
endloop:
	nextCompVal = compVal
	return nextNodeS, nextCompVal, isValid
}

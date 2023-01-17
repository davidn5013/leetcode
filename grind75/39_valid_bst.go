package grind75

import (
	"log"

	"github.com/davidn5013/leetcode/tools/ds"
)

// Leetcode 98 Validate that a binary three is binary search tree
// Grind75 number 39

// IsValidBST valid binary tree as binary search tree
func IsValidBST(root *ds.TreeNode) bool {
	var (
		// BUG We are not storing the max value last used not a list of node values
		// nodes does not only need to be bigger then there parent they need to be
		// bigger then all previews nodes. So we to store only current max
		parVal  = make([]int, 1)          // queue of the node values from last nodes
		nodeS   = make([]*ds.TreeNode, 1) // node queue
		isValid bool                      // start false
	)

	parVal[0] = root.Val // start
	nodeS[0] = root      // start

	// in current node get children node values and check
	// drop node value store children value
	// return new node list, new values that the childre w had (new parVal)
	// and if the node was not valid
	for len(nodeS) > 0 {
		nodeS, parVal, isValid = getChildNodeValues(nodeS, parVal)

		// nodes most be same length as parval, my be use struct here
		// assert len node==len parval
		if len(nodeS) != len(parVal) {
			log.Fatalf("Miss match len(nodes)=%d and len(parval)=%d\n", len(nodeS), len(parVal))
		}

		if !isValid {
			break
		}
	}

	return isValid
}

func getChildNodeValues(nodeS []*ds.TreeNode, compValues []int) (nextNodeS []*ds.TreeNode, nextCompVal []int, isValid bool) {
	var leftVal, rightVal int
	isValid = true

	for len(nodeS) > 0 {
		// poping from queue
		node := nodeS[0]
		nodeS = nodeS[1:]
		compVal := compValues[0]
		compValues = compValues[1:]

		if node.Left != nil {
			leftVal = node.Left.Val
			leftnode := node.Left

			nextCompVal = append(nextCompVal, leftVal)
			nextNodeS = append(nextNodeS, leftnode)

			if leftVal >= compVal {
				isValid = false
				break
			}
		}

		if node.Right != nil {
			rightVal = node.Right.Val
			rightnode := node.Right

			nextCompVal = append(nextCompVal, rightVal)
			nextNodeS = append(nextNodeS, rightnode)

			if rightVal <= compVal {
				isValid = false
				break
			}
		}

	}
	return nextNodeS, nextCompVal, isValid
}

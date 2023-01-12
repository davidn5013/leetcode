package grind75

import (
	"github.com/davidn5013/leetcode/tools/ds"
)

// Leetcode 102. Binary Tree Level Order Traversal, Grind75 31

// LevelOrder return TreeNode as [][]int
func LevelOrder(t *ds.TreeNode) (res [][]int) {
	// t.Val == 00 && t.Left .. is only so we not get any value back on a empty
	// if this second part after || (.. is remove we get [][]int{{0}} back on empty
	// This is for compliance to LeetCod
	if t.Val == 0 && t.Left == nil && t.Right == nil {
		return [][]int{}
	}

	var (
		queue = []*ds.TreeNode{t} // hold this levels node pointers
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

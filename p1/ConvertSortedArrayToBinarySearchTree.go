package p1

import (
	"fmt"

	"github.com/emirpasic/gods/trees/btree"
)

func treeTest() {
	tree := btree.NewWithIntComparator(3) // empty (keys are of type int)

	//nums := []int{-10, -3, 0, 5, 9}
	// input need to be
	nums := []int{0, -10, 5, -3, 9}

	for i := 0; i < len(nums); i++ {
		tree.Put(i, nums[i])
	}

	fmt.Println(tree)
	fmt.Println(tree.Keys(), tree.Values())
}

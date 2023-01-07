package grind75

import (
	t "github.com/davidn5013/leetcode/tools"
	"github.com/davidn5013/leetcode/tools/ds"
)

func MaxDepth(root *ds.TreeNode) int {
	if root == nil {
		return 0
	}
	maxHeight := t.MaxInt(MaxDepth(root.Left)+1, MaxDepth(root.Right)+1)
	return maxHeight
}

/*

func MaxDepth(root *TreeNode) int {
    return Dfs(root, 1)
}

func Dfs(root *TreeNode, lvl int) int {
    if root == nil {
        return 0
    }
    l := dfs(root.Left, lvl+1)
    r := dfs(root.Right, lvl+1)
    if l > lvl && l >= r {
        return l
    }
    if r > lvl && r > l {
        return r
    }
    return lvl
}

}
*/

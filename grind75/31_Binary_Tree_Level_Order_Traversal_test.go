package grind75

import (
	"reflect"
	"testing"

	"github.com/davidn5013/leetcode/tools/ds"
)

func TestLevelOrder(t *testing.T) {
	// Input: root = {3,9,20,null,null,15,7}

	// Two version of the same TreeNode just to compare
	// how to insert new data, very lite different but a version is easier
	tree := ds.NewTreeNode()
	tree.Set(3)
	tree.InsertLeft(9)
	tree.InsertRight(20)
	right := tree.Right
	right.InsertLeft(15)
	right.InsertRight(7)

	// Same tree different definition method
	tree2 := &ds.TreeNode{}
	tree2.Val = 3
	tree2.Left = &ds.TreeNode{Val: 9}
	tree2.Right = &ds.TreeNode{Val: 20, Right: &ds.TreeNode{}, Left: &ds.TreeNode{}}
	tree2.Right.Right.Val = 7
	tree2.Right.Left.Val = 15

	// Test
	got1 := LevelOrder(tree)
	got2 := LevelOrder(tree2)
	want := [][]int{{3}, {9, 20}, {15, 7}}

	if !reflect.DeepEqual(got1, want) {
		t.Errorf("tree got %v; wanted %v", got1, want)
	}

	if !reflect.DeepEqual(got2, want) {
		t.Errorf("tree got %v; wanted %v", got2, want)
	}
}

func TestLevelOrder2(t *testing.T) {
	tree := ds.NewTreeNode()
	tree.Set(1)
	got := LevelOrder(tree)
	want := [][]int{{1}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("tree got %v; wanted %v", got, want)
	}
}

func TestLevelOrder3(t *testing.T) {
	tree := ds.NewTreeNode()
	got := LevelOrder(tree)
	want := [][]int{}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("tree got %v; wanted %v", got, want)
	}
}

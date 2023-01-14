package grind75

import (
	"fmt"
	"testing"
)

func ExampleCloneGraph() {
	// Input: adjList = [[2,4],[1,3],[2,4],[1,3]]
	// Output: [[2,4],[1,3],[2,4],[1,3]]

	// Clonked declaration with out methods...
	var n [8]Node

	for i := range n {
		n[i].Val = i + 2
	}

	var m Node
	m.Neighbors = []*Node{
		&n[0],
		&n[1],
		&n[2],
		&n[3],
		&n[4],
		&n[5],
		&n[6],
		&n[7],
	}

	// fmt.Println(m)
	o := CloneGraph(&m)
	for i := range o.Neighbors {
		fmt.Printf("%d ", o.Neighbors[i].Val)
	}
	fmt.Println()

	// OutPut: 2 3 4 5 6 7 8 9
}

func BenchmarkCloneGraph(b *testing.B) {
	// Input: adjList = [[2,4],[1,3],[2,4],[1,3]]
	// Output: [[2,4],[1,3],[2,4],[1,3]]

	// Clonked declaration with out methods...
	var n [8]Node

	for i := range n {
		n[i].Val = i + 2
	}

	var m Node
	m.Neighbors = []*Node{
		&n[0],
		&n[1],
		&n[2],
		&n[3],
		&n[4],
		&n[5],
		&n[6],
		&n[7],
	}

	for i := 0; i < b.N; i++ {
		CloneGraph(&m)
	}
}

func BenchmarkCloneGraph2(b *testing.B) {
	// Input: adjList = [[2,4],[1,3],[2,4],[1,3]]
	// Output: [[2,4],[1,3],[2,4],[1,3]]

	// Clonked declaration with out methods...
	var n [8]Node

	for i := range n {
		n[i].Val = i + 2
	}

	var m Node
	m.Neighbors = []*Node{
		&n[0],
		&n[1],
		&n[2],
		&n[3],
		&n[4],
		&n[5],
		&n[6],
		&n[7],
	}

	for i := 0; i < b.N; i++ {
		cloneGraph2(&m)
	}
}

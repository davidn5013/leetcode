package grind75

// Leetcode 133 Grind75 number 32 Clone Graph

/*
Test case format:

For simplicity, each node's value is the same as the node's index (1-indexed). For example, the
first node with val == 1, the second node with val == 2, and so on. The graph is represented in
the test case using an adjacency list.

An adjacency list is a collection of unordered lists used to represent a finite graph. Each list
describes the set of neighbors of a node in the graph.

The given node will always be the first node with val = 1. You must return the copy of the given
node as a reference to the cloned graph.
*/

// Node Test case graph. Using adjacency list is a collection of
// unordered lists used to represent a finite graph.
type Node struct {
	Val       int
	Neighbors []*Node
}

// CloneGraph return a deep copy (clone) of the graph.
func CloneGraph(node *Node) (res *Node) {
	visited := make(map[*Node]*Node)
	return dfsClone(node, visited)
}

// clones graph using a map stored in caller scope.
func dfsClone(node *Node, visited map[*Node]*Node) *Node {
	// End of node recursion
	if node == nil {
		return nil
	}
	// Using callers map to skip visited nodes
	if val, ok := visited[node]; ok {
		return val
	}

	// Copy part
	cloneNode := &Node{
		Val:       node.Val,
		Neighbors: make([]*Node, len(node.Neighbors)),
	}
	visited[node] = cloneNode

	// Recursive loop next copy using visited map as argument
	// This is method makes it possible store what done of the recursion
	for i := 0; i < len(node.Neighbors); i++ {
		cloneNode.Neighbors[i] = dfsClone(node.Neighbors[i], visited)
	}
	return cloneNode
}

/*
This function uses a queue to traverse through each node in the graph and a map to keep track
of the nodes that have already been visited. For each unvisited neighbor, it creates a new
node with the same value and adds it to the queue to be visited later. It also adds the newly
created node to the current node's list of neighbors, ensuring that the copy is a complete
replica of the original graph.

This bfs has been us with the same methodology before in this Grind 75 package for Binary
Trees and List so no real deep documentation is written here. It a queue that replace the
function recursion an sins there less overhead in loops then function calls it faster.

I think it easier to understand to. Maybe that's just me.
*/

// clonesGraph2 creates a deep copy of a Node using a breadth-first search (BFS) algorithm:
func cloneGraph2(node *Node) *Node {
	if node == nil {
		return nil
	}
	visited := map[*Node]*Node{}
	queue := []*Node{node}
	visited[node] = &Node{Val: node.Val}
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		for _, neighbor := range curr.Neighbors {
			if _, ok := visited[neighbor]; !ok {
				visited[neighbor] = &Node{Val: neighbor.Val}
				queue = append(queue, neighbor)
			}
			visited[curr].Neighbors = append(
				visited[curr].Neighbors, visited[neighbor])
		}
	}
	return visited[node]
}

/*
package main

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	visited := make(map[*Node]*Node)
	return dfsClone(node, visited)
}

func dfsClone(node *Node, visited map[*Node]*Node) *Node {
	if node == nil {
		return nil
	}
	if val, ok := visited[node]; ok {
		return val
	}
	cloneNode := &Node{
		Val:       node.Val,
		Neighbors: make([]*Node, len(node.Neighbors)),
	}
	visited[node] = cloneNode
	for i := 0; i < len(node.Neighbors); i++ {
		cloneNode.Neighbors[i] = dfsClone(node.Neighbors[i], visited)
	}
	return cloneNode
}

*/

/*

func cloneGraph(node *Node) *Node {
    if node == nil {
        return nil
    }
    mapping := make(map[int]*Node)
    clone(node, mapping)
    return mapping[node.Val]
}

func clone(node *Node, mapping map[int]*Node) *Node {
    if node == nil {
        return nil
    }

    if _, exists := mapping[node.Val]; !exists {
        mapping[node.Val] = &Node{Val: node.Val}
    }

    clonedNode := mapping[node.Val]

    for _, option := range node.Neighbors {
        if premapped, exists := mapping[option.Val]; exists {
            clonedNode.Neighbors = append(clonedNode.Neighbors, premapped)
        } else {
            clonedNeighbor := clone(option, mapping)
            clonedNode.Neighbors = append(clonedNode.Neighbors, clonedNeighbor)
        }
    }

    return clonedNode
}

*/

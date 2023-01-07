package grind75

import (
	"fmt"
	"log"
	r "math/rand"
	"testing"

	"github.com/davidn5013/leetcode/tools/ds"
)

// RandomLinkedList() return a linked list of 10 nodes
// No seed so value are always the same
// Exemple:
// Node: 0 Value: 81
// Node: 1 Value: 87
// Node: 2 Value: 47
// Node: 3 Value: 100
// Node: 4 Value: 81
// Node: 5 Value: 118
// Node: 6 Value: 25
// Node: 7 Value: 140
// Node: 8 Value: 50
// Node: 9 Value: 56
// Node: 10 Value: 100
func RandomLinkedList() *ds.ListNode {

	// Fill a list with 10 random values
	head := &ds.ListNode{Val: r.Intn(200)}
	node := head
	for i := 0; i < 9; i++ {
		node.Insert(r.Intn(200))
		node = node.Next
	}

	// Set the value of the fourth node
	node = head
	for i := 0; i < 3; i++ {
		node = node.Next
	}
	node.Set(100)

	// Get the value of the sixth node
	node = head
	for i := 0; i < 5; i++ {
		node = node.Next
	}
	fmt.Println(node.Get())

	// Insert a new node with value 50 after the eighth node
	node = head
	for i := 0; i < 7; i++ {
		node = node.Next
	}
	node.Insert(50)

	// Get the head node of the list
	log.Println(head)

	return head
}

func TestMiddleNode(t *testing.T) {
	n := RandomLinkedList()
	got := n.MiddleNode().Val
	// NOT RandomLinkedList has no seed so node number 5 is always 118
	// This problem failes compiled with new compiler or other platform
	// where rand.Intn gives another value. It work for now for test MiddleNode
	want := 118 // Can be failing spot
	if got != want {
		t.Errorf("got %d; wanted %d", got, want)
	}
}

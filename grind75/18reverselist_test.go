package grind75

import (
	"fmt"
	"log"

	ds "github.com/davidn5013/leetcode/tools/ds"
)

func ExampleReverseList() {
	l := ds.NewListNode(1)
	l.Next = ds.NewListNode(2)
	l.Next.Next = ds.NewListNode(3)

	nl := ReverseList(l)
	l = nil
	if l != nil {
		log.Println(l.Val, l.Next)
	}
	fmt.Println(nl)

	// Output:
	// Node: 0 Value: 3
	// Node: 1 Value: 2
	// Node: 2 Value: 1
}

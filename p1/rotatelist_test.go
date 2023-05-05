package p1

import (
	"fmt"
)

// Test rotateRight
func Example_rotateRight() {
	list := new(ListNode)

	head := list
	list.Val = 1
	list.Next = new(ListNode)
	current := list.Next
	current.Val = 2
	current.Next = new(ListNode)
	current = current.Next
	current.Val = 3
	current.Next = new(ListNode)
	current = current.Next
	current.Val = 4
	current.Next = new(ListNode)
	current = current.Next
	current.Val = 5

	current = head

	// for current != nil {
	// 	fmt.Println(current.Val)
	// 	current = current.Next
	// }

	// fmt.Printf("%#v\n", head)
	current = rotateRight(head, 2)
	// fmt.Printf("%#v\n", current)

	for current != nil {
		fmt.Println(current.Val)
		current = current.Next
	}

	// OutPut: 4
	// 5
	// 1
	// 2
	// 3
}

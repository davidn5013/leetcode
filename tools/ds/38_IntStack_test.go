package ds

import "fmt"

func ExampleMinStack() {

	s := NewMinStack()
	s.Push(-2)
	s.Push(0)
	s.Push(-3)
	fmt.Println(s.GetMin())
	s.Pop()
	fmt.Println(s.Top())
	fmt.Println(s.GetMin())

	fmt.Print(s)
	// OutPut:
	// -3
	// 0
	// -2
	// {0 [-2 0]}
}

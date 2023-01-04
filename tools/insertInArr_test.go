package tools

import "fmt"

func ExampleInsertIntAt() {
	a := []int{1, 2, 3}
	fmt.Println("Array before:", a)

	res, err := InsertIntAt(a, 0, []int{51, 56, 57})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("After insert tre int in pos 0:", res)

	// OutPut: Array before: [1 2 3]
	// After insert tre int in pos 0: [51 56 57 1 2 3]
}

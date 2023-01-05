package grind75

import "fmt"

func ExampleCanConstruct() {
	ransom := "Hello"
	magazine := "Hi there lolo"
	fmt.Printf("Value: %#v\n", CanConstruct(ransom, magazine))
	// OutPut: Value: true
}

package p1

import "fmt"

func ExampleLongestCommonPrefix() {
	strs := []string{"flower", "flower", "flower", "flower"}
	fmt.Println(LongestCommonPrefix(strs))
	// OutPut: flower
}

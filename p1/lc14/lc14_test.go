package lc14

import "fmt"

func ExampleLongestCommonPrefix() {
	strs := []string{"predog", "preracecar", "precar"}
	fmt.Println(LongestCommonPrefix(strs))
	// OutPut: fl
}

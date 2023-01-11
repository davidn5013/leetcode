package tools

import "fmt"

func ExampleBinStrAdd() {
	one := BinStrAdd("11101", "10")
	fmt.Println(BigIntToBinStr(one))
	two := BinStrAdd(BigIntToBinStr(one), "1")
	fmt.Println(BigIntToBinStr(two))

	// OutPut:
	// 11111
	// 100000
}

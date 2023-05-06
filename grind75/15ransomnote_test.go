package grind75

import (
	"fmt"
	"testing"
)

func Example_canConstruct() {
	fmt.Printf("%#v ", canConstruct("Hello", "Hi there lolo"))
	fmt.Printf("%#v ", canConstruct("HELLO", "Hi there lolo"))
	fmt.Printf("%#v ", canConstruct("abcdef", "Hi there lolo"))
	fmt.Printf("%#v ", canConstruct("Hello", "tere lolo"))
	fmt.Printf("%#v ", canConstruct2("Hello", "Hi there lolo"))
	fmt.Printf("%#v ", canConstruct2("HELLO", "Hi there lolo"))
	fmt.Printf("%#v ", canConstruct2("abcdef", "Hi there lolo"))
	fmt.Printf("%#v ", canConstruct2("Hello", "tere lolo"))
	fmt.Printf("%#v ", canConstruct3("Hello", "Hi there lolo"))
	fmt.Printf("%#v ", canConstruct3("HELLO", "Hi there lolo"))
	fmt.Printf("%#v ", canConstruct3("abcdef", "Hi there lolo"))
	fmt.Printf("%#v ", canConstruct3("Hello", "tere lolo"))
	// OutPut: true true false false true true false false true true false false
}

func Example_canConstruct2() {
	ransom := "Hello"
	magazine := "Hi there lolo"
	fmt.Printf("Value: %#v\n", canConstruct2(ransom, magazine))
	// OutPut: Value: true
}

func Benchmark_canConstruct(b *testing.B) {
	b.Run("using byte array", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			ransom := "Hello"
			magazine := "Hi there lolo"
			canConstruct(ransom, magazine)
		}
	})
	b.Run("using strings", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			ransom := "Hello"
			magazine := "Hi there lolo"
			canConstruct2(ransom, magazine)
		}
	})
	b.Run("using strings", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			ransom := "Hello"
			magazine := "Hi there lolo"
			canConstruct3(ransom, magazine)
		}
	})
}

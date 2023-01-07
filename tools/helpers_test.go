package tools

import (
	"fmt"
	"log"
	"math/rand"
	"testing"
)

func ExampleInsertIntAt() {
	a := []int{1, 2, 3}
	fmt.Println("Array before:", a)

	res, err := InsertIntAt(a, 0, []int{51, 56, 57})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("After insert 3 int in pos 0:", res)

	// OutPut: Array before: [1 2 3]
	// After insert 3 int in pos 0: [51 56 57 1 2 3]
}

func TestInsertAtInRune(t *testing.T) {
	a := []rune{'a', 'b', 'd'}
	fmt.Println("Array before:", a)

	res, err := InsertAt(a, 0, []rune{'c'})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("After insert on rune pos 3:", res)

	// OutPut: Array before: ['a' 'b' 'd']
	// After insert on rune pos 3: ['a' 'b' 'c' 'd']
}

// Yes InsertIntAt() has some small performance gain compare to InsertAt() so in a tight loop it
// good to use InsertIntAt() but everywhere else you can just use the generic InsertAt()

func BenchmarkInsertIntAt(b *testing.B) {
	slice := rand.Perm(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		newslice, err := InsertIntAt(slice, 500, []int{66666})
		if err != nil || len(slice)+1 != len(newslice) {
			log.Printf("Error insert int in slice new or wrong len size %d %d", len(slice), len(newslice))
		}
	}
}

func BenchmarkInsertAt(b *testing.B) {
	slice := rand.Perm(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		newslice, err := InsertAt(slice, 500, []int{66666})
		if err != nil || len(slice)+1 != len(newslice) {
			log.Printf("Error insert int in slice new or wrong len size %d %d", len(slice), len(newslice))
		}
	}
}

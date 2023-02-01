package grind75

import (
	"log"
	"testing"
)

func TestMerge(t *testing.T) {
	log.Println(merge([][]int{{1, 3}, {2, 6}, {8, 19}, {9, 18}}))
}

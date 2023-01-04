package leetcode

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestMaxProfit(t *testing.T) {
	prices := []int{7, 1, 5, 3, 6, 4}
	want := 5
	ans := MaxProfit(prices)
	if ans != want {
		t.Errorf("got %d, want %d", ans, want)
	}
}

func TestMaxProfitMulti(t *testing.T) {
	var tests = []struct {
		prices []int
		want   int
	}{
		{
			prices: []int{7, 1, 5, 3, 6, 4},
			want:   5,
		},
		{
			prices: []int{2, 10, 4, 3, 1, 8, 9, 11},
			want:   10,
		},
		{
			prices: []int{7, 6, 4, 3, 1},
			want:   0,
		},
		{
			prices: []int{1, 2, 3, 4, 5},
			want:   4,
		},
		{
			prices: []int{7, 6},
			want:   0,
		},
		{
			prices: []int{7},
			want:   0,
		},
		{
			prices: []int{1, 2},
			want:   1,
		},
		{
			prices: []int{2, 4, 1},
			want:   2,
		},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v", tt.prices)
		t.Run(testname, func(t *testing.T) {
			ans := MaxProfit(tt.prices)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}

func TestMaxProfitBigTest(t *testing.T) {
	b, err := os.ReadFile("testdata/stockprice_testfile.txt")
	if err != nil {
		log.Fatal(err)
	}
	fields := strings.Split(string(b), ",")
	prices := make([]int, 10000)
	for _, cell := range fields {
		i, _ := strconv.Atoi(cell)
		// if err != nil {
		// 	log.Fatalf("Err converting %s ", cell)
		// }
		prices = append(prices, i)
	}
	want := 999
	ans := MaxProfit(prices)
	if ans != want {
		t.Errorf("Wrong answer, want %d", want)
	}
}

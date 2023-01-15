package grind75

import "testing"

/*
Input: coins = [1,2,5], amount = 11
Output: 3
Explanation: 11 = 5 + 5 + 1

Example 2:

Input: coins = [2], amount = 3
Output: -1

Example 3:

Input: coins = [1], amount = 0
Output: 0
*/

func TestCoinChange(t *testing.T) {
	got := CoinChange([]int{1, 2, 5}, 11)
	want := 3
	if got != want {
		t.Errorf("Testing []int{1, 2, 5}, 11 got %d ; wanted %d", got, want)
	}
}

func TestCoinChange2(t *testing.T) {
	got := CoinChange([]int{2}, 3)
	want := -1
	if got != want {
		t.Errorf("Testing []int{2}, 3 got %d ; wanted %d", got, want)
	}
}

func TestCoinChang3(t *testing.T) {
	got := CoinChange([]int{1}, 0)
	want := 0
	if got != want {
		t.Errorf("Testing []int{1}, 0 got %d ; wanted %d", got, want)
	}
}

func TestCoinChang4(t *testing.T) {
	got := CoinChange([]int{186, 419, 83, 408}, 6249)
	want := 20
	if got != want {
		t.Errorf("Testing []int{1}, 0 got %d ; wanted %d", got, want)
	}
}

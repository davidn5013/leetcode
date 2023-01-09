package grind75

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	inp := "pwwkew"
	want := 3
	got := lengthOfLongestSubstring(inp)
	if got != want {
		t.Errorf("got %d; wanted %d", got, want)
	}
}

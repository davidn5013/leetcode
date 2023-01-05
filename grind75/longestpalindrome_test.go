package grind75

import (
	"reflect"
	"testing"
)

func TestSplitOnLetter(t *testing.T) {
	inp := "abccccdd"
	want := []string{"a", "b", "cccc", "dd"}
	res := SplitOnLetter(inp)
	if !reflect.DeepEqual(want, res) {
		t.Errorf("SplitOnLetter got = %#v; want %#v", res, want)
	}
}

func TestLongestPalindrom(t *testing.T) {
	got := LongestPalindrom("abccccdd")
	want := 7
	if got != want {
		t.Errorf("LongestPlaindrom got = %d; want %d", got, want)
	}
	got = LongestPalindrom("ccc")
	want = 3
	if got != want {
		t.Errorf("LongestPlaindrom got = %d; want %d", got, want)
	}
}

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

func TestLongestPalindrome(t *testing.T) {
	got := LongestPalindrome("abccccdd")
	want := 7
	if got != want {
		t.Errorf("LongestPlaindrom1 got = %d; want %d", got, want)
	}
	got = LongestPalindrome("ccc")
	want = 3
	if got != want {
		t.Errorf("LongestPlaindrom2 got = %d; want %d", got, want)
	}

	got = LongestPalindrome("ababababa")
	want = 9
	if got != want {
		t.Errorf("LongestPlaindrom3 got = %d; want %d", got, want)
	}
}

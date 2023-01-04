package leetcode

import (
	"sort"
	"strings"
)

// 242. Valid Anagram

// IsAnagram Two word have same letters but diffent order.
// "nagaram" is anagram of "anagram".
func IsAnagram(s, t string) bool {
	if len(s) != len(t) {
		return false
	}
	s = strings.ToLower(s)
	t = strings.ToLower(t)
	s = SortString(s)
	t = SortString(t)
	return strings.Compare(s, t) == 0
}

// SortString Sort a string alphanumeric order
func SortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

package p1

import (
	"strings"
)

func findWords(words []string) []string {
	var typed []string
	for _, word := range words {
		if contain(word, "qwertyuiop") || contain(word, "asdfghjkl") || contain(word, "zxcvbnm") {
			typed = append(typed, word)
		}
	}
	return typed
}

func contain(a string, b string) bool {
	a = strings.ToLower(a)
	b = strings.ToLower(b)
	for _, runeA := range a {
		if !strings.Contains(b, string(runeA)) {
			return false
		}
	}
	return true
}

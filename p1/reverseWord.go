package p1

import "strings"

func reverseWords(s string) string {
	words := strings.Split(s, " ")
	for i := range words {
		words[i] = reverseStr(words[i])
	}
	return strings.Join(words, " ")
}

func reverseStr(s string) string {
	sByte := []byte(s)
	i, j := 0, len(sByte)-1
	for i < j {
		sByte[i], sByte[j] = sByte[j], sByte[i]
		i++
		j--
	}
	return string(sByte)
}

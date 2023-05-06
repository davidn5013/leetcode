package grind75

import (
	"strings"
)

// CanConstruct 0383 can string a letter be taken out of string b
// a taken from b false
// aa taken from ab false
// aa taken from aab true
func canConstruct(ransom string, magazine string) bool {
	magBytearr := []byte(strings.ToLower(magazine))
	for _, r := range []byte(strings.ToLower(ransom)) {
		found := false
		for i, s := range magBytearr {
			if r >= 'a' && r <= 'z' && r == s {
				found = true
				magBytearr[i] = 0
				break
			}
		}
		if !found {
			return false
		}

	}
	return true
}

func canConstruct2(ransom string, magazine string) bool {
	magaLcase := strings.ToLower(magazine)
	for _, r := range strings.ToLower(ransom) {
		found := false
		if strings.Contains(magaLcase, string(r)) {
			found = true
			magaLcase = strings.Replace(magaLcase, string(r), "", 1)
			break
		}
		if !found {
			return false
		}
	}
	return true
}

func canConstruct3(ransom string, magazine string) bool {
	ransom = strings.ToUpper(ransom)
	magazine = strings.ToUpper(magazine)
	letterCount := [25]int{}

	for _, v := range magazine {
		if v >= 'A' && v <= 'Z' {
			letterCount[v-'A']++
		}
	}

	for _, v := range ransom {
		if v >= 'A' && v <= 'Z' {
			if letterCount[v-'A'] == 0 {
				return false
			}
			letterCount[v-'A']--
		}
	}

	return true
}

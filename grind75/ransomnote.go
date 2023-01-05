package grind75

import (
	"strings"
)

// CanConstruct 0383 can string a letter be taken out of string b
// a taken from b false
// aa taken from ab false
// aa taken from aab true
func CanConstruct(a string, b string) bool {
	strip := strings.ToLower(b)
	for _, r := range strings.ToLower(a) {
		prev := strip
		strip = strings.Replace(strip, string(r), "", 1)
		if strings.Compare(prev, strip) == 0 {
			return false
		}
	}
	return true
}

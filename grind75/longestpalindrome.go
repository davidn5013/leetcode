package grind75

import (
	"sort"
	"strings"
)

/* 409. Longest Palindrom

What det longest palindrome from example "abccccdd" it "dccaccd"
so return value is 7
it case sensitiv so "Aa" is not paln drom.

longest palindrome for "a" is 1

How?

abccccdd
moving single letter to middle work for one single letter
but second single has to be discarded.
Then we have cccc that vi can split in "cc"
Then we have dd that vi can split in d
if we get ccdadcc or dccaccd does not matter we are loook for
number of letters 7

sort the string

So first singel to the middel vi can do that last

singel:=getFirstSingel(string)

for every group that is even split in haft
 cccc = cc  dd=d
for not even group bigger >= 3 letter sub one and split in two
 if we hade like eee that giv e
comban and add singel

result is multi by 2
*/

// LongestPalingdrome how many letter in palindrome char mix made by they letter, case senitive
func LongestPalindrom(s string) (res int) {
	// store one singel
	one := false
	// split string in words here?
	grps := SplitOnLetter(s)
	if len(grps) == 1 {
		return len(grps[0])
	}
	for _, grp := range grps {
		switch {
		case len(grp) == 1 && one == false:
			one = true
			res++
		case len(grp)%2 == 0:
			res += len(grp)
		case len(grp)%2 != 0:
			if len(grp) > 2 {
				res += len(grp) - 1
			}
		}
	}
	return res

}

// SplitOnLetter split a string on new letter
// "abccccdd" will result in [a b ccc dd]
// "abcabc" will split in "aa","bb","cc"
// "KalleKula" will split in to "aa", "e" , "KK", "lll" ,"u" ** Is this right?
func SplitOnLetter(s string) (res []string) {
	// Need a sorted string
	sorted := func(w string) string {
		newstring := strings.Split(w, "")
		sort.Strings(newstring)
		return strings.Join(newstring, "")
	}(s)

	addstring := ""

	for len(sorted) > 0 {
		// Get the first character of the string
		char := sorted[:1]
		// Check if the next character is the same
		if len(sorted) > 1 && sorted[1:2] == char {
			// If it is, find the longest run of the same character
			for i, r := range sorted[1:] {
				if strings.Compare(string(r), char) == 0 {
					addstring += sorted[:i+1]
					sorted = sorted[i+1:]
					break
				}
			}
		} else {
			// If it is not, add the character to the result and remove it from the string
			addstring += char
			sorted = sorted[1:]
			res = append(res, addstring)
			addstring = ""
		}
	}

	return res
	// Output: [a b cccc d d]
}

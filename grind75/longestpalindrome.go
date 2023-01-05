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

// LongestPalingdrome 0409 how many letter in palindrome char mix made by they letter, case senitive
// Support only rune until 256 skip char above
func LongestPalindrome(s string) (res int) {
	// Solution from https://github.com/chai2010/LeetCode-in-Go/blob/master/Algorithms/0409.longest-palindrome/longest-palindrome.go
	// Many thank to chai2010
	const maxRunevalue = 256
	a := make([]int, maxRunevalue) // Ascii number until 256 with count for ever char
	for i := range s {
		// What is this? s[i] = rune ascii number so a = a[65]++ how many a is inte list
		if int(s[i]) < maxRunevalue+1 {
			a[s[i]]++
		}
	}

	hasOdd := 0
	for i := range a {
		if a[i] == 0 {
			continue
		}
		// Check for even (9&1=1, 8&1=0)
		if a[i]&1 == 0 {
			res += a[i]
		} else {
			// Odd exampel 7 then add 6
			res += a[i] - 1
			hasOdd = 1
		}
	}

	return res + hasOdd
}

/*
  Did not work. Can take string like "ccc" is much more ivolved then the version above

func defektLongestPalindrome(s string) (res int) {
	// store one singel
	one := false
	// split string in words here?
	grps := SplitOnLetter(s)
	if len(grps) == 1 {
		return len(grps[0])
	}
	for _, grp := range grps {
		r := res
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
		log.Printf("grp %s %d\n", grp, res-r)
	}
	return res

}
*/

/* This SplitOnLetter was made for LongestPalindrome but can be use full so I keepit
even do the LongestPalindrome did not work */

// SplitOnLetter split a string on new letter
// "abccccdd" will result in [a b ccc dd]
// "abcabc" will split in "aa","bb","cc"
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

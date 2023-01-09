package grind75

import (
	"github.com/davidn5013/leetcode/tools"
)

// leetcode 3 Grind75 number 29

// lengthOfLongestSubString return n length of substring without repeating characters.
func lengthOfLongestSubstring(s string) int {
	var (
		visited      = make(map[uint8]int)
		maxLen, sIdx = 0, 0
	)

	// Sliding window check storing byte of string in uint8
	for sIdx < len(s) {
		// unvisited character? Save if not
		if charIdx, ok := visited[s[sIdx]]; !ok {
			visited[s[sIdx]] = sIdx
			sIdx++
		} else {
			// (idx++) + char index [expand window...]
			sIdx = charIdx + 1
			maxLen = tools.MaxInt(maxLen, len(visited))
			visited = make(map[uint8]int)
		}
	}

	maxLen = tools.MaxInt(maxLen, len(visited))
	visited = nil
	return maxLen
}

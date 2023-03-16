package lc171

// titleToNumber excel column letter to numbers
func titleToNumber(columnTitle string) int {
	carry := 1
	result := 0
	for i := len(columnTitle) - 1; i >= 0; i-- {
		result += carry * (int(columnTitle[i]-'A') + 1)
		carry = carry * 26
	}
	return result
}

/* my solution
func titleToNumber(columnTitle string) int {
	var (
		sum    int
		lenCol = len(columnTitle)
	)
	for i, s := range columnTitle {
		// Add power 26 to columnTitle from position excluding the las on
		p := IntPow(26, lenCol-i-1)
		sum += toInt(s) * p
	}
	return sum
}
*/

// IntPow calculates n to the mth power. Since the result is an int, it is assumed that m is a positive power
func IntPow(n, m int) int {
	if m == 0 {
		return 1
	}
	result := n
	for i := 2; i <= m; i++ {
		result *= n
	}
	return result
}

// toInt rune to in 'A' = 1
func toInt(r rune) int {
	return int(r) - 64
}

package p1

// CountOdds count odd number in range
func CountOdds(low int, high int) int {
	diff := high - low
	if low%2 == 1 || high%2 == 1 {
		return diff/2 + 1
	}
	return diff / 2
}

/* Simple but slow version
func CountOdds(low int, high int) int {
    diff := high - low

    if low % 2 == 1 || high % 2 == 1 {
        return diff / 2 + 1
    }
    return diff / 2
}
*/

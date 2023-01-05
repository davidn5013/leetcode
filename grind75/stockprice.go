package grind75

import "math"

/*
 * prices = [7,1,5,3,6,4] return 5
 * Explanation: Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit = 6-1 = 5.
 * Note that buying on day 2 and selling on day 1 is not allowed because you must buy before
 * you sell.
 */

// MaxProfit 0121 Return the maximum profit you can achieve from this transaction.
// If you cannot achieve any profit, return 0.
func MaxProfit(prices []int) (maxprof int) {
	// maxLen := len(prices)

	// if maxLen <= 1 {
	// 	return 0
	// }

	// for i := 0; i < maxLen; i++ {
	// 	for j := i; j < maxLen; j++ {
	// 		prof := prices[j] - prices[i]
	// 		if maxprof < prof {
	// 			maxprof = prof
	// 		}
	// 	}
	// }

	// Realy smart solution from :
	// https://leetcode.com/problems/best-time-to-buy-and-sell-stock/solutions/385299/Go-golang-fast-and-slow-solutions/
	// Check for the smallest price and check max proffit on ever price
	// only one loop O(n)
	minPrice := math.MaxInt32
	for _, price := range prices {
		if price < minPrice {
			minPrice = price
		} else if price-minPrice > maxprof {
			maxprof = price - minPrice
		}
	}
	return maxprof
}

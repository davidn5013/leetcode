package grind75

import "math"

// leetcode 322. Coin Change
// Grind 75 number 36

/*
You are given an integer array coins representing coins of different denominations and an integer amount representing a total amount of money.

Return the fewest number of coins that you need to make up that amount. If that amount of money cannot be made up by any combination of the coins, return -1.

You may assume that you have an infinite number of each kind of coin.

Example 1:

Input: coins = [1,2,5], amount = 11
Output: 3
Explanation: 11 = 5 + 5 + 1

Example 2:

Input: coins = [2], amount = 3
Output: -1

Example 3:

Input: coins = [1], amount = 0
Output: 0
*/

// TODO Go back try again to understand this CoinChange solution from
// https://github.com/frankegoesdown/LeetCode-in-Go/blob/master/Algorithms/0322.coin-change/coin-change.go

// CoinChange return best coin denominations mix
func CoinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		dp[i] = math.MaxInt
	}

	lenCoins := len(coins)

	for i := 1; i <= amount; i++ {
		for j := 0; j < lenCoins; j++ {

			if coins[j] <= i {
				rest := dp[i-coins[j]]
				if rest != math.MaxInt && rest+1 < dp[i] {
					dp[i] = rest + 1
				}
			}

		}
	}

	res := dp[amount]
	if res == math.MaxInt {
		res = -1
	}

	return res
}

/*
// Fastest on leetcode
// https://leetcode.com/problems/coin-change/submissions/888940711/

func coinChange(coins []int, amount int) int {
	// Create a dp array with size of amount + 1 and initialize it with a large number (amount + 1)
	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = amount + 1
	}
	dp[0] = 0

	// Iterate through the coins
	for _, coin := range coins {
		// For each coin, starting from that coin value to the amount,
		for i := coin; i <= amount; i++ {
			// update the dp array by taking the minimum number of coins needed to make up that amount using that coin
			dp[i] = min(dp[i], dp[i-coin]+1)
		}
	}

	// check the last element of the dp array, if it is still a large number, return -1 as it is not possible to make up that amount using the given coins
	if dp[amount] > amount {
		return -1
	}
	// Else, return the last element as it will be the minimum number of coins needed
	return dp[amount]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
*/

package grind75

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
		dp[i] = amount + 1
		for _, c := range coins {
			if c <= i && dp[i] > dp[i-c]+1 {
				dp[i] = dp[i-c] + 1
			}
		}
	}

	if dp[amount] > amount {
		return -1
	}

	return dp[amount]
}

// coinChange return best coin denominations mix
/* My solution that did not work it take no consideration for using lover denotation to
get correct exchange of coins
func coinChange(coins []int, amount int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(coins)))
	coinUse := 0
	for _, v := range coins {
		for amount >= v {
			log.Println(amount, v, coinUse)
			amount -= v
			coinUse++
		}
	}
	if amount < 0 {
		return -1
	}
	return coinUse
}
*/

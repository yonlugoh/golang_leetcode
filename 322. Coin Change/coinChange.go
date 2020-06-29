package _22__Coin_Change

import "math"

func coinChange(coins []int, amount int) int {
	// slice for dynamic programming to store results for minimum number to reach each value from 0 to amount
	dp := make([]int, amount+1)
	// initialize, 0 coins required to reach 0 amount
	dp[0] = 0

	for i := 1; i <= amount; i++ {
		// initialize result to infinity
		dp[i] = math.MaxInt32
		for _, coin := range coins {

			if i-coin >= 0 {
				// if this condition is true, we can set the current result to be minimum of current amount and taking previously reachable amount+1
				dp[i] = min(dp[i-coin]+1, dp[i])
			}
		}
	}
	// unreachable
	if dp[amount] == math.MaxInt32 {
		return -1
	} else {
		return int(dp[amount])
	}
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

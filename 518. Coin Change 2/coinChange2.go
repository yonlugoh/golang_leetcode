package _18__Coin_Change_2

func change2(amount int, coins []int) int {

	// Time complexity: O(amount* len(coins))
	// Space complexity: O(amount* len(coins))
	n := len(coins)
	if amount == 0 {
		return 1
	}
	if n == 0 {
		return 0
	}
	dp := make([][]int, n+1)
	for row := 0; row < n+1; row++ {
		dp[row] = make([]int, amount+1)
		if row == 0 {
			for i := 0; i < amount+1; i++ {
				if i == 0 {
					dp[0][i] = 1
				} else {
					dp[0][i] = 0
				}
			}

		} else {
			for col := 0; col < amount+1; col++ {
				coin := coins[row-1]
				if coin > col {
					dp[row][col] = dp[row-1][col]

				} else {
					dp[row][col] = dp[row-1][col] + dp[row][col-coin]
				}
			}
		}
	}
	return dp[n][amount]
}

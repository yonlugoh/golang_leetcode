package _22__Best_Time_to_Buy_and_Sell_Stock_II

func maxProfit2(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	profit := 0

	for i := 1; i < len(prices); i++ {
		diff := prices[i] - prices[i-1]
		if diff > 0 {
			profit += diff
		}
	}

	return profit
}

package _21__Best_Time_to_Buy_and_Sell_Stock

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	lowestPrice := prices[0]
	maxProfitSoFar := 0

	for _, price := range prices {
		if price-lowestPrice > maxProfitSoFar {
			maxProfitSoFar = price - lowestPrice
		}
		if price < lowestPrice {
			lowestPrice = price
		}

	}

	return maxProfitSoFar
}

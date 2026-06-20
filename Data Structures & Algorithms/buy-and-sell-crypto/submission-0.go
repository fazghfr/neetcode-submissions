func maxProfit(prices []int) int {
	maxProfit := 0
	minBuy := math.MaxInt32

	for _, curr := range prices {
		profit := curr - minBuy
		maxProfit = max(maxProfit, profit)
		minBuy = min(minBuy, curr)
	}

	return maxProfit
}


func coinChange(coins []int, amount int) int {
    dp := make([]int, amount + 1)
	for i, _ := range dp {
		dp[i] = amount + 1
	}
	dp[0] = 0
	for i := 1; i <= amount ; i++ {
		for _, c := range coins {
			if c <= i {
				dp[i] = min(dp[i-c] + 1, dp[i])
			}
		}
	}

	if dp[amount] == amount + 1 {
		return -1
	}
	return dp[amount]
}


// 1 5 10 , 1 -> dp[1] = 1
// 1 5 10 , 2 -> dp[2] = use coin(1) + dp[1] -> 2
// ..
// ..
// ..            dp[5] = 5
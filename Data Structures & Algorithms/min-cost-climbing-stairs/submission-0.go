func minCostClimbingStairs(cost []int) int {
    dp := make([]int, len(cost)+1)

	dp[0], dp[1] = 0, 0

	for i := 2; i < len(cost) + 1; i++ {
		dp[i] = min(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2])
	}

	return dp[len(dp)-1]
}


// dp[i] -> min cost to reach i, where i is assumed top


// 1 2 3 ->
// 0 1 2 3


// start could be 0 or 1 index so base cases are these
// dp[0] -> 0
// dp[1] -> 0

// dp[n] for n [2, target] is
// dp[n] = min(cost[n-2] + dp[n-2], cost[n-1] + dp[n-1])

// idea -> min cost to reach n is either 
// min cost to reach n-1 or n-2 + the cost on either of those
func climbStairs(n int) int {
    dp := make([]int, n+1)

	dp[0] = 1
	dp[1] = 1

	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}



// so to get to x,
// either 1 step meaning dp[x-1] or 2 step dp[x-2], accumulate


// dp[x] = dp[x-1] + dp[x-2]
// base dp[0] -> 1
// dp[1] -> 1
// dp[2] 2
// dp[3] 3
// dp[4] 5




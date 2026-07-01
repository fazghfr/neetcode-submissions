func rob(nums []int) int {
	dp := make([]int, len(nums))

	dp[0] = nums[0]
	if len(nums) < 2 {
		return nums[0]
	}
	dp[1] = max(dp[0], nums[1])

	for i := 2;  i < len(nums); i++ {
		dp[i] = max(nums[i] + dp[i-2], dp[i-1])
	}

	return dp[len(nums)-1]
}

// 1 1 3 3

// // bottom up approach
// // dp[i] -> maximum money until ith position
// // return dp[n]

// dp[i] = max(nums[i] + dp[i-2], dp[i - 1])

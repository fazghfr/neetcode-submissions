func rob(nums []int) int {
    dp := make([]int, len(nums))
	dp[0] = nums[0]
	if len(nums) < 2 {
		return dp[0]
	} 

	dp[1] = max(dp[0], nums[1])

	for i, _ := range nums {
		if i < 2 {
			continue
		}

		dp[i] = max(nums[i] + dp[i-2], dp[i-1])
	}

	return dp[len(nums) - 1]
}

// cannot rob two adjacent houses

// dp[i] -> max value on this house

// on i, you can either take i, hence cant take i-1
// or dont take i, hence took i-1
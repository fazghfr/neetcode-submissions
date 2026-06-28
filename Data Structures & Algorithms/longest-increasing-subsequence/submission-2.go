func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	for i, _ := range dp {
		dp[i] = 1
	}
	maxLen := 0
	for i, v := range nums {
		for j := i-1; j >= 0; j--{
			if nums[j] < v {
				temp := dp[j] + 1
				dp[i] = max(dp[i], temp)
			} 
		}
		maxLen = max(maxLen, dp[i])

	}
	return maxLen
}


// 1 3 6 7 9 4 10 5 6
// 1 2 3 4 5 3 6

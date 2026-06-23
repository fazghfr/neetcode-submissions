func findMin(nums []int) int {
	// the o(n) time solution

	minVal := nums[0]

	for _, v := range nums {
		minVal = min(minVal, v)
	}

	return minVal
}


func maxSubArray(nums []int) int {
	runningSum := 0
	maxSub := -1
	for _, v := range nums {
		if runningSum < 0 {
			runningSum = 0
		}
		runningSum += v

		maxSub = max(maxSub, runningSum)
	}
	return maxSub
}

// if sum is negative use new sub instead


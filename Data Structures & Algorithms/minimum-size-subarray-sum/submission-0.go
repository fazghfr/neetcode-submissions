func minSubArrayLen(target int, nums []int) int {
	l := 0
	minLen := len(nums) + 1
	runningSum := 0
	for r := 0; r < len(nums); r++ {
		runningSum += nums[r]

		for runningSum >= target {
			minLen = min(minLen, r - l + 1)
			runningSum -= nums[l]
			l++
		}
	}
	if minLen == len(nums) + 1 {
		return 0
	}
	return minLen
}

// minimal len sub array with sum >= target
// all positive : kadane's not required
// sliding window


// sliding window rule
// shrink when sum bigger or equal
// calculate on shrinking

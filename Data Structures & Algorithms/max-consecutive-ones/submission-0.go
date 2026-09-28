func findMaxConsecutiveOnes(nums []int) int {
	counting := 0
	maxC := 0
	for _, v := range nums {
		if v == 1 {
			counting++
		} else {
			maxC = max(maxC, counting)
			counting = 0
		}
	}
	// trailing handling
	maxC = max(maxC, counting)
	return maxC
}

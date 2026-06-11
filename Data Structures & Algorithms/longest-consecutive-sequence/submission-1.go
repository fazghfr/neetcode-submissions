func longestConsecutive(nums []int) int {
	// possible exists
	// for each start, extend if possible

	exists := make(map[int]bool)

	for _, v := range nums {
		exists[v] = true
	}
	maxCount := 0
	for _, v := range nums {
		if _, ok := exists[v-1]; !ok {
			// v is a possible start since v-1 doesnt exist
			counted := 0
			myNum := v
			for exists[myNum] {
				counted++
				myNum++
			}
			maxCount = max(maxCount, counted)
		}
	}

	return maxCount
}

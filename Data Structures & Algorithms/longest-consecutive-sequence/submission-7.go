func longestConsecutive(nums []int) int {
	exists := make(map[int]bool)
	for _, v := range nums {
		exists[v] = true
	}

	// finding starts
	var starts []int
	for _, v := range nums {
		if !exists[v-1] {
			starts = append(starts, v)
		}
	}

	// using starts to iterate using the exists map
	maxSeq := 0
	for _, v := range starts {
		copyV := v
		c := 0
		for exists[copyV] {
			copyV++
			c++
		}
		maxSeq = max(maxSeq, c)
	}
	return maxSeq
}

// possible starts
// definition :  exist[val - 1] does not exist

// with that starts each, try to find++\
// record the max iteration count
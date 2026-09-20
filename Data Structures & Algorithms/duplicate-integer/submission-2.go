func hasDuplicate(nums []int) bool {
    seen := make(map[int]bool)

	for _, v := range nums {
		if !seen[v] {
			seen[v] = true
		} else {
			return true
		}
	}

	return false
}

// iterate,
// populate map on unseen
// on seen break return true

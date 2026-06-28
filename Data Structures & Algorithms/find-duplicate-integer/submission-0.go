func findDuplicate(nums []int) int {
	freq := make(map[int]bool)
	for _, v := range nums {
		if freq[v] {
			return v
		} else {
			freq[v] = true
		}
	}

	return -1
}

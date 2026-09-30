func removeDuplicates(nums []int) int {
	seen := make(map[int]bool)

	var set []int
	for _, v := range nums {
		if !seen[v] {
			set = append(set, v)
		}
		seen[v] = true
	}

	copy(nums, set)
	return len(set)
}
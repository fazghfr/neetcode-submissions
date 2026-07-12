func search(nums []int, target int) int {
	for i, v := range nums {
		if v == target {
			return i
		}
	}

	return -1
}

// naive approach would be just search by iteration 
// which leads into o(n) time complexity

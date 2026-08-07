func search(nums []int, target int) int {
	for i, v := range nums {
		if v == target {
			return i
		}
	}
	return -1
}		

// nums, sorted. target, return index not found -1


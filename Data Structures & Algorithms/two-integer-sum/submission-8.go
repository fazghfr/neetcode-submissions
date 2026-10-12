func twoSum(nums []int, target int) []int {
    seen := make(map[int]int)

	for i, v := range nums {
		complement := target - v
		if _, ok := seen[complement]; ok {
			return []int{seen[complement], i}
		} else {
			seen[v] = i
		}
	}

	return nil
}

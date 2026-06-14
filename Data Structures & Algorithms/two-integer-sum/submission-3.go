func twoSum(nums []int, target int) []int {
    lookup := make(map[int]int)

	for i, v := range nums {
		complement := target - v
		if _, ok := lookup[complement]; ok {
			return []int{lookup[complement], i}
		} else {
			lookup[v] = i
		}
	}

	return nil
}

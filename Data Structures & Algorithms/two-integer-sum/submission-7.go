func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, v := range nums {
		if _, ok := m[target - v]; ok {
			return []int{m[target - v], i}
		}
		m[v] = i
	}
	return nil
}

// store map
// if target - v in map, return val of target - v maps and cur index

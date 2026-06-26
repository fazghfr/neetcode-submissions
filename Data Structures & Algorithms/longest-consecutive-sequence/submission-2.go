func longestConsecutive(nums []int) int {
	seen := make(map[int]bool)
	ans := 0
	for _, v := range nums {
		seen[v] = true
	}

	for _, v := range nums {
		if !seen[v-1] {
			tempV := v
			count := 0
			for seen[tempV] {
				count++
				tempV++
			}

			ans = max(ans, count)
		}
	}

	return ans
}

// using hashmaps
// find all the possible starts
// loop through nums -> v
// if v is a possible start, increment v and calculate (calc when incremented exist only)

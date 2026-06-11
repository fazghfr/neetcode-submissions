func hasDuplicate(nums []int) bool {
    myMap := make(map[int]int)

	for _, v := range nums {
		myMap[v]++
		if myMap[v] > 1 {
			return true
		}
	}

	return false
}

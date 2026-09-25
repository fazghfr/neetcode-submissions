func replaceElements(arr []int) []int {
	// naive approach
	ans := make([]int, len(arr))
	for i, _ := range ans {
		curMax := -1
		for j := i + 1; j < len(arr); j++ {
			curMax = max(curMax, arr[j])
		}
		ans[i] = curMax
	}
	return ans
}	
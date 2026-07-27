func findLucky(arr []int) int {
	freq := make(map[int]int)

	for _, v := range arr {
		freq[v]++
	}
	maxNum := -1
	for k, v := range freq {
		if k == v {
			maxNum = max(maxNum, k)
		}
	}

	return maxNum
}

// frequency count
// then iterate through the keys
// sort by its value

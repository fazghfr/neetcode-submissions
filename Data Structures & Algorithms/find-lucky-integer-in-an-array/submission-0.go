func findLucky(arr []int) int {
	freq := make(map[int]int)

	for _, v := range arr {
		freq[v]++
	}
	var temp []int
	for k, v := range freq {
		if k == v {
			temp = append(temp, v)
		}
	}
	sort.Slice(temp, func(i, j int)bool {
		return freq[temp[i]] > freq[temp[j]]
	})
	if len(temp) != 0 {
		return temp[0]
	}

	return -1
}

// frequency count
// then iterate through the keys
// sort by its value

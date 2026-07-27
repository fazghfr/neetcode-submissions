func findLucky(arr []int) int {
	freq := make(map[int]int)

	for _, v := range arr {
		freq[v]++
	}
	var lucky []int
	for k, v := range freq {
		if k == v {
			lucky = append(lucky, v)
		}
	}
	sort.Slice(lucky, func(i, j int)bool {
		return freq[lucky[i]] > freq[lucky[j]]
	})
	if len(lucky) != 0 {
		return lucky[0]
	}

	return -1
}

// frequency count
// then iterate through the keys
// sort by its value

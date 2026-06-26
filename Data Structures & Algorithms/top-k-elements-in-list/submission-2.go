func topKFrequent(nums []int, k int) []int {
	freq := make(map[int]int)

	for _, v := range nums {
		freq[v]++
	}

	var tempList []int
	for k, _ := range freq {
		tempList = append(tempList, k)
	}
	sort.Slice(tempList, func(i, j int) bool {
		return freq[tempList[i]] > freq[tempList[j]]
	})
	return tempList[:k]
}

// frequency Map
// sort by value descending
// extract k elements from start, resulting the answer

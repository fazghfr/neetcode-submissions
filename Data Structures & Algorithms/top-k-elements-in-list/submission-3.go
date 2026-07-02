func topKFrequent(nums []int, k int) []int {
	freq := make(map[int]int)
	var keys []int
	for _, v := range nums {
		if _, ok := freq[v]; !ok {
			keys = append(keys, v)
		}
		freq[v]++
	}

	// sort the keys array referencing to the map
	sort.Slice(keys, func(i, j int)bool {
		return freq[keys[i]] > freq[keys[j]]
	})

	return keys[:k]
}

// frequency map nums[i] as the key, and value is the freq count of the key
// sort by value descending
// return top k from 0

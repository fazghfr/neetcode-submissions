func majorityElement(nums []int) int {
    freq := make(map[int]int)
	target := len(nums) / 2
	for _, v := range nums {
		freq[v]++
		if freq[v] > target {
			return v
		}
	}

	return (-1) * 1000000 - 1
}

// freq map
// store target freq
// calculate, if more than target freq, return that nums

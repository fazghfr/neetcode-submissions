func majorityElement(nums []int) []int {
	freq := make(map[int]int)

	for _, v := range nums {
		freq[v]++
	}

	target := len(nums) / 3
	var ans []int
	for k, v := range freq {
		if v > target {
			ans = append(ans, k)
		}
	}
	return ans
}

// populate frequency map
// define a new array
// populate this array 

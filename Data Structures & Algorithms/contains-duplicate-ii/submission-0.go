func containsNearbyDuplicate(nums []int, k int) bool {
	l := 0
	freq := make(map[int]int)

	for r := 0; r < len(nums); r++ {
		freq[nums[r]]++

		if r - l > k {
			freq[nums[l]]--
			l++
		}

		if freq[nums[r]] == 2 {
			return true
		}
	}

	return false
}

// restate -> is there a duplicate within k window

// expand by freq[v]++
// shrink when r - l + 1 > k
// check for if freq[r] == 2 -> return true
// default false

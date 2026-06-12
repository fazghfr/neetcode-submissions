func lengthOfLongestSubstring(s string) int {
	// idea : sliding window
	// valid window : when disctinct
	// on invalid, shrink until valid

	left := 0
	freqMap := make(map[byte]int)

	ans := 0
	for right := 0; right < len(s); right++ {
		freqMap[s[right]]++

		for freqMap[s[right]] > 1 {
			freqMap[s[left]]--
			left++
		}

		ans = max(ans, right - left + 1)
	}
	return ans
}



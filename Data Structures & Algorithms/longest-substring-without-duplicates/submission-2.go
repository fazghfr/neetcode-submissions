func lengthOfLongestSubstring(s string) int {
	seen := make(map[byte]int)

	l := 0
	ans := 0
	for r := 0; r < len(s); r++ {
		seen[s[r]]++

		for seen[s[r]] > 1 {
			seen[s[l]]--
			l++
		}

		ans = max(ans, r - l + 1)
	}
	return ans
}

// sliding window approach

// always expand
// shrink when found duplicate
// calculate length, record max

// hash maps is required to find what is duplicate or not
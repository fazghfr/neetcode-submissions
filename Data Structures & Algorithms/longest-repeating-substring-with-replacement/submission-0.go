func characterReplacement(s string, k int) int {
	l := 0
	freq := make(map[byte]int)

	runMax := -1
	ans := -1
	for r := 0; r < len(s); r++ {
		freq[s[r]]++
		runMax = max(runMax, freq[s[r]])

		for r - l + 1 - runMax > k {
			freq[s[l]]--
			l++
		}

		ans = max(ans, r-l+1)
	}

	return ans
}







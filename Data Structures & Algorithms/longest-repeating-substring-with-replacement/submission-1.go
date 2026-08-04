func characterReplacement(s string, k int) int {
	l := 0
	rMax := -1
	freq := make(map[rune]int)
	ans := 0
	for r:=0;r<len(s);r++{
		freq[rune(s[r])]++
		rMax = max(rMax, freq[rune(s[r])])

		for (r - l + 1) - rMax > k {
			freq[rune(s[l])]--
			l++
		}

		ans = max(ans, r - l + 1)
	}
	return ans
}

// on expand, add freq++
// shrink when attempted replacement is exceeding k

// calculate the majority char
// minority = curLen - majority
// shrink when minority is more than k
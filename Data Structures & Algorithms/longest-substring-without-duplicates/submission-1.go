func lengthOfLongestSubstring(s string) int {
	freq := make(map[rune]int)

	l:=0
	maxLen := 0
	for r:=0;r<len(s);r++{
		freq[rune(s[r])]++

		for freq[rune(s[r])] > 1 {
			freq[rune(s[l])]--
			l++
		}	
		maxLen = max(maxLen, r-l+1)
	}
	return maxLen
}


// frequency map
// shrink when non distinct are found

// freq[r]++
// if freq[r] non distinct shrink
// calculate after, return the r - l + 1
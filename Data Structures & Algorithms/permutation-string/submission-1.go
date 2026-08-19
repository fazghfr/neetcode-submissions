func checkInclusion(s1 string, s2 string) bool {

	isAnagram := func (s string, pivot string) bool {
		freq := make(map[rune]int)
		for _, v := range pivot {
			freq[v]++
		}	

		for _, v := range s {
			freq[v]--
			if freq[v] == 0 {
				delete(freq, v)
			}
		}

		return len(freq) == 0
	}
	l := 0
	for r := 0; r < len(s2); r++ {
		if r - l + 1 == len(s1) {
			if isAnagram(s2[l:r+1], s1) == true {
				return true
			}
			l++
		}

		if r - l + 1 > len(s1) {
			l++
		}
	}
	return false
}


// fixed sliding window with s1 length
// s1 length -> n

// on every n window, check if it is an anagram with s1
// return true if it is
// return false as default
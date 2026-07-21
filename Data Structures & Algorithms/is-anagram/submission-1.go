func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

 	freq := make(map[rune]int)
	for _, v := range s {
		freq[rune(v)]++
	}

	for _, v := range t {
		freq[rune(v)]--
	}

	for _, v := range freq {
		if v != 0 {
			return false
		}
	}

	return true
}

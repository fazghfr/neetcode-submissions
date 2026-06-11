func isAnagram(s string, t string) bool {
	if len(s) != len(t) {return false}
	// frequency map approach
	freqMap := make(map[rune]int)

	for _, v := range s {
		freqMap[v]++
	}

	for _, v := range t {
		if _, ok := freqMap[v]; ok {
			freqMap[v]--
		} else {
			return false
		}
	}

	for _, v := range freqMap {
		if v != 0 {
			return false
		}
	}

	return true
}

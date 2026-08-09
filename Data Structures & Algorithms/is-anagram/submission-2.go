func isAnagram(s string, t string) bool {
	if len(s) !=  len(t) {
		return false
	}
	freq := make(map[rune]int)

	for _, v := range s {
		freq[v]++
	}

	for _, v := range t {
		if _, ok := freq[v]; !ok {
			return false
		} else {
			freq[v]--
			if freq[v] == 0 {
				delete(freq, v)
			}
		}
	}

	return len(freq) == 0
}


// key idea
// populate freq on s
// decrement freq on t
// when found a nonexistant character as key -> false instantly
// on key value of 0, delete that key
// length of map freq has to be zero to determine s and t as anagram


func checkInclusion(s1 string, s2 string) bool {
	freq := make(map[byte]int)
	windowLen := len(s1)
	for i, _ := range s1 {
		freq[s1[i]]++
	}

	l := 0
	for r := 0; r < len(s2); r++ {
		if _, ok := freq[s2[r]]; ok {
			freq[s2[r]]--
		}

		for r - l + 1 > windowLen {
			if _, ok := freq[s2[l]]; ok {
				freq[s2[l]]++
			}
			l++
		}

		test := 0
		for _, v := range freq {
			if v != 0 {
				test++
				break
			} else {
				continue
			}
		}

		if test == 0 {
			return true
		}
	}

	return false
}


// abc -> window len is 3

// lec, eca, cab
// on each window, check the frequency count

func minWindow(s string, t string) string {
	if len(s) < len(t) {
		return ""
	}

	l := 0

	freq := make(map[rune]int)
	for _, v := range t {
		freq[v]++
	}

	missing := len(t)

	fL, fR := -1, -1

	for r := 0; r < len(s); r++ {
		if _, ok := freq[rune(s[r])]; ok {
			freq[rune(s[r])]--
			if freq[rune(s[r])] >= 0 {
				missing--
			}
		}

		for missing == 0 {
			if fL == -1 || r-l < fR-fL {
				fL = l
				fR = r
			}

			if _, ok := freq[rune(s[l])]; ok {
				freq[rune(s[l])]++
				if freq[rune(s[l])] > 0 {
					missing++
				}
			}

			l++
		}
	}

	if fL == -1 {
		return ""
	}

	return s[fL : fR+1]
}
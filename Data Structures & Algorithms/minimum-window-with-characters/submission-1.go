func minWindow(s string, t string) string {
    left := 0

	myMap := make(map[rune]int)
	missing := len(t)
	for _, v := range t {
		myMap[rune(v)]++
	}

	var candAns string
	for right := 0; right < len(s); right++ {
		// expand 
		if myMap[rune(s[right])] > 0 {
			missing--
		}
		myMap[rune(s[right])]--

		// shrink when all missing are found
		for missing == 0 {
			myMap[rune(s[left])]++
			if myMap[rune(s[left])] > 0 { 
				missing++ 
			}

			if len(candAns) == 0 {
				candAns = s[left:right+1]
			} else {
				if right - left + 1 < len(candAns) {
					candAns = s[left:right+1]
				}
			}
			left++
		}
		
	}

	return candAns
}

// XYZ

// OUZODYXAZV

// // EXTEND UNTIL ALL IS THERE
// OUZODYX

// // SHRINK UNTIL IT IS ALL THERE
// ZODYX -> candidate

// // HSRI
// ODYXAZ

// EXTEND WHEN FREQ NOT ENOUGH
// MAX SHRINK
// CANDIDATE HERE 

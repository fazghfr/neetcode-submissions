func totalFruit(fruits []int) int {
	l := 0

	freq := make(map[int]int)
	ftypes := 0
	maxLen := -1
	for r := 0; r < len(fruits); r++ {
		// expansion
		if freq[fruits[r]] == 0 {
			ftypes++
		} 
		freq[fruits[r]]++

		// shrink
		for ftypes > 2 {
			freq[fruits[l]]--
			if freq[fruits[l]] == 0 {
				ftypes--
			} 
			l++
		}

		maxLen = max(maxLen, r - l + 1)
	} 

	return maxLen
}

// restating the problem

// maximum sequence length where the sequence have max 2 distinct type

// sliding window
// always expand
// shrink when type > 2

// calculate max length
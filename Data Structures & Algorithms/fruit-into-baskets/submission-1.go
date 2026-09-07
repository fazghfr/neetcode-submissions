func totalFruit(fruits []int) int {
	freq := make(map[int]int)
	treeType := 0
	l := 0

	maxTrees := 0
	for r := 0; r < len(fruits); r++ {
		if freq[fruits[r]] == 0 {
			treeType++
		}
		freq[fruits[r]]++

		for treeType > 2 {
			freq[fruits[l]]--
			if freq[fruits[l]] == 0 {
				treeType--
			}
			l++
		}

		// calculation
		maxTrees = max(maxTrees, r - l + 1)
	}

	return maxTrees
}


// sliding window approach
// valid window is when treeType is less than 2
// invalid is when above is false

// on expansion
// if tree is unseen before, treeType++
// add freq[tree]++

// on shrink (shrink until type is less than two)
// using the l pointer




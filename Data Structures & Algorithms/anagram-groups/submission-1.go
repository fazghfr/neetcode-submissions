func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]string)

	for _, words := range strs {
		rWords := []rune(words)
		sort.Slice(rWords, func(i, j int) bool {
			return rWords[i] < rWords[j]
		})

		m[string(rWords)] = append(m[string(rWords)], words)
	}

	var ans [][]string
	for _, arr := range m {
		ans = append(ans, arr)
	} 

	return ans
}

// strs -> array of words
// output array of array, each array is an anagrams

// iterate to each word and sort
// after sorted, append to map (map[strs]int[])
// after that, run each key
// populate the array

// revision : map[strs]string[] to make it easier to populate the answer variable
func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]string)

	for _, v := range strs {
		runeV := []rune(v)
		sort.Slice(runeV, func(i, j int) bool {
			return runeV[i] < runeV[j]
		})
		m[string(runeV)] = append(m[string(runeV)], v)
	}
	
	var ans [][]string
	for _, v := range m {
		ans = append(ans, v)
	}

	return ans
}

// hashing 

// on iterate
// o: sort, check key
// o: append the presorted

// append ans with the vals of each key

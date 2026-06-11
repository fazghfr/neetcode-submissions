func groupAnagrams(strs []string) [][]string {
	freqs := make(map[string][]string)

	for _, v := range strs {
		runify := []rune(v)
		sort.Slice(runify, func (i, j int) bool {
			return runify[i] < runify[j]
		})

		key := runify

		freqs[string(key)] = append(freqs[string(key)], v)
	}

	var ans [][]string
	for _, v := range freqs {
		var temps []string
		for _, words := range v {
			temps = append(temps, words)
		}

		ans = append(ans, temps)
	}
	return ans
}

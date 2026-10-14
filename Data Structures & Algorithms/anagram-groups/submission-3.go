func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]string)

	for _, v := range strs {
		temp := []rune(v)

		sort.Slice(temp, func(i, j int) bool {
			return temp[i] < temp[j]
		})
		stemp := string(temp)
		m[stemp] = append(m[stemp], v)
	}

	var ans [][]string
	for _, lists := range m {
		ans = append(ans, lists)
	}
	return ans
}

// sort the str -> temp
// check existance
// append on temp key

// loop through ans and append the strs

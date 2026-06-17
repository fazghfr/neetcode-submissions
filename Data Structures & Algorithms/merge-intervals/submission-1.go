func merge(intervals [][]int) [][]int {
    // sorting the intervals first
	sort.Slice(intervals, func (i, j int)bool {
		return intervals[i][0] < intervals[j][0]
	})

	// grow when new start less than cur end
	var ans [][]int
	cur := intervals[0]
	for i, v := range intervals {
		if i == 0 {continue}

		if v[0] <= cur[1] {
			cur = []int{
				min(v[0], cur[0]),
				max(v[1], cur[1]),
			}
		} else {
			ans = append(ans, cur)
			cur = v
		}
	}

	ans = append(ans, cur)
	return ans
}



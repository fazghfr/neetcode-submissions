func merge(intervals [][]int) [][]int {
    // sorting
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] < intervals[j][1]
		}

		return intervals[i][0] < intervals[j][0]
	})

	var temp []int
	var ans [][]int
	for i, arr := range intervals {
		if i == 0 {
			temp = arr
			continue
		}

		if arr[0] <= temp[1] {
			temp[0] = min(temp[0], arr[0]) 
			temp[1] = max(temp[1], arr[1]) 
		} else {
			ans = append(ans, temp)
			temp = arr
		}
	} 
	// trailing item appending
	ans = append(ans, temp)

	return ans
}
// sort
// merge the overlapping interval
// if start is less than prev.end -> overlap
// merge it, but keep try to merge since its still possible to merge the next










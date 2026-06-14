func insert(intervals [][]int, newInterval []int) [][]int {
    // 3 possible cases
	// the whole newInterval lies before all of the intervals (non overlapping)
	// the whole newInterval lies after all of the intervals (non overlapping)
	// somewhere in the middle

	var ans [][]int
	for i, v := range intervals {
		if newInterval[1] < v[0] {
			ans = append(ans, newInterval)
			ans = append(ans, intervals[i:]...)
			return ans
		}else if newInterval[0] > v[1] {
			ans = append(ans, v)
		} else {
			newInterval = []int{min(newInterval[0], v[0]), max(newInterval[1], v[1])}
		}
	}

	ans = append(ans, newInterval)
	return ans
}

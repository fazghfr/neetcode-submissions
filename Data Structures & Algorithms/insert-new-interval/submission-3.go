func insert(intervals [][]int, newInterval []int) [][]int {
	maxInterval := 0
	for _, v := range intervals {
		maxInterval = max(maxInterval, v[1])
	}

	for _, v := range newInterval {
		maxInterval = max(maxInterval, v)
	}

	diff := make([]int, maxInterval*2+3)

	tempInterval := append(intervals, newInterval)
	for _, v := range tempInterval {
		diff[v[0]*2]++
		diff[v[1]*2+1]--
	}

	prefixSum := make([]int, maxInterval*2+3)
	for i, v := range diff {
		if i == 0 {
			prefixSum[i] = v
		} else {
			prefixSum[i] = prefixSum[i-1] + v
		}
	}

	isStart := false
	var ans [][]int
	var curStart int

	for i, v := range prefixSum {
		if v != 0 && !isStart {
			isStart = true
			curStart = i / 2
		} else if isStart && v == 0 {
			ans = append(ans, []int{curStart, (i-1)/2})
			isStart = false
		}
	}

	return ans
}

// initial solution, diff array,
// calculation too complicated, with a little assist from AI in currStart and end determination

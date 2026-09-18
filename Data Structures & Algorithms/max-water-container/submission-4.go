func maxArea(heights []int) int {
	l := 0
	r := len(heights) - 1

	maxVol := -1
	for l < r {
		lwall := heights[l]
		rwall := heights[r]

		curWall := min(lwall, rwall)

		maxVol = max(maxVol, curWall * (r - l))

		if lwall < rwall {
			l++
		} else {
			r--
		}
	}

	return maxVol
}

// converging
// shift the smaller ones 


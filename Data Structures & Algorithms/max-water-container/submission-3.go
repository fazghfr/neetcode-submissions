func maxArea(heights []int) int {
	l, r := 0, len(heights) - 1
	maxCon := 0
	for l < r {
		maxCon = max(maxCon, (r-l) * min(heights[l], heights[r]))
		if heights[l] < heights[r] {
			l++
		} else {
			r--
		}
	}

	return maxCon
}


// converging two pointer approach
// shift the lowest bar of either left/right
// reason : converging shrinks the width, so we need to find
// a potential bigger "base". 
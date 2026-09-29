func maxArea(heights []int) int {
	r := len(heights) - 1
	l := 0
	maxc := 0
	for l < r {
		curc := (r - l) * min(heights[l], heights[r])
		maxc = max(maxc, curc)
		if heights[l] < heights[r] {
			l++
		} else {
			r--
		}
	}
	return maxc
}
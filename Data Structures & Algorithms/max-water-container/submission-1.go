func maxArea(heights []int) int {
	r := len(heights) - 1
	l := 0

	maxVolume := 0
	for l < r {
		curVolume := (r - l) * min(heights[l], heights[r])
		maxVolume = max(maxVolume, curVolume)

		if heights[l] < heights[r] {
			l++
		} else {
			r--
		}
	}

	return maxVolume
}


// goal : return max volume in a 2d container

// ideally : maxwidth with maxheight
// so we traverse on the biggest width, and shrink the width
// because shrinking the width decrease the w variable, we need to choose the highest height


// so
// check left and right
// shift left if left is less than right
// otherwise shift right
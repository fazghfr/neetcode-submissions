func maxArea(heights []int) int {
	// two pointer approach
	// idea
	// shift right-- if right height is smaller
	// shift left++ if left height is smnaller
	// reason : always try to get the bigger height 
	// to compensate the shrinking width 

	left := 0
	right := len(heights)-1

	maxVolume := 0
	for left < right {
		maxVolume = max(maxVolume, (right - left) * min(heights[left], heights[right]))

		if heights[left] > heights[right] {
			right--
		} else {
			left++
		}
	}
	return maxVolume
}

func findMin(nums []int) int {
	l := 0
	r := len(nums) - 1 

	for l < r {
		m := l + (r - l) / 2

		if nums[m] > nums[r] {
			l = m + 1
		} else {
			r = m
		}
	}

	return nums[l]
}

// binary search
// if middle bigger than right -> minimum is on the right side
// if middle is not bigger than right -> minimum on the left side

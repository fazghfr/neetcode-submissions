func search(nums []int, target int) int {
	l := 0
	r := len(nums) - 1

	for l <= r {
		m := l + (r - l) / 2

		if nums[m] == target {
			return m
		} else if target > nums[m] {
			l = m + 1
		} else {
			r = m - 1
		}
	}

	return -1
} 

// mid calculation = l + (r-l) / 2
// if target bigger than nums at mid -> left is now at mid
// else right is now at mid

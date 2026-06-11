func productExceptSelf(nums []int) []int {
	leftProduct := make([]int, len(nums))
	rightProduct := make([]int, len(nums))

	for i, v := range nums {
		if i == 0 {
			leftProduct[i] = v
			continue
		}
		leftProduct[i] = leftProduct[i-1] * v
	}

	for i := len(nums)-1; i >= 0; i-- {
		v := nums[i]
		if i == len(nums)-1 {
			rightProduct[i] = v
			continue
		}
		rightProduct[i] = rightProduct[i+1] * v
	}

	ans := make([]int, len(nums))
	for i, _ := range ans {
		if i == 0 {
			ans[i] = rightProduct[i+1]
			continue
		}
		if i == len(nums) - 1 {
			ans[i] = leftProduct[i-1]
			continue
		}

		ans[i] = rightProduct[i+1] * leftProduct[i-1]
	}
	return ans
}


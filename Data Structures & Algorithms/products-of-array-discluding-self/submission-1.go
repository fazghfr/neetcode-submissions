func productExceptSelf(nums []int) []int {
	left := make([]int, len(nums))
	right := make([]int, len(nums))

	for i, v := range nums {
		if i == 0 {
			left[i] = v
		} else {
			left[i] = left[i-1] * v
		}
	}

	for i := len(nums)-1; i >=0; i-- {
		v := nums[i]
		if i == len(nums) - 1{
			right[i] = v
		} else {
			right[i] = right[i+1] * v
		}
	}

	ans := make([]int, len(nums))
	for i := 0; i < len(nums); i++{
		if i == 0 {
			ans[i] = right[i+1]
		} else if i == len(nums)-1 {
			ans[i] = left[i-1]
		} else {
			ans[i] = left[i-1] * right[i+1]
		}
	}

	return ans
}



func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left < right {
		mid := left + (right - left) / 2

		if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			right = mid
		}
	}

	if nums[left] == target {
		return left
	}

	newStart := left
	if target >= nums[newStart] && target <= nums[len(nums)-1] {
		left = newStart
		right = len(nums) - 1
	} else {
		left = 0
		right = newStart - 1
	}

	for left <= right {
		mid := left + (right - left) / 2

		if nums[mid] == target {
			return mid
		} else if target > nums[mid] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	
	return -1
}

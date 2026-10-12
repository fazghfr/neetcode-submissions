func twoSum(numbers []int, target int) []int {
	left := 0
	right := len(numbers) - 1

	for left < right {
		if numbers[left] + numbers[right] == target {
			return []int{left+1, right+1}
		} else if numbers[left] + numbers[right] > target {
			right--
		} else {
			left++
		}
	}

	return nil
}

// two pointer approach

// since numbers array is sorted, we can use a converging pointer approach
// if target == left + right, return left and right

// else, check if left + right > target, on true, shift left
// else shift right

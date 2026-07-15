func twoSum(numbers []int, target int) []int {
	left, right := 0, len(numbers) - 1

	for left < right {
		if numbers[left] + numbers[right] == target {
			return []int{left + 1, right + 1}
		} else if numbers[left] + numbers[right] > target {
			right--
		} else {
			left++
		}
	}

	return nil
}

// input array is sorted ascending
// meaning left always smaller than right

// brute force
// for each element, iterate the rest (this is o(n^2))

// two pointer converging approach

// check left + right, if bigger, then we need less, shift right-1
// if we need more, shift left + 1

// there is always an answer so no edge case for non existant solution

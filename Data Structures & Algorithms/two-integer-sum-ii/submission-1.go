func twoSum(numbers []int, target int) []int {
	l := 0
	r := len(numbers) - 1

	for l < r {
		temp :=  numbers[l] + numbers[r]
		if temp == target {
			break
		} else if temp < target {
			l++
		} else {
			r--
		}
	}

	return []int{l+1, r+1}
}









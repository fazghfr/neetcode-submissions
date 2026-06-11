func twoSum(nums []int, target int) []int {
    lookupTable := make(map[int]int)

	for i, v := range nums {
		lookupTable[v] = i
	}

	answers := make([]int, 2)
	for i, v := range nums {
		if _, ok := lookupTable[target - v]; ok && lookupTable[target - v] != i{
			answers[0] = lookupTable[target - v]
			answers[1] = i

			if answers[0] > answers[1] {
				answers[0], answers[1] = answers[1], answers[0]
			}
		}
	}

	return answers
}

func removeElement(nums []int, val int) int {
    nVal := 0
	for i, v := range nums {
		if v == val {
			nVal++
			nums[i] = 51
		}
	}

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	return len(nums) - nVal
}

// o(n) first pass to -> change all to max + 1, calculate how much

// sort them (log n) or n log n smething like tht

// return nums from 0 to n 
// 0 to n + how much = len(nums)

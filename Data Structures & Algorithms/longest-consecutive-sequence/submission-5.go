func longestConsecutive(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}

	exist := make(map[int]bool)
	for _, v := range nums {
		exist[v] = true
	}

	starts := make(map[int]bool)
	for _, v := range nums {
		if !exist[v-1] {
			starts[v] = true
		}
	}

	maxLen := 0
	for k, _ := range starts {
		cur := 1
		ktemp := k + 1

		for exist[ktemp] {
			cur++
			ktemp++
		}
		maxLen = max(maxLen, cur)
	}

	return maxLen
}



// we need to find elements that could be the start
// for each starting point, increment runningcur if and only if start+1 exist somewhere
// stop the loop when x+1 doesnt exist where x is the running var in the loop


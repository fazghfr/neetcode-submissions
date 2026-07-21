func longestConsecutive(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}

	exist := make(map[int]bool)
	for _, v := range nums {
		exist[v] = true
	}

	maxLen := 0
	for _, v := range nums {
		if !exist[v-1] {
			cur := 1
			vtemp := v + 1

			for exist[vtemp] {
				cur++
				vtemp++
			}
			maxLen = max(maxLen, cur)
		}
	}

	return maxLen
}



// we need to find elements that could be the start
// for each starting point, increment runningcur if and only if start+1 exist somewhere
// stop the loop when x+1 doesnt exist where x is the running var in the loop


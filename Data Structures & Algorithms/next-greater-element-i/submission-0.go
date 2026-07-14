func nextGreaterElement(nums1 []int, nums2 []int) []int {
	index1 := make(map[int]int)
	for i, v := range nums1 {
		index1[v] = i 
	}
	

	type pair struct {
		index int
		val int
	}
	var stack []pair
	ans := make([]int, len(nums1))
	for i, _ := range ans {
		ans[i] = -1
	}

	for _, curr := range nums2 {
		for len(stack) > 0 && curr > stack[len(stack)-1].val {
			popped := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			ans[popped.index] = curr
		}
		if _, ok := index1[curr]; ok {
			stack = append(stack, pair{
				index: index1[curr],
				val:   curr,
			})
		}
	}

	return ans
}

// map[value]index -> to lookup from nums1 to nums2
// nums1 = [4,1,2], nums2 = [1,3,4,2]
// // sort nums1 by index inside nums2
// [1,4,2]


// stack -> numbers need to be resolved (only push if exist in nums1)
// if current is greater than top -> pop.
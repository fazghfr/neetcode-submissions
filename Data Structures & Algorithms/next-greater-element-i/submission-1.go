func nextGreaterElement(nums1 []int, nums2 []int) []int {
	var stack []int
	exist := make(map[int]int)

	for i, v := range nums1 {
		exist[v] = i
	}

	ans := make([]int, len(nums1))
	for i, _ := range ans {
		ans[i] = -1
	}
	for _, v := range nums2 {
		for len(stack) > 0 && v > stack[len(stack)-1] {
			popped := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			ans[exist[popped]] = v
		} 
		if _, ok := exist[v]; ok {
			stack = append(stack, v)
		}
	}

	return ans
}

// find the next greater element of nums1 in nums2

// push to stack if exist in nums1
// pop when we find greater than top
// on pop, populate the answer with the greater element

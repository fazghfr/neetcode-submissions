func dailyTemperatures(temperatures []int) []int {
	var stack []int

	// this defaults to zer0
	ans := make([]int, len(temperatures))

	for i, temp := range temperatures {
		for len(stack) != 0 && temperatures[stack[len(stack)-1]] < temp {
			popped := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			ans[popped] = i - popped
		}

		stack = append(stack, i)
	}
	return ans
}

// stack approach
// dry run attempt
// incoming -> c
// 

// c
// 28
// s
// 30 solved, pop -> []
// 38 30
// 30 solved, pop -> [38]
// 38 36 35
// 35 solved, pop -> 38 36
// 36 solved, pop -> 38
// 38 solved, pop -> []
// 40 28

// END OF ARRAY

// assign all left inside stack to be 0. DONE
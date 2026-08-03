func getConcatenation(nums []int) []int {
	ans := make([]int, 2 * len(nums))
	j := 0
	for i := range ans {
		if j >= len(nums) {
			j = 0
		}
		ans[i] = nums[j]
		j++ 
	}
	return ans
}

func merge(intervals [][]int) [][]int {
    // sort the intervals first

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	fmt.Println(intervals)

	cur := intervals[0]
	var ans [][]int
	for i := 1; i < len(intervals); i++{
		test := intervals[i]
		// non overlap
		if cur[1] < test[0]  {
			ans = append(ans, cur)
			cur = test
		} else {
			// merge algo
			merging := []int{
				min(cur[0], test[0]),
				max(cur[1], test[1]),
			}
			cur = merging
		}
	}
	ans = append(ans, cur)	
	return ans
}

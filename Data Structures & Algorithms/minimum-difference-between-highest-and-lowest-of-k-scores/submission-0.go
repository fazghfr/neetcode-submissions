func minimumDifference(nums []int, k int) int {
	sort.Slice(nums, func(i, j int)bool {
		return nums[i] < nums[j]
	})

	l := 0
	ans := 100001
	for r := 0; r < len(nums); r++ {

		for r - l + 1 > k {
			l++
		}

		if r - l + 1 == k {
			ans = min(ans, nums[r] - nums[l])
		}
	}

	return ans
}

// sort ascending 
// sliding window
// shrink when k is more than three
// calculate when actual length of window is k


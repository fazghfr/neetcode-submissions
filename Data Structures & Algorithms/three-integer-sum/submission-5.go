func threeSum(nums []int) [][]int {
	sort.Slice(nums, func(i, j int)bool {
		return nums[i] < nums[j]
	})

	var ans [][]int
	for currentIdx, pivot := range nums {
		if currentIdx > 0 && nums[currentIdx] == nums[currentIdx-1]{
			continue
		}
		target := pivot * -1

		l := currentIdx + 1
		r := len(nums) - 1

		for l < r {
			if nums[l] + nums[r] == target {
				ans = append(ans, []int{pivot, nums[l], nums[r]})
				l++
				r--

				for l < r && nums[l] == nums[l-1] {
					l++
				}

				for l < r && nums[r] == nums[r+1] {
					r--
				}
			} else if nums[l] + nums[r] > target {
				r--
			} else {
				l++
			}
		}
	}
	return ans
}

// each element will be a new target elem * -1
// if found target of that element, append
// if not continue

// on continue, if same as current element, skip


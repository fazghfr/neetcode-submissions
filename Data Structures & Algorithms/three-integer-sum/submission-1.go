func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	var ans [][]int
	for i, v := range nums {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		low := i + 1
		high := len(nums) - 1

		target := v * -1

		for low < high {
			if nums[low] + nums[high] == target {
				ans = append(ans, []int{v, nums[low], nums[high]})
				low++
				high--

				for low < high && nums[low] == nums[low-1] {
					low++
				}
				for low < high && nums[high] == nums[high+1] {
					high--
				}
			} else if nums[low] + nums[high] > target {
				high--
			} else {
				low++
			}
		}
	}
	return ans
}


// two pointer approach
// sort

// [-4 -1 -1 0 1 2]
// -4 (-1, 2)
// -1 (-1, 2)  -> save [-1, -1, 2]
// -1 (0, 1)
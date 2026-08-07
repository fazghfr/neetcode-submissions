func twoSum(nums []int, target int) []int {
	onIndex := make(map[int]int)

	for i, v := range nums {
		complement := target - v

		if _, ok := onIndex[complement]; ok {
			return []int{onIndex[complement], i}
		}
		
		onIndex[v] = i
	}

	return nil
}


// hash approach

// setiap iterasi, kalo complementnya gaada simpen ke hash
// cek setiap iterasi, komplemennya ada ga. 

// kalo ada, return komplemen itu dan index sekarang



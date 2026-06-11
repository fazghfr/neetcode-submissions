func dailyTemperatures(temperatures []int) []int {
	var unresolved []int

	ans := make([]int, len(temperatures))
	for i, temp := range temperatures {
		for len(unresolved) > 0 && temp > temperatures[unresolved[len(unresolved)-1]] {
			popped := unresolved[len(unresolved)-1]
			unresolved = unresolved[:len(unresolved)-1]

			ans[popped] = i - popped
		}		
		unresolved = append(unresolved, i)
	}
	return ans
}




func carFleet(target int, position []int, speed []int) int {

	m := make(map[int]int)
	for i := range position {
		m[position[i]] = speed[i]
	}

	sort.Slice(position, func(i, j int) bool {
		return position[i] > position[j]
	})

	var myStack []float64
	for _, pos := range position {
		timeArrival := float64(target-pos) / float64(m[pos])

		if len(myStack) == 0 || timeArrival > myStack[len(myStack)-1] {
			myStack = append(myStack, timeArrival)
		}
	}

	return len(myStack)
}
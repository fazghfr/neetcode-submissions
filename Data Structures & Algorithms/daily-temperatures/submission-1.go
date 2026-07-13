func dailyTemperatures(temperatures []int) []int {
	type Pair struct {
		Temp  int
		Index int
	}

	var monoStack []Pair
	ans := make([]int, len(temperatures))

	for i, temp := range temperatures {
		for len(monoStack) > 0 && monoStack[len(monoStack)-1].Temp < temp {
			// pop
			poppedPair := monoStack[len(monoStack)-1]
			monoStack = monoStack[:len(monoStack)-1]

			ans[poppedPair.Index] = i - poppedPair.Index
		}
		monoStack = append(monoStack, Pair{
			Temp: temp,
			Index: i,
		})
	} 

	return ans
}

// monotonic stack approach
// reason : for i, the most recent bigger will be the top stack

// pop the stack when we find something bigger
// stack will consist the days that we havent found something bigger

// when bigger is found check top with for loop
// until it is smaller, push to stack
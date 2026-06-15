func compress(chars []byte) int {
    var write int
	var currNum byte
	var counting int
	for i, v := range chars {
		if i == 0 {
			// set curNum
			currNum = v
			counting++
			continue
		}

		if currNum == v {
			counting++
		} else {
			// emit
			chars[write] = currNum
			write++
			if counting > 1 {
				for _, d := range strconv.Itoa(counting) {
					chars[write] = byte(d)
					write++
				}
			}
			currNum = v
			counting = 1
		}
	}
	chars[write] = currNum
	write++
	if counting > 1 {
		for _, d := range strconv.Itoa(counting) {
			chars[write] = byte(d)
			write++
		}
	}
	return write
}

// accumulator is the string ans
// if currNum != v emit
// else, accumulate.
// return len(accumulator)

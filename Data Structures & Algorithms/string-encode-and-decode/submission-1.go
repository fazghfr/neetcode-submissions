type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	var ans string

	for _, str := range strs {
		strLen := strconv.Itoa(len(str))

		encoded := make([]rune, len(strLen)+1+len(str))

		i := 0
		for i < len(strLen) {
			encoded[i] = rune(strLen[i])
			i++
		}

		encoded[i] = '#'

		it := len(strLen) + 1
		for _, v := range str {
			encoded[it] = v
			it++
		}

		ans += string(encoded)
	}

	return ans
}

func (s *Solution) Decode(encoded string) []string {
	
	// read the number until first #
	// use that number to loop n number of times
	// 
	i:=0
	var ans []string
	for i < len(encoded) {
		var numStr string 
		for encoded[i] != '#' {
			numStr += string(encoded[i])
			i++
		} 

		num, _ := strconv.Atoi(numStr)
		i++
		start := i

		temp := make([]rune, num)
		it := 0
		for i - start + 1 <= num {
			temp[it] = rune(encoded[i])
			it++
			i++
		}

		ans = append(ans, string(temp))
	}

	return ans
	
}

// 5#hello

// 5#{ANYSTRING}

func evalRPN(tokens []string) int {
	var stack []int

	for _, v := range tokens {
		if v != "+" && v != "-" && v != "*" && v != "/" {
			num, _ := strconv.Atoi(v)
			stack = append(stack, num)
		} else {
			// pop 2 angka
			b := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			a := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			var result int

			if v == "+" {
				result = a + b
			} else if v == "-" {
				result = a - b
			} else if v == "*" {
				result = a * b
			} else {
				result = a / b
			}

			// push hasil kembali ke stack
			stack = append(stack, result)
		}
	}

	return stack[0]
}
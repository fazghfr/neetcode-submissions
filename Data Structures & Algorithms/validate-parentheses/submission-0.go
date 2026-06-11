func isValid(s string) bool {
    // push to stack if we have open brackets (, [, {
	// pop stack when we found top of stack with matching closed brackets

	var stack []rune

	for _, v := range s {
		if v == '{' || v == '[' || v == '(' {
			stack = append(stack, v)
		} else if len(stack) != 0 {
			popped := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			if popped == '{'  && v != '}' {
				return false
			} 
			if popped == '('  && v != ')' {
				return false
			}
			if popped == '['  && v != ']' {
				return false
			}
		} else {
			return false
		}
	}

	if len(stack) != 0 {return false}

	return true
}



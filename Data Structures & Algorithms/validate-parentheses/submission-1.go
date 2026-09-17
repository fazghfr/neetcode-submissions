func isValid(s string) bool {
	var st []rune

	for _, v := range s {
		if v == '{' || v == '[' || v == '(' {
			st = append(st, v)
		} else if len(st) != 0 {
			popped := st[len(st)-1]
			st = st[:len(st)-1]

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

	if len(st) != 0 {return false}
	return true
}



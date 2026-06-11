func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	s = strings.TrimSpace(s)

	var str []rune
	for _, v := range s {
		if (v >= 'a' && v <= 'z') ||  (v >= '0' && v <= '9') {
			str = append(str, rune(v))
		} 
	}

	fmt.Println(string(str))
	right := len(str) - 1
	left := 0

	for left < right {
		if str[left] != str[right] {
			return false
		}
		left++
		right--
	}
	return true
}

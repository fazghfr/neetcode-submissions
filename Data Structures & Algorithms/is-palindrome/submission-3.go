func isPalindrome(s string) bool {
	lowered := strings.ToLower(s)
	trimmed := strings.TrimSpace(lowered)
	var runifiedTrimmed []rune

	for _, v := range trimmed {
		if (rune(v) >= 'a' && rune(v) <= 'z') || (rune(v) >= '0' && rune(v) <= '9') {
			runifiedTrimmed = append(runifiedTrimmed, rune(v))
		}
	}

	finalString := string(runifiedTrimmed)

	fmt.Println(finalString)
	middle := len(finalString) / 2
	if len(finalString) % 2 == 0 {
		return expandFromMiddle(finalString, middle-1, middle)
	}
	return expandFromMiddle(finalString, middle-1, middle+1)
}

func expandFromMiddle(s string, left int, right int) bool {
	for left < right && left >= 0 && right < len(s) {
		if s[left] != s[right] {
			return false
		}
		left--
		right++
	}

	return true
}



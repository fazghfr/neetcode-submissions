func isSubsequence(s string, t string) bool {
	spointer := 0
	tpointer := 0

	if len(s) == 0 && len(t) >= 0 {
		return true
	}

	for tpointer < len(t) && spointer < len(s) {
		if s[spointer] == t[tpointer] {
			spointer++
		} 
		tpointer++
	}
	return spointer == len(s)
}
func isSubsequence(s string, t string) bool {
	spointer := 0
	tpointer := 0

	for tpointer < len(t) && spointer < len(s) {
		if s[spointer] == t[tpointer] {
			spointer++
		} 
		tpointer++
	}
	return spointer == len(s)
}
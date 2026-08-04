func minimumRecolors(blocks string, k int) int {
	freq := make(map[rune]int)

	l:=0
	minNum := len(blocks)
	for r:=0;r<len(blocks);r++{
		freq[rune(blocks[r])]++

		if r - l + 1 > k {
			freq[rune(blocks[l])]--
			l++
		} 

		if r - l + 1 == k {
			allW := freq['W']
			minNum = min(minNum, allW)
		}
	}

	return minNum
}

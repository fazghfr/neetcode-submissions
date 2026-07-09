func hoursNeeded(piles []int, k int) int {
	total := 0
	for _, pile := range piles {
		total += int(math.Ceil(float64(pile) / float64(k)))
	}

	return int(total)
}

func minEatingSpeed(piles []int, h int) int {
	maxK := 0
	for _, pile := range piles {
		if pile >= maxK {
			maxK = pile
		}
	}

	l := 1
	r := maxK

	for l < r {
		mid := l + (r - l) / 2
		if hoursNeeded(piles, mid) <= h {
			r = mid
		} else {
			l = mid + 1
		}
	}

	return l
}


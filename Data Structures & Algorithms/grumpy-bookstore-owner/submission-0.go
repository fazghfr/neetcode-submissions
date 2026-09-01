func maxSatisfied(customers []int, grumpy []int, minutes int) int {
	winMaxSat := -1
	l := 0
	runMax := 0
	lMax, rMax := -1, -1

	for r := 0; r < len(grumpy); r++ {
		if grumpy[r] == 1 {
			runMax += customers[r]
		}

		if r-l+1 > minutes {
			if grumpy[l] == 1 {
				runMax -= customers[l]
			}
			l++
		} 
		if r-l+1 == minutes {
			if runMax > winMaxSat {
				winMaxSat = runMax
				lMax, rMax = l, r
			}
		}
	}

	maxSatisfied := 0
	for i, v := range customers {
		if (i >= lMax && i <= rMax) || grumpy[i] == 0 {
			maxSatisfied += v
		}
	}

	return maxSatisfied
}


// sliding window on grumpy array
// minutes : is the window size
// customers is the "value" array
// right approach, struggle to implement
// always have clear path of your solution instead of jumping the gun

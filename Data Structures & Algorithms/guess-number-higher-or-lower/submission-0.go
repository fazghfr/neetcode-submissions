/** 
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guessNumber(n int) int {
    found := false
	l := 1
	r := n
	var ans int
	for !found {
		m := l + (r - l) / 2

		if guess(m) == 0 {
			ans = m
			found = true
		} else if guess(m) == 1 {
			l = m + 1
		} else {
			r = m - 1
		}
	}

	return ans
}


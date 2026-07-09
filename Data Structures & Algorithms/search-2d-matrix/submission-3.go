func isThisRow(row []int, target int) bool {
	return target >= row[0] && target <= row[len(row)-1]
}

func searchMatrix(matrix [][]int, target int) bool {
	l := 0
	r := len(matrix) - 1

	targetRow := -1
	for l <= r {
		mid := l + (r-l)/2

		if isThisRow(matrix[mid], target) {
			targetRow = mid
			break
		} else if target < matrix[mid][0] {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	if targetRow == -1 {
		return false
	}

	theRow := matrix[targetRow]
	l = 0
	r = len(theRow) - 1

	for l <= r {
		mid := l + (r-l)/2

		if theRow[mid] == target {
			return true
		} else if target < theRow[mid] {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return false
}
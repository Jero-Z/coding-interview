package searchingIn2DArrays

func Find(numbers [][]int, target int) bool {
	matrixWight := len(numbers)

	if matrixWight < 1 {
		return false
	}
	matrixHeight := len(numbers[0])
	if matrixHeight < 1 {

	}
	r, c := 0, matrixHeight-1
	for r < matrixWight && c >= 0 {
		if numbers[r][c] == target {
			return true
		}
		if numbers[r][c] > target {
			c--

		} else {
			r++
		}

	}
	return false
}

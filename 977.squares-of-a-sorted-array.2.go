func sortedSquares(A []int) []int {
	idxL, idxR := 0, len(A)-1
	squares := []int{}

	for idxL != idxR {
		l, r := A[idxL], A[idxR]

		if l*l > r*r {
			squares = append([]int{l * l}, squares...)
			idxL++
			continue
		}

		squares = append([]int{r * r}, squares...)
		idxR--
	}

	squares = append([]int{A[idxL] * A[idxL]}, squares...)
	return squares
}

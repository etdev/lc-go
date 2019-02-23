func sortedSquares(A []int) []int {
	idxL, idxR := 0, len(A)-1
	squares := make([]int, len(A))

	for i := len(A) - 1; i >= 0; i-- {
		l, r := (A[idxL] * A[idxL]), (A[idxR] * A[idxR])

		if l > r {
			squares[i] = l
			idxL++
			continue
		}

		squares[i] = r
		idxR--
	}
	return squares
}

func minDeletionSize(A []string) int {
	count := 0

	if len(A) == 0 {
		return 0
	}

	for i := 0; i < len([]rune(A[0])); i++ {
		for j := range A {
			if j == 0 {
				continue
			}
			if []rune(A[j-1])[i] > []rune(A[j])[i] {
				count++
				break
			}
		}
	}

	return count
}

func flipAndInvertImage(A [][]int) [][]int {
	for _, row := range A {
		// flip horiz
		for i := 0; i < len(row)/2; i++ {
			row[i], row[len(row)-1-i] = row[len(row)-1-i], row[i]
		}

		for i := 0; i < len(row); i++ {
			// invert
			if row[i] == 1 {
				row[i] = 0
				continue
			}
			row[i] = 1
		}
	}

	return A
}

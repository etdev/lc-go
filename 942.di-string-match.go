func diStringMatch(S string) []int {
	chars := []rune(S)
	indices := []int{}
	l, r := 0, len(chars)

	for _, char := range chars {
		if char == 'D' {
			indices = append(indices, r)
			r--
		} else {
			indices = append(indices, l)
			l++
		}
	}

	// Fix last char
	indices = append(indices, l)

	return indices
}

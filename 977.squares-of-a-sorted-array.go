func sortedSquares(A []int) []int {
	neg, pos := []int{}, []int{}
	ret := []int{}

	for _, a := range A {
		if a <= 0 {
			neg = append([]int{a * a}, neg...)
		} else {
			pos = append(pos, a*a)
		}
	}

	posIdx, negIdx := 0, 0
	for posIdx < len(pos) || negIdx < len(neg) {
		if posIdx >= len(pos) || negIdx >= len(neg) {
			break
		}
		if pos[posIdx] > neg[negIdx] {
			ret = append(ret, neg[negIdx])
			negIdx++
			continue
		}

		ret = append(ret, pos[posIdx])
		posIdx++
	}

	ret = append(ret, append(pos[posIdx:], neg[negIdx:]...)...)

	return ret
}

func intMax(a, b int) int {
	if a > b {
		return a
	}

	return b
}

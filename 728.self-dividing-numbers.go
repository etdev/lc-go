func selfDividingNumbers(left int, right int) []int {
	digits := []int{}

	for i := left; i <= right; i++ {
		iDigs := []int{}
		if i == 0 {
			continue
		}

		j := i
		for j > 0 {
			iDigs = append(iDigs, j%10)
			j /= 10
		}

		found := true
		for _, d := range iDigs {
			if d == 0 {
				found = false
				break
			}

			if i%d != 0 {
				found = false
			}
		}

		if found {
			digits = append(digits, i)
		}
	}

	return digits
}

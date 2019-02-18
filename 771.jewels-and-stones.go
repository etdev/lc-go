package main

func numJewelsInStones(J string, S string) int {
	count := 0
	for _, jChar := range []rune(J) {
		for _, sChar := range []rune(S) {
			if jChar == sChar {
				count++
			}
		}
	}

	return count
}

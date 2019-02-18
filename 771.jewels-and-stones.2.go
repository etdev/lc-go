package main

func numJewelsInStones(J string, S string) int {
	jewels := map[rune]bool{}
	count := 0

	// insert into jewels
	for _, chr := range []rune(J) {
		jewels[chr] = true
	}

	for _, chr := range []rune(S) {
		if jewels[chr] {
			count++
		}
	}

	return count
}

var morseCodes = []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}

func uniqueMorseRepresentations(words []string) int {
	seen := map[string]bool{}

	for _, word := range words {
		seen[toMorse(word)] = true
	}

	return len(seen)
}

func toMorse(str string) string {
	var morse string

	for _, chr := range str {
		morse += morseCodes[chr-'a']
	}

	return morse
}

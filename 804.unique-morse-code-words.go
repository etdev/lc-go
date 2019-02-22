var mappings = map[rune]string{
	'a': ".-",
	'b': "-...",
	'c': "-.-.",
	'd': "-..",
	'e': ".",
	'f': "..-.",
	'g': "--.",
	'h': "....",
	'i': "..",
	'j': ".---",
	'k': "-.-",
	'l': ".-..",
	'm': "--",
	'n': "-.",
	'o': "---",
	'p': ".--.",
	'q': "--.-",
	'r': ".-.",
	's': "...",
	't': "-",
	'u': "..-",
	'v': "...-",
	'w': ".--",
	'x': "-..-",
	'y': "-.--",
	'z': "--..",
}

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
		morse += mappings[chr]
	}

	return morse
}

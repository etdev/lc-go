func toLowerCase(str string) string {
	new := []rune{}

	for _, chr := range str {
		if chr >= 'A' && chr <= 'Z' {
			new = append(new, chr+32)
			continue
		}

		new = append(new, chr)
	}

	return string(new)
}

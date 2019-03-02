func judgeCircle(moves string) bool {
	chars := []rune(moves)
	horiz, vert := 0, 0

	for _, char := range chars {
		switch char {
		case 'L':
			horiz--
		case 'R':
			horiz++
		case 'D':
			vert--
		case 'U':
			vert++
		}
	}

	return (horiz == 0 && vert == 0)
}

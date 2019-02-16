package main

func isValid(str string) bool {
	parens := map[rune]rune{')': '(', ']': '[', '}': '{'}
	var stack []rune

	for _, char := range str {
		if char == '(' || char == '[' || char == '{' {
			stack = append(stack, char)
			continue
		}

		if len(stack) > 0 && parens[char] == stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}

	return len(stack) == 0
}

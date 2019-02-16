package main

type stack []rune

func (s stack) push(chr rune) stack {
	return append(s, chr)
}

func (s stack) pop() (stack, rune) {
	l := len(s)
	if l < 1 {
		return s, ' '
	}

	return s[:l-1], s[l-1]
}

func isValid(str string) bool {
	stk := stack{}
	found := ' '

	for _, chr := range str {
		if chr == '(' || chr == '[' || chr == '{' {
			stk = stk.push(chr)
		}

		if chr == ')' {
			stk, found = stk.pop()
			if found != '(' {
				return false
			}
		}
		if chr == ']' {
			stk, found = stk.pop()
			if found != '[' {
				return false
			}
		}
		if chr == '}' {
			stk, found = stk.pop()
			if found != '{' {
				return false
			}
		}
	}

	if len(stk) != 0 {
		return false
	}

	return true
}

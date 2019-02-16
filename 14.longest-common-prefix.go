package main

import (
	"strings"
)

func longestCommonPrefix(strs []string) string {
	var current string
	for _, str := range strs {
		if str == "" {
			return ""
		}

		if current == "" {
			current = str
			continue
		}

		for !strings.HasPrefix(str, current) {
			runeCur := []rune(current)
			current = string(runeCur[:len(runeCur)-1])

			if current == "" {
				return ""
			}
		}
	}

	return current
}

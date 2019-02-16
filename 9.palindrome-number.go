package main

import "fmt"

func isPalindrome(x int) bool {
	xStr := fmt.Sprintf("%d", x)
	return reverseStr(xStr) == xStr
}

func reverseStr(str string) string {
	runes := []rune(str)

	for i := 0; i < len(runes)/2; i++ {
		runes[i], runes[len(runes)-1-i] = runes[len(runes)-1-i], runes[i]
	}

	return string(runes)
}

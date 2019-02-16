package main

import (
	"strings"
)

type mapping struct {
	code  string
	value int
}

func romanToInt(s string) int {
	mappings := []mapping{
		{"M", 1000},
		{"CM", 900},
		{"D", 500},
		{"CD", 400},
		{"C", 100},
		{"XC", 90},
		{"L", 50},
		{"XL", 40},
		{"X", 10},
		{"IX", 9},
		{"V", 5},
		{"IV", 4},
		{"I", 1},
	}

	sum := 0
	runes := []rune(s)

	for _, mapping := range mappings {
		for strings.HasPrefix(string(runes), mapping.code) {
			sum += mapping.value
			runes = append(runes[:0], runes[len(mapping.code):]...)
		}
	}

	return sum
}

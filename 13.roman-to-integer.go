package main

import (
	"strings"
)

func romanToInt(s string) int {
	mappings := map[string]int{
		"M":  1000,
		"CM": 900,
		"D":  500,
		"CD": 400,
		"C":  100,
		"XC": 90,
		"L":  50,
		"XL": 40,
		"X":  10,
		"IX": 9,
		"V":  5,
		"IV": 4,
		"I":  1,
	}

	letterGroups := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	sum := 0
	runes := []rune(s)

	for _, letters := range letterGroups {
		for strings.HasPrefix(string(runes), letters) {
			sum += mappings[letters]
			runes = append(runes[:0], runes[len(letters):]...)
		}
	}

	return sum
}

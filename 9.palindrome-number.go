package main

// Version without string conversion
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	x2 := x

	y := 0
	for {
		if x2 <= 0 {
			break
		}

		y = y*10 + (x2 % 10)
		x2 /= 10
	}

	return x == y
}

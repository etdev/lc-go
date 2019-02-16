package main

import "math"

func reverse(x int) int {
	y := 0
	for {
		if x == 0 {
			break
		}
		y = (y * 10) + x%10
		x /= 10
	}

	if willOverflow(y) {
		return 0
	}

	return y
}

func willOverflow(x int) bool {
	if (float64(x) < -math.Pow(2, 31)) || (float64(x) > math.Pow(2, 31)-1) {
		return true
	}
	return false
}

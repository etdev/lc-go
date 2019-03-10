import "fmt"

func hammingDistance(x int, y int) int {
	xord := fmt.Sprintf("%b", x^y)

	count := 0
	for _, chr := range xord {
		if chr == '1' {
			count++
		}
	}

	return count
}

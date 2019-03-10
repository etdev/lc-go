import (
	"fmt"
	"strconv"
)

func hammingDistance(x int, y int) int {
	xbin, ybin := fmt.Sprintf("%b", x), fmt.Sprintf("%b", y)

	maxLen := len(xbin)
	if len(ybin) > maxLen {
		maxLen = len(ybin)
	}

	maxLens := strconv.Itoa(maxLen)
	xbin, ybin = fmt.Sprintf("%0"+maxLens+"b", x), fmt.Sprintf("%0"+maxLens+"b", y)

	count := 0
	xbinr, ybinr := []rune(xbin), []rune(ybin)
	for i := 0; i < maxLen; i++ {
		if xbinr[i] != ybinr[i] {
			count++
		}
	}

	return count
}

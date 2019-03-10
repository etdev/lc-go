import "strings"

func commonChars(A []string) []string {
	strs := []string{}

	for _, chr := range A[0] {
		str := string(chr)

		count := 0
		for _, x := range strs {
			if string(x) == str {
				count++
			}
		}

		found := true
		for i := 1; i < len(A); i++ {
			if strings.Count(A[i], str) <= count {
				found = false
			}
		}

		if found {
			strs = append(strs, str)
		}
	}

	return strs
}

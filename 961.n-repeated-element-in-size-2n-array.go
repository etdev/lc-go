import "sort"

func repeatedNTimes(A []int) int {
	sort.Ints(A)
	last := A[0]

	for i, a := range A {
		if i < 1 {
			continue
		}

		if a == last {
			return a
		}
		last = a
	}

	return 0
}

func topKFrequent(nums []int, k int) []int {
	// create map of nums to count (counts)
	counts := map[int]int{}
	for _, num := range nums {
		counts[num]++
	}

	// create buckets of count => array_of_nums_that_appeared_count_times (buckets)
	buckets := map[int][]int{}

	// go through counts and add each num to the correct bucket
	for num, count := range counts {
		// caveat: must be unique
		if !containsInt(buckets[count], num) {
			buckets[count] = append(buckets[count], num)
		}
	}

	// add bucket results to an array
	returnInts := []int{}

	for i := len(nums); i >= 0; i-- {
		returnInts = append(returnInts, buckets[i]...)
	}

	return returnInts[:k]
}

func containsInt(nums []int, k int) bool {
	for _, num := range nums {
		if num == k {
			return true
		}
	}

	return false
}

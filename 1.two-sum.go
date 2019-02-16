package main

func twoSum(nums []int, target int) []int {
	// num => idx
	seen := map[int]int{}

	for i, num := range nums {
		if _, ok := seen[target-num]; ok {
			return []int{seen[target-num], i}
		}

		seen[num] = i
	}

	return []int{}
}

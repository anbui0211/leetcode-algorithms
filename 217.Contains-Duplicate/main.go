package main



func containsDuplicate(nums []int) bool {
	mv := make(map[int]int)
	for _, num := range nums {
		mv[num]++
		if mv[num] > 1 {
			return true
		}
	}

	return false
}

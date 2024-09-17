package main

func main() {
}

func twoSumV1(nums []int, target int) []int {
	for i := 0; i <= len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}

	return []int{}
}

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, num := range nums {
		complement := target - num
		// Check j:= m[complement] có trong map không
		// Nếu có => return i, j
		// Nếu không , add m[num] = i (lưu lại index)

		if j, ok := m[complement]; ok {
			return []int{j, i}
		}

		m[num] = i
	}

	return []int{}
}

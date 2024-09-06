# Code

```go
func intersect(nums1 []int, nums2 []int) []int {
	mp := make(map[int]int)

	// Step 1: count accurences in nums1
	for _, v := range nums1 {
		mp[v]++
	}

	// Step 2: Find the intersection with nums2
	var result []int
	for _, v := range nums2 {
		if mp[v] > 0 {
			result = append(result, v)
			mp[v]--
		}
	}

	return result
}
```

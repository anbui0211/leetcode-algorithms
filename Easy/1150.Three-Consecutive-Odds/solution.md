# Code

#### Solution 1
```go
func threeConsecutiveOdds(arr []int) bool {
	mp := make(map[int]bool)
	for i := 0; i < len(arr); i++ {
		if arr[i]%2 != 0 {
			mp[arr[i]] = true
		}

		if i > 1 && mp[arr[i]] == true {
			if mp[arr[i-1]] && mp[arr[i-2]] {
				return true
			}
		}
	}
	return false
}
```
#### Solution 2
```go
func threeConsecutiveOddsV2(arr []int) bool {
	for i := 0; i < len(arr); i++ {
		if i < 2 {
			continue
		}
		if arr[i]%2 != 0 && arr[i-1]%2 != 0 && arr[i-2]%2 != 0 {
			return true
		}
	}
	return false
}
```
#### Solution 3
```go
func threeConsecutiveOddsV3(arr []int) bool {
	streak := 0

	for i := 0; i < len(arr); i++ {
		if arr[i]%2 != 0 {
			streak++
			if streak == 3 {
				return true
			}
		} else {
			streak = 0
		}
	}

	return false
}
```

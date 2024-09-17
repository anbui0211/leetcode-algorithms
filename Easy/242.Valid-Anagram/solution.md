# Code

#### Solution 1

```go
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	mp := make(map[rune]int)

	for _, v := range s {
		mp[v]++
	}
	for _, v := range t {
		mp[v]--
	}
	for _, v := range mp {
		if v != 0 {
			return false
		}
	}
	return true
}
```

#### Solution 2

```go
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	charMap := make(map[rune]int, len(s))

	for _, val := range s {
		charMap[val]++
	}

	for _, val := range t {
		if count, exist := charMap[val]; exist && count > 0 {
			charMap[val]--
		} else {
			return false
		}
	}

	return true
}

```

#### Solution 3

```go
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	count := make(map[byte]int, len(s))

	for i := 0; i < len(s); i++ {
		count[s[i]]++
		count[t[i]]--
	}

	for _, v := range count {
		if v != 0 {
			return false
		}
	}

	return true
}
```

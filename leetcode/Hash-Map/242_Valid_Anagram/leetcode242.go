package leetcode242

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	if s == t {
		return true
	}

	counter := make(map[rune]int)

	// Count the occurrences of each character within `s`
	for _, char := range s {
		counter[char]++
	}

	// Subtract the occurrences while iterating through `t`
	for _, char := range t {
		counter[char]--
		if counter[char] < 0 {
			return false
		}
	}
	return true
}

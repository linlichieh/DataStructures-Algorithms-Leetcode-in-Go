package leetcode3

func lengthOfLongestSubstring(s string) int {
	if len(s) < 2 {
		return len(s)
	}

	// For checking if the character is repeated within the window
	dup := make(map[rune]bool)

	// Use 2-pointer and sliding window approach
	var longest, ptr0 int
	for ptr1, char := range s {
		for dup[char] {
			// if char is repeated, remove the value from the map by key `ptr0` and move `ptr0` forward
			delete(dup, rune(s[ptr0]))
			ptr0++
		}

		// Save the longest length of the window
		length := ptr1 - ptr0 + 1
		if length > longest {
			longest = length
		}

		// Mark the char as existed and move the `ptr1` forward
		dup[char] = true
		ptr1++
	}

	return longest
}

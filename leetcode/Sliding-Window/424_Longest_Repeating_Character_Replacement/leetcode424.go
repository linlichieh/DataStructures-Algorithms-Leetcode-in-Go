package leetcode424

func characterReplacement(s string, k int) int {
	var longest int
	sMap := make(map[rune]int)
	left := 0
	mostFreq := 0
	for right, char := range s {
		// increase the counter by char
		sMap[char]++
		// find the most frequent char
		if sMap[char] > mostFreq {
			mostFreq = sMap[char]
		}

		// make sure the window has to pass the rule
		length := right - left + 1

		// NOTE Don't need `loop` here for passing the rule and the `mostFreq` doesn't need to be updated here as well.
		//      Becasue the window will be shorted by moving the `left` forward and the length is also shortened by 1
		//      that makes the longest remain the same.
		//
		//      Even though the window doesn't pass the rule, it's fine, as the longest won't be updated,
		//      it will only be updated until the rule pass and more than the current longest.
		//		For example, the `A B A C D E`
		if length-mostFreq > k {
			// decrease the counter by char
			sMap[rune(s[left])]--
			left++
			length--

		}

		// record the longest window length
		if length > longest {
			longest = length
		}
	}
	return longest
}

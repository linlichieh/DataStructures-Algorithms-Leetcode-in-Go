package leetcode76

func minWindow(s string, t string) string {
	sMap, tMap := make(map[rune]int), make(map[rune]int)
	for _, char := range t {
		tMap[char]++
	}

	// Save the shortest substring
	var shortest string

	// Save the position of matched characters
	queue := []int{}

	// the num of matched characters
	matched := 0

	curr := 0
	for curr < len(s) {
		char := rune(s[curr])

		// if the char is in the t string
		tVal, ok := tMap[char]
		if ok {
			sMap[char]++

			// if the counter of char in the window map is same as tMap, increase 1
			if sMap[char] == tVal {
				matched++
			}

			// save the char into the queue
			queue = append(queue, curr)
		}

		// if matched number achieves the critiria, get the shortest substring within the window
		for matched == len(tMap) {
			// Pop the first matched characters from the queue
			ptr := queue[0]
			queue = queue[1:]

			// Campare the current substring with the shortest substring
			tmpStr := s[ptr : curr+1]
			if shortest == "" || len(tmpStr) < len(shortest) {
				shortest = tmpStr
			}

			// if char counter is same as tMap counterpart, decrease 1 because it's going to be lower than critiria in later operation
			if sMap[rune(s[ptr])] == tMap[rune(s[ptr])] {
				matched--
			}
			// Update sMap
			sMap[rune(s[ptr])]--
		}

		curr++
	}
	return shortest
}

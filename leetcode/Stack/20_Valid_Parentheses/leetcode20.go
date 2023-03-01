package leetcode20

func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}
	return validate(s, 0, &[]rune{})
}

func validate(s string, pos int, stack *[]rune) bool {
	// If 'pos' has reached the end, which means the whole process is done, check if there are any items left in the stack.
	if pos == len(s) {
		return len(*stack) == 0
	}

	char := rune(s[pos])
	pos++

	// Check if the string contains any forbidden characters
	switch char {
	case '(', '[', '{':
		// If it's open bracket, put it into the stack
		*stack = append(*stack, char)
		return validate(s, pos, stack)
	case ')', ']', '}':
		// If there is no any open brackets in tht stack or the open bracket in the stack doesn't match the `char`
		if len(*stack) == 0 || pair((*stack)[len(*stack)-1]) != char {
			return false
		}

		// If `pos` hasn't reached the end, pop the item from the stack
		*stack = (*stack)[:len(*stack)-1]
		return validate(s, pos, stack)
	}
	return false
}

func pair(char rune) rune {
	switch char {
	case '(':
		return ')'
	case '[':
		return ']'
	case '{':
		return '}'
	}
	return 0
}

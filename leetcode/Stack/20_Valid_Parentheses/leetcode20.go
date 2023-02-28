package leetcode20

const (
	OPEN  = int(1)
	CLOSE = int(0)
)

var m = map[string]int{
	"(": OPEN,
	"[": OPEN,
	"{": OPEN,
	")": CLOSE,
	"]": CLOSE,
	"}": CLOSE,
}

var pair = map[string]string{
	"(": ")",
	"[": "]",
	"{": "}",
}

func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}
	return validate(s, 0, &[]string{})
}

func validate(s string, pos int, stack *[]string) bool {
	if pos == len(s) {
		return false
	}

	char := string(s[pos])
	pos++

	// Check if the string contains any forbidden characters
	typ, ok := m[char]
	if !ok {
		return false
	}

	// If it's open bracket, put it into the stack
	if typ == OPEN {
		*stack = append(*stack, char)
		return validate(s, pos, stack)
	}

	// If it's a close bracket, but there are no items in the stack.
	if len(*stack) < 1 {
		return false
	}

	// The topmost item on the stack should be the matching pair., for example `(` -> `)`, can't be `(` -> `]`
	peek := (*stack)[len(*stack)-1]
	if pair[peek] != char {
		return false
	}

	// If 'pos' has reached the end, check if there are any items left on the stack.
	if pos == len(s) {
		if len(*stack) > 1 {
			return false
		}
		return true
	}

	// If `pos` hasn't reached the end, pop the item from the stack
	*stack = (*stack)[:len(*stack)-1]
	return validate(s, pos, stack)
}

package leetcode739

// hint: save the index of the array into the stack to get the days
// by subtracting the values in stack from the value of current temperature if it's bigger than the value of Peek()
func dailyTemperatures(temperatures []int) []int {
	var stack []int
	answer := make([]int, len(temperatures))
	for i, temp := range temperatures {
		// If the current temperature is greater than the most recent day in the stack (the topmost value)
		for len(stack) > 0 && temp > temperatures[stack[len(stack)-1]] {
			// pop the value
			lastTempIdx := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			// To get the gap of days by subtracting last index from current index
			answer[lastTempIdx] = i - lastTempIdx
		}

		// Push the index instead of the value
		stack = append(stack, i)
	}
	return answer
}

package leetcode11

func maxArea(height []int) int {
	// Set 2 pointers at both ends
	left, right := 0, len(height)-1

	var maximum int
	for left != right {
		// calculate the area
		dimension := calculate(&height, left, right)
		if dimension > maximum {
			maximum = dimension
		}

		// Move the smaller height as it can never lead to a larger area
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return maximum
}

func calculate(arr *[]int, l, r int) int {
	height := (*arr)[l]
	if height > (*arr)[r] {
		height = (*arr)[r]
	}
	width := r - l
	return height * width
}

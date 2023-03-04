package countingSort

func counting(arr []int) []int {
	if len(arr) == 0 {
		return []int{}
	}

	// Find the minimum and maximum
	var min, max int
	for _, v := range arr {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}

	// Fill count array with the number of occurrences
	countSize := max - min + 1
	countArr := make([]int, countSize, countSize)
	for _, val := range arr {
		// index = value - minimum (value)
		countArr[val-min]++
	}

	// Sum up values with predecessor values
	// NOTE The sum values in the array is related to the index of sorted array
	for i := 1; i < countSize; i++ {
		countArr[i] += countArr[i-1]
	}

	// Create the sorted array
	sortedArr := make([]int, len(arr), len(arr))
	for _, val := range arr {
		// Get the the index of count array by subtracting min from value of original array
		countIdx := val - min
		// Get the value of count arry
		// NOTE Subtract 1 in order to get the correct index of sorted array
		sortedIdx := countArr[countIdx] - 1
		sortedArr[sortedIdx] = val
		// value of count array minus 1
		// NOTE It's for assigning the value of original array to the correct index of sorted array if the value exists 1+ times
		countArr[countIdx]--
	}

	return sortedArr
}

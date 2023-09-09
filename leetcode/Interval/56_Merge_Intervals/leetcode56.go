package leetcode56

import "sort"

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return [][]int{}
	}

	// Sort the intervals based on the starting value
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	// Initialise the merged intervals with the first interval
	merged := [][]int{intervals[0]}

	// Iterate through all the sorted intervals, starting from the second interval
	for i := 1; i < len(intervals); i++ {
		// Get the last interval from the merged
		lastIntFromMerged := merged[len(merged)-1]
		// Get the current interval
		currInt := intervals[i]
		// Check if the starting value of the current interval is within the range of the last intervals from the merged
		// If yes, it means they are overlapped
		if currInt[0] <= lastIntFromMerged[1] {
			// If the ending value of the current interval is greater than the ending value of the last intervals from the merged
			if currInt[1] > lastIntFromMerged[1] {
				lastIntFromMerged[1] = currInt[1]
			}
			continue
		}
		// If the starting value of the current interval is out of the range of the last intervals from the merged, it means that it's non-overlapping interval
		merged = append(merged, currInt)
	}

	return merged
}

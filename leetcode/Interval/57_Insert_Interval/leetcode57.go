package leetcode57

func insert(intervals [][]int, newInterval []int) [][]int {
	var merged [][]int
	if len(intervals) == 0 {
		return append(merged, newInterval)
	}

	// Put the intervals that don't overlap with the new interval into the merged list
	curr := 0
	for curr < len(intervals) && intervals[curr][1] < newInterval[0] {
		merged = append(merged, intervals[curr])
		curr++
	}

	// Iterate through intervals while merging the current interval with the new interval
	for curr < len(intervals) && intervals[curr][0] <= newInterval[1] {
		if newInterval[0] > intervals[curr][0] {
			newInterval[0] = intervals[curr][0]
		}
		if newInterval[1] < intervals[curr][1] {
			newInterval[1] = intervals[curr][1]
		}
		curr++
	}
	// Add the merged interval that is done by the previous step
	merged = append(merged, newInterval)

	// Add the remaining intevals into the merged list
	merged = append(merged, intervals[curr:]...)
	return merged
}

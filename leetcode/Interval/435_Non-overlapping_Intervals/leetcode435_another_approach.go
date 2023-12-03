package leetcode435

import "sort"

// sort by start point
func anotherApproach(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}

	// Sort intervals based on start point
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	var result int
	prevEnd := intervals[0][1]
	for _, interval := range intervals[1:] {
		// if previous end is less than current start
		if prevEnd <= interval[0] {
			// update the end
			prevEnd = interval[1]
		} else {
			// overlapping case
			result++
			// tackle overlapping cases, please see README
			if interval[1] < prevEnd {
				prevEnd = interval[1]
			}
		}
	}

	return result
}

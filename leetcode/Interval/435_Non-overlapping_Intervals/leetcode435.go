package leetcode435

import "sort"

// sort by end point
func eraseOverlapIntervals(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}

	// Sort intervals based on end point
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})

	var result int
	prevEnd := intervals[0][1]
	for _, interval := range intervals[1:] {
		// if previous end is less than current start
		if interval[0] < prevEnd {
			result++
		} else {
			// tackle overlapping cases, please see README
			prevEnd = interval[1]
		}
	}

	return result
}

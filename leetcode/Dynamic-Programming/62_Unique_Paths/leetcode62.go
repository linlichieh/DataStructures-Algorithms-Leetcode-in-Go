package leetcode62

func uniquePaths(m int, n int) int {
	if m <= 0 || n <= 0 {
		return 0
	}

	// set the default value as it's on the edge so they are all 1
	cache := make([]int, n)
	for i := range cache {
		cache[i] = 1
	}

	// Skip the last row, because it's initialised as default cache
	for i := m - 2; i >= 0; i-- {
		// NOTE skip the last one column (n-1), because it's always 1
		for j := n - 2; j >= 0; j-- {
			// add current value with its next one
			cache[j] += cache[j+1]
		}
	}

	// the first value will be the total number of unique path
	return cache[0]
}

package leetcode70

func climbStairs(n int) int {
	return add(n, &[]int{0, 1, 2, 3})
}

func add(idx int, cache *[]int) int {
	switch {
	case idx < len(*cache):
		return (*cache)[idx]
	}
	// add the result into the cache slice, so the same number won't be recomputed again
	*cache = append(*cache, add(idx-1, cache)+add(idx-2, cache))
	return (*cache)[idx]
}

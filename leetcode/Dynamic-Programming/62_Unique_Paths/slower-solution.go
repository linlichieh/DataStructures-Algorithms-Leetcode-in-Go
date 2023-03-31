package leetcode62

func slowerSolution(m int, n int) int {
	if m == 0 || n == 0 {
		return 0
	}
	if m == 1 || n == 1 {
		return 1
	}

	// Initialise the cache table
	cache := make([][]int, m)
	for i := 0; i < m; i++ {
		cache[i] = make([]int, n)
	}
	return dfs(0, 0, cache)
}

func dfs(i int, j int, cache [][]int) int {
	// out of range
	if i >= len(cache) || j >= len(cache[0]) {
		return 0
	}
	// if it's on the edge, the route is always one
	if i == len(cache)-1 || j == len(cache[0])-1 {
		return 1
	}
	// add right and down up and cache it
	if cache[i][j] == 0 {
		cache[i][j] = dfs(i+1, j, cache) + dfs(i, j+1, cache)
	}
	return cache[i][j]
}

package leetcode63

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	row, col := len(obstacleGrid), len(obstacleGrid[0])
	if row == 0 || col == 0 {
		return 0
	}

	// Initialise an array with the length of the number of columns
	cache := make([]int, col)

	// Set the default value to the end of columns as it's on the edge and there's only one way to go
	cache[col-1] = 1

	// Iterate through the grid in reverse order
	for i := row - 1; i >= 0; i-- {
		for j := col - 1; j >= 0; j-- {
			// If it's an obstacle, set the value to 0
			if obstacleGrid[i][j] == 1 {
				cache[j] = 0
				continue
			}
			// If it's on the edge, skip it as it's always 1 unless it's an obstacle
			if j+1 == col {
				continue
			}
			// Add current value with the next one, what it really does is similar to `grid[i][j] = grid[i+1][j] + grid[i][j+1]`
			cache[j] += cache[j+1]
		}
	}
	return cache[0]
}

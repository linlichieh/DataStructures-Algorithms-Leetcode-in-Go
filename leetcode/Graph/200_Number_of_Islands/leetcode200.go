package leetcode200

func numIslands(grid [][]byte) int {
	var islands int
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[0]); col++ {
			if grid[row][col] == '1' {
				dfs(row, col, grid)

				// The visited elements will be changed to `0`, so that it won't be counted twice
				islands++
			}
		}
	}
	return islands
}

func dfs(row, col int, grid [][]byte) {
	if row < 0 || col < 0 || row >= len(grid) || col >= len(grid[0]) || grid[row][col] == '0' {
		return
	}

	// Mark the element as visited
	grid[row][col] = '0'

	// To traverse other `1` in adjancent elements
	dfs(row-1, col, grid)
	dfs(row+1, col, grid)
	dfs(row, col-1, grid)
	dfs(row, col+1, grid)
}

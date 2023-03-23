package leetcode417

func pacificAtlantic(heights [][]int) [][]int {
	if len(heights) == 0 {
		return [][]int{}
	}
	rows, cols := len(heights), len(heights[0])

	// Create visited matrices for pacific ocean and atlantic ocean
	pacVisited := make([][]bool, rows)
	atlVisited := make([][]bool, rows)

	// Initialise the matrices
	for row := 0; row < rows; row++ {
		pacVisited[row] = make([]bool, cols)
		atlVisited[row] = make([]bool, cols)
	}

	for col := 0; col < cols; col++ {
		// top edge
		dfs(0, col, pacVisited, heights, 0)

		// bottom edge
		dfs(rows-1, col, atlVisited, heights, 0)
	}

	for row := 0; row < rows; row++ {
		// left edge
		dfs(row, 0, pacVisited, heights, 0)

		// right edge
		dfs(row, cols-1, atlVisited, heights, 0)
	}

	var result [][]int
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			// If the coordination exists in both matrices, the rain can flow from the cell to both oceans.
			if pacVisited[row][col] && atlVisited[row][col] {
				result = append(result, []int{row, col})
			}
		}
	}

	return result
}
func dfs(row, col int, visited [][]bool, heights [][]int, prevHei int) {
	if row < 0 || col < 0 || row >= len(heights) || col >= len(heights[0]) || visited[row][col] || heights[row][col] < prevHei {
		return
	}
	visited[row][col] = true
	dfs(row+1, col, visited, heights, heights[row][col])
	dfs(row-1, col, visited, heights, heights[row][col])
	dfs(row, col+1, visited, heights, heights[row][col])
	dfs(row, col-1, visited, heights, heights[row][col])
}

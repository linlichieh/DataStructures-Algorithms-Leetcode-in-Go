package leetcode79

func exist(board [][]byte, word string) bool {
	if len(board) == 0 {
		return false
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if backtrack(i, j, board, 0, word) {
				return true
			}
		}
	}
	return false
}

func backtrack(i, j int, board [][]byte, charIdx int, word string) bool {
	// If the charIdx reaches the end of the word, it means that the word exists
	if charIdx >= len(word) {
		return true
	}
	if i < 0 || j < 0 || i >= len(board) || j >= len(board[0]) || board[i][j] == '-' || board[i][j] != byte(word[charIdx]) {
		return false
	}

	// Save the char for later restoration
	char := board[i][j]

	// Mark it as visited instead of using extra matrix to store
	board[i][j] = '-'

	result := backtrack(i+1, j, board, charIdx+1, word) ||
		backtrack(i-1, j, board, charIdx+1, word) ||
		backtrack(i, j+1, board, charIdx+1, word) ||
		backtrack(i, j-1, board, charIdx+1, word)

	// Restore the character as unvisited
	board[i][j] = char

	return result
}

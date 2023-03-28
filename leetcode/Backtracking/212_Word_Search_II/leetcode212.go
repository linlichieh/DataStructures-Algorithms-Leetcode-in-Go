package leetcode212

import "fmt"

type Trie struct {
	children [26]*Trie // array is slightly faster than `map[rune]*Trie`
	word     string    // to store the whole word for better performance
}

func (t *Trie) insert(s string) {
	var idx int
	curr := t
	for _, char := range s {
		idx = int(char - 'a')
		// If the character doesn't exist in the children, then new a Trie
		if curr.children[idx] == nil {
			curr.children[idx] = &Trie{}
		}
		// Point to its children
		curr = curr.children[idx]
	}
	curr.word = s
}

// For debugging
func (t *Trie) print(prefix string) {
	for char, tr := range t.children {
		if tr == nil {
			continue
		}
		var endOfWord bool
		if tr.word != "" {
			endOfWord = true
		}
		newPrefix := fmt.Sprintf("%s -> %s(%v)", prefix, string(rune(char+'a')), endOfWord)
		tr.print(newPrefix)
		if endOfWord {
			fmt.Printf("%s\n", newPrefix)
		}
	}
}

func findWords(board [][]byte, words []string) []string {
	if len(board) == 0 {
		return []string{}
	}

	// Add all words into the trie
	trie := &Trie{}
	for _, word := range words {
		trie.insert(word)
	}

	result := []string{}

	// Traverse each cell
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			backtrack(i, j, board, trie, &result)
		}
	}

	return result
}

func backtrack(row, col int, board [][]byte, tr *Trie, result *[]string) {
	if row < 0 || col < 0 || row >= len(board) || col >= len(board[0]) || board[row][col] == '-' {
		return
	}

	// If it doesn't have any children, then exit
	if len(tr.children) == 0 {
		return
	}

	// If it doesn't match any character, then exit
	char := board[row][col]
	tr = tr.children[int(char-'a')]
	if tr == nil {
		return
	}

	// This way significantly improves the performance as it doesn't need to store the path and pass it into the DFS call each time.
	if tr.word != "" {
		// If this attribute isn't empty, it means it's the end of word.
		*result = append(*result, tr.word)

		// Mark as visited for avoiding the same path being counted twice like `[oa oa oaa]`
		tr.word = ""
	}

	// Prevent the cell from been accessed twice
	board[row][col] = '-'
	backtrack(row+1, col, board, tr, result)
	backtrack(row-1, col, board, tr, result)
	backtrack(row, col+1, board, tr, result)
	backtrack(row, col-1, board, tr, result)
	// Change it back at the end so it won't affect the next cell traversal
	board[row][col] = char
}

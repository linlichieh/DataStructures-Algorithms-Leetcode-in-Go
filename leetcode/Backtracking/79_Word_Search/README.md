# Idea

The solution is similar to `200. Number of Islands` or `417. Pacific Atlantic Water Flow`.

Begin by iterating over each cell in the board by starting a backtracking search (DFS) from that cell.
Check if the current character is matched and the adjacent cells (in all four directions) to find the next letter in the word.
If the next letter is not found in a given direction, backtrack.
Otherwise, continue the process until all characters are found.

To avoid revisiting cells, mark cells that have already been visited by changing their value to a special character,
and change them back to their original value after the search is complete.


Time: `O( m * n * dfs)` = `O(m*n*3^n)`

* n indicates the lengh of the word
* We only explore up to 3 directions for each of the `n` characters in the word, because we are searching for the next character of the word in the adjacent cells (not including the cell that it comes from)

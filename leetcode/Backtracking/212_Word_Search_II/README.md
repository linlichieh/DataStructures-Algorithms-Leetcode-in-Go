# Idea

 Use a Trie data structure to store all the words in the given dictionary.
 Then, use a DFS algorithm to traverse the given board and search for the words in the Trie.
 If the character doesn't exist in the children of current path or no children, then backtrack.
 Once the word is found, then aad it to the result and mark as visited to avoid the same path being counted twice.
 Mark the cell as visited to prevent it from being accessed twice. Change it back to the original character after each DFS call.

* Time complexity: `O(M * N * 4^L + K * L)` = `O(M * N * 4^L)`
    * The additional term of `K * L` comes from building the Trie
    * Should it be 3 instead of `4` due to the cell it comes from has been marked as visited?
    * `4^L` instead of `4^L*(word length)`, is it because it goes through the trie in one path for searching words each time?
* Space complexity: `O(K * L)`

> M is the number of rows in the board, N is the number of columns in the board, L is the maximum length of the words, and K is the number of words in the dictionary.

Keys for better performance

* Use `word string` instead of `endOfWord bool`.
    * It significantly improves the performance as the DFS call doesn't need to pass the `path` each time.
    * significantly faster by 450ms
* Use `[26]*Trie` instead of `map[rune]*Trie`
    * slightly faster by 4ms

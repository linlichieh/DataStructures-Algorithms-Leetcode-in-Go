package trie

import "fmt"

type Trie struct {
	children  map[rune]*Trie
	endOfWord bool
}

func NewTrie() *Trie {
	return &Trie{
		children:  make(map[rune]*Trie),
		endOfWord: false,
	}
}

func (t *Trie) Insert(word string) {
	current := t
	for _, char := range word {
		if current.children[char] == nil {
			current.children[char] = NewTrie()
		}
		current = current.children[char]
	}
	current.endOfWord = true
}

func (t *Trie) Delete(word string) {
	if len(word) == 0 {
		return
	}

	// Call the helper function to delete the word from the Trie
	t.del(word, 0)

	// DEBUG
	// t.Print()
}

func (t *Trie) del(word string, idx int) {
	// Get the character for the current index
	char := rune(word[idx])
	// Get the node for the current character
	node, ok := t.children[char]
	// If the character doesn't exist in the Trie, return
	if !ok {
		return
	}

	// If the current index is the last character of the word
	if idx == len(word)-1 {
		// Set the endOfWord flag to false
		node.endOfWord = false
	} else {
		// Call the delete function for the next character
		node.del(word, idx+1)
	}

	// DEBUG
	// If the node doesn't have any children and it's not the end of any word
	// fmt.Printf("%c children: %d, endOfWord: %t\n", char, len(node.children), node.endOfWord)
	if len(node.children) == 0 && !node.endOfWord {
		// DEBUG
		// fmt.Printf("del %c\n", char)

		// Remove it from the map
		delete(t.children, char)
	}
}

// whether the string is a word
func (t *Trie) Search(word string) bool {
	current := t
	for _, char := range word {
		if current.children[char] == nil {
			return false
		}
		current = current.children[char]
	}
	return current.endOfWord
}

// whether the prefix exists
func (t *Trie) HasPrefix(word string) bool {
	if len(word) == 0 {
		return false
	}

	current := t
	for _, char := range word {
		if current.children[char] == nil {
			return false
		}
		current = current.children[char]
	}
	return true
}

func (t *Trie) Print() {
	t.print("")
}

func (t *Trie) print(prefix string) {
	for char, node := range t.children {
		newPrefix := fmt.Sprintf("%s %c(%t)", prefix, char, node.endOfWord)
		if node.endOfWord {
			fmt.Println(newPrefix)
		}
		node.print(newPrefix)
	}
}

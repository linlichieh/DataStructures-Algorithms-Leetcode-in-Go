package leetcode208

type TrieNode struct {
	children  [26]*TrieNode
	endOfWord bool
}

type Trie struct {
	root *TrieNode
}

func Constructor() Trie {
	return Trie{root: &TrieNode{}}
}

func (this *Trie) Insert(word string) {
	curr := this.root
	var index rune
	for _, char := range word {
		index = char - 'a'
		if curr.children[index] == nil {
			curr.children[index] = &TrieNode{}
		}
		curr = curr.children[index]
	}
	curr.endOfWord = true
}

func (this *Trie) Search(word string) bool {
	curr := this.root
	var index rune
	for _, char := range word {
		index = char - 'a'
		if curr.children[index] == nil {
			return false
		}
		curr = curr.children[index]
	}
	return curr.endOfWord
}

func (this *Trie) StartsWith(prefix string) bool {
	curr := this.root
	var index rune
	for _, char := range prefix {
		index = char - 'a'
		if curr.children[index] == nil {
			return false
		}
		curr = curr.children[index]
	}
	return true
}

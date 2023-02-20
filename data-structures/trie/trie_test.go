package trie

import "testing"

func TestInsertAndSearch(t *testing.T) {
	cases := []struct {
		words  []string
		find   string
		search bool
	}{
		{
			words:  []string{},
			find:   "",
			search: false,
		},
		{
			words:  []string{"ddd"},
			find:   "d",
			search: false,
		},
		{
			words:  []string{"ddd"},
			find:   "ddd",
			search: true,
		},
		{
			words:  []string{"ddd", "ddd"},
			find:   "ddd",
			search: true,
		},
		{
			words:  []string{"cat", "care"},
			find:   "ca",
			search: false,
		},
		{
			words:  []string{"cat", "care"},
			find:   "cat",
			search: true,
		},
		{
			words:  []string{"cat", "care"},
			find:   "care",
			search: true,
		},
		{
			words:  []string{"cat", "care"},
			find:   "brilliant",
			search: false,
		},
		{
			words:  []string{"cat", "care", "brilliant"},
			find:   "care",
			search: true,
		},
		{
			words:  []string{"cat", "care", "brilliant"},
			find:   "brilliant",
			search: true,
		},
	}

	for _, c := range cases {
		tr := NewTrie()
		for _, word := range c.words {
			tr.Insert(word)
		}
		result := tr.Search(c.find)
		if result != c.search {
			t.Errorf("expect '%v', but got '%v'", c.search, result)
		}
	}
}

func TestHasPrefix(t *testing.T) {

	cases := []struct {
		words     []string
		find      string
		hasPrefix bool
	}{
		{
			words:     []string{},
			find:      "",
			hasPrefix: false,
		},
		{
			words:     []string{"ddd"},
			find:      "d",
			hasPrefix: true,
		},
		{
			words:     []string{"ddd"},
			find:      "ddd",
			hasPrefix: true,
		},
		{
			words:     []string{"ddddddd"},
			find:      "ddd",
			hasPrefix: true,
		},
		{
			words:     []string{"cat", "care"},
			find:      "ca",
			hasPrefix: true,
		},
		{
			words:     []string{"cat", "care"},
			find:      "cat",
			hasPrefix: true,
		},
		{
			words:     []string{"cat", "care"},
			find:      "care",
			hasPrefix: true,
		},
		{
			words:     []string{"cat", "care"},
			find:      "caree",
			hasPrefix: false,
		},
	}

	for i, c := range cases {
		tr := NewTrie()
		for _, word := range c.words {
			tr.Insert(word)
		}
		result := tr.HasPrefix(c.find)
		if result != c.hasPrefix {
			t.Errorf("(%d) expect '%v', but got '%v'", i, c.hasPrefix, result)
		}
	}
}

func TestDelete(t *testing.T) {
	cases := []struct {
		words       []string
		del         string
		find        string
		wordExisted bool
	}{
		{
			words:       []string{},
			del:         "",
			find:        "",
			wordExisted: false,
		},
		{
			words:       []string{"ddd"},
			del:         "",
			find:        "dd",
			wordExisted: false,
		},
		{
			words:       []string{"ddd"},
			del:         "",
			find:        "ddd",
			wordExisted: true,
		},
		{
			words:       []string{"ddd"},
			del:         "d",
			find:        "d",
			wordExisted: false,
		},
		{
			words:       []string{"ddd"},
			del:         "d",
			find:        "ddd",
			wordExisted: true,
		},
		{
			words:       []string{"ddd"},
			del:         "ddd",
			find:        "d",
			wordExisted: false,
		},
		{
			words:       []string{"ddd"},
			del:         "ddd",
			find:        "ddd",
			wordExisted: false,
		},
		{
			words:       []string{"ddd", "dd"},
			del:         "ddd",
			find:        "dd",
			wordExisted: true,
		},
		{
			words:       []string{"ddd", "dd"},
			del:         "dd",
			find:        "dd",
			wordExisted: false,
		},
		{
			words:       []string{"ddd", "dd"},
			del:         "dd",
			find:        "ddd",
			wordExisted: true,
		},
		{
			words:       []string{"cat", "tr", "trid", "trie", "trief"},
			del:         "trie",
			find:        "trie",
			wordExisted: false,
		},
		{
			words:       []string{"cat", "tr", "trid", "trie", "trief"},
			del:         "trie",
			find:        "trid",
			wordExisted: true,
		},
		{
			words:       []string{"cat", "tr", "trid", "trie", "trief"},
			del:         "trief",
			find:        "trie",
			wordExisted: true,
		},
		{
			words:       []string{"cat", "tr", "trid", "trie", "trief"},
			del:         "trief",
			find:        "trief",
			wordExisted: false,
		},
		{
			words:       []string{"cat", "tr", "trid", "trie", "trief"},
			del:         "tr",
			find:        "trie",
			wordExisted: true,
		},
	}

	for i, c := range cases {
		tr := NewTrie()
		for _, word := range c.words {
			tr.Insert(word)
		}
		tr.Delete(c.del)
		wordExisted := tr.Search(c.find)
		if c.wordExisted != wordExisted {
			t.Errorf("(%d) expect '%v', but got '%v'", i, c.wordExisted, wordExisted)
		}
	}
}

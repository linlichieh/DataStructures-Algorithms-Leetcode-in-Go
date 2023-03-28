package leetcode212

import (
	"reflect"
	"testing"
)

func TestFindWords(t *testing.T) {
	cases := []struct {
		board  [][]byte
		words  []string
		output []string
	}{
		{
			board:  [][]byte{},
			words:  []string{},
			output: []string{},
		},
		{
			board: [][]byte{
				{'o', 'a', 'a', 'n'},
				{'e', 't', 'a', 'e'},
				{'i', 'h', 'k', 'r'},
				{'i', 'f', 'l', 'v'},
			},
			words:  []string{"oath", "pea", "eat", "rain"},
			output: []string{"oath", "eat"},
		},
		{
			board:  [][]byte{{'a', 'b', 'c', 'd'}},
			words:  []string{"abcd"},
			output: []string{"abcd"},
		},
		{
			board:  [][]byte{{'a', 'b', 'c', 'd'}},
			words:  []string{"ab", "ac", "ad", "cd"},
			output: []string{"ab", "cd"},
		},
		{
			board: [][]byte{
				{'a', 'b'},
				{'c', 'd'},
			},
			words:  []string{"abd", "abcd", "abc", "adc", "adb", "acdb"},
			output: []string{"acdb", "abd"}, // it fails sometimes due to the order
		},
		{
			board: [][]byte{
				{'o', 'a', 'b', 'n'},
				{'o', 't', 'a', 'e'},
				{'a', 'h', 'k', 'r'},
				{'a', 'f', 'l', 'v'},
			},
			words:  []string{"oa", "oaa"},
			output: []string{"oa", "oaa"},
		},
	}

	for i, c := range cases {
		result := findWords(c.board, c.words)
		if !reflect.DeepEqual(c.output, result) {
			t.Errorf("(%d) expect '%v', but got '%v'", i, c.output, result)
		}
	}
}

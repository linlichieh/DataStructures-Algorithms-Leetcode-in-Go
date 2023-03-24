package leetcode79

import (
	"reflect"
	"testing"
)

func TestExist(t *testing.T) {
	cases := []struct {
		board [][]byte
		word  string
		exist bool
	}{
		{
			board: [][]byte{
				{'A', 'B', 'C', 'E'},
				{'S', 'F', 'C', 'S'},
				{'A', 'D', 'E', 'E'},
			},
			word:  "AFE",
			exist: false,
		},
		{
			board: [][]byte{
				{'A', 'B', 'C', 'E'},
				{'S', 'F', 'C', 'S'},
				{'A', 'D', 'E', 'E'},
			},
			word:  "ABCB",
			exist: false,
		},
		{
			board: [][]byte{
				{'A', 'B', 'C', 'E'},
				{'S', 'F', 'C', 'S'},
				{'A', 'D', 'E', 'E'},
			},
			word:  "ABCED",
			exist: false,
		},
		{
			board: [][]byte{
				{'A', 'B', 'C', 'E'},
				{'S', 'F', 'C', 'S'},
				{'A', 'D', 'E', 'E'},
			},
			word:  "ABCES",
			exist: true,
		},
		{
			board: [][]byte{
				{'A', 'B', 'C', 'E'},
				{'S', 'F', 'C', 'S'},
				{'A', 'D', 'E', 'E'},
			},
			word:  "ABCCED",
			exist: true,
		},
		{
			board: [][]byte{
				{'A', 'B', 'C', 'E'},
				{'S', 'F', 'C', 'S'},
				{'A', 'D', 'E', 'E'},
			},
			word:  "ASFDECCESE",
			exist: true,
		},
		{
			board: [][]byte{
				{'A', 'B', 'C', 'E'},
				{'S', 'F', 'C', 'S'},
				{'A', 'D', 'E', 'E'},
			},
			word:  "SEE",
			exist: true,
		},
		{
			board: [][]byte{
				{'A', 'B', 'C', 'E'},
				{'S', 'F', 'C', 'S'},
				{'A', 'D', 'E', 'E'},
			},
			word:  "CBFB",
			exist: false,
		},
		{
			board: [][]byte{
				{'A', 'B', 'C', 'E'},
				{'S', 'F', 'C', 'S'},
				{'A', 'D', 'E', 'E'},
			},
			word:  "CBFDAS",
			exist: true,
		},
		{
			board: [][]byte{
				{'C', 'E', 'S', 'D'},
				{'C', 'E', 'B', 'S'},
			},
			word:  "CEES",
			exist: true,
		},
	}

	for i, c := range cases {
		result := exist(c.board, c.word)
		if !reflect.DeepEqual(c.exist, result) {
			t.Errorf("(%d) expect '%v', but got '%v'", i, c.exist, result)
		}
	}

}
